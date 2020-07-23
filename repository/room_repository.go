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

type RoomRepository interface {
	Add(roomName string, inGame game.Game) error
	Remove(roomName string) error
	Find(roomName string) (game.Game, error)
	Rooms() map[string]game.Game
}

type RoomRepositoryImpl struct {
	rooms map[string]game.Game
}

func NewGameRepository() RoomRepository {
	return &RoomRepositoryImpl{
		rooms: map[string]game.Game{},
	}
}

func (r *RoomRepositoryImpl) Add(roomName string, inGame game.Game) error {
	if roomName == "" {
		return RoomNameIsEmpry
	}

	if inGame == nil {
		return GameIsNil
	}
	r.rooms[roomName] = inGame
	return nil
}

func (r *RoomRepositoryImpl) Remove(roomName string) error {
	if r.rooms[roomName] == nil {
		return GameNotFoundErr
	}
	delete(r.rooms, roomName)
	return nil
}

func (r *RoomRepositoryImpl) Find(roomName string) (game.Game, error) {
	if r.rooms[roomName] == nil {
		return nil, GameNotFoundErr
	}
	return r.rooms[roomName], nil
}

func (r *RoomRepositoryImpl) Rooms() map[string]game.Game {
	return r.rooms
}
