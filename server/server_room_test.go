// +build integration

package server

import (
	"aws-mahjong/game"
	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"
	"aws-mahjong/testutil"
	"aws-mahjong/usecase"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func TestRooms(t *testing.T) {

	cases := []struct {
		Description  string
		CurrentRooms []handler.CreateRoomRequest
		OutRooms     []handler.RoomsResponse
	}{
		{
			Description:  "valid case, single room",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Fatima.Reilly", RoomCapacity: 3, UserName: "Rosalyn King"}},
			OutRooms:     []handler.RoomsResponse{{RoomName: "Fatima.Reilly", RoomCapacity: 3, RoomMemberCount: 1}},
		},
		{
			Description: "valid case, multi rooms",
			CurrentRooms: []handler.CreateRoomRequest{
				{RoomName: "Wilford30", RoomCapacity: 3, UserName: "Rosalyn King"},
				{RoomName: "Carmen.Turcotte", RoomCapacity: 4, UserName: "Rosalyn King"},
				{RoomName: "Vincent62", RoomCapacity: 3, UserName: "Rosalyn King"},
			},
			OutRooms: []handler.RoomsResponse{
				{RoomName: "Carmen.Turcotte", RoomCapacity: 4, RoomMemberCount: 1},
				{RoomName: "Vincent62", RoomCapacity: 3, RoomMemberCount: 1},
				{RoomName: "Wilford30", RoomCapacity: 3, RoomMemberCount: 1},
			},
		},
	}

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			testutil.CreateRooms(client, c.CurrentRooms)

			response, err := http.Get(uri + "/rooms")
			assert.NoError(t, err)

			bytes, err := ioutil.ReadAll(response.Body)
			resBody := []handler.RoomsResponse{}
			err = json.Unmarshal(bytes, &resBody)

			// really bad checking
			for _, room := range c.OutRooms {
				isExist := false
				for _, resRoom := range resBody {
					if resRoom.RoomName != room.RoomName {
						continue
					}
					isExist = true
					assert.Equal(t, room.RoomCapacity, resRoom.RoomCapacity)
					assert.Equal(t, room.RoomMemberCount, resRoom.RoomMemberCount)
				}
				assert.Equal(t, true, isExist)
			}

		})
	}
}

func TestCreateRoom(t *testing.T) {

	cases := []struct {
		Description  string
		CurrentRooms []handler.CreateRoomRequest
		InBody       string
		OutError     string
	}{
		{
			Description:  "valid case",
			CurrentRooms: []handler.CreateRoomRequest{},
			InBody:       `{"user_name": "Malcolm Ferry", "room_name": "repellendus", "room_capacity": 3}`,
			OutError:     "",
		},
		{
			Description:  "invalid case, name already taken",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "aut", UserName: "Ansel66", RoomCapacity: 3}},
			InBody:       `{"user_name": "Mireya VonRueden", "room_name": "aut", "room_capacity": 3}`,
			OutError:     usecase.RoomAlraedyTakenErr.Error(),
		},
		{
			Description:  "invalid case, invalid json",
			CurrentRooms: []handler.CreateRoomRequest{},
			InBody:       `{user_name": "Mireya VonRueden", "room_name": "Emerald.Haley", "room_capacity": 3}`,
			OutError:     "invalid character 'u' looking for beginning of object key string",
		},
		{
			Description:  "invalid case, invalid capacity",
			CurrentRooms: []handler.CreateRoomRequest{},
			InBody:       `{"user_name": "Mireya VonRueden", "room_name": "Ullrich.Neha", "room_capacity": 0}`,
			OutError:     game.GameCapacityInvalid.Error(),
		},
		{
			Description:  "invalid case, invalid capacity",
			CurrentRooms: []handler.CreateRoomRequest{},
			InBody:       `{"user_name": "Mireya VonRueden", "room_name": "Phoebe.Abshire", "room_capacity": -1}`,
			OutError:     game.GameCapacityInvalid.Error(),
		},
	}

	client, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		panic(err)
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			outError := ""
			testutil.CreateRooms(client, c.CurrentRooms)

			client.On(event.RoomError, func(payload string) {
				resBody := getErrorBody(payload)
				outError = resBody.ErrorMessage
			})

			client.Emit(event.CreateRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, c.OutError, outError)
		})
	}
}

func TestJoinRoom(t *testing.T) {

	cases := []struct {
		Description  string
		CurrentRooms []handler.CreateRoomRequest
		InBody       string
		OutError     string
	}{
		{
			Description:  "valid case",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Burdette75", UserName: "Christ Konopelski DDS", RoomCapacity: 4}},
			InBody:       `{"user_name": "Malcolm Ferry", "room_name": "Burdette75"}`,
			OutError:     "",
		},
		{
			Description:  "invalid case, room does not exist",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Rowe.Madilyn", UserName: "Christ Konopelski DDS", RoomCapacity: 3}},
			InBody:       `{"user_name": "Malcolm Ferry", "room_name": "wWhite"}`,
			OutError:     usecase.RoomNotFound.Error(),
		},
		{
			Description:  "invalid case, invalid json",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Kathryn06", UserName: "Noble Heidenreich II", RoomCapacity: 4}},
			InBody:       `"user_name": "Malcolm Ferry", "room_name": "Kathryn06"}`,
			OutError:     "invalid character ':' after top-level value",
		},
		{
			Description:  "invalid case, capacity over",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Vance40", UserName: "Noble Heidenreich II", RoomCapacity: 3}},
			InBody:       `{"user_name": "Malcolm Ferry", "room_name": "Vance40"}`,
			OutError:     "",
		},
	}

	client, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			outError := ""
			testutil.CreateRooms(client, c.CurrentRooms)

			client.On(event.RoomError, func(payload string) {
				resBody := getErrorBody(payload)
				outError = resBody.ErrorMessage
			})
			client.Emit(event.JoinRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, c.OutError, outError)
		})
	}
}

func TestLeaveRoom(t *testing.T) {
	client1, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client2, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		Description     string
		CurrentClient   *socketio_client.Client
		CurrentRooms    []handler.CreateRoomRequest
		CurrentJoinRoom handler.JoinRoomRequest
		InBody          string
		OutError        string
	}{
		{
			Description:     "valid case",
			CurrentClient:   client2,
			CurrentRooms:    []handler.CreateRoomRequest{{RoomName: "uArmstrong", UserName: "Frances Schamberger", RoomCapacity: 3}},
			CurrentJoinRoom: handler.JoinRoomRequest{RoomName: "uArmstrong", UserName: "Virgie Ankunding III"},
			InBody:          `{"room_name": "uArmstrong"}`,
			OutError:        "",
		},
		{
			Description:     "invalid case, roomName not found",
			CurrentClient:   client2,
			CurrentRooms:    []handler.CreateRoomRequest{{RoomName: "Bud.Kirlin", UserName: "Dr. Lucas Simonis Sr.", RoomCapacity: 3}},
			CurrentJoinRoom: handler.JoinRoomRequest{RoomName: "Bud.Kirlin", UserName: "Virgie Ankunding III"},
			InBody:          `{"room_name": "Roel.Cummerata"}`,
			OutError:        usecase.RoomNotFound.Error(),
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			testutil.CreateRooms(client1, c.CurrentRooms)
			testutil.JoinRoom(c.CurrentClient, c.CurrentJoinRoom)
			outError := ""

			c.CurrentClient.On(event.RoomError, func(payload string) {
				resBody := getErrorBody(payload)
				outError = resBody.ErrorMessage
			})

			c.CurrentClient.Emit(event.LeaveRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, c.OutError, outError)
		})
	}
}

func TestNewRoomStatus(t *testing.T) {
	client1, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client2, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		Description       string
		CurrentClient     *socketio_client.Client
		CurrentRooms      []handler.CreateRoomRequest
		CurrentJoinClient *socketio_client.Client
		CurrentJoinRoom   handler.JoinRoomRequest
		InClient          *socketio_client.Client
		InEventName       string
		InBody            string
		OutClient         *socketio_client.Client
		OutBody           string
	}{
		{
			Description:       "valid case, create room",
			CurrentClient:     client1,
			CurrentRooms:      []handler.CreateRoomRequest{},
			CurrentJoinClient: nil,
			CurrentJoinRoom:   handler.JoinRoomRequest{},
			InClient:          client1,
			InEventName:       event.CreateRoom,
			InBody:            `{"user_name": "Celine Marks", "room_name": "Roger.Predovic", "room_capacity": 3}`,
			OutClient:         client1,
			OutBody:           `{"room_name":"Roger.Predovic","room_member_count":1,"room_capacity":3}`,
		},
		{
			Description:       "valid case, join room",
			CurrentClient:     client2,
			CurrentRooms:      []handler.CreateRoomRequest{{RoomName: "Charlotte.Ritchie", UserName: "Mr. Pablo Langworth", RoomCapacity: 4}},
			CurrentJoinClient: nil,
			CurrentJoinRoom:   handler.JoinRoomRequest{},
			InClient:          client1,
			InEventName:       event.JoinRoom,
			InBody:            `{"user_name": "Malcolm Ferry", "room_name": "Charlotte.Ritchie"}`,
			OutClient:         client2,
			OutBody:           `{"room_name":"Charlotte.Ritchie","room_member_count":2,"room_capacity":4}`,
		},
		{
			Description:       "valid case, leave room",
			CurrentClient:     client2,
			CurrentRooms:      []handler.CreateRoomRequest{{RoomName: "Christ.Ondricka", UserName: "Mr. Pablo Langworth", RoomCapacity: 4}},
			CurrentJoinClient: client1,
			CurrentJoinRoom:   handler.JoinRoomRequest{UserName: "Stanton Fay Sr.", RoomName: "Christ.Ondricka"},
			InClient:          client1,
			InEventName:       event.LeaveRoom,
			InBody:            `{"room_name": "Christ.Ondricka"}`,
			OutClient:         client2,
			OutBody:           `{"room_name":"Christ.Ondricka","room_member_count":1,"room_capacity":4}`,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			testutil.CreateRooms(c.CurrentClient, c.CurrentRooms)
			if c.CurrentJoinClient != nil {
				testutil.JoinRoom(c.CurrentJoinClient, c.CurrentJoinRoom)
			}
			outBody := ""
			c.OutClient.On(event.NewRoomStatus, func(payload string) {
				outBody = payload
			})
			c.InClient.Emit(c.InEventName, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, c.OutBody, outBody)

		})
	}
}

func getErrorBody(body string) handler.RoomErrorResponse {

	resBody := handler.RoomErrorResponse{}
	err := json.Unmarshal([]byte(body), &resBody)
	if err != nil {
		panic(err)
	}
	return resBody
}
