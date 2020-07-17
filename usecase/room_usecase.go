package usecase

import (
	"aws-mahjong/game"
	"aws-mahjong/repository"
	"errors"

	socketio "github.com/googollee/go-socket.io"
)

var (
	RoomAlraedyTokenErr = errors.New("room already token")
	RoomNotFound        = errors.New("room is not found")
	RoomReachMaxMember  = errors.New("room already fulled")
)

type RoomUsecase interface {
	Rooms() []*RoomInfo
	CreateRoom(s socketio.Conn, username string, roomName string, roomCapacity int) error
	JoinRoom(s socketio.Conn, username string, roomName string) error
	LeaveRoom(s socketio.Conn, roomName string) error
	LeaveAllRoom(s socketio.Conn) error
}

type RoomUsecaseImpl struct {
	gameRepo repository.GameRepository
	roomRepo *repository.RoomRepository
}

func NewRoomUsecase(gameRepo repository.GameRepository, roomRepo *repository.RoomRepository) RoomUsecase {
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

	for _, roomName := range u.roomRepo.Rooms() {
		g, err := u.gameRepo.Find(roomName)
		if err != nil {
			continue
		}
		rooms = append(rooms, &RoomInfo{
			Name:     roomName,
			Len:      u.roomRepo.RoomLen(roomName),
			Capacity: g.Capacity(),
		})
	}
	return rooms
}

func (u *RoomUsecaseImpl) CreateRoom(s socketio.Conn, username string, roomName string, roomCapacity int) error {

	if u.roomRepo.RoomLen(roomName) != 0 {
		return RoomAlraedyTokenErr
	}

	user := &game.User{ID: s.ID(), Name: username}
	err := u.gameRepo.Add(roomName, game.NewGame(roomCapacity, user))
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
	u.roomRepo.LeaveRoom(s, roomName)
	err := u.gameRepo.Remove(roomName)
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
