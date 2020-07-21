package server

import (
	"fmt"
	"net/http"

	"aws-mahjong/server/handler"
	"aws-mahjong/usecase"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

func AttachHandlerAndEvent(router *mux.Router, wsserver *socketio.Server, roomUsecase usecase.RoomUsecase, gameUsecase usecase.GameUsecase) {
	// api handlers
	router.HandleFunc("/rooms", handler.Rooms(roomUsecase)).Methods(http.MethodGet)

	// room events
	router.Handle("/socket.io/", wsserver)

	wsserver.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		return nil
	})

	wsserver.OnDisconnect("/", func(s socketio.Conn, _ string) {
		fmt.Println("disconnected:", s.ID())
		roomUsecase.LeaveAllRoom(s)
	})
}
