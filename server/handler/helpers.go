package handler

import (
	"aws-mahjong/usecase"
	"encoding/json"
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
