package board

import (
	"aws-mahjong/deck"
	"aws-mahjong/player"
	"aws-mahjong/tile"
)

type Board struct {
	bakaze  *tile.Tile
	deck    *deck.Deck
	oya     int
	players []*player.Player
	turn    int
}

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

// func (b *Board) TurnPlayerDahai() {
// 	outTile, err := b.players[b.turn].Dahai()
// }
