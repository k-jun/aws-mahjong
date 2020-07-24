package view

import "aws-mahjong/usecase"

type RoomResponse struct {
	RoomName     string `json:"room_name"`
	RoomCapacity int    `json:"room_capacity"`
	RoomLen      int    `json:"room_len"`
}

func NewRoomResponse(rooms []*usecase.RoomStatus) []*RoomResponse {
	response := []*RoomResponse{}
	for _, roomInfo := range rooms {
		response = append(response, &RoomResponse{
			RoomName:     roomInfo.Name,
			RoomLen:      roomInfo.Len,
			RoomCapacity: roomInfo.Capacity,
		})
	}

	return response
}

type CreateRoomRequest struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	RoomName     string `json:"room_name"`
	RoomCapacity int    `json:"room_capacity"`
}

func NewStatusRoomResponse(status *usecase.RoomStatus) *RoomResponse {
	return &RoomResponse{
		RoomName:     status.Name,
		RoomCapacity: status.Capacity,
		RoomLen:      status.Len,
	}
}

type JoinRoomRequest struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type LeaveRoomRequest struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}
