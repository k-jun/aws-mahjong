package board

import (
	"aws-mahjong/naki"
	"aws-mahjong/player"
	"aws-mahjong/tile"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnPlayerTsumo(t *testing.T) {
	cases := []struct {
		Description    string
		CurrentPlayers []player.Player
		CurrentTurn    int
		OutError       error
	}{
		{
			Description: "valid case",
			CurrentPlayers: []player.Player{&player.PlayerMock{
				ExpectedTile:  &tile.Chun,
				ExpectedError: nil,
			}},
			CurrentTurn: 0,
			OutError:    nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			board := BoardImpl{
				players: c.CurrentPlayers,
				turn:    c.CurrentTurn,
			}
			err := board.TurnPlayerTsumo()
			assert.Equal(t, c.OutError, err)
		})
	}
}

func TestNextTurn(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentTurn   int
		CurrentPlayer []player.Player
		OutTurn       int
	}{
		{
			Description:   "0 -> 1",
			CurrentTurn:   0,
			CurrentPlayer: []player.Player{&player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}},
			OutTurn:       1,
		},
		{
			Description:   "3 -> 0",
			CurrentTurn:   3,
			CurrentPlayer: []player.Player{&player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}},
			OutTurn:       0,
		},
		{
			Description:   "2 -> 0",
			CurrentTurn:   2,
			CurrentPlayer: []player.Player{&player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}},
			OutTurn:       0,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			board := BoardImpl{turn: c.CurrentTurn, players: c.CurrentPlayer}
			board.NextTurn()
			assert.Equal(t, c.OutTurn, board.turn)

		})
	}
}

func TestChangeTurn(t *testing.T) {
	cases := []struct {
		Description    string
		CurrentPlayers []player.Player
		InPlayerIdx    int
		OutError       error
		OutTurn        int
	}{
		{
			Description:    "valid case",
			CurrentPlayers: []player.Player{&player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}},
			InPlayerIdx:    0,
			OutError:       nil,
			OutTurn:        0,
		},
		{
			Description:    "invalid case",
			CurrentPlayers: []player.Player{&player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}},
			InPlayerIdx:    -1,
			OutError:       BoardTurnOutOfRange,
			OutTurn:        0,
		},
		{
			Description:    "invalid case",
			CurrentPlayers: []player.Player{&player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}, &player.PlayerMock{}},
			InPlayerIdx:    4,
			OutError:       BoardTurnOutOfRange,
			OutTurn:        0,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			board := BoardImpl{players: c.CurrentPlayers}
			err := board.ChangeTurn(c.InPlayerIdx)
			if err != nil && err == c.OutError {
				return
			}
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutTurn, board.turn)
		})
	}
}

func TestTurnPlayerDahai(t *testing.T) {
	cases := []struct {
		Description         string
		CurrentNakiTile     *tile.Tile
		CurrentFirstPlayer  *player.PlayerMock
		CurrentSecondPlayer *player.PlayerMock
		CurrentTurn         int
		InTile              *tile.Tile
		OutError            error
		OutTurn             int
	}{
		{
			Description:         "valid case",
			CurrentNakiTile:     nil,
			CurrentFirstPlayer:  &player.PlayerMock{ExpectedTile: &tile.Chun},
			CurrentSecondPlayer: &player.PlayerMock{ExpectedNakiActions: []*naki.NakiAction{}},
			CurrentTurn:         0,
			InTile:              &tile.Chun,
			OutError:            nil,
			OutTurn:             1,
		},
		{
			Description:         "valid case",
			CurrentNakiTile:     nil,
			CurrentFirstPlayer:  &player.PlayerMock{ExpectedTile: &tile.Chun},
			CurrentSecondPlayer: &player.PlayerMock{ExpectedNakiActions: []*naki.NakiAction{&naki.Pon}},
			CurrentTurn:         0,
			InTile:              &tile.Chun,
			OutError:            nil,
			OutTurn:             0,
		},
		{
			Description:         "invalid case",
			CurrentNakiTile:     &tile.Chun,
			CurrentFirstPlayer:  &player.PlayerMock{},
			CurrentSecondPlayer: &player.PlayerMock{},
			CurrentTurn:         0,
			InTile:              &tile.Chun,
			OutError:            BoardNakiTileAlreadyExist,
			OutTurn:             0,
		},
		{
			Description:         "invalid case",
			CurrentNakiTile:     nil,
			CurrentFirstPlayer:  &player.PlayerMock{ExpectedError: errors.New("")},
			CurrentSecondPlayer: &player.PlayerMock{},
			CurrentTurn:         0,
			InTile:              &tile.Chun,
			OutError:            errors.New(""),
			OutTurn:             0,
		},
		{
			Description:         "invalid case",
			CurrentNakiTile:     nil,
			CurrentFirstPlayer:  &player.PlayerMock{},
			CurrentSecondPlayer: &player.PlayerMock{ExpectedError: errors.New("")},
			CurrentTurn:         0,
			InTile:              &tile.Chun,
			OutError:            errors.New(""),
			OutTurn:             1,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			board := &BoardImpl{
				nakiTile: c.CurrentNakiTile,
				players:  []player.Player{c.CurrentFirstPlayer, c.CurrentSecondPlayer},
				turn:     c.CurrentTurn,
			}

			err := board.TurnPlayerDahai(c.InTile)
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutTurn, board.turn)
		})
	}
}

func TestCanOtherPlayersNaki(t *testing.T) {
	cases := []struct {
		Description         string
		InTile              *tile.Tile
		CurrentTurn         int
		CurrentFirstPlayer  *player.PlayerMock
		CurrentSecondPlayer *player.PlayerMock
		OutResult           bool
	}{
		{
			Description:         "true case",
			InTile:              &tile.Manzu1,
			CurrentTurn:         0,
			CurrentFirstPlayer:  &player.PlayerMock{},
			CurrentSecondPlayer: &player.PlayerMock{ExpectedNakiActions: []*naki.NakiAction{&naki.Pon}},
			OutResult:           true,
		},
		{
			Description:         "false case",
			InTile:              &tile.Manzu1,
			CurrentTurn:         0,
			CurrentFirstPlayer:  &player.PlayerMock{},
			CurrentSecondPlayer: &player.PlayerMock{ExpectedNakiActions: []*naki.NakiAction{}},
			OutResult:           false,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			board := BoardImpl{
				players: []player.Player{c.CurrentFirstPlayer, c.CurrentSecondPlayer},
				turn:    c.CurrentTurn,
			}

			result := board.CanOtherPlayersNaki(c.InTile)
			assert.Equal(t, c.OutResult, result)
		})

	}

}
