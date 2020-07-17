// +build integration

package server

import (
	"aws-mahjong/repository"
	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"
	"aws-mahjong/testutil"
	"aws-mahjong/usecase"
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
	uri         = "http://localhost:8000"
	socketUri   = uri + "/socket.io/"
	roomUsecase usecase.RoomUsecase
)

func TestMain(m *testing.M) {
	wsserver, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	roomRepo := repository.NewRoomRepository(wsserver)
	gameRepo := repository.NewGameRepository()
	roomUsecase = usecase.NewRoomUsecase(roomRepo, gameRepo)

	AttachHandlerAndEvent(wsserver, roomUsecase)

	go wsserver.Serve()
	defer wsserver.Close()
	go http.ListenAndServe(":8000", nil)
	os.Exit(m.Run())
}

func TestRooms(t *testing.T) {
	// client, err := socketio_client.NewClient(uri, opts)
	// testRooms := []string{"perspiciatis", "eius", "molestiae"}
	// counter := map[string]bool{}
	//
	// response, err := http.Get(uri + "/rooms")
	// assert.NoError(t, err)
	//
	// bytes, err := ioutil.ReadAll(response.Body)
	// resBody := []handler.RoomsResponse{}
	// err = json.Unmarshal(bytes, &resBody)
	//
	// for _, room := range resBody {
	// 	counter[room.RoomName] = true
	// }
	//
	// for _, roomName := range testRooms {
	// 	assert.Equal(t, true, counter[roomName])
	//
	// }
	//
}

func TestOnConnect(t *testing.T) {
	_, err := socketio_client.NewClient(uri, opts)
	assert.NoError(t, err)
}

func TestCreateRoom(t *testing.T) {

	cases := []struct {
		Description  string
		CurrentRooms []handler.CreateRoomRequest
		InBody       string
		OutError     bool
	}{
		{
			Description:  "valid case",
			CurrentRooms: []handler.CreateRoomRequest{},
			InBody:       `{"user_name": "Malcolm Ferry", "room_name": "repellendus", "room_capacity": 1}`,
			OutError:     false,
		},
		{
			Description:  "invalid case, name already taken",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "aut", UserName: "Ansel66", RoomCapacity: 2}},
			InBody:       `{"user_name": "Mireya VonRueden", "room_name": "aut", "room_capacity": 2}`,
			OutError:     true,
		},
		{
			Description:  "invalid case, invalid json",
			CurrentRooms: []handler.CreateRoomRequest{},
			InBody:       `{user_name": "Mireya VonRueden", "room_name": "repellendus", "room_capacity": 2}`,
			OutError:     true,
		},
		{
			Description:  "invalid case, invalid capacity",
			CurrentRooms: []handler.CreateRoomRequest{},
			InBody:       `{user_name": "Mireya VonRueden", "room_name": "repellendus", "room_capacity": 0}`,
			OutError:     true,
		},
		{
			Description:  "invalid case, invalid capacity",
			CurrentRooms: []handler.CreateRoomRequest{},
			InBody:       `{user_name": "Mireya VonRueden", "room_name": "repellendus", "room_capacity": -1}`,
			OutError:     true,
		},
	}

	client, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		panic(err)
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			isError := false
			testutil.CreateRooms(client, c.CurrentRooms)

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

	cases := []struct {
		Description  string
		CurrentRooms []handler.CreateRoomRequest
		InBody       string
		OutError     bool
	}{
		{
			Description:  "valid case",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Burdette75", UserName: "Christ Konopelski DDS", RoomCapacity: 2}},
			InBody:       `{"user_name": "Malcolm Ferry", "room_name": "Burdette75"}`,
			OutError:     false,
		},
		{
			Description:  "invalid case, room does not exist",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Rowe.Madilyn", UserName: "Christ Konopelski DDS", RoomCapacity: 2}},
			InBody:       `{"user_name": "Malcolm Ferry", "room_name": "wWhite"}`,
			OutError:     true,
		},
		{
			Description:  "invalid case, invalid json",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Kathryn06", UserName: "Noble Heidenreich II", RoomCapacity: 2}},
			InBody:       `"user_name": "Malcolm Ferry", "room_name": "Kathryn06"}`,
			OutError:     true,
		},
		{
			Description:  "invalid case, capacity over",
			CurrentRooms: []handler.CreateRoomRequest{{RoomName: "Vance40", UserName: "Noble Heidenreich II", RoomCapacity: 1}},
			InBody:       `{"user_name": "Malcolm Ferry", "room_name": "Vance40"}`,
			OutError:     true,
		},
	}

	client, err := socketio_client.NewClient(socketUri, opts)
	if err != nil {
		panic(err)
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			isError := false
			testutil.CreateRooms(client, c.CurrentRooms)

			client.On(event.JoinRoomError, func(payload string) {
				isError = true
			})
			client.Emit(event.JoinRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, c.OutError, isError)
		})
	}
}
