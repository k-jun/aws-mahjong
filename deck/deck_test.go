package deck

import (
	"aws-mahjong/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	counter := map[string]int{}

	for _, i := range deck.tiles {
		counter[i.Name()] += 1
	}

	for _, i := range tile.All {
		assert.Equal(t, tile.Count, counter[i.Name()])
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	size := len(deck.tiles)

	counter := map[string]int{}

	for i := 0; i < size; i++ {
		tile, err := deck.Draw()
		assert.NoError(t, err)
		counter[tile.Name()] += 1
	}
	_, err := deck.Draw()
	assert.Equal(t, RunOutOfTileErr, err)

	for _, i := range tile.All {
		assert.Equal(t, tile.Count, counter[i.Name()])
	}
}
