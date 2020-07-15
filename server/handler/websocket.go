package handler

import (
	"aws-mahjong/server/event"
	"encoding/json"
	"errors"
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

var (
	RoomAlraedyTokenErr = errors.New("room already token")
	RoomNotFound        = errors.New("room is not found")
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
			s.Emit(event.CreateRoomError, err.Error())
			return
		}

		if roomLen(wsserver, body.RoomName) != 0 {
			fmt.Println("room_name already token")
			s.Emit(event.CreateRoomError, "")
			return
		}
		joinRoom(s, body.RoomName)
	}
}

type JoinRoomRequest struct {
	UserName string `json:"user_name"`
	RoomName string `json:"room_name"`
}

func JoinRoom(wsserver *socketio.Server) func(socketio.Conn, string) {
	return func(s socketio.Conn, bodyStr string) {
		var body JoinRoomRequest
		err := json.Unmarshal([]byte(bodyStr), &body)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.JoinRoomError, err.Error())
			return
		}

		if roomLen(wsserver, body.RoomName) == 0 {
			fmt.Println(RoomNotFound)
			s.Emit(event.JoinRoomError, RoomNotFound.Error())
			return
		}

		joinRoom(s, body.RoomName)
	}
}
