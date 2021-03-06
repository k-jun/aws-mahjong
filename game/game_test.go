package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUsername(t *testing.T) {

	cases := []struct {
		Description     string
		CurrentCapacity int
		CurrentUsers    []*User
		InUser          *User
		OutError        error
	}{
		{
			Description:     "valid case",
			CurrentCapacity: 2,
			CurrentUsers:    []*User{{}},
			InUser:          &User{ID: "accusantium", Name: "Marianna Gorczany"},
			OutError:        nil,
		},
		{
			Description:     "invalid case",
			CurrentCapacity: 0,
			CurrentUsers:    []*User{{}},
			InUser:          nil,
			OutError:        UserIsEmptyErr,
		},
		{
			Description:     "invalid case",
			CurrentCapacity: 1,
			CurrentUsers:    []*User{{ID: "libero", Name: "Mrs. Ivah Hauck"}},
			InUser:          &User{ID: "accusantium", Name: "Marianna Gorczany"},
			OutError:        GameReachMaxMemberErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			game := &GameImpl{
				capacity: c.CurrentCapacity,
				users:    c.CurrentUsers,
			}
			err := game.AddUser(c.InUser)
			assert.Equal(t, c.OutError, err)
		})
	}
}

func TestRemoveUser(t *testing.T) {
	cases := []struct {
		Description     string
		CurrentCapacity int
		CurrentUsers    []*User
		InUser          *User
		OutError        error
	}{
		{
			Description:     "valid case",
			CurrentCapacity: 2,
			CurrentUsers:    []*User{{ID: "123"}},
			InUser:          &User{ID: "123"},
			OutError:        nil,
		},
		{
			Description:     "invalid case, user not found",
			CurrentCapacity: 2,
			CurrentUsers:    []*User{{ID: "123"}},
			InUser:          &User{ID: "122"},
			OutError:        UserNotFound,
		},
		{
			Description:     "invalid case, user nil",
			CurrentCapacity: 2,
			CurrentUsers:    []*User{{ID: "123"}},
			InUser:          nil,
			OutError:        UserIsEmptyErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {

			game := GameImpl{
				users: c.CurrentUsers,
			}

			err := game.RemoveUser(c.InUser)
			assert.Equal(t, c.OutError, err)
		})
	}

}

func TestGameStart(t *testing.T) {
	cases := []struct {
		Description     string
		CurrentCapacity int
		CurrentUsers    []*User
		OutError        error
	}{
		{
			Description:     "valid case",
			CurrentCapacity: 4,
			CurrentUsers: []*User{
				{ID: "1", Name: "Angela Hudson"},
				{ID: "2", Name: "Angela Hudson"},
				{ID: "3", Name: "Angela Hudson"},
				{ID: "4", Name: "Angela Hudson"},
			},
			OutError: nil,
		},
		{
			Description:     "invalid case",
			CurrentCapacity: 3,
			CurrentUsers: []*User{
				{ID: "1", Name: "Angela Hudson"},
				{ID: "2", Name: "Angela Hudson"},
				{ID: "3", Name: "Angela Hudson"},
				{ID: "4", Name: "Angela Hudson"},
			},
			OutError: GameMemberInvalid,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			game := GameImpl{
				users:    c.CurrentUsers,
				capacity: c.CurrentCapacity,
			}

			err := game.GameStart()
			assert.Equal(t, c.OutError, err)
		})

	}
}
