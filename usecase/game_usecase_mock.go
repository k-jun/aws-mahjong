package usecase

import (
	"aws-mahjong/board"
	"aws-mahjong/tile"
)

var _ GameUsecase = &GameUsecaseMock{}

type GameUsecaseMock struct {
	ExpectedError       error
	ExpectedBoardStatus *board.BoardStatus
}

func (g *GameUsecaseMock) Dahai(userId string, roomName string, hai *tile.Tile) (*board.BoardStatus, error) {
	return g.ExpectedBoardStatus, g.ExpectedError

}
