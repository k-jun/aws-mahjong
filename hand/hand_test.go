package hand

import (
	"aws-mahjong/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHand(t *testing.T) {
	hand := NewHand()
	assert.Equal(t, []*tile.Tile{}, hand.Tiles())
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
			hand := HandImpl{tiles: c.CurrentTiles}
			err := hand.Add(c.InTile)
			if err != nil && err == c.ExpectedError {
				return
			}
			assert.Equal(t, c.ExpectedError, err)
			assert.Equal(t, c.ExpectedTiles, hand.Tiles())

		})
	}
}

func TestAdds(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*tile.Tile
		InTiles       []*tile.Tile
		ExpectedTiles []*tile.Tile
		ExpectedError error
	}{
		{
			Description:   "valid case",
			CurrentTiles:  []*tile.Tile{&tile.Chun, &tile.Chun, &tile.Chun},
			InTiles:       []*tile.Tile{&tile.Haku, &tile.Haku, &tile.Haku},
			ExpectedTiles: []*tile.Tile{&tile.Chun, &tile.Chun, &tile.Chun, &tile.Haku, &tile.Haku, &tile.Haku},
			ExpectedError: nil,
		},
		{
			Description:   "invalid case",
			CurrentTiles:  []*tile.Tile{&tile.Chun, &tile.Chun, &tile.Chun},
			InTiles:       []*tile.Tile{&tile.Haku, &tile.Haku, &tile.Haku, &tile.Manzu1, &tile.Manzu2, &tile.Manzu3, &tile.Manzu4, &tile.Manzu5, &tile.Chun, &tile.East, &tile.Haku},
			ExpectedTiles: []*tile.Tile{},
			ExpectedError: TileCountErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := HandImpl{tiles: c.CurrentTiles}
			err := hand.Adds(c.InTiles)
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
			hand := HandImpl{tiles: c.CurrentTiles}
			outTile, err := hand.Remove(c.OutTile)
			if err != nil && err == c.ExpectedError {
				return
			}
			assert.Equal(t, c.OutTile, outTile)
		})
	}
}

func TestRemoves(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*tile.Tile
		OutTiles      []*tile.Tile
		ExpectedTiles []*tile.Tile
		ExpectedError error
	}{
		{
			Description:   "valid case",
			CurrentTiles:  []*tile.Tile{&tile.Chun, &tile.Chun},
			OutTiles:      []*tile.Tile{&tile.Chun, &tile.Chun},
			ExpectedTiles: []*tile.Tile{},
			ExpectedError: nil,
		},
		{
			Description:   "invalid case",
			CurrentTiles:  []*tile.Tile{&tile.Chun},
			OutTiles:      []*tile.Tile{&tile.Chun, &tile.Chun},
			ExpectedTiles: []*tile.Tile{},
			ExpectedError: TileNotFoundErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {

			hand := HandImpl{tiles: c.CurrentTiles}
			_, err := hand.Removes(c.OutTiles)
			if err != nil && err == c.ExpectedError {
				return
			}
			assert.Equal(t, c.ExpectedError, err)
			assert.Equal(t, c.ExpectedTiles, hand.Tiles())
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
			hand := HandImpl{tiles: c.CurrentTiles}
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

func TestFindPonPair(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*tile.Tile
		InTile        *tile.Tile
		ExpectedPairs [][2]*tile.Tile
	}{
		{
			Description:  "suhai pair",
			CurrentTiles: []*tile.Tile{&tile.West, &tile.West},
			InTile:       &tile.West,
			ExpectedPairs: [][2]*tile.Tile{
				[2]*tile.Tile{&tile.West, &tile.West},
			},
		},
		{
			Description:  "zihai pair",
			CurrentTiles: []*tile.Tile{&tile.Manzu1, &tile.Manzu1, &tile.Manzu1, &tile.West},
			InTile:       &tile.Manzu1,
			ExpectedPairs: [][2]*tile.Tile{
				[2]*tile.Tile{&tile.Manzu1, &tile.Manzu1},
			},
		},
		{
			Description:  "zihai 5 pair",
			CurrentTiles: []*tile.Tile{&tile.Manzu5, &tile.Manzu5Aka, &tile.Manzu5, &tile.West},
			InTile:       &tile.Manzu5,
			ExpectedPairs: [][2]*tile.Tile{
				[2]*tile.Tile{&tile.Manzu5, &tile.Manzu5},
				[2]*tile.Tile{&tile.Manzu5, &tile.Manzu5Aka},
			},
		},
		{
			Description:   "zihai no pair",
			CurrentTiles:  []*tile.Tile{&tile.Manzu4, &tile.Manzu5Aka, &tile.Manzu5, &tile.West},
			InTile:        &tile.Manzu3,
			ExpectedPairs: [][2]*tile.Tile{},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := HandImpl{tiles: c.CurrentTiles}
			pairs := hand.FindPonPair(c.InTile)
			assert.Equal(t, c.ExpectedPairs, pairs)

		})
	}
}

func TestFindChiiPair(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*tile.Tile
		InTile        *tile.Tile
		ExpectedPairs [][2]*tile.Tile
	}{
		{
			Description:   "zihai no pair",
			CurrentTiles:  []*tile.Tile{&tile.West, &tile.North, &tile.South, &tile.East},
			InTile:        &tile.East,
			ExpectedPairs: [][2]*tile.Tile{},
		},
		{
			Description:  "suhai pair",
			CurrentTiles: []*tile.Tile{&tile.Manzu1, &tile.Manzu2},
			InTile:       &tile.Manzu3,
			ExpectedPairs: [][2]*tile.Tile{
				[2]*tile.Tile{&tile.Manzu1, &tile.Manzu2},
			},
		},
		{
			Description:   "suhai no pair",
			CurrentTiles:  []*tile.Tile{&tile.Souzu1, &tile.Manzu2},
			InTile:        &tile.Manzu3,
			ExpectedPairs: [][2]*tile.Tile{},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := HandImpl{tiles: c.CurrentTiles}
			pairs := hand.FindChiiPair(c.InTile)
			assert.Equal(t, c.ExpectedPairs, pairs)
		})

	}

}

func TestFindKanPair(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTiles  []*tile.Tile
		InTile        *tile.Tile
		ExpectedPairs [][3]*tile.Tile
	}{
		{
			Description:  "zihai pair",
			CurrentTiles: []*tile.Tile{&tile.West, &tile.East, &tile.West, &tile.West},
			InTile:       &tile.West,
			ExpectedPairs: [][3]*tile.Tile{
				[3]*tile.Tile{&tile.West, &tile.West, &tile.West},
			},
		},
		{
			Description:  "suhai pair",
			CurrentTiles: []*tile.Tile{&tile.Manzu1, &tile.Manzu1, &tile.Manzu1},
			InTile:       &tile.Manzu1,
			ExpectedPairs: [][3]*tile.Tile{
				[3]*tile.Tile{&tile.Manzu1, &tile.Manzu1, &tile.Manzu1},
			},
		},
		{
			Description:  "suhai 5 pair",
			CurrentTiles: []*tile.Tile{&tile.Manzu5, &tile.Manzu5Aka, &tile.Manzu5},
			InTile:       &tile.Manzu5,
			ExpectedPairs: [][3]*tile.Tile{
				[3]*tile.Tile{&tile.Manzu5, &tile.Manzu5, &tile.Manzu5Aka},
			},
		},
		{
			Description:   "suhai no pair",
			CurrentTiles:  []*tile.Tile{&tile.Manzu4, &tile.Manzu5Aka, &tile.Manzu5},
			InTile:        &tile.Manzu5,
			ExpectedPairs: [][3]*tile.Tile{},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := HandImpl{tiles: c.CurrentTiles}
			pairs := hand.FindKanPair(c.InTile)
			assert.Equal(t, c.ExpectedPairs, pairs)

		})

	}

}

func TestStatus(t *testing.T) {
	cases := []struct {
		Description  string
		CurrentTiles []*tile.Tile
		OutStatus    []string
	}{
		{
			Description:  "valid case",
			CurrentTiles: []*tile.Tile{&tile.Manzu5, &tile.Manzu5Aka, &tile.Manzu5},
			OutStatus:    []string{"manzu5", "manzu5", "manzu5aka"},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			hand := &HandImpl{tiles: c.CurrentTiles}
			assert.Equal(t, c.OutStatus, hand.Status())
		})
	}
}
