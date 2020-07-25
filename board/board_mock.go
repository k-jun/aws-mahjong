package board

import "aws-mahjong/tile"

var _ Board = &BoardMock{}

type BoardMock struct {
	ExpectedError       error
	ExpectedBoardStatus *BoardStatus
	ExpectedBool        bool
}

func (b *BoardMock) TurnPlayerTsumo() error {
	return b.ExpectedError

}

func (b *BoardMock) TurnPlayerDahai(playerID string, outTile *tile.Tile) error {
	return b.ExpectedError
}

func (b *BoardMock) NextTurn() {

}

func (b *BoardMock) ChangeTurn(playerIdx int) error {
	return b.ExpectedError

}

func (b *BoardMock) Start() error {
	return b.ExpectedError
}

func (b *BoardMock) Status(playerId string) *BoardStatus {
	return b.ExpectedBoardStatus
}

func (b *BoardMock) IsTurnPlayer(playerId string) bool {
	return b.ExpectedBool
}
