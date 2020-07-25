package server

import (
	"aws-mahjong/board"
	"aws-mahjong/testutil"
	"aws-mahjong/usecase"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDahai(t *testing.T) {
	cases := []struct {
		Description    string
		CurrentUsecase usecase.GameUsecase
		InRoomName     string
		InBody         string
		OutCode        int
		OutStatus      board.BoardStatus
	}{
		{
			Description: "valid case",
			CurrentUsecase: &usecase.GameUsecaseMock{
				ExpectedBoardStatus: &board.BoardStatus{
					Bakaze:  "east",
					DeckLen: 98,
				},
			},
			InRoomName: "bProsacco",
			InBody:     `{"user_id": "67664e29-6983-31e2-9469-29668841baa5", "hai": "haku"}`,
			OutCode:    200,
			OutStatus: board.BoardStatus{
				Bakaze:  "east",
				DeckLen: 98,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			router := makeServer(nil, c.CurrentUsecase)
			res := testutil.MakeRequest(router, http.MethodPost, "/rooms/"+c.InRoomName+"/dahai", c.InBody)
			if res.Code != 200 && c.OutCode == res.Code {
				return
			}
			assert.Equal(t, c.OutCode, res.Code)
			resBody := board.BoardStatus{}
			err := json.Unmarshal(res.Body.Bytes(), &resBody)
			assert.NoError(t, err)
			assert.Equal(t, c.OutStatus, resBody)
		})
	}
}
