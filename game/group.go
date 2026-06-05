package game

func FindGroup(b *Board, p Point) []Point {
	group := []Point{}
	color := b.Get(p)
	if color == Empty {
		return group
	}
	visited := make(map[Point]bool)
	var dfs func(Point)
	dfs = func(pt Point) {
		if visited[pt] == true {
			return
		}
		visited[pt] = true

		if b.Get(pt) != color {
			return
		}
		group = append(group, pt)
		for _, neighbor := range b.getNeighbors(pt) {
			dfs(neighbor)
		}	
	}
	dfs(p)
	return group
}

func CountLiberties(b *Board, group []Point) int {
	neighborSet := make(map[Point]bool)

	for _, pt := range group {
		for _, neighbor := range b.getNeighbors(pt) {
			if b.Get(neighbor) == Empty {
				neighborSet[neighbor] = true
			}
		}
	}
	return len(neighborSet)
}

func IsCaptured(b *Board, group []Point) bool {
	return false
}