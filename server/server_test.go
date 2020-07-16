// +build integration

package server

import (
	"aws-mahjong/repository"
	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"
	"aws-mahjong/testutil"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/stretchr/testify/assert"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

var (
	opts = &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	uri       = "http://localhost:8000"
	socketUri = uri + "/socket.io/"
)

func TestMain(m *testing.M) {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	stg := repository.NewRoomRepository(wsserver)

	AttachHandlerAndEvent(wsserver, stg)

	go wsserver.Serve()
	defer wsserver.Close()
	go http.ListenAndServe(":8000", nil)
	os.Exit(m.Run())
}

func TestRooms(t *testing.T) {
	client, err := socketio_client.NewClient(uri, opts)
	testRooms := []string{"perspiciatis", "eius", "molestiae"}
	counter := map[string]bool{}
	for _, room := range testRooms {
		testutil.SampleRoomCreate(client, room)
	}

	response, err := http.Get(uri + "/rooms")
	assert.NoError(t, err)

	bytes, err := ioutil.ReadAll(response.Body)
	resBody := []handler.RoomsResponse{}
	err = json.Unmarshal(bytes, &resBody)

	for _, room := range resBody {
		counter[room.RoomName] = true
	}

	for _, roomName := range testRooms {
		assert.Equal(t, true, counter[roomName])

	}

}

func TestOnConnect(t *testing.T) {
	_, err := socketio_client.NewClient(uri, opts)
	assert.NoError(t, err)
}

func TestCreateRoom(t *testing.T) {

	cases := []struct {
		Description string
		InBody      string
		OutError    bool
	}{
		{
			Description: "valid case",
			InBody:      `{"user_name": "Malcolm Ferry", "room_name": "repellendus", "room_capacity": 1}`,
			OutError:    false,
		},
		{
			Description: "invalid case",
			InBody:      `{"user_name": "Mireya VonRueden", "room_name": "repellendus", "room_capacity": 2}`,
			OutError:    true,
		},
		{
			Description: "invalid case",
			InBody:      `{user_name": "Mireya VonRueden", "room_name": "repellendus", "room_capacity": 2}`,
			OutError:    true,
		},
	}

	client, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		panic(err)
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			isError := false

			client.On(event.CreateRoomError, func(payload string) {
				isError = true
			})

			client.Emit(event.CreateRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, c.OutError, isError)
		})
	}
}

func TestJoinRoom(t *testing.T) {

	testRoomName := "rerum"
	cases := []struct {
		Description string
		InBody      string
		OutError    bool
	}{
		{
			Description: "valid case",
			InBody:      `{"user_name": "Malcolm Ferry", "room_name": "` + testRoomName + `"}`,
			OutError:    false,
		},
		{
			Description: "invalid case",
			InBody:      `{"user_name": "Malcolm Ferry", "room_name": "does_not_exist_room"}`,
			OutError:    true,
		},
		{
			Description: "invalid case",
			InBody:      `"user_name": "Malcolm Ferry", "room_name": "` + testRoomName + `"}`,
			OutError:    true,
		},
	}

	client, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		panic(err)
	}

	testutil.SampleRoomCreate(client, testRoomName)

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			isError := false

			client.On(event.JoinRoomError, func(payload string) {
				isError = true
			})
			client.Emit(event.JoinRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, c.OutError, isError)
		})
	}
}
