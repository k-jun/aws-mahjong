package handler

import (
	"aws-mahjong/server/view"
	"aws-mahjong/usecase"
	"encoding/json"
	"net/http"
)

func Rooms(roomUsecase usecase.RoomUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms := roomUsecase.Rooms()
		resBody := view.NewRoomResponse(rooms)

		bytes, err := json.Marshal(&resBody)
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
