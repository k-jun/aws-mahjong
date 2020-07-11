package player

import (
	"aws-mahjong/deck"
	"aws-mahjong/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTsumo(t *testing.T) {
	cases := []struct {
		Description   string
		PlayerName    string
		CurrentTsumo  *tile.Tile
		CurrentDeck   *deck.Deck
		ExpectedError error
	}{
		{
			Description:   "valid tsumo",
			PlayerName:    "Alba Abshire",
			CurrentTsumo:  nil,
			CurrentDeck:   deck.NewDeck(),
			ExpectedError: nil,
		},
		{
			Description:   "invalid tsumo",
			PlayerName:    "Kaci Larkin",
			CurrentTsumo:  &tile.West,
			CurrentDeck:   deck.NewDeck(),
			ExpectedError: TsumoAlreadyExistErr,
		},
		{
			Description:   "invalid deck",
			PlayerName:    "Frankie Schumm",
			CurrentTsumo:  nil,
			CurrentDeck:   blankDeck(),
			ExpectedError: deck.RunOutOfTileErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			player := NewPlayer(c.PlayerName, c.CurrentDeck)
			player.tsumo = c.CurrentTsumo
			err := player.Tsumo()
			assert.Equal(t, c.ExpectedError, err)
		})
	}

}

func blankDeck() *deck.Deck {
	deck := deck.NewDeck()
	for {
		_, err := deck.Draw()
		if err != nil {
			break
		}
	}
	return deck
}
