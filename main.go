package main

import (
	"net/http"
	"os"

	"aws-mahjong/repository"
	"aws-mahjong/server"
	"aws-mahjong/usecase"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	roomRepo := repository.NewRoomRepository(wsserver)
	gameRepo := repository.NewGameRepository()

	roomUsecase := usecase.NewRoomUsecase(roomRepo, gameRepo)
	server.AttachHandlerAndEvent(wsserver, roomUsecase)

	go wsserver.Serve()
	defer wsserver.Close()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
