package usecase

// func newGameStatus(roomRepo *repository.RoomRepository, roomName string, roomGame game.Game) {
//
// 	roomRepo.ForEach(roomName, func(s socketio.Conn) {
// 		status := roomGame.Board().Status(s.ID())
// 		bytes, err := json.Marshal(status)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		s.Emit(event.NewGameStatus, string(bytes))
// 	})
// }
