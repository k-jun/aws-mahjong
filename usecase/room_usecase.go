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

type RoomUsecase struct {
	gameRepo repository.GameRepository
	roomRepo repository.RoomRepository
}

func NewRoomUsecase(gameRepo repository.GameRepository, roomRepo repository.RoomRepository) *RoomUsecase {
	return &RoomUsecase{
		gameRepo: gameRepo,
		roomRepo: roomRepo,
	}
}

func (u *RoomUsecase) CreateRoom(s socketio.Conn, username string, roomName string, roomCapacity int) error {

	if u.roomRepo.RoomLen(roomName) != 0 {
		return RoomAlraedyTokenErr
	}
	u.roomRepo.JoinRoom(s, roomName)
	err := u.gameRepo.Add(roomName, game.NewGame(roomCapacity, username))
	return err
}

func (u *RoomUsecase) JoinRoom(s socketio.Conn, username string, roomName string) error {

	if u.roomRepo.RoomLen(roomName) == 0 {
		return RoomNotFound
	}

	game, err := u.gameRepo.Find(roomName)
	if err != nil {
		return RoomNotFound
	}
	if u.roomRepo.RoomLen(roomName) >= game.Capacity() {
		return RoomReachMaxMember
	}

	u.roomRepo.JoinRoom(s, roomName)
	return game.AddUser(username)
}

func (u *RoomUsecase) LeaveRoom(s socketio.Conn, roomName string) error {
	u.roomRepo.LeaveRoom(s, roomName)
	err := u.gameRepo.Remove(roomName)
	return err
}

func (u *RoomUsecase) LeaveAllRoom(s socketio.Conn) error {
	for _, roomName := range s.Rooms() {
		u.roomRepo.LeaveRoom(s, roomName)
		err := u.gameRepo.Remove(roomName)
		if err != nil {
			return err
		}
	}
	return nil
}
