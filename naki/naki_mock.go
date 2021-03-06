package naki

import "aws-mahjong/tile"

var _ Naki = &NakiMock{}

type NakiMock struct {
	ExpectedSets   [][]*NakiTile
	ExpectedError  error
	ExpectedStatus [][]*NakiStatus
}

func (n *NakiMock) AddSet(tiles []*tile.Tile, cha NakiFrom) error {
	return n.ExpectedError
}

func (n *NakiMock) AddTileToSet(tile *tile.Tile) error {
	return n.ExpectedError
}

func (n *NakiMock) Sets() [][]*NakiTile {
	return n.ExpectedSets
}

func (n *NakiMock) Status() [][]*NakiStatus {
	return n.ExpectedStatus
}
