package player

import (
	"aws-mahjong/hand"
	"aws-mahjong/naki"
	"aws-mahjong/tile"
)

var _ Player = &PlayerMock{}

type PlayerMock struct {
	ExpectedHand        hand.Hand
	ExpectedError       error
	ExpectedTile        *tile.Tile
	ExpectedNakiActions []*naki.NakiAction
}

func (p *PlayerMock) Hand() hand.Hand {
	return p.ExpectedHand
}

func (p *PlayerMock) Tsumo() error {
	return p.ExpectedError
}

func (p *PlayerMock) Dahai(outTile *tile.Tile) (*tile.Tile, error) {
	return p.ExpectedTile, p.ExpectedError
}

func (p *PlayerMock) DahaiDone(deadTile *tile.Tile, isSide bool) error {
	return p.ExpectedError
}

func (p *PlayerMock) Naki(inTile *tile.Tile, fromHandTiles []*tile.Tile, cha naki.NakiFrom) error {
	return p.ExpectedError
}

func (p *PlayerMock) CanNakiActions(inTile *tile.Tile) []*naki.NakiAction {
	return p.ExpectedNakiActions

}
