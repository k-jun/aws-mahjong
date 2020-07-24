package handler

import (
	"aws-mahjong/server/view"
	"aws-mahjong/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func Rooms(roomUsecase usecase.RoomUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms := roomUsecase.Rooms()
		resBody := view.NewRoomResponse(rooms)

		bytes, err := json.Marshal(&resBody)
		if err != nil {
			log.Println(err)
			InternalServerError(w, r)
			return
		}

		if _, err := w.Write(bytes); err != nil {
			log.Println(err)
			InternalServerError(w, r)
			return
		}
	}
}

func CreateRoom(roomUsecase usecase.RoomUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := view.CreateRoomRequest{}
		err := ExtractBody(r, &body)
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}

		status, err := roomUsecase.CreateRoom(body.UserId, body.UserName, body.RoomName, body.RoomCapacity)
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}

		resBody := view.NewStatusRoomResponse(status)
		bytes, err := json.Marshal(resBody)
		if err != nil {
			log.Println(err)
			InternalServerError(w, r)
			return
		}
		if _, err = w.Write(bytes); err != nil {
			log.Println(err)
			InternalServerError(w, r)
		}

	}
}

func JoinRoom(roomUsecase usecase.RoomUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := view.JoinRoomRequest{}
		err := ExtractBody(r, &body)
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}
		params := ExtractPathParams(r)
		status, err := roomUsecase.JoinRoom(body.UserId, body.UserName, params[RoomName])
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}

		resBody := view.NewStatusRoomResponse(status)
		bytes, err := json.Marshal(resBody)
		if err != nil {
			log.Println(err)
			InternalServerError(w, r)
			return
		}
		if _, err = w.Write(bytes); err != nil {
			log.Println(err)
			InternalServerError(w, r)
		}

	}
}

func LeaveRoom(roomUsecase usecase.RoomUsecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := view.LeaveRoomRequest{}
		err := ExtractBody(r, &body)
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}
		params := ExtractPathParams(r)
		err = roomUsecase.LeaveRoom(body.UserId, body.UserName, params[RoomName])
		if err != nil {
			log.Println(err)
			BadRequest(w, r)
			return
		}
	}
}
