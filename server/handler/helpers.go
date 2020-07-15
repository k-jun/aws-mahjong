package handler

import (
	"regexp"

	socketio "github.com/googollee/go-socket.io"
)

var (
	roomPrefix = "aws-mahjong/"
	re         = regexp.MustCompile(`aws-mahjong.+`)
)

func joinRoom(s socketio.Conn, roomName string) {
	s.Join(roomPrefix + roomName)

}

func leaveRoom(s socketio.Conn, roomName string) {
	s.Leave(roomPrefix + roomName)

}

func rooms(wsserver *socketio.Server) []string {
	names := []string{}
	for _, name := range wsserver.Rooms("/") {
		if re.MatchString(name) {
			names = append(names, name[len(roomPrefix):])
		}
	}
	return names
}

func roomLen(wsserver *socketio.Server, roomName string) int {
	return wsserver.RoomLen("/", roomPrefix+roomName)

}
