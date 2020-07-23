package repository

import (
	"aws-mahjong/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		Description string
		InRoomName  string
		InBoard     game.Game
		OutError    error
	}{
		{
			Description: "valid case",
			InRoomName:  "beatae",
			InBoard:     &game.GameMock{},
			OutError:    nil,
		},
		{
			Description: "invalid case",
			InRoomName:  "",
			InBoard:     &game.GameMock{},
			OutError:    RoomNameIsEmpry,
		},
		{
			Description: "invalid case",
			InRoomName:  "libero",
			InBoard:     nil,
			OutError:    GameIsNil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			repo := RoomRepositoryImpl{rooms: map[string]game.Game{}}
			err := repo.Add(c.InRoomName, c.InBoard)
			assert.Equal(t, c.OutError, err)
		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		Description string
		InRoomName  string
		OutError    error
	}{
		{
			Description: "valid case",
			InRoomName:  "iusto",
			OutError:    nil,
		},
		{
			Description: "invalid case",
			InRoomName:  "not_exist_room_name",
			OutError:    GameNotFoundErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			repo := RoomRepositoryImpl{rooms: map[string]game.Game{"iusto": &game.GameMock{}}}
			err := repo.Remove(c.InRoomName)
			assert.Equal(t, c.OutError, err)
		})
	}

}

func TestFind(t *testing.T) {
	cases := []struct {
		Description  string
		CurrentGames map[string]game.Game
		InRoomName   string
		OutError     error
		OutGame      game.Game
	}{
		{
			Description:  "valid case",
			CurrentGames: map[string]game.Game{"quis": &game.GameMock{}},
			InRoomName:   "quis",
			OutError:     nil,
			OutGame:      &game.GameMock{},
		},
		{
			Description:  "invalid case",
			CurrentGames: map[string]game.Game{"quis": &game.GameMock{}},
			InRoomName:   "omnis",
			OutError:     GameNotFoundErr,
			OutGame:      nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {

			repo := RoomRepositoryImpl{rooms: c.CurrentGames}
			outGame, err := repo.Find(c.InRoomName)
			if err != nil && c.OutError == err {
				return
			}
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutGame, *outGame)
		})
	}
}
