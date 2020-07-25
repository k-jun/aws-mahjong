package handler

import (
	"aws-mahjong/server/view"
	"aws-mahjong/tile"
	"aws-mahjong/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func Dahai(roomUsecase usecase.GameUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := view.DahaiRequest{}
		err := ExtractBody(r, &body)
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}

		hai, err := tile.TileFromString(body.Hai)
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}
		params := ExtractPathParams(r)
		status, err := roomUsecase.Dahai(body.UserId, params[RoomName], hai)
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}
		bytes, err := json.Marshal(status)
		if err != nil {
			log.Println(err)
			InternalServerError(w, r)
			return
		}
		if _, err = w.Write(bytes); err != nil {
			log.Println(err)
			InternalServerError(w, r)
			return
		}

	}
}
