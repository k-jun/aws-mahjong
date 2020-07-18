package repository

import (
	"aws-mahjong/game"
	"errors"
)

var (
	GameNotFoundErr = errors.New("game not found")
	GameIsNil       = errors.New("game is nil")
	RoomNameIsEmpry = errors.New("roomName is empty")
)

type GameRepository interface {
	Add(roomName string, inGame game.Game) error
	Remove(roomName string) error
	Find(roomName string) (game.Game, error)
}

type GameRepositoryImpl struct {
	games map[string]game.Game
}

func NewGameRepository() GameRepository {
	return &GameRepositoryImpl{
		games: map[string]game.Game{},
	}
}

func (r *GameRepositoryImpl) Add(roomName string, inGame game.Game) error {
	if roomName == "" {
		return RoomNameIsEmpry
	}

	if inGame == nil {
		return GameIsNil
	}
	r.games[roomName] = inGame
	return nil
}

func (r *GameRepositoryImpl) Remove(roomName string) error {
	if r.games[roomName] == nil {
		return GameNotFoundErr
	}
	delete(r.games, roomName)
	return nil
}

func (r *GameRepositoryImpl) Find(roomName string) (game.Game, error) {
	if r.games[roomName] == nil {
		return nil, GameNotFoundErr
	}
	return r.games[roomName], nil
}
