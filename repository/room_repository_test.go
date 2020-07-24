package repository

import (
	"aws-mahjong/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		Description string
		InRoomName  string
		InBoard     game.Game
		OutError    error
	}{
		{
			Description: "valid case",
			InRoomName:  "beatae",
			InBoard:     &game.GameMock{},
			OutError:    nil,
		},
		{
			Description: "invalid case",
			InRoomName:  "",
			InBoard:     &game.GameMock{},
			OutError:    RoomNameIsEmpry,
		},
		{
			Description: "invalid case",
			InRoomName:  "libero",
			InBoard:     nil,
			OutError:    GameIsNil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			repo := RoomRepositoryImpl{rooms: map[string]game.Game{}}
			err := repo.Add(c.InRoomName, c.InBoard)
			assert.Equal(t, c.OutError, err)
		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		Description string
		InRoomName  string
		OutError    error
	}{
		{
			Description: "valid case",
			InRoomName:  "iusto",
			OutError:    nil,
		},
		{
			Description: "invalid case",
			InRoomName:  "not_exist_room_name",
			OutError:    GameNotFoundErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			repo := RoomRepositoryImpl{rooms: map[string]game.Game{"iusto": &game.GameMock{}}}
			err := repo.Remove(c.InRoomName)
			assert.Equal(t, c.OutError, err)
		})
	}

}

func TestFind(t *testing.T) {
	cases := []struct {
		Description  string
		CurrentGames map[string]game.Game
		InRoomName   string
		OutError     error
		OutGame      game.Game
	}{
		{
			Description:  "valid case",
			CurrentGames: map[string]game.Game{"quis": &game.GameMock{}},
			InRoomName:   "quis",
			OutError:     nil,
			OutGame:      &game.GameMock{},
		},
		{
			Description:  "invalid case",
			CurrentGames: map[string]game.Game{"quis": &game.GameMock{}},
			InRoomName:   "omnis",
			OutError:     GameNotFoundErr,
			OutGame:      nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {

			repo := RoomRepositoryImpl{rooms: c.CurrentGames}
			outGame, err := repo.Find(c.InRoomName)
			if err != nil && c.OutError == err {
				return
			}
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutGame, outGame)
		})
	}
}

func TestAddUserToRoom(t *testing.T) {
	cases := []struct {
		Description  string
		CurrentGames map[string]game.Game
		InUser       *game.User
		InRoomName   string
		OutError     error
		OutUsers     []*game.User
	}{
		{
			Description:  "valid case",
			CurrentGames: map[string]game.Game{"quis": &game.GameMock{}},
			InUser:       &game.User{},
			InRoomName:   "quis",
			OutError:     nil,
			OutUsers:     []*game.User{&game.User{}},
		},
		{
			Description:  "invalid case",
			CurrentGames: map[string]game.Game{"quis": &game.GameMock{ExpectedUsers: []*game.User{}}},
			InUser:       &game.User{},
			InRoomName:   "quiz",
			OutError:     GameNotFoundErr,
			OutUsers:     []*game.User{},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			repo := RoomRepositoryImpl{rooms: c.CurrentGames}
			err := repo.AddUserToRoom(c.InRoomName, c.InUser)
			if err != nil && err == c.OutError {
				return
			}
			assert.Equal(t, c.OutError, err)
			assert.Equal(t, c.OutUsers, repo.rooms[c.InRoomName].Users())
		})
	}
}

func TestRemoveUserFromRoom(t *testing.T) {
	cases := []struct {
		Description   string
		CurrentGames  map[string]game.Game
		InUser        *game.User
		InRoomName    string
		OutError      error
		OutRoomDelete bool
		OutUsers      []*game.User
	}{
		{
			Description: "valid case",
			CurrentGames: map[string]game.Game{
				"Douglas.Maribel": &game.GameMock{
					ExpectedUsers: []*game.User{
						&game.User{
							"f24e0aab-337c-383e-82cf-2c32a930d73e",
							"Lonnie Mertz",
						},
						&game.User{
							"a2d6bb27-029d-37e2-98d8-ec6df5b14444",
							"Lonnie Mertz",
						},
					},
				},
			},
			InRoomName:    "Douglas.Maribel",
			InUser:        &game.User{ID: "f24e0aab-337c-383e-82cf-2c32a930d73e"},
			OutError:      nil,
			OutRoomDelete: false,
			OutUsers: []*game.User{
				&game.User{
					"a2d6bb27-029d-37e2-98d8-ec6df5b14444",
					"Lonnie Mertz",
				},
			},
		},
		{
			Description: "valid case, delete room",
			CurrentGames: map[string]game.Game{
				"Douglas.Maribel": &game.GameMock{
					ExpectedUsers: []*game.User{
						&game.User{
							"f24e0aab-337c-383e-82cf-2c32a930d73e",
							"Lonnie Mertz",
						},
					},
				},
			},
			InRoomName:    "Douglas.Maribel",
			InUser:        &game.User{ID: "f24e0aab-337c-383e-82cf-2c32a930d73e"},
			OutError:      nil,
			OutRoomDelete: true,
			OutUsers:      []*game.User{},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			repo := RoomRepositoryImpl{rooms: c.CurrentGames}
			err := repo.RemoveUserFromRoom(c.InRoomName, c.InUser)
			if err != nil && err == c.OutError {
				return
			}
			assert.Equal(t, c.OutError, err)
			if c.OutRoomDelete {
				assert.Equal(t, nil, repo.rooms[c.InRoomName])
				return
			}
			assert.Equal(t, c.OutUsers, repo.rooms[c.InRoomName].Users())
		})
	}
}
