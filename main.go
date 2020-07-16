package main

import (
	"net/http"
	"os"

	"aws-mahjong/server"
	"aws-mahjong/storage"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	stg := storage.NewStorage(wsserver)
	server.AttachHandlerAndEvent(wsserver, stg)

	go wsserver.Serve()
	defer wsserver.Close()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
