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
		OutSets     [][]*NakiTile
	}{
		{
			Description: "valid pon case",
			InTiles:     []*tile.Tile{&tile.Chun, &tile.Chun, &tile.Chun},
			InCha:       Jicha,
			OutError:    nil,
			OutSets: [][]*NakiTile{
				[]*NakiTile{
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
			OutSets: [][]*NakiTile{
				[]*NakiTile{
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
			OutSets: [][]*NakiTile{
				[]*NakiTile{
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
			OutSets: [][]*NakiTile{
				[]*NakiTile{
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
			assert.Equal(t, c.OutSets, naki.Sets())

		})
	}
}

func TestAddTileToSet(t *testing.T) {
	cases := []struct {
		Description string
		CurrentSets [][]*NakiTile
		InTile      *tile.Tile
		OutError    error
		OutSets     [][]*NakiTile
	}{
		{
			Description: "valid case",
			CurrentSets: [][]*NakiTile{
				{
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
				},
			},
			InTile:   &tile.Chun,
			OutError: nil,
			OutSets: [][]*NakiTile{
				{
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
				},
			},
		},
		{
			Description: "invalid case",
			CurrentSets: [][]*NakiTile{
				{
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
				},
				{
					&NakiTile{tile: &tile.Hatu, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Hatu, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Hatu, isOpen: true, isSide: false},
				},
			},
			InTile:   &tile.Haku,
			OutError: SetNotFoundErr,
			OutSets:  [][]*NakiTile{},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			naki := &NakiImpl{
				sets: c.CurrentSets,
			}
			err := naki.AddTileToSet(c.InTile)
			if err != nil && c.OutError == err {
				return
			}
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutSets, naki.Sets())
		})
	}
}

func TestStatus(t *testing.T) {
	cases := []struct {
		Description string
		CurrentSets [][]*NakiTile
		OutStatus   [][]*NakiStatus
	}{
		{
			Description: "valid case",
			CurrentSets: [][]*NakiTile{
				{
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Chun, isOpen: true, isSide: false},
				},
				{
					&NakiTile{tile: &tile.Hatu, isOpen: true, isSide: false},
					&NakiTile{tile: &tile.Hatu, isOpen: true, isSide: true},
					&NakiTile{tile: &tile.Hatu, isOpen: true, isSide: false},
				},
			},
			OutStatus: [][]*NakiStatus{
				{
					&NakiStatus{Name: tile.Chun.Name(), IsOpen: true, IsSide: false},
					&NakiStatus{Name: tile.Chun.Name(), IsOpen: true, IsSide: true},
					&NakiStatus{Name: tile.Chun.Name(), IsOpen: true, IsSide: false},
				},
				{
					&NakiStatus{Name: tile.Hatu.Name(), IsOpen: true, IsSide: false},
					&NakiStatus{Name: tile.Hatu.Name(), IsOpen: true, IsSide: true},
					&NakiStatus{Name: tile.Hatu.Name(), IsOpen: true, IsSide: false},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			naki := NakiImpl{sets: c.CurrentSets}
			assert.Equal(t, c.OutStatus, naki.Status())
		})
	}
}
