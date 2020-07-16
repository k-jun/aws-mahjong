package storage

import (
	"regexp"

	socketio "github.com/googollee/go-socket.io"
)

var (
	roomPrefix = "aws-mahjong/"
	re         = regexp.MustCompile(`aws-mahjong.+`)
)

type Storage struct {
	wsserver *socketio.Server
}

func NewStorage(wsserver *socketio.Server) *Storage {
	return &Storage{wsserver: wsserver}
}

func (s *Storage) JoinRoom(conn socketio.Conn, roomName string) {
	conn.Join(roomPrefix + roomName)

}

func (s *Storage) LeaveRoom(conn socketio.Conn, roomName string) {
	conn.Leave(roomPrefix + roomName)

}

func (s *Storage) Rooms() []string {
	names := []string{}
	for _, name := range s.wsserver.Rooms("/") {
		if re.MatchString(name) {
			names = append(names, name[len(roomPrefix):])
		}
	}
	return names
}

func (s *Storage) RoomLen(roomName string) int {
	return s.wsserver.RoomLen("/", roomPrefix+roomName)

}
