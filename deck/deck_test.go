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
