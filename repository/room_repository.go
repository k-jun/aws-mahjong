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
	Find(roomName string) (*game.Game, error)
	Rooms() []*RoomStatus
}

type RoomRepositoryImpl struct {
	games map[string]game.Game
}

func NewGameRepository() RoomRepository {
	return &RoomRepositoryImpl{
		games: map[string]game.Game{},
	}
}

func (r *RoomRepositoryImpl) Add(roomName string, inGame game.Game) error {
	if roomName == "" {
		return RoomNameIsEmpry
	}

	if inGame == nil {
		return GameIsNil
	}
	r.games[roomName] = inGame
	return nil
}

func (r *RoomRepositoryImpl) Remove(roomName string) error {
	if r.games[roomName] == nil {
		return GameNotFoundErr
	}
	delete(r.games, roomName)
	return nil
}

func (r *RoomRepositoryImpl) Find(roomName string) (*game.Game, error) {
	if r.games[roomName] == nil {
		return nil, GameNotFoundErr
	}
	g := r.games[roomName]
	return &g, nil
}

type RoomStatus struct {
	Name     string
	Len      int
	Capacity int
}

func (r *RoomRepositoryImpl) Rooms() []*RoomStatus {
	rooms := []*RoomStatus{}
	for key, game := range r.games {
		rooms = append(rooms, &RoomStatus{
			Name:     key,
			Len:      len(game.Users()),
			Capacity: game.Capacity(),
		})
	}

	return rooms
}
