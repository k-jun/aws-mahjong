package player

import (
	"aws-mahjong/deck"
	"aws-mahjong/hand"
	"aws-mahjong/kawa"
	"aws-mahjong/naki"
	"aws-mahjong/tile"
	"errors"
)

var (
	TsumoAlreadyExistErr = errors.New("tsumo already exist")
)

type Player struct {
	// user info
	id   string
	name string

	// game info
	deck   deck.Deck
	bakaze *tile.Tile
	zikaze *tile.Tile
	isOya  bool

	// player info
	tsumo *tile.Tile
	hand  hand.Hand
	kawa  kawa.Kawa
	naki  naki.Naki
}

func NewPlayer(
	id string,
	playername string,
	bakaze *tile.Tile,
	zikaze *tile.Tile,
	isOya bool,
	deck deck.Deck,
	hand hand.Hand,
	kawa kawa.Kawa,
	naki naki.Naki,
) *Player {
	return &Player{
		id:   id,
		name: playername,
		deck: deck,
		hand: hand,
		kawa: kawa,
		naki: naki,
	}
}

func (p *Player) Tsumo() error {
	if p.tsumo != nil {
		return TsumoAlreadyExistErr
	}

	tsumo, err := p.deck.Draw()
	if err != nil {
		return err
	}
	p.tsumo = tsumo
	return nil
}

func (p *Player) Dahai(outTile *tile.Tile) (*tile.Tile, error) {
	if outTile != p.tsumo {
		_, err := p.hand.Replace(p.tsumo, outTile)
		if err != nil {
			return nil, err
		}
	}
	p.tsumo = nil

	return outTile, nil
}

func (p *Player) DahaiDone(deadTile *tile.Tile, isSide bool) error {
	return p.kawa.Add(deadTile, isSide)
}

func (p *Player) Naki(inTile *tile.Tile, fromHandTiles []*tile.Tile, cha naki.NakiFrom) error {
	set, err := p.hand.Removes(fromHandTiles)
	if err != nil {
		return err
	}
	set = append(set, inTile)
	tile.SortTiles(set)

	err = p.naki.AddSet(set, cha)
	return err
}

func (p *Player) CanNaki(inTile *tile.Tile) bool {
	return p.canPon(inTile) || p.canChii(inTile) || p.canKan(inTile)
}

func (p *Player) canPon(inTile *tile.Tile) bool {
	pairs := p.hand.FindPonPair(inTile)
	return len(pairs) != 0
}

func (p *Player) canChii(inTile *tile.Tile) bool {
	pairs := p.hand.FindChiiPair(inTile)
	return len(pairs) != 0
}

func (p *Player) canKan(inTile *tile.Tile) bool {
	pairs := p.hand.FindKanPair(inTile)
	return len(pairs) != 0
}
