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
		InBoard     *game.Game
		OutError    error
	}{
		{
			Description: "valid case",
			InRoomName:  "beatae",
			InBoard:     &game.Game{},
			OutError:    nil,
		},
		{
			Description: "invalid case",
			InBoard:     nil,
			OutError:    GameIsNil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			repo := GameRepository{games: map[string]*game.Game{}}
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
			repo := GameRepository{games: map[string]*game.Game{"iusto": &game.Game{}}}
			err := repo.Remove(c.InRoomName)
			assert.Equal(t, c.OutError, err)
		})
	}

}

func TestFind(t *testing.T) {
	cases := []struct {
		Description string
		InRoomName  string
		OutError    error
	}{
		{
			Description: "valid case",
			InRoomName:  "quis",
			OutError:    nil,
		},
		{
			Description: "invalid case",
			InRoomName:  "omnis",
			OutError:    GameNotFoundErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			repo := GameRepository{games: map[string]*game.Game{"quis": &game.Game{}}}
			outGame, err := repo.Find(c.InRoomName)
			if err != nil && c.OutError == err {
				return
			}
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, &game.Game{}, outGame)

		})
	}

}
