package main

import (
	"net/http"
	"os"

	"aws-mahjong/repository"
	"aws-mahjong/server"
	"aws-mahjong/usecase"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

func main() {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	roomRepo := repository.NewRoomRepository()

	roomUsecase := usecase.NewRoomUsecase(roomRepo)
	router := mux.NewRouter()
	server.AttachHandlerAndEvent(router, wsserver, roomUsecase)

	go wsserver.Serve()
	defer wsserver.Close()

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
