package testutil

import (
	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"
	"encoding/json"
	"time"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func CreateRoom(client *socketio_client.Client, room handler.CreateRoomRequest) {
	body, _ := json.Marshal(&room)

	client.Emit(event.CreateRoom, string(body))
	time.Sleep(1 * time.Second)
}

func CreateRooms(client *socketio_client.Client, rooms []handler.CreateRoomRequest) {
	for _, room := range rooms {
		CreateRoom(client, room)
	}
}

func JoinRoom(client *socketio_client.Client, room handler.JoinRoomRequest) {
	body, _ := json.Marshal(&room)
	client.Emit(event.JoinRoom, string(body))
	time.Sleep(1 * time.Second)
}

func JoinRooms(clients []*socketio_client.Client, room handler.JoinRoomRequest) {
	for _, client := range clients {
		JoinRoom(client, room)
	}
}
