package server

import (
	"aws-mahjong/board"
	"aws-mahjong/deck"
	"aws-mahjong/hand"
	"aws-mahjong/server/event"
	"aws-mahjong/server/handler"
	"aws-mahjong/testutil"
	"aws-mahjong/tile"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func TestGameStart(t *testing.T) {
	client1, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client2, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client3, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client4, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		Description         string
		CurrentCreateClient *socketio_client.Client
		CurrentCreateRoom   handler.CreateRoomRequest
		CurrentJoinClients  []*socketio_client.Client
		CurrentJoinRoom     handler.JoinRoomRequest
		InClient            *socketio_client.Client
		InBody              string
		OutClient           *socketio_client.Client
	}{
		{
			Description:         "valid case",
			CurrentCreateClient: client1,
			CurrentCreateRoom:   handler.CreateRoomRequest{RoomName: "Shields.Isai", UserName: "Victor Lynch DDS", RoomCapacity: 4},
			CurrentJoinClients:  []*socketio_client.Client{client2, client3},
			CurrentJoinRoom:     handler.JoinRoomRequest{RoomName: "Shields.Isai", UserName: "Lisa DuBuque"},
			InClient:            client4,
			InBody:              `{"room_name": "Shields.Isai", "user_name": "Miss Hertha Casper V"}`,
			OutClient:           client2,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			testutil.CreateRoom(c.CurrentCreateClient, c.CurrentCreateRoom)
			testutil.JoinRooms(c.CurrentJoinClients, c.CurrentJoinRoom)

			IsFired := false

			c.OutClient.On(event.GameStart, func(payload string) {
				IsFired = true
			})
			c.InClient.Emit(event.JoinRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, true, IsFired)
		})
	}
}

func TestNewGameStatus(t *testing.T) {
	client1, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client2, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client3, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	client4, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		Description         string
		CurrentCreateClient *socketio_client.Client
		CurrentCreateRoom   handler.CreateRoomRequest
		CurrentJoinClients  []*socketio_client.Client
		CurrentJoinRoom     handler.JoinRoomRequest
		InClient            *socketio_client.Client
		InBody              string
		OutClient           *socketio_client.Client
		OutZikaze           string
		OutPlayerName       string
	}{
		{
			Description:         "valid case",
			CurrentCreateClient: client1,
			CurrentCreateRoom:   handler.CreateRoomRequest{RoomName: "Wilhelmine23", UserName: "Victor Lynch DDS", RoomCapacity: 4},
			CurrentJoinClients:  []*socketio_client.Client{client2, client3},
			CurrentJoinRoom:     handler.JoinRoomRequest{RoomName: "Wilhelmine23", UserName: "Lisa DuBuque"},
			InClient:            client4,
			InBody:              `{"room_name": "Wilhelmine23", "user_name": "Miss Hertha Casper V"}`,
			OutClient:           client1,
			OutPlayerName:       "Victor Lynch DDS",
			OutZikaze:           tile.East.Name(),
		},
		{
			Description:         "valid case",
			CurrentCreateClient: client1,
			CurrentCreateRoom:   handler.CreateRoomRequest{RoomName: "yBalistreri", UserName: "Victor Lynch DDS", RoomCapacity: 4},
			CurrentJoinClients:  []*socketio_client.Client{client2, client3},
			CurrentJoinRoom:     handler.JoinRoomRequest{RoomName: "yBalistreri", UserName: "Lisa DuBuque"},
			InClient:            client4,
			InBody:              `{"room_name": "yBalistreri", "user_name": "Miss Hertha Casper V"}`,
			OutClient:           client4,
			OutPlayerName:       "Miss Hertha Casper V",
			OutZikaze:           tile.North.Name(),
		},
	}
	// testTime, _ := time.Parse("2006-01-02", "1997-12-21")
	// deck.TimeNowUnix = testTime.Unix()

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			testutil.CreateRoom(c.CurrentCreateClient, c.CurrentCreateRoom)
			testutil.JoinRooms(c.CurrentJoinClients, c.CurrentJoinRoom)

			IsFired := false
			resBody := board.BoardStatus{}
			c.OutClient.On(event.NewGameStatus, func(payload string) {
				err := json.Unmarshal([]byte(payload), &resBody)
				if err != nil {
					return
				}
				IsFired = true
			})
			c.InClient.Emit(event.JoinRoom, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, true, IsFired)
			assert.Equal(t, tile.East.Name(), resBody.Bakaze)
			assert.Equal(t, 83, resBody.DeckCound)
			assert.Equal(t, c.OutZikaze, resBody.Jicha.Zikaze)
			assert.Equal(t, c.OutPlayerName, resBody.Jicha.Name)
			assert.Equal(t, hand.HandCount, len(resBody.Jicha.Hand))
		})
	}
}

func TestGameDahai(t *testing.T) {
	cases := []struct {
		Description         string
		CurrentRoomName     string
		CurrentRoomCapaticy int
		InClientIdx         int
		InBody              string
	}{
		{
			Description:         "valid case",
			CurrentRoomName:     "Carmella28",
			CurrentRoomCapaticy: 3,
			InClientIdx:         0,
			InBody:              `{"room_name": "Carmella28", "dahai": "manzu1"}`,
		},
	}

	testTime, _ := time.Parse("2006-01-02", "1997-12-21")
	deck.TimeNowUnix = testTime.Unix()

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			resBody := board.BoardStatus{}
			IsFired := false
			clients := testutil.CreateAndJoinClientsToRoom(uri, opts, c.CurrentRoomName, c.CurrentRoomCapaticy)

			clients[c.InClientIdx].On(event.NewGameStatus, func(payload string) {
				err := json.Unmarshal([]byte(payload), &resBody)
				if err != nil {
					return
				}
				IsFired = true
			})
			clients[c.InClientIdx].Emit(event.GameDahai, c.InBody)
			time.Sleep(1 * time.Second)
			assert.Equal(t, true, IsFired)
			assert.Equal(t, "pinzu6", resBody.Shimocha.Tsumo)
		})
	}

	deck.TimeNowUnix = time.Now().Unix()
}
