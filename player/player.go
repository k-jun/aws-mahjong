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
	name  string
	deck  *deck.Deck
	tsumo *tile.Tile
	hand  *hand.Hand
	kawa  *kawa.Kawa
	naki  *naki.Naki
}

func NewPlayer(playername string, deck *deck.Deck) *Player {
	return &Player{
		name: playername,
		deck: deck,
		hand: hand.NewHand(),
		kawa: kawa.NewKawa(),
		naki: naki.NewNaki(),
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

func (p *Player) DahaiDone(deadTile *tile.Tile) error {
	// TODO when reach isSide=true
	err := p.kawa.Add(deadTile, false)
	return err
}
