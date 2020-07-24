package testutil

// func CreateRoom(client *socketio_client.Client, room handler.CreateRoomRequest) {
// 	body, _ := json.Marshal(&room)
//
// 	client.Emit(event.CreateRoom, string(body))
// 	time.Sleep(1 * time.Second)
// }
//
// func CreateRooms(client *socketio_client.Client, rooms []handler.CreateRoomRequest) {
// 	for _, room := range rooms {
// 		CreateRoom(client, room)
// 	}
// }
//
// func JoinRoom(client *socketio_client.Client, room handler.JoinRoomRequest) {
// 	body, _ := json.Marshal(&room)
// 	client.Emit(event.JoinRoom, string(body))
// 	time.Sleep(1 * time.Second)
// }
//
// func JoinRooms(clients []*socketio_client.Client, room handler.JoinRoomRequest) {
// 	for _, client := range clients {
// 		JoinRoom(client, room)
// 	}
// }
//
// func CreateAndJoinClientsToRoom(uri string, opts *socketio_client.Options, roomName string, roomCapacity int) []*socketio_client.Client {
// 	client1, err := socketio_client.NewClient(uri, opts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	client2, err := socketio_client.NewClient(uri, opts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	client3, err := socketio_client.NewClient(uri, opts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	client4, err := socketio_client.NewClient(uri, opts)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	clients := []*socketio_client.Client{client1, client2, client3, client4}
// 	createRequest := handler.CreateRoomRequest{
// 		UserName:     "Delores Okuneva",
// 		RoomName:     roomName,
// 		RoomCapacity: roomCapacity,
// 	}
// 	CreateRoom(clients[0], createRequest)
// 	joinRequest := handler.JoinRoomRequest{
// 		UserName: "Dina Rosenbaum",
// 		RoomName: roomName,
// 	}
// 	JoinRooms(clients[1:roomCapacity], joinRequest)
// 	return clients
// }
