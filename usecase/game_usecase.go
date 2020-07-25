package usecase

import (
	"aws-mahjong/board"
	"aws-mahjong/repository"
	"aws-mahjong/tile"
)

type GameUsecase interface {
	Dahai(userId string, roomName string, hai *tile.Tile) (*board.BoardStatus, error)
}

type GameUsecaseImpl struct {
	roomRepo repository.RoomRepository
}

func NewGameUsecaseImpl(roomRepo repository.RoomRepository) GameUsecase {
	return &GameUsecaseImpl{roomRepo: roomRepo}
}

func (u *GameUsecaseImpl) Dahai(userId string, roomName string, hai *tile.Tile) (*board.BoardStatus, error) {
	g, err := u.roomRepo.Find(roomName)
	if err != nil {
		return nil, err
	}
	status, err := g.Dahai(userId, hai)
	return status, err
}
