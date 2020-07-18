// +build integration

package server

import (
	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"
	"aws-mahjong/testutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func TestGameStart(t *testing.T) {
	client1, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client2, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client3, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client4, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		Description         string
		CurrentCreateClient *socketio_client.Client
		CurrentCreateRoom   handler.CreateRoomRequest
		CurrentJoinClients  []*socketio_client.Client
		CurrentJoinRoom     handler.JoinRoomRequest
		InClient            *socketio_client.Client
		InBody              string
		OutClient           *socketio_client.Client
	}{
		{
			Description:         "valid case",
			CurrentCreateClient: client1,
			CurrentCreateRoom:   handler.CreateRoomRequest{RoomName: "Shields.Isai", UserName: "Victor Lynch DDS", RoomCapacity: 4},
			CurrentJoinClients:  []*socketio_client.Client{client2, client3},
			CurrentJoinRoom:     handler.JoinRoomRequest{RoomName: "Shields.Isai", UserName: "Lisa DuBuque"},
			InClient:            client4,
			InBody:              `{"room_name": "Shields.Isai", "user_name": "Miss Hertha Casper V"}`,
			OutClient:           client2,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			testutil.CreateRoom(c.CurrentCreateClient, c.CurrentCreateRoom)
			testutil.JoinRooms(c.CurrentJoinClients, c.CurrentJoinRoom)

			IsFired := false

			c.OutClient.On(event.GameStart, func(payload string) {
				IsFired = true
			})
			c.InClient.Emit(event.JoinRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, true, IsFired)
		})
	}
}
