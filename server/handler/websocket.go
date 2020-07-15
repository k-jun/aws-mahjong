package handler

import (
	"encoding/json"
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

type CreateRoomRequest struct {
	UserName     string `json:"user_name"`
	RoomName     string `json:"room_name"`
	RoomCapacity int    `json:"room_capacity"`
}

func CreateRoom(wsserver *socketio.Server) func(socketio.Conn, string) {
	return func(s socketio.Conn, bodyStr string) {
		var body CreateRoomRequest

		err := json.Unmarshal([]byte(bodyStr), &body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(body)

		if wsserver.RoomLen("/", body.RoomName) != 0 {
			fmt.Println("room_name already token")
			return
		}

		s.Join(body.RoomName)
	}

}
