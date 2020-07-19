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
	CanNakiActions(inTile *tile.Tile) []*naki.NakiAction
	Status(inTile *tile.Tile) *PlayerStatus
}

var (
	TsumoAlreadyExistErr = errors.New("tsumo already exist")
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
		id:     id,
		name:   playername,
		bakaze: bakaze,
		zikaze: zikaze,
		deck:   deck,
		hand:   hand,
		kawa:   kawa,
		naki:   naki,
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

func (p *PlayerImpl) CanNakiActions(inTile *tile.Tile) []*naki.NakiAction {
	actions := []*naki.NakiAction{}
	if p.canChii(inTile) {
		actions = append(actions, &naki.Chii)
	}
	if p.canPon(inTile) {
		actions = append(actions, &naki.Pon)
	}
	if p.canKan(inTile) {
		actions = append(actions, &naki.Kan)
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

type PlayerStatus struct {
	ID          string
	Name        string               `json:"name"`
	Zikaze      string               `json:"zikaze"`
	Tsumo       string               `json:"tsumo"`
	Hand        []string             `json:"hand"`
	Kawa        []*kawa.KawaStatus   `json:"kawa"`
	NakiActions *NakiActions         `json:"naki_actions"`
	Naki        [][]*naki.NakiStatus `json:"naki"`
}

func (p *PlayerImpl) Status(nakiTile *tile.Tile) *PlayerStatus {
	return &PlayerStatus{
		ID:          p.id,
		Name:        p.name,
		Zikaze:      p.safeZihaiName(),
		Tsumo:       p.safeTsumoName(),
		Hand:        p.hand.Status(),
		Kawa:        p.kawa.Status(),
		NakiActions: p.NakiActionStatus(nakiTile),
		Naki:        p.naki.Status(),
	}
}

func (p *PlayerImpl) safeTsumoName() string {
	if p.tsumo == nil {
		return ""
	}
	return p.tsumo.Name()
}

func (p *PlayerImpl) safeZihaiName() string {
	if p.zikaze == nil {
		return ""
	}
	return p.zikaze.Name()

}

type NakiActions struct {
	Pon  [][2]*tile.Tile `json:"pon"`
	Kan  [][3]*tile.Tile `json:"kan"`
	Chii [][2]*tile.Tile `json:"chii"`
}

func (p *PlayerImpl) NakiActionStatus(inTile *tile.Tile) *NakiActions {
	return &NakiActions{
		Pon:  p.hand.FindPonPair(inTile),
		Kan:  p.hand.FindKanPair(inTile),
		Chii: p.hand.FindPonPair(inTile),
	}
}
