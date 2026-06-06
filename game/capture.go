package game

func RemoveGroup(b *Board, group []Point) {
	for _, pt := range group {
		b.Set(pt, Empty)
	}
}

func CheckCaptures(b *Board, p Point, color Color) int {
	captures := 0
	for _, neighbor := range b.getNeighbors(p) {
		if b.Get(neighbor) == color {
			continue
		}
		neighborGroup := FindGroup(b, neighbor)
		if IsCaptured(b, neighborGroup) {
			captures += len(neighborGroup)
			RemoveGroup(b, neighborGroup)
		}
	}
	return captures
}
