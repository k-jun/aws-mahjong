package kawa

import (
	"aws-mahjong/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKawa(t *testing.T) {
	kawa := NewKawa()
	assert.Equal(t, []*KawaTile{}, kawa.Tiles())
}

func TestAdd(t *testing.T) {
	cases := []struct {
		Description  string
		CurrentTiles []*KawaTile
		InTile       *tile.Tile
		InIsSide     bool
		OutError     error
		OutTiles     []*KawaTile
	}{
		{
			Description:  "valid case",
			CurrentTiles: []*KawaTile{},
			InIsSide:     false,
			InTile:       &tile.Chun,
			OutError:     nil,
			OutTiles:     []*KawaTile{&KawaTile{tile: &tile.Chun, isSide: false}},
		},
		{
			Description: "valid case",
			CurrentTiles: []*KawaTile{
				&KawaTile{tile: &tile.Chun, isSide: false},
				&KawaTile{tile: &tile.Souzu1, isSide: false},
				&KawaTile{tile: &tile.Souzu5, isSide: false},
			},
			InIsSide: true,
			InTile:   &tile.Haku,
			OutError: nil,
			OutTiles: []*KawaTile{
				&KawaTile{tile: &tile.Chun, isSide: false},
				&KawaTile{tile: &tile.Souzu1, isSide: false},
				&KawaTile{tile: &tile.Souzu5, isSide: false},
				&KawaTile{tile: &tile.Haku, isSide: true},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			kawa := KawaImpl{tiles: c.CurrentTiles}
			err := kawa.Add(c.InTile, c.InIsSide)
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutTiles, kawa.Tiles())
		})

	}

}
