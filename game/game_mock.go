package game

var _ Game = &GameMock{}

type GameMock struct {
	ExpectedCapacity int
	ExpectedError    error
}

func (g *GameMock) Capacity() int {
	return g.ExpectedCapacity

}

func (g *GameMock) AddUsername(username string) error {
	return g.ExpectedError
}
