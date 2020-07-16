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
			OutError:    BoardIsNil,
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
			OutError:    BoardNotFoundErr,
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
