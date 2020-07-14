package hand

import (
	"aws-mahjong/tile"
	"errors"
	"strconv"
)

var (
	HandCount = 13
)

type Hand interface {
	Tiles() []*tile.Tile
	Add(inTile *tile.Tile) error
	Adds(inTiles []*tile.Tile) error
	Remove(outTile *tile.Tile) (*tile.Tile, error)
	Removes(outTiles []*tile.Tile) ([]*tile.Tile, error)
	Replace(inTile *tile.Tile, outTile *tile.Tile) (*tile.Tile, error)
	FindChiiPair(inTile *tile.Tile) [][2]*tile.Tile
	FindPonPair(inTile *tile.Tile) [][2]*tile.Tile
	FindKanPair(inTile *tile.Tile) [][3]*tile.Tile
}

var (
	TileNotFoundErr = errors.New("specified tile does not exist")
	TileCountErr    = errors.New("invalid hand count")
)

type HandImpl struct {
	tiles []*tile.Tile
}

func NewHand() Hand {
	return &HandImpl{
		tiles: []*tile.Tile{},
	}
}

func (h *HandImpl) Tiles() []*tile.Tile {
	tiles := h.tiles
	tile.SortTiles(tiles)
	return tiles
}

func (h *HandImpl) Add(inTile *tile.Tile) error {
	if len(h.tiles) > HandCount-1 {
		return TileCountErr
	}
	h.tiles = append(h.tiles, inTile)
	return nil
}

func (h *HandImpl) Adds(inTiles []*tile.Tile) error {
	for _, tile := range inTiles {
		err := h.Add(tile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *HandImpl) Remove(outTile *tile.Tile) (*tile.Tile, error) {
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

func (h *HandImpl) Removes(outTiles []*tile.Tile) ([]*tile.Tile, error) {
	tiles := []*tile.Tile{}
	for _, tile := range outTiles {
		tile, err := h.Remove(tile)
		if err != nil {
			return tiles, err
		}
		tiles = append(tiles, tile)
	}

	return tiles, nil
}

func (h *HandImpl) Replace(inTile *tile.Tile, outTile *tile.Tile) (*tile.Tile, error) {
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

func (h *HandImpl) FindPonPair(inTile *tile.Tile) [][2]*tile.Tile {
	pairs := [][2]*tile.Tile{}

	hitTiles := []*tile.Tile{}

	for _, tile := range h.tiles {
		if tile.IsSame(inTile) {
			hitTiles = append(hitTiles, tile)
		}
	}

	tile.SortTiles(hitTiles)

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

func (h *HandImpl) FindChiiPair(inTile *tile.Tile) [][2]*tile.Tile {
	pairs := [][2]*tile.Tile{}

	if inTile.IsZihai() {
		return pairs
	}

	num := inTile.Number()
	smpKind := inTile.KindSMP()

	lambda := func(xkind *tile.TileKind, ykind *tile.TileKind) {
		for _, i := range h.findByKinds([]*tile.TileKind{xkind, smpKind}) {
			for _, j := range h.findByKinds([]*tile.TileKind{ykind, smpKind}) {
				pairs = append(pairs, [2]*tile.Tile{i, j})
			}
		}
	}
	// inTile is on left
	if num <= 7 {
		centerKind := tile.TileKindFromString(strconv.Itoa(num + 1))
		rightKind := tile.TileKindFromString(strconv.Itoa(num + 2))
		lambda(centerKind, rightKind)
	}

	// inTile is on center
	if num <= 8 && num >= 2 {
		leftKind := tile.TileKindFromString(strconv.Itoa(num - 1))
		rightKind := tile.TileKindFromString(strconv.Itoa(num + 1))
		lambda(leftKind, rightKind)
	}

	// inTile is on right
	if num <= 3 {
		leftKind := tile.TileKindFromString(strconv.Itoa(num - 2))
		centerKind := tile.TileKindFromString(strconv.Itoa(num - 1))
		lambda(leftKind, centerKind)
	}

	return pairs
}

func (h *HandImpl) FindKanPair(inTile *tile.Tile) [][3]*tile.Tile {
	pairs := [][3]*tile.Tile{}

	hitTiles := []*tile.Tile{}

	for _, tile := range h.tiles {
		if tile.IsSame(inTile) {
			hitTiles = append(hitTiles, tile)
		}
	}

	tile.SortTiles(hitTiles)

	if len(hitTiles) == 3 {
		pairs = append(pairs, [3]*tile.Tile{hitTiles[0], hitTiles[1], hitTiles[2]})
	}
	return pairs
}

func (h *HandImpl) findByKinds(kinds []*tile.TileKind) []*tile.Tile {
	tiles := []*tile.Tile{}

	for _, tile := range h.tiles {
		hasAllKind := true
		for _, kind := range kinds {
			hasKind := false
			for _, tKind := range tile.Kinds() {
				if *tKind == *kind {
					hasKind = true
				}
			}
			if !hasKind {
				hasAllKind = false
			}
		}
		if hasAllKind {
			tiles = append(tiles, tile)
		}
	}

	return tiles
}
