package kawa

import "aws-mahjong/tile"

type KawaTile struct {
	tile   *tile.Tile
	isSide bool
}

type Kawa struct {
	tiles []*KawaTile
}

func NewKawa() *Kawa {
	return &Kawa{
		tiles: []*KawaTile{},
	}
}

func (k *Kawa) Tiles() []*KawaTile {
	return k.tiles
}

func (k *Kawa) Add(inTile *tile.Tile, isSide bool) error {
	k.tiles = append(k.tiles, &KawaTile{
		tile:   inTile,
		isSide: isSide,
	})
	return nil
}
