package naki

import (
	"aws-mahjong/tile"
	"errors"
)

var (
	SetNotFoundErr = errors.New("specified tile's set does not found")
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

func NewNakiTile(inTile *tile.Tile, isOpen bool, isSide bool) *NakiTile {
	return &NakiTile{
		tile:   inTile,
		isOpen: isOpen,
		isSide: isSide,
	}
}

type Naki struct {
	sets [][]*NakiTile
}

func NewNaki() *Naki {
	return &Naki{
		sets: [][]*NakiTile{},
	}
}

func (n *Naki) AddSet(tiles []*tile.Tile, cha NakiFrom) error {
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

func (n *Naki) AddTileToSet(inTile *tile.Tile) error {
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
