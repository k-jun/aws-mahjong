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