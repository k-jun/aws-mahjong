package server

import (
	"fmt"
	"net/http"

	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"
	"aws-mahjong/usecase"

	socketio "github.com/googollee/go-socket.io"
)

func AttachHandlerAndEvent(wsserver *socketio.Server, roomUsecase usecase.RoomUsecase, gameUsecase usecase.GameUsecase) {

	// api handlers
	http.HandleFunc("/rooms", handler.Rooms(roomUsecase))

	// events
	http.Handle("/socket.io/", wsserver)
	wsserver.OnEvent("/", event.CreateRoom, handler.CreateRoom(roomUsecase))
	wsserver.OnEvent("/", event.JoinRoom, handler.JoinRoom(roomUsecase))
	wsserver.OnEvent("/", event.LeaveRoom, handler.LeaveRoom(roomUsecase))

	wsserver.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		return nil
	})

	wsserver.OnDisconnect("/", func(s socketio.Conn, _ string) {
		fmt.Println("disconnected:", s.ID())
		roomUsecase.LeaveAllRoom(s)
	})
}
