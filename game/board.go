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

func (p Point) Neighbors() []Point {
	return []Point{
		{p.Row - 1, p.Col}, 
		{p.Row + 1, p.Col}, 
		{p.Row, p.Col - 1}, 
		{p.Row, p.Col + 1}, 
	}
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
	snapshot := make([][]Color, b.Size)
	for i := range snapshot {
		snapshot[i] = make([]Color, b.Size)
		copy(snapshot[i], b.Cells[i])
	}
	return snapshot
}


