package usecase

import (
	"aws-mahjong/naki"
	"aws-mahjong/repository"
	"aws-mahjong/tile"
)

type GameUsecase interface {
	Dahai(roomName string, dahai *tile.Tile) error
	Naki(roomName string, action naki.NakiAction, tiles []*tile.Tile) error
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

func (u *GameUsecaseImpl) Dahai(roomName string, dahai *tile.Tile) error {
	roomGame, err := u.gameRepo.Find(roomName)
	if err != nil {
		return err
	}
	if err = roomGame.Board().TurnPlayerDahai(dahai); err != nil {
		return err
	}
	newGameStatus(u.roomRepo, roomName, roomGame)
	return nil
}

func (u *GameUsecaseImpl) Naki(roomName string, action naki.NakiAction, tiles []*tile.Tile) error {
	return nil
}
