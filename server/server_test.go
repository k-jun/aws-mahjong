// +build integration

package server

import (
	"aws-mahjong/repository"
	"aws-mahjong/usecase"
	"net/http"
	"os"
	"testing"

	socketio "github.com/googollee/go-socket.io"
	"github.com/stretchr/testify/assert"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

var (
	opts = &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	uri         = "http://localhost:8000"
	socketUri   = uri + "/socket.io/"
	roomUsecase usecase.RoomUsecase
)

func TestMain(m *testing.M) {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	roomRepo := repository.NewRoomRepository(wsserver)
	gameRepo := repository.NewGameRepository()
	roomUsecase = usecase.NewRoomUsecase(roomRepo, gameRepo)

	AttachHandlerAndEvent(wsserver, roomUsecase)

	go wsserver.Serve()
	defer wsserver.Close()
	go http.ListenAndServe(":8000", nil)
	os.Exit(m.Run())
}

func TestOnConnect(t *testing.T) {
	_, err := socketio_client.NewClient(uri, opts)
	assert.NoError(t, err)
}
