package kawa

import "aws-mahjong/tile"

type Kawa interface {
	Tiles() []*KawaTile
	Add(inTile *tile.Tile, isSide bool) error
}

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

type KawaImpl struct {
	tiles []*KawaTile
}

func NewKawa() Kawa {
	return &KawaImpl{
		tiles: []*KawaTile{},
	}
}

func (k *KawaImpl) Tiles() []*KawaTile {
	return k.tiles
}

func (k *KawaImpl) Add(inTile *tile.Tile, isSide bool) error {
	k.tiles = append(k.tiles, NewKawaTile(inTile, isSide))
	return nil
}
