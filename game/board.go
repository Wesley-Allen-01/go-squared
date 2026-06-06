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

//Point.isValid => returns whether a point is a valid point in a board
func (p Point) isValid(boardSize int) bool {
	if p.Row < 0 || p.Row >= boardSize || p.Col < 0 || p.Col >= boardSize {
		return false
	}
	return true
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

//Get: returns the color of the point provided as arg
// func (b *Board) Get(p Point) Color {
//	 return b.Cells[p.Row][p.Col]
// }
func (b *Board) Get(p Point) Color {
	return b.Cells[p.Row][p.Col]
}

func (b *Board) Set(p Point, c Color) {
	b.Cells[p.Row][p.Col] = c
}

// Returns the number of valid neighbors on a board
func (b *Board) getNeighbors(p Point) []Point {
	neighbors := []Point{}

	for _, neighbor := range p.Neighbors() {
		if neighbor.isValid(b.Size) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (b *Board) Snapshot() [][]Color {
	snapshot := make([][]Color, b.Size)
	for i := range snapshot {
		snapshot[i] = make([]Color, b.Size)
		copy(snapshot[i], b.Cells[i])
	}
	return snapshot
}


