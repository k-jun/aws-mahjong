package board

import (
	"aws-mahjong/deck"
	"aws-mahjong/hand"
	"aws-mahjong/player"
	"aws-mahjong/tile"
	"errors"
)

type Board interface {
	TurnPlayerTsumo() error
	TurnPlayerDahai(outTile *tile.Tile) error
	NextTurn()
	ChangeTurn(playerIdx int) error
	Start() error
	Status(playerID string) *BoardStatus
}

type BoardImpl struct {
	bakaze  *tile.Tile
	deck    deck.Deck
	oya     int
	players []player.Player
	turn    int

	nakiTile *tile.Tile
}

var (
	BoardNakiTileAlreadyExist = errors.New("board naki tile already exist")
	BoardTurnOutOfRange       = errors.New("specified turn is out of range")
	GameAlreadyStarted        = errors.New("game have already started")
)

func NewBoard(deck deck.Deck, players []player.Player) Board {
	return &BoardImpl{
		bakaze:  tile.Bakazes[0],
		deck:    deck,
		oya:     0,
		players: players,
		turn:    0,
	}
}

func (b *BoardImpl) Start() error {
	for _, player := range b.players {
		if len(player.Hand().Tiles()) != 0 {
			return GameAlreadyStarted
		}
	}

	for _, player := range b.players {
		tiles := []*tile.Tile{}
		for i := 0; i < hand.HandCount; i++ {
			tile, err := b.deck.Draw()
			if err != nil {
				return err
			}
			tiles = append(tiles, tile)
		}

		err := player.Hand().Adds(tiles)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *BoardImpl) TurnPlayerTsumo() error {
	err := b.players[b.turn].Tsumo()
	return err
}

func (b *BoardImpl) CanOtherPlayersNaki(nakiTile *tile.Tile) bool {
	for idx, player := range b.players {
		if idx == b.turn {
			continue
		}
		actions := player.CanNakiActions(nakiTile)
		if len(actions) != 0 {
			b.nakiTile = nakiTile
			return true
		}
	}
	return false

}

func (b *BoardImpl) TurnPlayerDahai(outTile *tile.Tile) error {
	if b.nakiTile != nil {
		return BoardNakiTileAlreadyExist
	}
	outTile, err := b.players[b.turn].Dahai(outTile)
	if err != nil {
		return err
	}

	// check all player's naki
	// if b.CanOtherPlayersNaki(outTile) {
	// 	b.nakiTile = outTile
	// 	return nil
	// }

	// change turn
	b.NextTurn()
	err = b.TurnPlayerTsumo()
	if err != nil {
		return err
	}

	return nil
}

func (b *BoardImpl) NextTurn() {
	b.turn = (b.turn + 1) % len(b.players)
}

func (b *BoardImpl) ChangeTurn(playerIdx int) error {
	if playerIdx < 0 || playerIdx >= len(b.players) {
		return BoardTurnOutOfRange
	}
	b.turn = playerIdx
	return nil
}

type BoardStatus struct {
	Bakaze    string               `json:"bakaze"`
	DeckCound int                  `json:"deck_count"`
	Oya       string               `json:"oya"`
	Turn      string               `json:"turn"`
	Jicha     *player.PlayerStatus `json:"jicha"`
	Kamicha   *player.PlayerStatus `json:"kamicha"`
	Toimen    *player.PlayerStatus `json:"toimen"`
	Shimocha  *player.PlayerStatus `json:"shimocha"`
}

type Cha string

var (
	Kamicha  Cha = "kamicha"
	Toimen   Cha = "toimen"
	Shimocha Cha = "shimocha"
	Jicha    Cha = "jicha"
	Nil      Cha = ""
)

func (b *BoardImpl) Status(playerID string) *BoardStatus {
	status := &BoardStatus{
		Bakaze:    b.bakaze.Name(),
		DeckCound: b.deck.Count(),
	}

	myIdx := 0
	playerStatuses := []*player.PlayerStatus{}
	for idx, player := range b.players {
		playerStatus := player.Status(b.nakiTile)
		if playerStatus.ID == playerID {
			myIdx = idx
		}
		playerStatuses = append(playerStatuses, playerStatus)
	}

	status.Oya = string(getCha(myIdx, b.oya))
	status.Turn = string(getCha(myIdx, b.turn))

	for idx, s := range playerStatuses {
		switch getCha(myIdx, idx) {
		case Kamicha:
			status.Kamicha = s
		case Toimen:
			status.Toimen = s
		case Shimocha:
			status.Shimocha = s
		case Jicha:
			status.Jicha = s
		default:
			continue
		}
	}

	return status
}

func getCha(jichaIdx int, tachaIdx int) Cha {
	switch jichaIdx - tachaIdx {
	case 3, -1:
		return Shimocha
	case 2, -2:
		return Toimen
	case 1, -3:
		return Kamicha
	case 0:
		return Jicha
	default:
		return Nil
	}
}
