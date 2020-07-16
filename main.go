package main

import (
	"net/http"
	"os"

	"aws-mahjong/repository"
	"aws-mahjong/server"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRoomRepository(wsserver)
	server.AttachHandlerAndEvent(wsserver, repo)

	go wsserver.Serve()
	defer wsserver.Close()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
