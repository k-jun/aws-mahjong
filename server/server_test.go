// +build integration

package server

import (
	"aws-mahjong/usecase"
	"errors"
	"os"
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
	wsserver  *socketio.Server
)

func TestMain(m *testing.M) {
	err := errors.New("")
	wsserver, err = socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	go wsserver.Serve()
	defer wsserver.Close()
	os.Exit(m.Run())
}

func makeServer(roomUsecase usecase.RoomUsecase) *mux.Router {
	router := mux.NewRouter()
	AttachHandlerAndEvent(router, wsserver, roomUsecase)
	return router
}
