package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUsername(t *testing.T) {

	cases := []struct {
		Description      string
		CurrentCapacity  int
		CurrentUsernames []string
		InUsername       string
		OutError         error
	}{
		{
			Description:      "valid case",
			CurrentCapacity:  2,
			CurrentUsernames: []string{},
			InUsername:       "Dr. Zoie Funk MD",
			OutError:         nil,
		},
		{
			Description:      "invalid case",
			InUsername:       "",
			CurrentCapacity:  0,
			CurrentUsernames: []string{},
			OutError:         UsernameIsEmpty,
		},
		{
			Description:      "invalid case",
			InUsername:       "Casandra Maggio",
			CurrentCapacity:  1,
			CurrentUsernames: []string{"Allen Hane"},
			OutError:         GameReachMaxMemberErr,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			game := &GameImpl{
				capacity:  c.CurrentCapacity,
				usernames: c.CurrentUsernames,
			}
			err := game.AddUsername(c.InUsername)
			assert.Equal(t, c.OutError, err)
		})
	}
}
