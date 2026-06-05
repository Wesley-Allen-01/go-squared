package game

type Color int

const (
	Empty Color = iota
	Black
	White
)

type Point struct {
	Row int
	Col int
}
type Board struct {
	Size  int
	Cells [][]Color
}

func NewBoard(size int) *Board {
	cells := make([][]Color, size)
	for i := range cells {
		cells[i] = make([]Color, size)
	}
	return &Board{Size: size, Cells: cells}
}

func (b *Board) Get(p Point) Color {
	return b.Cells[p.Row][p.Col]
}

func (b * Board) Set(p Point, c Color) {
	b.Cells[p.Row][p.Col] = c
}

func (b *Board) Snapshot() [][]Color {
	return b.Cells
}


