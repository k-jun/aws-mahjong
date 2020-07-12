package naki

import "aws-mahjong/tile"

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

type Naki struct {
	setMap map[int][]*NakiTile
}

func NewNaki() *Naki {
	return &Naki{
		setMap: map[int][]*NakiTile{},
	}
}

func (n *Naki) AddSet(tiles []*tile.Tile, cha NakiFrom) error {
	set := []*NakiTile{}

	for _, t := range tiles {
		set = append(set, &NakiTile{
			tile:   t,
			isOpen: true,
			isSide: false,
		})
	}

	// ankan
	if len(set) == 4 && cha == "jicha" {
		set[1].isOpen = false
		set[2].isOpen = false
	}

	// minkan
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

	n.setMap[n.nextIndex()] = set
	return nil

}

func (n *Naki) nextIndex() int {
	return len(n.setMap)
}
