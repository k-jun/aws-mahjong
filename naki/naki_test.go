package naki

import (
	"aws-mahjong/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSet(t *testing.T) {
	cases := []struct {
		Description string
		InTiles     []*tile.Tile
		InCha       NakiFrom
		OutError    error
		OutMapSet   map[int][]*NakiTile
	}{
		{
			Description: "valid pon case",
			InTiles:     []*tile.Tile{&tile.Chun, &tile.Chun, &tile.Chun},
			InCha:       Jicha,
			OutError:    nil,
			OutMapSet: map[int][]*NakiTile{
				0: []*NakiTile{
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
				},
			},
		},
		{
			Description: "valid pon case",
			InTiles:     []*tile.Tile{&tile.Chun, &tile.Chun, &tile.Chun},
			InCha:       Toimen,
			OutError:    nil,
			OutMapSet: map[int][]*NakiTile{
				0: []*NakiTile{
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
				},
			},
		},
		{
			Description: "valid kan case",
			InTiles:     []*tile.Tile{&tile.Chun, &tile.Chun, &tile.Chun, &tile.Chun},
			InCha:       Jicha,
			OutError:    nil,
			OutMapSet: map[int][]*NakiTile{
				0: []*NakiTile{
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: false, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: false, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
				},
			},
		},
		{
			Description: "valid kan case",
			InTiles:     []*tile.Tile{&tile.Chun, &tile.Chun, &tile.Chun, &tile.Chun},
			InCha:       Kamicha,
			OutError:    nil,
			OutMapSet: map[int][]*NakiTile{
				0: []*NakiTile{
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			naki := NewNaki()
			err := naki.AddSet(c.InTiles, c.InCha)
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutMapSet, naki.setMap)

		})
	}
}
