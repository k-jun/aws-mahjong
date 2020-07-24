// +build integration

package server

import (
	"aws-mahjong/repository"
	"aws-mahjong/testutil"
	"aws-mahjong/usecase"
	"fmt"
	"net/http"
	"testing"
)

func TestRooms(t *testing.T) {

	cases := []struct {
		Description    string
		CurrentUsecase usecase.RoomUsecase
	}{
		{
			Description: "valid case, single room",
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			roomRepo := repository.NewRoomRepository()
			roomUsecase := usecase.NewRoomUsecase(roomRepo)
			router := makeServer(roomUsecase)
			response := testutil.MakeRequest(router, http.MethodGet, "/rooms", "")
			fmt.Println(response.Body)
		})
	}
}
