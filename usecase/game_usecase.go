package usecase

import (
	"aws-mahjong/naki"
	"aws-mahjong/repository"
)

type GameUsecase interface {
	Dahai(roomName string, dahai string) error
	Naki(roomName string, action naki.NakiAction, tiles []string) error
}

type GameUsecaseImpl struct {
	gameRepo repository.GameRepository
	roomRepo *repository.RoomRepository
}

func NewGameUsecase(roomRepo *repository.RoomRepository, gameRepo repository.GameRepository) GameUsecase {
	return &GameUsecaseImpl{
		roomRepo: roomRepo,
		gameRepo: gameRepo,
	}
}

func (u *GameUsecaseImpl) Dahai(roomName string, dahai string) error {
	return nil
}

func (u *GameUsecaseImpl) Naki(roomName string, action naki.NakiAction, tiles []string) error {
	return nil
}
