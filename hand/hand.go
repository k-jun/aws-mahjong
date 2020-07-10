package hand

import (
	"aws-mahjong/tile"
	"errors"
	"sort"
)

var (
	handCount = 13
)

var (
	TileNotFoundErr = errors.New("specified tile does not exist")
	TileCountErr    = errors.New("invalid hand count")
)

type Hand struct {
	tiles []*tile.Tile
}

func NewHand() *Hand {
	return &Hand{
		tiles: []*tile.Tile{},
	}
}

func (h *Hand) Tiles() []*tile.Tile {
	tiles := h.tiles
	sort.Slice(tiles, func(i int, j int) bool { return tiles[i].Name() < tiles[j].Name() })
	return tiles
}

func (h *Hand) Add(inTile *tile.Tile) error {
	if len(h.tiles) > handCount-1 {
		return TileCountErr

	}
	h.tiles = append(h.tiles, inTile)
	return nil

}

func (h *Hand) Remove(outTile *tile.Tile) (*tile.Tile, error) {
	for idx, tile := range h.tiles {
		if tile.Name() == outTile.Name() {
			tile := h.tiles[idx]
			h.tiles[idx] = h.tiles[0]
			h.tiles = h.tiles[1:]
			return tile, nil
		}
	}
	return nil, TileNotFoundErr

}

func (h *Hand) Replace(inTile *tile.Tile, outTile *tile.Tile) (*tile.Tile, error) {
	for idx, tile := range h.tiles {
		if tile.Name() == outTile.Name() {
			tile = h.tiles[idx]
			h.tiles[idx] = h.tiles[0]
			h.tiles[0] = inTile
			return tile, nil
		}
	}

	return nil, TileNotFoundErr
}
