package player

import (
	"aws-mahjong/deck"
	"aws-mahjong/hand"
	"aws-mahjong/kawa"
	"aws-mahjong/naki"
	"aws-mahjong/tile"
	"errors"
)

type Player interface {
	Hand() hand.Hand
	Tsumo() error
	Dahai(outTile *tile.Tile) (*tile.Tile, error)
	DahaiDone(deadTile *tile.Tile, isSide bool) error
	Naki(inTile *tile.Tile, fromHandTiles []*tile.Tile, cha naki.NakiFrom) error
	CanNakiActions(inTile *tile.Tile) []*NakiAction
}

var (
	TsumoAlreadyExistErr = errors.New("tsumo already exist")
)

type NakiAction = string

var (
	Pon  NakiAction = "pon"
	Kan  NakiAction = "kan"
	Chii NakiAction = "chii"
)

type PlayerImpl struct {
	// user info
	id   string
	name string

	// game info
	deck   deck.Deck
	bakaze *tile.Tile
	zikaze *tile.Tile

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
	deck deck.Deck,
	hand hand.Hand,
	kawa kawa.Kawa,
	naki naki.Naki,
) Player {
	return &PlayerImpl{
		id:   id,
		name: playername,
		deck: deck,
		hand: hand,
		kawa: kawa,
		naki: naki,
	}
}

func (p *PlayerImpl) Hand() hand.Hand {
	return p.hand
}

func (p *PlayerImpl) Tsumo() error {
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

func (p *PlayerImpl) Dahai(outTile *tile.Tile) (*tile.Tile, error) {
	if outTile != p.tsumo {
		_, err := p.hand.Replace(p.tsumo, outTile)
		if err != nil {
			return nil, err
		}
	}
	p.tsumo = nil

	return outTile, nil
}

func (p *PlayerImpl) DahaiDone(deadTile *tile.Tile, isSide bool) error {
	return p.kawa.Add(deadTile, isSide)
}

func (p *PlayerImpl) Naki(inTile *tile.Tile, fromHandTiles []*tile.Tile, cha naki.NakiFrom) error {
	set, err := p.hand.Removes(fromHandTiles)
	if err != nil {
		return err
	}
	set = append(set, inTile)
	tile.SortTiles(set)

	err = p.naki.AddSet(set, cha)
	return err
}

func (p *PlayerImpl) CanNakiActions(inTile *tile.Tile) []*NakiAction {
	actions := []*NakiAction{}
	if p.canChii(inTile) {
		actions = append(actions, &Chii)
	}
	if p.canPon(inTile) {
		actions = append(actions, &Pon)
	}
	if p.canKan(inTile) {
		actions = append(actions, &Kan)
	}

	return actions
}

func (p *PlayerImpl) canPon(inTile *tile.Tile) bool {
	pairs := p.hand.FindPonPair(inTile)
	return len(pairs) != 0
}

func (p *PlayerImpl) canChii(inTile *tile.Tile) bool {
	pairs := p.hand.FindChiiPair(inTile)
	return len(pairs) != 0
}

func (p *PlayerImpl) canKan(inTile *tile.Tile) bool {
	pairs := p.hand.FindKanPair(inTile)
	return len(pairs) != 0
}
