package handler

type NewRoomStatus struct {
	RoomName        string `json:"room_name"`
	RoomMemberCount int    `json:"room_member_count"`
	RoomCapacity    int    `json:"room_capacity"`
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
