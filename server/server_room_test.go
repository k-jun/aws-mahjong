package server

import (
	"aws-mahjong/server/view"
	"aws-mahjong/testutil"
	"aws-mahjong/usecase"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRooms(t *testing.T) {

	cases := []struct {
		Description    string
		CurrentUsecase usecase.RoomUsecase
		OutCode        int
		OutStatuses    []view.RoomResponse
	}{
		{
			Description:    "valid case, single room",
			CurrentUsecase: &usecase.RoomUsecaseMock{ExpectedRoomStatuses: []*usecase.RoomStatus{&usecase.RoomStatus{Len: 2, Capacity: 3, Name: "mLarkin"}}},
			OutCode:        200,
			OutStatuses:    []view.RoomResponse{{RoomName: "mLarkin", RoomCapacity: 3, RoomLen: 2}},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			router := makeServer(c.CurrentUsecase)
			res := testutil.MakeRequest(router, http.MethodGet, "/rooms", "")
			if res.Code != 200 && c.OutCode == res.Code {
				return
			}
			assert.Equal(t, c.OutCode, res.Code)
			resBody := []view.RoomResponse{}
			err := json.Unmarshal(res.Body.Bytes(), &resBody)
			assert.NoError(t, err)
			assert.Equal(t, c.OutStatuses, resBody)
		})
	}
}

func TestCreateRoom(t *testing.T) {
	cases := []struct {
		Description    string
		CurrentUsecase usecase.RoomUsecase
		InBody         string
		OutCode        int
		OutStatus      view.RoomResponse
	}{
		{
			Description: "valid case",
			CurrentUsecase: &usecase.RoomUsecaseMock{
				ExpectedRoomStatus: &usecase.RoomStatus{Name: "Zboncak.Cole", Len: 1, Capacity: 3},
			},
			InBody:    `{"user_id": "711ce37b-6174-3aa3-b86d-e18ec3a77300", "user_name": "Dr. Manuela Jones", "room_name": "Dimitri71", "room_capacity": 4}`,
			OutCode:   200,
			OutStatus: view.RoomResponse{RoomName: "Zboncak.Cole", RoomCapacity: 3, RoomLen: 1},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			router := makeServer(c.CurrentUsecase)
			res := testutil.MakeRequest(router, http.MethodPost, "/rooms", c.InBody)
			if res.Code != 200 && c.OutCode == res.Code {
				return
			}
			assert.Equal(t, c.OutCode, res.Code)
			resBody := view.RoomResponse{}
			err := json.Unmarshal(res.Body.Bytes(), &resBody)
			assert.NoError(t, err)
			assert.Equal(t, c.OutStatus, resBody)
		})
	}

}

func TestJoinRoom(t *testing.T) {
	cases := []struct {
		Description    string
		CurrentUsecase usecase.RoomUsecase
		InRoomName     string
		InBody         string
		OutCode        int
		OutStatus      view.RoomResponse
	}{
		{
			Description: "valid case",
			CurrentUsecase: &usecase.RoomUsecaseMock{
				ExpectedRoomStatus: &usecase.RoomStatus{Name: "Gage Schoen", Len: 2, Capacity: 4},
			},
			InRoomName: "Zboncak.Cole",
			InBody:     `{"user_id": "67664e29-6983-31e2-9469-29668841baa5", "user_name": "Mr. Cory Goyette"}`,
			OutCode:    200,
			OutStatus:  view.RoomResponse{RoomName: "Gage Schoen", RoomLen: 2, RoomCapacity: 4},
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			router := makeServer(c.CurrentUsecase)
			res := testutil.MakeRequest(router, http.MethodPut, "/rooms/"+c.InRoomName+"/join", c.InBody)
			if res.Code != 200 && c.OutCode == res.Code {
				return
			}
			assert.Equal(t, c.OutCode, res.Code)
			resBody := view.RoomResponse{}
			err := json.Unmarshal(res.Body.Bytes(), &resBody)
			assert.NoError(t, err)
			assert.Equal(t, c.OutStatus, resBody)
		})
	}
}

func TestLeave(t *testing.T) {

	cases := []struct {
		Description    string
		CurrentUsecase usecase.RoomUsecase
		InRoomName     string
		InBody         string
		OutCode        int
		OutStatus      view.RoomResponse
	}{
		{
			Description:    "valid case",
			CurrentUsecase: &usecase.RoomUsecaseMock{},
			InRoomName:     "Zboncak.Cole",
			InBody:         `{"user_id": "67664e29-6983-31e2-9469-29668841baa5", "user_name": "Mr. Cory Goyette"}`,
			OutCode:        200,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			router := makeServer(c.CurrentUsecase)
			res := testutil.MakeRequest(router, http.MethodPut, "/rooms/"+c.InRoomName+"/leave", c.InBody)
			assert.Equal(t, c.OutCode, res.Code)
		})

	}
}
