package usecase

import (
	"aws-mahjong/game"
	"aws-mahjong/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRooms(t *testing.T) {
	cases := []struct {
		Description  string
		CurrentRooms map[string]game.Game
		OutRooms     []*RoomStatus
	}{
		{
			Description: "valid case",
			CurrentRooms: map[string]game.Game{
				"Ted.Koelpin": &game.GameMock{
					ExpectedUsers: []*game.User{
						&game.User{},
					},
					ExpectedCapacity: 3,
				},
			},
			OutRooms: []*RoomStatus{
				&RoomStatus{Name: "Ted.Koelpin", Len: 1, Capacity: 3},
			},
		},
		{
			Description:  "valid case, no-room",
			CurrentRooms: map[string]game.Game{},
			OutRooms:     []*RoomStatus{},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			usecase := RoomUsecaseImpl{roomRepo: &repository.RoomRepositoryMock{ExpectedRooms: c.CurrentRooms}}
			assert.Equal(t, c.OutRooms, usecase.Rooms())
		})
	}
}

func TestCreateRoom(t *testing.T) {
	cases := []struct {
		Description    string
		CurrentRepo    repository.RoomRepository
		InUserId       string
		InUserName     string
		InRoomName     string
		InRoomCapacity int
		OutError       error
		OutRoomStatus  *RoomStatus
	}{
		{
			Description: "valid case",
			CurrentRepo: &repository.RoomRepositoryMock{
				ExpectedGame: &game.GameMock{
					ExpectedUsers: []*game.User{
						&game.User{},
					},
					ExpectedCapacity: 4,
				},
			},
			InUserId:       "327158ad-ad00-3990-b594-874d958d3675",
			InUserName:     "Mrs. Lorna Schmidt",
			InRoomName:     "Ledner.Wilhelmine",
			InRoomCapacity: 4,
			OutError:       nil,
			OutRoomStatus: &RoomStatus{
				Name:     "Ledner.Wilhelmine",
				Len:      1,
				Capacity: 4,
			},
		},
		{
			Description:    "invalid case",
			CurrentRepo:    &repository.RoomRepositoryMock{},
			InUserId:       "327158ad-ad00-3990-b594-874d958d3675",
			InUserName:     "Mrs. Lorna Schmidt",
			InRoomName:     "Ledner.Wilhelmine",
			InRoomCapacity: 2,
			OutError:       game.GameCapacityInvalid,
			OutRoomStatus:  nil,
		},
		{
			Description: "invalid case",
			CurrentRepo: &repository.RoomRepositoryMock{
				ExpectedError: errors.New(""),
			},
			InUserId:       "327158ad-ad00-3990-b594-874d958d3675",
			InUserName:     "Mrs. Lorna Schmidt",
			InRoomName:     "Ledner.Wilhelmine",
			InRoomCapacity: 3,
			OutError:       errors.New(""),
			OutRoomStatus:  nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			usecase := RoomUsecaseImpl{roomRepo: c.CurrentRepo}
			status, err := usecase.CreateRoom(c.InUserId, c.InUserName, c.InRoomName, c.InRoomCapacity)
			if err != nil && err == c.OutError {
				return
			}
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutRoomStatus, status)
		})
	}
}

func TestJoinRoom(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentRepo   repository.RoomRepository
		InUserId      string
		InUserName    string
		InRoomName    string
		OutError      error
		OutRoomStatus *RoomStatus
	}{
		{
			Description: "valid case",
			CurrentRepo: &repository.RoomRepositoryMock{
				ExpectedGame: &game.GameMock{
					ExpectedUsers: []*game.User{
						&game.User{},
						&game.User{},
						&game.User{},
					},
					ExpectedCapacity: 4,
				},
			},
			InUserId:   "c0c6f73f-8cc3-3226-90c3-074c61a9568c",
			InUserName: "Demarcus Schmitt",
			InRoomName: "Julio69",
			OutError:   nil,
			OutRoomStatus: &RoomStatus{
				Name:     "Julio69",
				Len:      3,
				Capacity: 4,
			},
		},
		{
			Description: "invalid case",
			CurrentRepo: &repository.RoomRepositoryMock{
				ExpectedError: errors.New(""),
			},
			InUserId:      "c0c6f73f-8cc3-3226-90c3-074c61a9568c",
			InUserName:    "Demarcus Schmitt",
			InRoomName:    "Julio69",
			OutError:      errors.New(""),
			OutRoomStatus: nil,
		},
	}

	for _, c := range cases {
		usecase := RoomUsecaseImpl{roomRepo: c.CurrentRepo}
		status, err := usecase.JoinRoom(c.InUserId, c.InUserName, c.InRoomName)
		if err != nil && err == c.OutError {
			return
		}
		assert.Equal(t, c.OutError, err)
		assert.Equal(t, c.OutRoomStatus, status)
	}

}

func TestLeaveRoom(t *testing.T) {
	cases := []struct {
		Description string
		CurrentRepo repository.RoomRepository
		InUserId    string
		InUserName  string
		InRoomName  string
		OutError    error
	}{
		{
			Description: "valid case",
			CurrentRepo: &repository.RoomRepositoryMock{},
			InUserId:    "5a45a8a0-8b76-3734-bceb-63693110f3d8",
			InUserName:  "Astrid Eichmann",
			InRoomName:  "Lafayette77",
			OutError:    nil,
		},
		{
			Description: "invalid case",
			CurrentRepo: &repository.RoomRepositoryMock{ExpectedError: errors.New("")},
			InUserId:    "5a45a8a0-8b76-3734-bceb-63693110f3d8",
			InUserName:  "Astrid Eichmann",
			InRoomName:  "Lafayette77",
			OutError:    errors.New(""),
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			usecase := RoomUsecaseImpl{roomRepo: c.CurrentRepo}
			err := usecase.LeaveRoom(c.InUserId, c.InUserName, c.InRoomName)
			assert.Equal(t, c.OutError, err)
		})
	}

}
