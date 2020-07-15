package handler

import (
	"encoding/json"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

type RoomsResponse struct {
	RoomName        string `json:"room_name"`
	RoomCapacity    int    `json:"room_capacity"`
	RoomMemberCount int    `json:"room_member_count"`
}

func Rooms(wsserver *socketio.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			MethodNotAllowed(w, r)
			return
		}
		resBodyRooms := []RoomsResponse{}

		for _, roomName := range rooms(wsserver) {
			resBodyRooms = append(resBodyRooms, RoomsResponse{
				RoomName:        roomName,
				RoomMemberCount: roomLen(wsserver, roomName),
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
