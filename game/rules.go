package game



func InBounds(b *Board, p Point) bool {
	return p.isValid(b.Size)
}

func IsEmpty(b *Board, p Point) bool {
	return b.Get(p) == Empty
}

// note to self, this function can be cleaned with inversion
func CheckSuicide(b *Board, p Point, color Color) bool {
	b.Set(p, color)
	if IsCaptured(b, FindGroup(b, p)) {
		for _, neighbor := range b.getNeighbors(p) {
			if IsCaptured(b, FindGroup(b, neighbor)) {
				return false
			}
		}
		return true
	}
	b.Set(p, Empty)
	return false
}

func CheckKo(g *Game, snap [][]Color) bool {
	if g.History == nil {
		return false
	}
	gameSize := g.Board.Size
	
	history := g.History

	for _, prevBoard := range history {
		for i := 0; i < gameSize; i++ {
			for j := 0; j < gameSize; j++ {
				if prevBoard[i][j] != snap[i][j] {
					return false
				}
			}
		}
	}
	return true
}
