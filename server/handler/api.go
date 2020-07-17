package handler

import (
	"aws-mahjong/usecase"
	"encoding/json"
	"net/http"
)

type RoomsResponse struct {
	RoomName        string `json:"room_name"`
	RoomCapacity    int    `json:"room_capacity"`
	RoomMemberCount int    `json:"room_member_count"`
}

func Rooms(roomUsecase *usecase.RoomUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			MethodNotAllowed(w, r)
			return
		}
		resBodyRooms := []RoomsResponse{}

		for _, roomInfo := range roomUsecase.Rooms() {
			resBodyRooms = append(resBodyRooms, RoomsResponse{
				RoomName:        roomInfo.Name,
				RoomMemberCount: roomInfo.Len,
				RoomCapacity:    roomInfo.Capacity,
			})
		}

		bytes, err := json.Marshal(&resBodyRooms)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		if _, err := w.Write(bytes); err != nil {
			InternalServerError(w, r)
			return
		}
	}
}
