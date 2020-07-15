package server

import (
	"fmt"
	"net/http"

	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"

	socketio "github.com/googollee/go-socket.io"
)

func AttachHandlerAndEvent(wsserver *socketio.Server) {
	// api handlers
	http.HandleFunc("/rooms", handler.Rooms(wsserver))

	// events
	http.Handle("/socket.io/", wsserver)
	wsserver.OnEvent("/", event.CreateRoom, handler.CreateRoom(wsserver))
	wsserver.OnEvent("/", event.JoinRoom, handler.JoinRoom(wsserver))

	wsserver.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		return nil
	})

	wsserver.OnDisconnect("/", func(s socketio.Conn, _ string) {
		fmt.Println("disconnected:", s.ID())
		s.LeaveAll()
	})
}
