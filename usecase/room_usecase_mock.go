package usecase

var _ RoomUsecase = &RoomUsecaseMock{}

type RoomUsecaseMock struct {
	ExpectedRoomStatuses []*RoomStatus
	ExpectedRoomStatus   *RoomStatus
	ExpectedError        error
}

func (u *RoomUsecaseMock) Rooms() []*RoomStatus {
	return u.ExpectedRoomStatuses
}

func (u *RoomUsecaseMock) CreateRoom(userId string, userName string, roomName string, roomCapacity int) error {
	return u.ExpectedError

}

func (u *RoomUsecaseMock) JoinRoom(userId string, userName string, roomName string) (*RoomStatus, error) {
	return u.ExpectedRoomStatus, u.ExpectedError

}

func (u *RoomUsecaseMock) LeaveRoom(userId string, userName string, roomName string) error {
	return u.ExpectedError

}
