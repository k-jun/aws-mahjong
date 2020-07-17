package server

import (
	"fmt"
	"net/http"

	"aws-mahjong/repository"
	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"
	"aws-mahjong/usecase"

	socketio "github.com/googollee/go-socket.io"
)

func AttachHandlerAndEvent(wsserver *socketio.Server, roomRepo *repository.RoomRepository, gameRepo repository.GameRepository) {

	roomUsecase := usecase.NewRoomUsecase(gameRepo, roomRepo)

	// api handlers
	http.HandleFunc("/rooms", handler.Rooms(roomUsecase))

	// events
	http.Handle("/socket.io/", wsserver)
	wsserver.OnEvent("/", event.CreateRoom, handler.CreateRoom(roomUsecase))
	wsserver.OnEvent("/", event.JoinRoom, handler.JoinRoom(roomUsecase))

	wsserver.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		return nil
	})

	wsserver.OnDisconnect("/", func(s socketio.Conn, _ string) {
		fmt.Println("disconnected:", s.ID())
		s.LeaveAll()
	})
}
