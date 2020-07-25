package usecase

import (
	"aws-mahjong/board"
	"aws-mahjong/game"
	"aws-mahjong/repository"
	"aws-mahjong/tile"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDahai(t *testing.T) {
	cases := []struct {
		Description       string
		CurrentRepository *repository.RoomRepositoryMock
		InUserId          string
		InHai             *tile.Tile
		InRoomName        string
		OutError          error
		OutStatus         *board.BoardStatus
	}{
		{
			Description: "valid case",
			CurrentRepository: &repository.RoomRepositoryMock{ExpectedGame: &game.GameMock{
				ExpectedBoardStatus: &board.BoardStatus{},
			}},
			OutStatus: &board.BoardStatus{},
		},
		{
			Description: "invalid case",
			CurrentRepository: &repository.RoomRepositoryMock{ExpectedGame: &game.GameMock{
				ExpectedError: errors.New(""),
			}},
			OutError: errors.New(""),
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			usecase := GameUsecaseImpl{roomRepo: c.CurrentRepository}
			status, err := usecase.Dahai(c.InUserId, c.InRoomName, c.InHai)
			if err != nil && err == c.OutError {
				return
			}
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutStatus, status)
		})
	}

}
