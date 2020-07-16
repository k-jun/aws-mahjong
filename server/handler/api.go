package handler

import (
	"aws-mahjong/storage"
	"encoding/json"
	"net/http"
)

type RoomsResponse struct {
	RoomName        string `json:"room_name"`
	RoomCapacity    int    `json:"room_capacity"`
	RoomMemberCount int    `json:"room_member_count"`
}

func Rooms(stg *storage.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			MethodNotAllowed(w, r)
			return
		}
		resBodyRooms := []RoomsResponse{}

		for _, roomName := range stg.Rooms() {
			resBodyRooms = append(resBodyRooms, RoomsResponse{
				RoomName:        roomName,
				RoomMemberCount: stg.RoomLen(roomName),
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
