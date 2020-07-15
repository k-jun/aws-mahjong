package testutil

import (
	"aws-mahjong/server/event"
	"time"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func SampleRoomCreate(client *socketio_client.Client, roomName string) {

	body := `{"user_name": "Elaina Prosacco IV", "room_name": "` + roomName + `", "room_capacity": 4}`

	client.Emit(event.CreateRoom, body)
	time.Sleep(1 * time.Second)
}
