package kawa

import "aws-mahjong/tile"

var _ Kawa = &KawaMock{}

type KawaMock struct {
	ExpectedTiles  []*KawaTile
	ExpectedError  error
	ExpectedStatus []*KawaStatus
}

func (k *KawaMock) Tiles() []*KawaTile {
	return k.ExpectedTiles
}

func (k *KawaMock) Add(inTile *tile.Tile, isSide bool) error {
	return k.ExpectedError
}

func (k *KawaMock) Status() []*KawaStatus {
	return k.ExpectedStatus
}
