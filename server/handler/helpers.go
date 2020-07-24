package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	RoomName = "room_name"
)

func ExtractBody(r *http.Request, body interface{}) error {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, body)
}

func ExtractPathParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// func roomStatus(roomUsecase usecase.RoomUsecase, roomName string) (string, error) {
//
// 	roomInfo, err := roomUsecase.Room(roomName)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	status := NewRoomStatus{
// 		RoomName:        roomInfo.Name,
// 		RoomMemberCount: roomInfo.Len,
// 		RoomCapacity:    roomInfo.Capacity,
// 	}
// 	resBody, err := json.Marshal(&status)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return string(resBody), nil
// }
//
// type RoomErrorResponse struct {
// 	EventName    string `json:"event_name"`
// 	ErrorMessage string `json:"error_message"`
// }
//
// func websocketError(eventName string, errorMessage string) string {
// 	resBody := RoomErrorResponse{
// 		EventName:    eventName,
// 		ErrorMessage: errorMessage,
// 	}
//
// 	bytes, _ := json.Marshal(&resBody)
// 	return string(bytes)
// }
