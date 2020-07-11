package player

import (
	"aws-mahjong/deck"
	"aws-mahjong/hand"
	"aws-mahjong/tile"
	"errors"
)

var (
	TsumoAlreadyExistErr = errors.New("tsumo already exist")
)

type NakiTile struct {
	tile   tile.Tile
	isOpen bool
}

type KawaTile struct {
	tile   tile.Tile
	isSide bool
}

type Player struct {
	name  string
	deck  *deck.Deck
	tsumo *tile.Tile
	hand  *hand.Hand
	kawa  []*KawaTile
	naki  []*NakiTile
}

func NewPlayer(playername string, deck *deck.Deck) *Player {
	return &Player{
		name: playername,
		deck: deck,
		hand: hand.NewHand(),
		kawa: []*KawaTile{},
		naki: []*NakiTile{},
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
			return outTile, err
		}
	}
	p.tsumo = nil

	return outTile, nil
}
