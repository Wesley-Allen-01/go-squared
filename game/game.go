package game

type Game struct {
	Board     *Board
	Current   Color
	Captures  map[Color]int
	Passes    int
	PrevBoard [][]Color
	Over      bool
}

func NewGame(size int) *Game {
	return nil
}

func (g *Game) PlaceStone(p Point) error {
	return nil
}

func (g *Game) Pass() {
}

func (g *Game) Resign() {
}

func (g *Game) SwitchPlayer() {
}
