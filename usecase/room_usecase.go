package usecase

import (
	"aws-mahjong/game"
	"aws-mahjong/repository"
	"errors"
)

var (
	RoomAlraedyTakenErr = errors.New("room already taken")
	RoomNotFound        = errors.New("room is not found")
	RoomReachMaxMember  = errors.New("room already fulled")
)

type RoomUsecase interface {
	Rooms() []*RoomStatus
	CreateRoom(userId string, userName string, roomName string, roomCapacity int) (*RoomStatus, error)
	JoinRoom(userId string, userName string, roomName string) (*RoomStatus, error)
	LeaveRoom(userId string, userName string, roomName string) error
}

type RoomUsecaseImpl struct {
	roomRepo repository.RoomRepository
}

func NewRoomUsecase(roomRepo repository.RoomRepository) RoomUsecase {
	return &RoomUsecaseImpl{roomRepo: roomRepo}
}

type RoomStatus struct {
	Name     string
	Len      int
	Capacity int
}

func (u *RoomUsecaseImpl) Rooms() []*RoomStatus {
	rooms := []*RoomStatus{}
	for key, game := range u.roomRepo.Rooms() {
		rooms = append(rooms, &RoomStatus{
			Name:     key,
			Len:      len(game.Users()),
			Capacity: game.Capacity(),
		})

	}
	return rooms
}

func (u *RoomUsecaseImpl) Room(roomName string) (*RoomStatus, error) {

	g, err := u.roomRepo.Find(roomName)
	if err != nil {
		return nil, RoomNotFound
	}
	foundRoom := &RoomStatus{
		Name:     roomName,
		Len:      len(g.Users()),
		Capacity: g.Capacity(),
	}

	return foundRoom, nil
}

func (u *RoomUsecaseImpl) CreateRoom(userId string, userName string, roomName string, roomCapacity int) (*RoomStatus, error) {
	user := &game.User{ID: userId, Name: userName}
	newGame, err := game.NewGame(roomCapacity, user)
	if err != nil {
		return nil, err
	}
	err = u.roomRepo.Add(roomName, newGame)
	if err != nil {
		return nil, err
	}
	return u.Room(roomName)
}

func (u *RoomUsecaseImpl) JoinRoom(userId string, userName string, roomName string) (*RoomStatus, error) {
	user := &game.User{ID: userId, Name: userName}
	err := u.roomRepo.AddUserToRoom(roomName, user)
	if err != nil {
		return nil, err
	}
	return u.Room(roomName)
}

func (u *RoomUsecaseImpl) LeaveRoom(userId string, userName string, roomName string) error {
	user := &game.User{ID: userId, Name: userName}
	return u.roomRepo.RemoveUserFromRoom(roomName, user)
}
