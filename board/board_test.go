package board

import (
	"aws-mahjong/deck"
	"aws-mahjong/hand"
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
			CurrentFirstPlayer:  &player.PlayerMock{ExpectedTile: &tile.Chun},
			CurrentSecondPlayer: &player.PlayerMock{ExpectedNakiActions: []*naki.NakiAction{}},
			CurrentTurn:         0,
			InTile:              &tile.Chun,
			OutError:            nil,
			OutTurn:             1,
		},
		// {
		// 	Description:         "valid case",
		// 	CurrentNakiTile:     nil,
		// 	CurrentFirstPlayer:  &player.PlayerMock{ExpectedTile: &tile.Chun},
		// 	CurrentSecondPlayer: &player.PlayerMock{ExpectedNakiActions: []*naki.NakiAction{&naki.Pon}},
		// 	CurrentTurn:         0,
		// 	InTile:              &tile.Chun,
		// 	OutError:            nil,
		// 	OutTurn:             0,
		// },
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

func TestStart(t *testing.T) {
	cases := []struct {
		Description  string
		CurrentHands []hand.Hand
		CurrentDeck  deck.Deck
		OutError     error
	}{
		{
			Description:  "valid case",
			CurrentHands: []hand.Hand{&hand.HandMock{ExpectedTiles: []*tile.Tile{}, ExpectedError: nil}},
			CurrentDeck:  deck.NewDeck(),
			OutError:     nil,
		},
		{
			Description:  "invalid case",
			CurrentHands: []hand.Hand{&hand.HandMock{ExpectedTiles: []*tile.Tile{&tile.Chun}, ExpectedError: nil}},
			CurrentDeck:  deck.NewDeck(),
			OutError:     GameAlreadyStarted,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			players := []player.Player{}
			for _, h := range c.CurrentHands {
				players = append(players, &player.PlayerMock{ExpectedHand: h})
			}
			board := BoardImpl{players: players, deck: c.CurrentDeck}
			err := board.Start()
			assert.Equal(t, c.OutError, err)
		})
	}
}

func TestStatus(t *testing.T) {

	status1 := &player.PlayerStatus{ID: "123"}
	status2 := &player.PlayerStatus{ID: "124"}
	status3 := &player.PlayerStatus{ID: "125"}
	status4 := &player.PlayerStatus{ID: "126"}

	cases := []struct {
		Description    string
		CurrentBakaze  *tile.Tile
		CurrentDeck    deck.Deck
		CurrentPlayers []player.Player
		CurrentTurn    int
		CurrentOya     int
		InPlayerID     string
		OutJicha       int
		OutOya         int
		OutTurn        int
		OutPlayers     []*player.PlayerStatus
	}{
		{
			Description:   "valid case",
			CurrentBakaze: &tile.East,
			CurrentDeck:   deck.NewDeck(),
			CurrentPlayers: []player.Player{
				&player.PlayerMock{ExpectedStatus: status1},
				&player.PlayerMock{ExpectedStatus: status2},
				&player.PlayerMock{ExpectedStatus: status3},
				&player.PlayerMock{ExpectedStatus: status4},
			},
			CurrentOya:  0,
			CurrentTurn: 0,
			InPlayerID:  status1.ID,
			OutJicha:    0,
			OutOya:      0,
			OutTurn:     0,
			OutPlayers:  []*player.PlayerStatus{status1, status2, status3, status4},
		},
		{
			Description:   "valid case, shimocha",
			CurrentBakaze: &tile.East,
			CurrentDeck:   deck.NewDeck(),
			CurrentPlayers: []player.Player{
				&player.PlayerMock{ExpectedStatus: status1},
				&player.PlayerMock{ExpectedStatus: status2},
				&player.PlayerMock{ExpectedStatus: status3},
			},
			CurrentOya:  2,
			CurrentTurn: 3,
			InPlayerID:  status2.ID,
			OutJicha:    1,
			OutOya:      2,
			OutTurn:     3,
			OutPlayers:  []*player.PlayerStatus{status1, status2, status3},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			board := BoardImpl{
				bakaze:  c.CurrentBakaze,
				deck:    c.CurrentDeck,
				players: c.CurrentPlayers,
				turn:    c.CurrentTurn,
				oya:     c.CurrentOya,
			}

			status := board.Status(c.InPlayerID)
			assert.Equal(t, c.CurrentBakaze.Name(), status.Bakaze)
			assert.Equal(t, c.CurrentDeck.Count(), status.DeckLen)
			assert.Equal(t, c.OutJicha, status.Jicha)
			assert.Equal(t, c.OutOya, status.Oya)
			assert.Equal(t, c.OutTurn, status.Turn)
			assert.Equal(t, c.OutPlayers, status.Players)
		})
	}
}

func TestIsTurnPlayer(t *testing.T) {
	cases := []struct {
		Description    string
		CurrentTurn    int
		CurrentPlayers []player.Player
		InPlayerID     string
		OutResult      bool
	}{
		{
			Description: "valid case",
			CurrentTurn: 0,
			CurrentPlayers: []player.Player{
				&player.PlayerMock{ExpectedID: "951b4115-cefe-336f-a65d-849e2c84169e"},
			},
			InPlayerID: "951b4115-cefe-336f-a65d-849e2c84169e",
			OutResult:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			board := BoardImpl{players: c.CurrentPlayers, turn: c.CurrentTurn}
			assert.Equal(t, c.OutResult, board.IsTurnPlayer(c.InPlayerID))
		})
	}
}
