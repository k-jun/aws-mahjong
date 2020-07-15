package handler

import (
	"fmt"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func Rooms(wsserver *socketio.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			MethodNotAllowed(w, r)
			return
		}
		fmt.Println(wsserver.Rooms("/"))
	}
}
