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
	TurnPlayerDahai(playerId string, outTile *tile.Tile) error
	NextTurn()
	ChangeTurn(playerIdx int) error
	Start() error
	Status(playerID string) *BoardStatus
	IsTurnPlayer(playerID string) bool
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
	NotTurnPlayerErr          = errors.New("this PlayerID is not turn PlayerID")
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

func (b *BoardImpl) TurnPlayerDahai(playerId string, outTile *tile.Tile) error {
	if !b.IsTurnPlayer(playerId) {
		return NotTurnPlayerErr
	}
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
	Bakaze  string                 `json:"bakaze"`
	DeckLen int                    `json:"deck_len"`
	Jicha   int                    `json:"jicha"`
	Oya     int                    `json:"oya"`
	Turn    int                    `json:"turn"`
	Players []*player.PlayerStatus `json:"players"`
}

func (b *BoardImpl) Status(playerID string) *BoardStatus {
	playerStatuses := []*player.PlayerStatus{}
	myIdx := 0
	for idx, player := range b.players {
		playerStatus := player.Status(b.nakiTile)
		if playerStatus.ID == playerID {
			myIdx = idx
		}
		playerStatuses = append(playerStatuses, playerStatus)
	}
	status := &BoardStatus{
		Bakaze:  b.bakaze.Name(),
		DeckLen: b.deck.Count(),
		Jicha:   myIdx,
		Oya:     b.oya,
		Turn:    b.turn,
		Players: playerStatuses,
	}

	return status
}

func (b *BoardImpl) IsTurnPlayer(playerID string) bool {
	for idx, player := range b.players {
		if player.ID() == playerID {
			return idx == b.turn
		}
	}

	return false

}
