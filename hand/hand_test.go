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
		InTile        *tile.Tile
		ExpectedTiles []*tile.Tile
		ExpectedError error
	}{
		{
			Description:   "valid case",
			CurrentTiles:  []*tile.Tile{&tile.North, &tile.East, &tile.West},
			InTile:        &tile.South,
			ExpectedTiles: []*tile.Tile{&tile.East, &tile.North, &tile.South, &tile.West},
			ExpectedError: nil,
		},
		{
			Description:   "invalid case",
			CurrentTiles:  []*tile.Tile{&tile.North, &tile.East, &tile.West, &tile.South, &tile.Manzu1, &tile.Manzu2, &tile.Manzu3, &tile.Manzu4, &tile.Manzu5, &tile.Manzu6, &tile.Manzu7, &tile.Manzu8, &tile.Manzu9},
			InTile:        &tile.South,
			ExpectedTiles: []*tile.Tile{},
			ExpectedError: TileCountErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := NewHand()
			hand.tiles = c.CurrentTiles
			err := hand.Add(c.InTile)
			if err != nil && err == c.ExpectedError {
				return
			}
			assert.Equal(t, c.ExpectedError, err)
			assert.Equal(t, c.ExpectedTiles, hand.Tiles())

		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*tile.Tile
		OutTile       *tile.Tile
		ExpectedTiles []*tile.Tile
		ExpectedError error
	}{
		{
			Description:   "valid case",
			CurrentTiles:  []*tile.Tile{&tile.East},
			OutTile:       &tile.East,
			ExpectedError: nil,
		},
		{
			Description:   "invalid case",
			CurrentTiles:  []*tile.Tile{&tile.East},
			OutTile:       &tile.West,
			ExpectedError: TileNotFoundErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := NewHand()
			hand.tiles = c.CurrentTiles
			outTile, err := hand.Remove(c.OutTile)
			if err != nil && err == c.ExpectedError {
				return

			}
			assert.Equal(t, c.OutTile, outTile)
		})

	}

}

func TestReplace(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*tile.Tile
		InTile        *tile.Tile
		OutTile       *tile.Tile
		ExpectedTiles []*tile.Tile
		ExpectedError error
	}{
		{
			Description:   "valid case",
			CurrentTiles:  []*tile.Tile{&tile.North, &tile.East},
			InTile:        &tile.West,
			OutTile:       &tile.East,
			ExpectedTiles: []*tile.Tile{&tile.North, &tile.West},
			ExpectedError: nil,
		},
		{
			Description:   "invalid case",
			CurrentTiles:  []*tile.Tile{&tile.North, &tile.East},
			InTile:        &tile.West,
			OutTile:       &tile.South,
			ExpectedTiles: []*tile.Tile{&tile.North, &tile.West},
			ExpectedError: TileNotFoundErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := NewHand()
			hand.tiles = c.CurrentTiles
			outTile, err := hand.Replace(c.InTile, c.OutTile)
			if err != nil && err == c.ExpectedError {
				return
			}
			assert.Equal(t, c.ExpectedError, err)
			assert.Equal(t, c.ExpectedTiles, hand.Tiles())
			assert.Equal(t, c.OutTile, outTile)
		})
	}
}
