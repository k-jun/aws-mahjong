package naki

import (
	"aws-mahjong/tile"
	"errors"
)

type NakiAction = string

var (
	Pon    NakiAction = "pon"
	Kan    NakiAction = "kan"
	Chii   NakiAction = "chii"
	Cancel NakiAction = "cancel"
)

type NakiFrom string

var (
	Kamicha NakiFrom = "kamicha"
	Simocha NakiFrom = "simacha"
	Toimen  NakiFrom = "toimen"
	Jicha   NakiFrom = "jicha"
)

type NakiTile struct {
	tile   *tile.Tile
	isOpen bool
	isSide bool
}

type Naki interface {
	AddSet(tiles []*tile.Tile, cha NakiFrom) error
	AddTileToSet(inTile *tile.Tile) error
	Sets() [][]*NakiTile
}

var (
	SetNotFoundErr = errors.New("specified tile's set does not found")
)

func NewNakiTile(inTile *tile.Tile, isOpen bool, isSide bool) *NakiTile {
	return &NakiTile{
		tile:   inTile,
		isOpen: isOpen,
		isSide: isSide,
	}
}

type NakiImpl struct {
	sets [][]*NakiTile
}

func NewNaki() Naki {
	return &NakiImpl{
		sets: [][]*NakiTile{},
	}
}

func (n *NakiImpl) Sets() [][]*NakiTile {
	return n.sets
}

func (n *NakiImpl) AddSet(tiles []*tile.Tile, cha NakiFrom) error {
	set := []*NakiTile{}

	for _, t := range tiles {
		set = append(set, NewNakiTile(t, true, false))
	}

	// ankan
	if len(set) == 4 && cha == "jicha" {
		set[1].isOpen = false
		set[2].isOpen = false
	}

	// min pon, kan, chii
	switch cha {
	case Kamicha:
		set[0].isSide = true
	case Toimen:
		set[1].isSide = true
	case Simocha:
		set[2].isSide = true
	default:
		break
	}

	n.sets = append(n.sets, set)
	return nil
}

func (n *NakiImpl) AddTileToSet(inTile *tile.Tile) error {
	for idx, set := range n.sets {
		if inTile.IsSame(set[0].tile) {

			// insert after isSide=true tile if exists
			insertIdx := len(set)
			for idx, t := range set {
				if t.isSide == true {
					insertIdx = idx + 1
					break
				}
			}
			set = append(set, NewNakiTile(inTile, true, true))
			set[insertIdx], set[len(set)-1] = set[len(set)-1], set[insertIdx]
			n.sets[idx] = set
			return nil
		}
	}

	return SetNotFoundErr
}
