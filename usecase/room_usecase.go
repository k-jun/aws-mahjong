package usecase

import (
	"aws-mahjong/game"
	"aws-mahjong/repository"
	"errors"
	"sort"

	socketio "github.com/googollee/go-socket.io"
)

var (
	RoomAlraedyTakenErr = errors.New("room already taken")
	RoomNotFound        = errors.New("room is not found")
	RoomReachMaxMember  = errors.New("room already fulled")
)

type RoomUsecase interface {
	Rooms() []*RoomInfo
	Room(roomName string) (*RoomInfo, error)
	CreateRoom(s socketio.Conn, username string, roomName string, roomCapacity int) error
	JoinRoom(s socketio.Conn, username string, roomName string) error
	LeaveRoom(s socketio.Conn, roomName string) error
	LeaveAllRoom(s socketio.Conn) error
}

type RoomUsecaseImpl struct {
	gameRepo repository.GameRepository
	roomRepo *repository.RoomRepository
}

func NewRoomUsecase(roomRepo *repository.RoomRepository, gameRepo repository.GameRepository) RoomUsecase {
	return &RoomUsecaseImpl{
		gameRepo: gameRepo,
		roomRepo: roomRepo,
	}
}

type RoomInfo struct {
	Name     string
	Len      int
	Capacity int
}

func (u *RoomUsecaseImpl) Rooms() []*RoomInfo {
	rooms := []*RoomInfo{}

	roomNames := u.roomRepo.Rooms()

	sort.Slice(roomNames, func(i int, j int) bool { return roomNames[i] < roomNames[j] })

	for _, roomName := range roomNames {
		r, err := u.Room(roomName)
		if err != nil {
			continue
		}
		rooms = append(rooms, r)
	}
	return rooms
}

func (u *RoomUsecaseImpl) Room(roomName string) (*RoomInfo, error) {

	g, err := u.gameRepo.Find(roomName)
	if err != nil {
		return nil, RoomNotFound
	}
	foundRoom := &RoomInfo{
		Name:     roomName,
		Len:      u.roomRepo.RoomLen(roomName),
		Capacity: g.Capacity(),
	}

	return foundRoom, nil
}

func (u *RoomUsecaseImpl) CreateRoom(s socketio.Conn, username string, roomName string, roomCapacity int) error {

	if u.roomRepo.RoomLen(roomName) != 0 {
		return RoomAlraedyTakenErr
	}

	user := &game.User{ID: s.ID(), Name: username}
	newGame, err := game.NewGame(roomCapacity, user)
	if err != nil {
		return err
	}
	err = u.gameRepo.Add(roomName, newGame)
	if err != nil {
		return err
	}
	u.roomRepo.JoinRoom(s, roomName)
	return nil
}

func (u *RoomUsecaseImpl) JoinRoom(s socketio.Conn, username string, roomName string) error {

	if u.roomRepo.RoomLen(roomName) == 0 {
		return RoomNotFound
	}

	roomGame, err := u.gameRepo.Find(roomName)
	if err != nil {
		return RoomNotFound
	}

	user := &game.User{ID: s.ID(), Name: username}
	err = roomGame.AddUser(user)
	if err != nil {
		return err
	}

	// if u.roomRepo.RoomLen(roomName) >= roomGame.Capacity() {
	// 	return RoomReachMaxMember
	// }
	u.roomRepo.JoinRoom(s, roomName)
	return nil
}

func (u *RoomUsecaseImpl) LeaveRoom(s socketio.Conn, roomName string) error {
	err := u.gameRepo.Remove(roomName)
	if err != nil {
		return err
	}
	u.roomRepo.LeaveRoom(s, roomName)
	return err
}

func (u *RoomUsecaseImpl) LeaveAllRoom(s socketio.Conn) error {
	for _, roomName := range s.Rooms() {
		u.roomRepo.LeaveRoom(s, roomName)
		err := u.gameRepo.Remove(roomName)
		if err != nil {
			return err
		}
	}
	return nil
}
