package hand

import (
	"aws-mahjong/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHand(t *testing.T) {
	hand := NewHand()
	assert.Equal(t, []*tile.Tile{}, hand.tiles)
}

func TestAdd(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*tile.Tile
		InputTile     *tile.Tile
		ExpectedError error
	}{
		{
			Description:   "valid case",
			CurrentTiles:  []*tile.Tile{&tile.North, &tile.East, &tile.West},
			InputTile:     &tile.East,
			ExpectedError: nil,
		},
		{
			Description: "invalid case",
			CurrentTiles: []*tile.Tile{
				&tile.North, &tile.East, &tile.West, &tile.Manzu1,
				&tile.Manzu1, &tile.Manzu2, &tile.Manzu3, &tile.Manzu4,
				&tile.Manzu5, &tile.Manzu6, &tile.Manzu7, &tile.Manzu8,
				&tile.Manzu9,
			},
			InputTile:     &tile.Manzu1,
			ExpectedError: TileCountErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := NewHand()
			hand.tiles = c.CurrentTiles
			err := hand.Add(c.InputTile)
			assert.Equal(t, c.ExpectedError, err)
		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		Description   string
		InputTiles    []*tile.Tile
		OutTile       *tile.Tile
		ExpectedError error
	}{
		{
			Description:   "valid case",
			InputTiles:    []*tile.Tile{&tile.East},
			OutTile:       &tile.East,
			ExpectedError: nil,
		},
		{
			Description:   "invalid case, not found",
			InputTiles:    []*tile.Tile{&tile.East},
			OutTile:       &tile.West,
			ExpectedError: TileNotFoundErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := NewHand()
			hand.tiles = c.InputTiles
			tile, err := hand.Remove(c.OutTile)
			if err != nil && err == c.ExpectedError {
				return

			}
			assert.Equal(t, c.OutTile, tile)
		})

	}

}
