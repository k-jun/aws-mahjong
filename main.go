package main

import (
	"net/http"
	"os"

	"aws-mahjong/server"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	server.AttachHandlerAndEvent(wsserver)

	go wsserver.Serve()
	defer wsserver.Close()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
