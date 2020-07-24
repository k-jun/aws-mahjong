package view

import "aws-mahjong/usecase"

type RoomsResponse struct {
	RoomName        string `json:"room_name"`
	RoomCapacity    int    `json:"room_capacity"`
	RoomMemberCount int    `json:"room_member_count"`
}

func NewRoomResponse(rooms []*usecase.RoomStatus) []*RoomsResponse {
	response := []*RoomsResponse{}
	for _, roomInfo := range rooms {
		response = append(response, &RoomsResponse{
			RoomName:        roomInfo.Name,
			RoomMemberCount: roomInfo.Len,
			RoomCapacity:    roomInfo.Capacity,
		})
	}

	return response

}
