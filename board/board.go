package board

import (
	"aws-mahjong/deck"
	"aws-mahjong/player"
	"aws-mahjong/tile"
	"errors"
)

type Board struct {
	bakaze  *tile.Tile
	deck    deck.Deck
	oya     int
	players []*player.Player
	turn    int

	// tmp data
	boardNakiTile *tile.Tile
	// TODO mutex
}

var (
	BoardNakiTileAlreadyExist = errors.New("baord naki tile already exist")
)

type UserInfo struct {
	ID   string
	Name string
}

func NewBoard(users []UserInfo) *Board {
	newDeck := deck.NewDeck()
	players := []*player.Player{}
	bakaze := tile.Bakazes[0]

	for idx, user := range users {
		players = append(players, player.NewPlayer(user.ID, user.Name, newDeck, bakaze, tile.Zikazes[idx], false))
	}

	return &Board{
		bakaze:  bakaze,
		deck:    newDeck,
		oya:     0,
		players: players,
		turn:    0,
	}
}

func (b *Board) TurnPlayerTsumo() error {
	err := b.players[b.turn].Tsumo()
	return err
}

func (b *Board) TurnPlayerDahai(outTile *tile.Tile) error {
	if b.boardNakiTile != nil {
		return BoardNakiTileAlreadyExist
	}
	outTile, err := b.players[b.turn].Dahai(outTile)
	if err != nil {
		return err
	}
	b.boardNakiTile = outTile

	// check all player's naki
	return nil
}

func (b *Board) Status() error {
	// TODO send all information to client
	// TODO create view layer to wrap information
	return nil
}
