package repository

import "aws-mahjong/game"

var _ RoomRepository = &RoomRepositoryMock{}

type RoomRepositoryMock struct {
	ExpectedGame  *game.Game
	ExpectedError error
	ExpectedRooms []*RoomStatus
}

func (r *RoomRepositoryMock) Add(roomName string, inGame game.Game) error {
	return r.ExpectedError
}

func (r *RoomRepositoryMock) Remove(roomName string) error {
	return r.ExpectedError
}

func (r *RoomRepositoryMock) Find(roomName string) (*game.Game, error) {
	return r.ExpectedGame, r.ExpectedError
}

func (r *RoomRepositoryMock) Rooms() []*RoomStatus {
	return r.ExpectedRooms
}
