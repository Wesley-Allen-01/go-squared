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
	board := &Board{size, make([][]Color, size)}
	for i := range board.Cells {
		board.Cells[i] = make([]Color, Empty)
	}
	return &Game{
		Board:     board,
		Current:   Black,
		Captures:  make(map[Color]int),
		Passes:    0,
		PrevBoard: nil,
		Over:      false,
	}
}

func (g *Game) PlaceStone(p Point) error {
	return nil
}

func (g *Game) Pass() {
	g.SwitchPlayer()
}

func (g *Game) Resign() {
	g.Over = true
	//NewGame(9)
}

func (g *Game) SwitchPlayer() {
	if g.Current == Black {
		g.Current = White
	} else {
		g.Current = Black
	}
}
