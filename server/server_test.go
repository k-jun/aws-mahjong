// +build integration

package server

import (
	"aws-mahjong/repository"
	"aws-mahjong/usecase"
	"context"
	"net/http"
	"testing"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

var (
	opts = &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	uri       = "http://localhost:8000"
	socketUri = uri + "/socket.io/"
)

func TestMain(m *testing.M) {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	roomRepo := repository.NewRoomRepository(wsserver)
	gameRepo := repository.NewGameRepository()
	roomUsecase := usecase.NewRoomUsecase(roomRepo, gameRepo)
	gameUsecase := usecase.NewGameUsecase(roomRepo, gameRepo)
	router := mux.NewRouter()

	AttachHandlerAndEvent(router, wsserver, roomUsecase, gameUsecase)

	go wsserver.Serve()
	defer wsserver.Close()
	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	go srv.ListenAndServe()
	defer srv.Shutdown(context.Background())
}
