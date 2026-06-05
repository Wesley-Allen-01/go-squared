package game

func FindGroup(b *Board, p Point) []Point {
	group := []Point{}
	color := b.Get(p)
	if color == Empty {
		return group
	}
	visited := make(map[Point]bool)
	var dfs func(Point)
	dfs = func(Point) []Point {
		if visi
	}


}

func CountLiberties(b *Board, group []Point) int {
	return 0
}

func IsCaptured(b *Board, group []Point) bool {
	return false
}






/*

	visited := make(map[Point]bool)
	var dfs func(Point)
	dfs = func(pt Point) {
		if visited[pt] {
			return
		}
		visited[pt] = true
		if b.Get(pt) != color {
			return
		}
		group = append(group, pt)
		for _, neighbor := range pt.Neighbors() {
			if neighbor.Row >= 0 && neighbor.Row < b.Size && neighbor.Col >= 0 && neighbor.Col < b.Size {
				dfs(neighbor)
			}
		}
	}
	dfs(p)
	return group
*/