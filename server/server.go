package server

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func rooms(wsserver *socketio.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			MethodNotAllowed(w, r)
			return
		}
		fmt.Println(wsserver.Rooms("/"))
	}
}

func AttachHandlerAndEvent(wsserver *socketio.Server) {
	// api handlers
	http.HandleFunc("/rooms", rooms(wsserver))

	// events
	http.Handle("/socket.io/", wsserver)
	// wsserver.OnEvent("/", event.LeaveChannel, LeaveChannel)

	wsserver.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		return nil
	})

	wsserver.OnDisconnect("/", func(s socketio.Conn, _ string) {
		fmt.Println("disconnected:", s.ID())
		s.LeaveAll()
	})
}
