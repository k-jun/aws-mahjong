package repository

import (
	"regexp"

	socketio "github.com/googollee/go-socket.io"
)

var (
	roomPrefix = "aws-mahjong/"
	re         = regexp.MustCompile(`aws-mahjong.+`)
)

type RoomRepository struct {
	wsserver *socketio.Server
}

func NewRoomRepository(wsserver *socketio.Server) *RoomRepository {
	return &RoomRepository{wsserver: wsserver}
}

func (s *RoomRepository) JoinRoom(conn socketio.Conn, roomName string) {
	conn.Join(roomPrefix + roomName)

}

func (s *RoomRepository) LeaveRoom(conn socketio.Conn, roomName string) {
	conn.Leave(roomPrefix + roomName)
}

func (s *RoomRepository) Rooms() []string {
	names := []string{}
	for _, name := range s.wsserver.Rooms("/") {
		if re.MatchString(name) {
			names = append(names, name[len(roomPrefix):])
		}
	}
	return names
}

func (s *RoomRepository) RoomLen(roomName string) int {
	return s.wsserver.RoomLen("/", roomPrefix+roomName)

}
