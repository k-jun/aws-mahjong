package deck

import (
	"aws-mahjong/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

var checker = map[string]int{
	"manzu5":    3,
	"manzu5aka": 1,
	"souzu5":    3,
	"souzu5aka": 1,
	"pinzu5":    3,
	"pinzu5aka": 1,
}

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	counter := map[string]int{}

	for _, i := range deck.tiles {
		counter[i.Name()] += 1
	}

	for _, i := range tile.AllTailKind {
		expectedNum := checker[i.Name()]
		if expectedNum == 0 {
			expectedNum = 4
		}
		assert.Equal(t, expectedNum, counter[i.Name()])
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

	for _, i := range tile.AllTailKind {
		expectedNum := checker[i.Name()]
		if expectedNum == 0 {
			expectedNum = 4
		}
		assert.Equal(t, expectedNum, counter[i.Name()])
	}
}
