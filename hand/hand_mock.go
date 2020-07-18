package hand

import "aws-mahjong/tile"

var _ Hand = &HandMock{}

type HandMock struct {
	ExpectedTile   *tile.Tile
	ExpectedTiles  []*tile.Tile
	ExpectedError  error
	ExpectedPair2  [][2]*tile.Tile
	ExpectedPair3  [][3]*tile.Tile
	ExpectedStatus []string
}

func (h *HandMock) Tiles() []*tile.Tile {
	return h.ExpectedTiles
}

func (h *HandMock) Add(inTile *tile.Tile) error {
	return h.ExpectedError
}

func (h *HandMock) Adds(inTiles []*tile.Tile) error {
	return h.ExpectedError
}

func (h *HandMock) Remove(outTile *tile.Tile) (*tile.Tile, error) {
	return h.ExpectedTile, h.ExpectedError
}

func (h *HandMock) Removes(outTiles []*tile.Tile) ([]*tile.Tile, error) {
	return h.ExpectedTiles, h.ExpectedError
}

func (h *HandMock) Replace(inTile *tile.Tile, outtile *tile.Tile) (*tile.Tile, error) {
	return h.ExpectedTile, h.ExpectedError
}

func (h *HandMock) FindChiiPair(inTile *tile.Tile) [][2]*tile.Tile {
	return h.ExpectedPair2
}

func (h *HandMock) FindPonPair(inTile *tile.Tile) [][2]*tile.Tile {
	return h.ExpectedPair2
}

func (h *HandMock) FindKanPair(inTile *tile.Tile) [][3]*tile.Tile {
	return h.ExpectedPair3
}

func (h *HandMock) Status() []string {
	return h.ExpectedStatus
}
