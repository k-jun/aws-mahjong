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
			CurrentUsers:    []*User{&User{}},
			InUser:          &User{ID: "accusantium", Name: "Marianna Gorczany"},
			OutError:        nil,
		},
		{
			Description:     "invalid case",
			CurrentCapacity: 0,
			CurrentUsers:    []*User{&User{}},
			InUser:          nil,
			OutError:        UserIsEmptyErr,
		},
		{
			Description:     "invalid case",
			CurrentCapacity: 1,
			CurrentUsers:    []*User{&User{ID: "libero", Name: "Mrs. Ivah Hauck"}},
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
