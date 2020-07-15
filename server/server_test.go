// +build integration

package server

import (
	"aws-mahjong/server/event"
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
	uri       = "http://localhost:8000"
	socketUri = uri + "/socket.io/"
)

func TestMain(m *testing.M) {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	AttachHandlerAndEvent(wsserver)

	go wsserver.Serve()
	defer wsserver.Close()
	go http.ListenAndServe(":8000", nil)
	os.Exit(m.Run())
}

func TestOnConnect(t *testing.T) {
	_, err := socketio_client.NewClient(uri, opts)
	assert.NoError(t, err)
}

func TestCreateRoom(t *testing.T) {
	client, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		panic(err)
	}

	cases := []struct {
		Description string
		InBody      string
	}{
		{
			Description: "single channel message",
			InBody:      `{"user_name": "Malcolm Ferry", "room_name": "repellendus", "room_capacity": 2}`,
		},
		{
			Description: "multi channel message",
			InBody:      `{"user_name": "Malcolm Ferry", "room_name": "repellendus", "room_capacity": 2}`,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			client.Emit(event.CreateRoom, c.InBody)
		})
	}
}
