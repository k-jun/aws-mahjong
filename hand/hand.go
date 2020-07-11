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

func (h *Hand) FindPonPair(inTile *tile.Tile) [][2]*tile.Tile {
	pairs := [][2]*tile.Tile{}

	hitTiles := []*tile.Tile{}

	for _, tile := range h.tiles {
		if tile.IsSame(inTile) {
			hitTiles = append(hitTiles, tile)
		}
	}

	// sort hit tiles
	sort.Slice(hitTiles, func(i int, j int) bool { return hitTiles[i].Name() < hitTiles[j].Name() })

	for i := 0; i < len(hitTiles); i++ {
		for j := i + 1; j < len(hitTiles); j++ {
			// check duplicate
			shouldAppend := true
			for _, pair := range pairs {
				if pair[0].Name() == hitTiles[i].Name() && pair[1].Name() == hitTiles[j].Name() {
					shouldAppend = false
				}
				if pair[1].Name() == hitTiles[i].Name() && pair[0].Name() == hitTiles[j].Name() {
					shouldAppend = false
				}
			}

			if shouldAppend {
				pairs = append(pairs, [2]*tile.Tile{hitTiles[i], hitTiles[j]})
			}
		}
	}

	return pairs
}

func (h *Hand) FindChiiPair(inTile *tile.Tile) [][2]*tile.Tile {
	pairs := [][2]*tile.Tile{}
	return pairs
}

func (h *Hand) FindKanPair(inTile *tile.Tile) [][3]*tile.Tile {
	pairs := [][3]*tile.Tile{}

	hitTiles := []*tile.Tile{}

	for _, tile := range h.tiles {
		if tile.IsSame(inTile) {
			hitTiles = append(hitTiles, tile)
		}
	}

	// sort hit tiles
	sort.Slice(hitTiles, func(i int, j int) bool { return hitTiles[i].Name() < hitTiles[j].Name() })

	if len(hitTiles) == 3 {
		pairs = append(pairs, [3]*tile.Tile{hitTiles[0], hitTiles[1], hitTiles[2]})
	}
	return pairs
}
