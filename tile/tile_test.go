package tile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	cases := []struct {
		Description string
		Tile        *Tile
	}{
		{
			Description: "west check",
			Tile:        &West,
		},
		{
			Description: "east check",
			Tile:        &East,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			assert.Equal(t, c.Tile.name, c.Tile.Name())
		})
	}
}

func TestNumber(t *testing.T) {
	cases := []struct {
		Description    string
		Tile           *Tile
		ExpectedNumber int
	}{
		{
			Description:    "suhai case",
			Tile:           &Manzu1,
			ExpectedNumber: 1,
		},
		{
			Description:    "zihai case",
			Tile:           &West,
			ExpectedNumber: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			assert.Equal(t, c.ExpectedNumber, c.Tile.Number())
		})
	}
}

func TestIsSame(t *testing.T) {
	cases := []struct {
		Description    string
		Tile           *Tile
		InTile         *Tile
		ExpectedResult bool
	}{
		{
			Description:    "suhai 1 case",
			Tile:           &Manzu1,
			InTile:         &Manzu1,
			ExpectedResult: true,
		},
		{
			Description:    "suhai 1 case",
			Tile:           &Manzu1,
			InTile:         &Pinzu1,
			ExpectedResult: false,
		},
		{
			Description:    "suhai 5 case",
			Tile:           &Souzu5,
			InTile:         &Souzu5Aka,
			ExpectedResult: true,
		},
		{
			Description:    "zihai 5 case",
			Tile:           &West,
			InTile:         &West,
			ExpectedResult: true,
		},
		{
			Description:    "zihai 5 case",
			Tile:           &West,
			InTile:         &East,
			ExpectedResult: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			assert.Equal(t, c.ExpectedResult, c.Tile.IsSame(c.InTile))
		})
	}
}

func TestIsSuhai(t *testing.T) {
	cases := []struct {
		Description    string
		Tile           *Tile
		ExpectedResult bool
	}{
		{
			Description:    "suhai case",
			Tile:           &Manzu1,
			ExpectedResult: true,
		},
		{
			Description:    "zihai case",
			Tile:           &South,
			ExpectedResult: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			assert.Equal(t, c.ExpectedResult, c.Tile.IsSuhai())
		})
	}
}

func TestIsZihai(t *testing.T) {
	cases := []struct {
		Description    string
		Tile           *Tile
		ExpectedResult bool
	}{
		{
			Description:    "suhai case",
			Tile:           &Manzu1,
			ExpectedResult: false,
		},
		{
			Description:    "zihai case",
			Tile:           &South,
			ExpectedResult: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			assert.Equal(t, c.ExpectedResult, c.Tile.IsZihai())
		})
	}
}

func TestTileKindFromString(t *testing.T) {
	cases := []struct {
		Description    string
		InputString    string
		ExpectedResult *TileKind
	}{
		{
			Description:    "valid string",
			InputString:    "1",
			ExpectedResult: &one,
		},
		{
			Description:    "valid string",
			InputString:    "suhai",
			ExpectedResult: &suhai,
		},
		{
			Description:    "invalid string",
			InputString:    "not_existing_key",
			ExpectedResult: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			assert.Equal(t, c.ExpectedResult, TileKindFromString(c.InputString))

		})
	}
}

func TestKindSMP(t *testing.T) {
	cases := []struct {
		Description    string
		InTile         *Tile
		ExpectedResult *TileKind
	}{
		{
			Description:    "souzu case",
			InTile:         &Souzu1,
			ExpectedResult: &souzu,
		},
		{
			Description:    "manzu case",
			InTile:         &Manzu1,
			ExpectedResult: &manzu,
		},
		{
			Description:    "pinzu case",
			InTile:         &Pinzu1,
			ExpectedResult: &pinzu,
		},
		{
			Description:    "invalid case",
			InTile:         &West,
			ExpectedResult: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			assert.Equal(t, c.ExpectedResult, c.InTile.KindSMP())
		})
	}
}

func TestSortTiles(t *testing.T) {
	cases := []struct {
		Description string
		InTiles     []*Tile
		OutTiles    []*Tile
	}{
		{
			Description: "only suhai",
			InTiles:     []*Tile{&Manzu9, &Manzu7, &Manzu8, &Manzu4},
			OutTiles:    []*Tile{&Manzu4, &Manzu7, &Manzu8, &Manzu9},
		},
		{
			Description: "only zihai",
			InTiles:     []*Tile{&West, &East, &North, &South},
			OutTiles:    []*Tile{&East, &North, &South, &West},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			tiles := c.InTiles
			SortTiles(tiles)
			assert.Equal(t, c.OutTiles, tiles)
		})

	}

}
