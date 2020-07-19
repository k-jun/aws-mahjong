package handler

import (
	"aws-mahjong/server/event"
	"aws-mahjong/usecase"
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

func CreateRoom(roomUsecase usecase.RoomUsecase) func(socketio.Conn, string) {
	return func(s socketio.Conn, bodyStr string) {
		var body CreateRoomRequest

		err := json.Unmarshal([]byte(bodyStr), &body)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.CreateRoom, err.Error()))
			return
		}

		if err = roomUsecase.CreateRoom(s, body.UserName, body.RoomName, body.RoomCapacity); err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.CreateRoom, err.Error()))
			return
		}

		resBody, err := roomStatus(roomUsecase, body.RoomName)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.NewRoomStatus, err.Error()))
			return
		}

		roomUsecase.NewRoomStatus(body.RoomName, resBody)
	}
}

type JoinRoomRequest struct {
	UserName string `json:"user_name"`
	RoomName string `json:"room_name"`
}

func JoinRoom(roomUsecase usecase.RoomUsecase) func(socketio.Conn, string) {
	return func(s socketio.Conn, bodyStr string) {
		var body JoinRoomRequest
		err := json.Unmarshal([]byte(bodyStr), &body)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.JoinRoom, err.Error()))
			return
		}
		if err = roomUsecase.JoinRoom(s, body.UserName, body.RoomName); err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.JoinRoom, err.Error()))
			return
		}

		resBody, err := roomStatus(roomUsecase, body.RoomName)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.NewRoomStatus, err.Error()))
			return
		}
		roomUsecase.NewRoomStatus(body.RoomName, resBody)
	}
}

type LeaveRoomRequest struct {
	RoomName string `json:"room_name"`
}

func LeaveRoom(roomUsecase usecase.RoomUsecase) func(socketio.Conn, string) {
	return func(s socketio.Conn, bodyStr string) {
		body := LeaveRoomRequest{}
		err := json.Unmarshal([]byte(bodyStr), &body)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.LeaveRoom, err.Error()))
			return
		}

		if err = roomUsecase.LeaveRoom(s, body.RoomName); err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.LeaveRoom, err.Error()))
			return
		}

		resBody, err := roomStatus(roomUsecase, body.RoomName)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.RoomError, websocketError(event.NewRoomStatus, err.Error()))
			return
		}
		roomUsecase.NewRoomStatus(body.RoomName, resBody)
	}
}
