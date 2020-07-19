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
		CurrentTsumo  *tile.Tile
		CurrentDeck   deck.Deck
		ExpectedError error
	}{
		{
			Description:   "valid tsumo",
			CurrentTsumo:  nil,
			CurrentDeck:   deck.NewDeck(),
			ExpectedError: nil,
		},
		{
			Description:   "invalid tsumo",
			CurrentTsumo:  &tile.West,
			CurrentDeck:   deck.NewDeck(),
			ExpectedError: TsumoAlreadyExistErr,
		},
		{
			Description:   "invalid deck",
			CurrentTsumo:  nil,
			CurrentDeck:   blankDeck(),
			ExpectedError: deck.RunOutOfTileErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			player := &PlayerImpl{
				tsumo: c.CurrentTsumo,
				deck:  c.CurrentDeck,
			}
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
			player := PlayerImpl{
				deck:  c.CurrentDeck,
				hand:  hand.NewHand(),
				kawa:  kawa.NewKawa(),
				naki:  &naki.NakiMock{},
				tsumo: c.CurrentTsumo,
			}
			if err := player.Hand().Adds(c.CurrentHandTiles); err != nil {
				t.Fatal()
			}

			outTile, err := player.Dahai(c.InputOutTile)
			if err != nil && err == c.ExpectedError {
				return
			}
			assert.Equal(t, c.ExpectedError, err)
			assert.Equal(t, c.ExpectedHandTiles, player.Hand().Tiles())
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
			player := PlayerImpl{kawa: &kawa.KawaMock{ExpectedError: c.OutError}}
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
		OutActions       []*naki.NakiAction
	}{
		{
			Description:      "found pon",
			CurrentHandTiles: []*tile.Tile{&tile.Manzu3, &tile.Manzu3},
			InTile:           &tile.Manzu3,
			OutActions:       []*naki.NakiAction{&naki.Pon},
		},
		{
			Description:      "found kan, and pon",
			CurrentHandTiles: []*tile.Tile{&tile.Manzu3, &tile.Manzu3, &tile.Manzu3},
			InTile:           &tile.Manzu3,
			OutActions:       []*naki.NakiAction{&naki.Pon, &naki.Kan},
		},
		{
			Description:      "found pon, and chii",
			CurrentHandTiles: []*tile.Tile{&tile.Manzu4, &tile.Manzu3, &tile.Manzu4, &tile.Manzu5},
			InTile:           &tile.Manzu4,
			OutActions:       []*naki.NakiAction{&naki.Chii, &naki.Pon},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			player := &PlayerImpl{
				deck: deck.NewDeck(),
				hand: hand.NewHand(),
				kawa: kawa.NewKawa(),
				naki: naki.NewNaki(),
			}
			if err := player.Hand().Adds(c.CurrentHandTiles); err != nil {
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
			player := &PlayerImpl{
				deck: deck.NewDeck(),
				hand: &handMock,
				kawa: kawa.NewKawa(),
				naki: &nakiMock,
			}

			err := player.Naki(c.InTile, c.InTiles, naki.Jicha)
			assert.Equal(t, c.OutError, err)
		})
	}
}

func TestStatus(t *testing.T) {
	cases := []struct {
		Description         string
		CurrentName         string
		CurrentHand         hand.Hand
		CurrentKawa         kawa.Kawa
		CurrentNaki         naki.Naki
		CurrentTsumo        *tile.Tile
		CurrentZihai        *tile.Tile
		InTile              *tile.Tile
		OutNakiActionStatus *NakiActions
		OutTsumo            string
	}{
		{
			Description:  "valid case",
			CurrentName:  "Edgar O'Connell I",
			CurrentZihai: &tile.East,
			CurrentTsumo: nil,
			CurrentHand: &hand.HandMock{
				ExpectedStatus: []string{"manzu1", "manzu2", "manzu3"},
				ExpectedPair2:  [][2]*tile.Tile{{&tile.Chun, &tile.East}},
				ExpectedPair3:  [][3]*tile.Tile{{&tile.Chun, &tile.East, &tile.Haku}},
			},
			CurrentKawa: &kawa.KawaMock{ExpectedStatus: []*kawa.KawaStatus{{Name: "manzu1", IsSide: true}, {Name: "east", IsSide: false}}},
			CurrentNaki: &naki.NakiMock{ExpectedStatus: [][]*naki.NakiStatus{
				{{Name: "pinzu1", IsOpen: true, IsSide: true}, {Name: "pinzu1", IsOpen: false, IsSide: true}, {Name: "pinzu1", IsOpen: false, IsSide: false}},
				{{Name: "souzu1", IsOpen: true, IsSide: true}, {Name: "souzu1", IsOpen: false, IsSide: true}, {Name: "souzu1", IsOpen: false, IsSide: false}},
			}},
			InTile: &tile.Manzu4,
			OutNakiActionStatus: &NakiActions{
				Pon:  [][2]*tile.Tile{{&tile.Chun, &tile.East}},
				Chii: [][2]*tile.Tile{{&tile.Chun, &tile.East}},
				Kan:  [][3]*tile.Tile{{&tile.Chun, &tile.East, &tile.Haku}},
			},
			OutTsumo: "",
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			player := &PlayerImpl{
				hand:   c.CurrentHand,
				kawa:   c.CurrentKawa,
				naki:   c.CurrentNaki,
				zikaze: c.CurrentZihai,
				tsumo:  c.CurrentTsumo,
			}
			status := player.Status(c.InTile)
			assert.Equal(t, c.CurrentHand.Status(), status.Hand)
			assert.Equal(t, c.CurrentKawa.Status(), status.Kawa)
			assert.Equal(t, c.CurrentNaki.Status(), status.Naki)
			assert.Equal(t, c.OutTsumo, status.Tsumo)
			assert.Equal(t, c.OutNakiActionStatus, status.NakiActions)

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
