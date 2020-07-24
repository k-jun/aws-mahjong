package handler

type GameDahaiRequest struct {
	RoomName string `json:"room_name"`
	Dahai    string `json:"dahai"`
}

// func GameDahai(gameUsecase usecase.GameUsecase) func(socketio.Conn, string) {
// 	return func(s socketio.Conn, bodyStr string) {
// 		body := GameDahaiRequest{}
// 		err := json.Unmarshal([]byte(bodyStr), &body)
// 		if err != nil {
// 			fmt.Println(err)
// 			s.Emit(event.GameError, websocketError(event.GameDahai, err.Error()))
// 		}
// 		dahai, err := tile.TileFromString(body.Dahai)
// 		if err != nil {
// 			fmt.Println(err)
// 			s.Emit(event.GameError, websocketError(event.GameDahai, err.Error()))
// 		}
// 		if err = gameUsecase.Dahai(body.RoomName, dahai); err != nil {
// 			fmt.Println(err)
// 			s.Emit(event.GameError, websocketError(event.GameDahai, err.Error()))
// 		}
// 	}
// }
