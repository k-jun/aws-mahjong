package handler

import (
	"aws-mahjong/server/event"
	"aws-mahjong/tile"
	"aws-mahjong/usecase"
	"encoding/json"
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

type GameDahaiRequest struct {
	RoomName string
	Dahai    string
}

func GameDahai(gameUsecase usecase.GameUsecase) func(socketio.Conn, string) {
	return func(s socketio.Conn, bodyStr string) {
		body := GameDahaiRequest{}
		err := json.Unmarshal([]byte(bodyStr), &body)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.GameError, websocketError(event.GameDahai, err.Error()))
		}
		dahai, err := tile.TileFromString(body.Dahai)
		if err != nil {
			fmt.Println(err)
			s.Emit(event.GameError, websocketError(event.GameDahai, err.Error()))
		}
		if err = gameUsecase.Dahai(body.RoomName, dahai); err != nil {
			fmt.Println(err)
			s.Emit(event.GameError, websocketError(event.GameDahai, err.Error()))
		}
	}
}
