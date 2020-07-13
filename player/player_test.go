package player

import (
	"aws-mahjong/deck"
	"aws-mahjong/hand"
	"aws-mahjong/kawa"
	"aws-mahjong/naki"
	"aws-mahjong/tile"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTsumo(t *testing.T) {
	cases := []struct {
		Description   string
		PlayerName    string
		CurrentTsumo  *tile.Tile
		CurrentDeck   deck.Deck
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
			player := NewPlayer(
				"test_id",
				c.PlayerName,
				nil,
				nil,
				false,
				c.CurrentDeck,
				&hand.HandMock{},
				&kawa.KawaMock{},
				&naki.NakiMock{},
			)
			player.tsumo = c.CurrentTsumo
			err := player.Tsumo()
			assert.Equal(t, c.ExpectedError, err)
		})
	}
}

func TestDahai(t *testing.T) {
	cases := []struct {
		Description       string
		PlayerName        string
		CurrentTsumo      *tile.Tile
		CurrentDeck       deck.Deck
		CurrentHandTiles  []*tile.Tile
		InputOutTile      *tile.Tile
		ExpectedTile      *tile.Tile
		ExpectedHandTiles []*tile.Tile
		ExpectedError     error
	}{
		{
			Description:       "valid case",
			PlayerName:        "Laury Schmeler",
			CurrentTsumo:      &tile.West,
			CurrentDeck:       deck.NewDeck(),
			CurrentHandTiles:  []*tile.Tile{&tile.East},
			InputOutTile:      &tile.West,
			ExpectedTile:      &tile.West,
			ExpectedHandTiles: []*tile.Tile{&tile.East},
			ExpectedError:     nil,
		},
		{
			Description:       "valid case",
			PlayerName:        "Mrs. Violet West MD",
			CurrentTsumo:      &tile.West,
			CurrentDeck:       deck.NewDeck(),
			CurrentHandTiles:  []*tile.Tile{&tile.East},
			InputOutTile:      &tile.East,
			ExpectedTile:      &tile.East,
			ExpectedHandTiles: []*tile.Tile{&tile.West},
			ExpectedError:     nil,
		},
		{
			Description:       "invalid case",
			PlayerName:        "Mrs. Mandy Thompson DVM",
			CurrentTsumo:      &tile.West,
			CurrentDeck:       deck.NewDeck(),
			CurrentHandTiles:  []*tile.Tile{&tile.East},
			InputOutTile:      &tile.South,
			ExpectedTile:      nil,
			ExpectedHandTiles: []*tile.Tile{},
			ExpectedError:     hand.TileNotFoundErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			player := NewPlayer(
				"test_id",
				c.PlayerName,
				nil,
				nil,
				false,
				c.CurrentDeck,
				hand.NewHand(),
				kawa.NewKawa(),
				&naki.NakiMock{},
			)
			player.tsumo = c.CurrentTsumo
			if err := player.hand.Adds(c.CurrentHandTiles); err != nil {
				t.Fatal()
			}

			outTile, err := player.Dahai(c.InputOutTile)
			if err != nil && err == c.ExpectedError {
				return
			}
			assert.Equal(t, c.ExpectedError, err)
			assert.Equal(t, c.ExpectedHandTiles, player.hand.Tiles())
			assert.Equal(t, c.ExpectedTile, outTile)
		})
	}
}

func TestDahaiDone(t *testing.T) {
	cases := []struct {
		Description string
		InTile      *tile.Tile
		InIsSide    bool
		OutError    error
	}{
		{
			Description: "valid case",
			InTile:      &tile.Pinzu5Aka,
			InIsSide:    false,
			OutError:    nil,
		},
		{
			Description: "invalid case",
			InTile:      &tile.Pinzu5Aka,
			InIsSide:    false,
			OutError:    errors.New(""),
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			player := Player{kawa: &kawa.KawaMock{ExpectedError: c.OutError}}
			err := player.DahaiDone(c.InTile, c.InIsSide)
			assert.Equal(t, c.OutError, err)
		})
	}
}

func TestCanNakiActions(t *testing.T) {
	cases := []struct {
		Description      string
		CurrentHandTiles []*tile.Tile
		InTile           *tile.Tile
		OutActions       []*NakiAction
	}{
		{
			Description:      "found pon",
			CurrentHandTiles: []*tile.Tile{&tile.Manzu3, &tile.Manzu3},
			InTile:           &tile.Manzu3,
			OutActions:       []*NakiAction{&Pon},
		},
		{
			Description:      "found kan, and pon",
			CurrentHandTiles: []*tile.Tile{&tile.Manzu3, &tile.Manzu3, &tile.Manzu3},
			InTile:           &tile.Manzu3,
			OutActions:       []*NakiAction{&Pon, &Kan},
		},
		{
			Description:      "found pon, and chii",
			CurrentHandTiles: []*tile.Tile{&tile.Manzu4, &tile.Manzu3, &tile.Manzu4, &tile.Manzu5},
			InTile:           &tile.Manzu4,
			OutActions:       []*NakiAction{&Chii, &Pon},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			player := NewPlayer(
				"test_id",
				"Genesis Gottlieb",
				nil,
				nil,
				false,
				deck.NewDeck(),
				hand.NewHand(),
				kawa.NewKawa(),
				naki.NewNaki(),
			)
			if err := player.hand.Adds(c.CurrentHandTiles); err != nil {
				t.Fatal()
			}

			actions := player.CanNakiActions(c.InTile)
			assert.Equal(t, c.OutActions, actions)
		})
	}
}

func TestNaki(t *testing.T) {
	cases := []struct {
		Description   string
		MockHandTiles []*tile.Tile
		MockHandError error
		MockNakiError error
		InTile        *tile.Tile
		InTiles       []*tile.Tile
		OutError      error
	}{
		{
			Description:   "valid case",
			MockHandTiles: []*tile.Tile{&tile.Manzu1, &tile.Manzu2},
			MockHandError: nil,
			MockNakiError: nil,
			InTile:        &tile.Manzu3,
			InTiles:       []*tile.Tile{&tile.Manzu1, &tile.Manzu2},
			OutError:      nil,
		},
		{
			Description:   "invalid case",
			MockHandTiles: []*tile.Tile{&tile.Manzu1, &tile.Manzu2},
			MockHandError: errors.New(""),
			MockNakiError: nil,
			InTile:        &tile.Manzu3,
			InTiles:       []*tile.Tile{&tile.Manzu1, &tile.Manzu2},
			OutError:      errors.New(""),
		},
		{
			Description:   "invalid case",
			MockHandTiles: []*tile.Tile{&tile.Manzu1, &tile.Manzu2},
			MockHandError: nil,
			MockNakiError: errors.New(""),
			InTile:        &tile.Manzu3,
			InTiles:       []*tile.Tile{&tile.Manzu1, &tile.Manzu2},
			OutError:      errors.New(""),
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			handMock := hand.HandMock{ExpectedTiles: c.MockHandTiles, ExpectedError: c.MockHandError}
			nakiMock := naki.NakiMock{ExpectedError: c.MockNakiError}
			player := NewPlayer(
				"test_id",
				"Mr. Claud Walker DVM",
				nil,
				nil,
				false,
				deck.NewDeck(),
				&handMock,
				kawa.NewKawa(),
				&nakiMock,
			)

			err := player.Naki(c.InTile, c.InTiles, naki.Jicha)
			assert.Equal(t, c.OutError, err)
		})
	}
}

func blankDeck() deck.Deck {
	deck := deck.NewDeck()
	for {
		_, err := deck.Draw()
		if err != nil {
			break
		}
	}
	return deck
}
