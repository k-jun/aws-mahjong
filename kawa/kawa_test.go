package kawa

import (
	"aws-mahjong/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKawa(t *testing.T) {
	kawa := NewKawa()
	assert.Equal(t, []*KawaTile{}, kawa.tiles)
}

func TestAdd(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*KawaTile
		InTile        *tile.Tile
		InIsSide      bool
		ExpectedError error
	}{
		{
			Description:   "valid case",
			CurrentTiles:  []*KawaTile{},
			InIsSide:      false,
			InTile:        &tile.Chun,
			ExpectedError: nil,
		},
		{
			Description:   "valid case",
			CurrentTiles:  []*KawaTile{},
			InIsSide:      true,
			InTile:        &tile.Haku,
			ExpectedError: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			kawa := NewKawa()
			kawa.tiles = c.CurrentTiles
			err := kawa.Add(c.InTile, c.InIsSide)
			assert.Equal(t, c.ExpectedError, err)
		})

	}

}
