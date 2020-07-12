package kawa

import "aws-mahjong/tile"

type KawaTile struct {
	tile   *tile.Tile
	isSide bool
}

func NewKawaTile(inTile *tile.Tile, isSide bool) *KawaTile {
	return &KawaTile{
		tile:   inTile,
		isSide: isSide,
	}
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
	k.tiles = append(k.tiles, NewKawaTile(inTile, isSide))
	return nil
}
