package handler

import (
	"aws-mahjong/usecase"
	"encoding/json"

	socketio "github.com/googollee/go-socket.io"
)

func roomStatus(roomUsecase usecase.RoomUsecase, roomName string) (string, error) {

	roomInfo, err := roomUsecase.Room(roomName)
	if err != nil {
		return "", err
	}

	status := NewRoomStatus{
		RoomName:        roomInfo.Name,
		RoomMemberCount: roomInfo.Len,
		RoomCapacity:    roomInfo.Capacity,
	}
	resBody, err := json.Marshal(&status)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}

type RoomErrorResponse struct {
	EventName    string `json:"event_name"`
	ErrorMessage string `json:"error_message"`
}

func roomError(s socketio.Conn, eventName string, errorMessage string) string {
	resBody := RoomErrorResponse{
		EventName:    eventName,
		ErrorMessage: errorMessage,
	}

	bytes, _ := json.Marshal(&resBody)
	return string(bytes)
}
