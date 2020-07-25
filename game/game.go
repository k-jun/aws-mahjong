package game

import (
	"aws-mahjong/board"
	"aws-mahjong/deck"
	"aws-mahjong/hand"
	"aws-mahjong/kawa"
	"aws-mahjong/naki"
	"aws-mahjong/player"
	"aws-mahjong/tile"
	"errors"
)

var (
	capacityMax = 4
	capacityMin = 3
)

var (
	UserIsEmptyErr        = errors.New("user can't be empty")
	UserNotFound          = errors.New("user is not found")
	GameReachMaxMemberErr = errors.New("game already fulled")
	GameCapacityInvalid   = errors.New("game capacity is invalid")
	GameMemberInvalid     = errors.New("game member is invalid")
	GameAlreadyStartedErr = errors.New("game already started")
	GameNotStartedErr     = errors.New("game not started")
	NotYourTurnErr        = errors.New("this is not your turn")
)

type User struct {
	ID   string
	Name string
}

type Game interface {
	AddUser(user *User) error
	RemoveUser(user *User) error
	Users() []*User
	Capacity() int
	GameStart() error
	Dahai(userId string, hai *tile.Tile) (*board.BoardStatus, error)
}

type GameImpl struct {
	capacity int
	users    []*User
	board    board.Board
}

func NewGame(capacity int, user *User) (Game, error) {
	if capacity > capacityMax || capacity < capacityMin {
		return nil, GameCapacityInvalid
	}

	newGame := &GameImpl{
		capacity: capacity,
		users:    []*User{user},
		board:    nil,
	}

	return newGame, nil
}

func (g *GameImpl) Capacity() int {
	return g.capacity
}

func (g *GameImpl) Users() []*User {
	return g.users
}

func (g *GameImpl) RemoveUser(user *User) error {
	if user == nil {
		return UserIsEmptyErr
	}

	for idx, roomUser := range g.users {
		if roomUser.ID == user.ID {
			g.users[idx], g.users[len(g.users)-1] = g.users[len(g.users)-1], g.users[idx]
			g.users = g.users[:len(g.users)-1]
			return nil
		}
	}
	return UserNotFound
}

func (g *GameImpl) AddUser(user *User) error {
	if user == nil {
		return UserIsEmptyErr
	}
	if len(g.users) >= g.capacity {
		return GameReachMaxMemberErr
	}
	g.users = append(g.users, user)
	return nil
}

func (g *GameImpl) GameStart() error {

	if len(g.users) != g.capacity {
		return GameMemberInvalid
	}

	newDeck := deck.NewDeck()
	players := []player.Player{}

	for idx, user := range g.users {
		newHand := hand.NewHand()
		newKawa := kawa.NewKawa()
		newNaki := naki.NewNaki()
		players = append(players, player.NewPlayer(
			user.ID,
			user.Name,
			tile.Bakazes[0],
			tile.Zikazes[idx],
			newDeck,
			newHand,
			newKawa,
			newNaki,
		))
	}
	newBoard := board.NewBoard(newDeck, players)
	g.board = newBoard
	return g.board.Start()
}

func (g *GameImpl) Dahai(userId string, hai *tile.Tile) (*board.BoardStatus, error) {
	if g.board == nil {
		return nil, GameNotStartedErr
	}
	if !g.board.IsTurnPlayer(userId) {
		return nil, NotYourTurnErr
	}
	err := g.board.TurnPlayerDahai(userId, hai)
	if err != nil {
		return nil, err
	}
	return g.board.Status(userId), nil
}
