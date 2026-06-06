package game

// Went with a much simpler scoring implementation, eventually will need to deal with dead stones.
func CountTerritory(b *Board) (int, int) {
	black_territory := 0
	white_territory := 0
	visited := make(map[Point]bool)

	for r := 0; r < b.Size; r++ {
		for c := 0; c < b.Size; c++ {
			p := Point{r, c}

			if visited[p] || b.Get(p) != Empty {
				continue
			}

			region := []Point{}
			borders := make(map[Color]bool)

			var dfs func(Point)
			dfs = func(pt Point) {
				if visited[pt] == true {
					return
				}
				visited[pt] = true
				region = append(region, pt)

				for _, neighbor := range b.getNeighbors(pt) {
					color := b.Get(neighbor)
					if color == Empty {
						dfs(neighbor)
					} else {
						borders[color] = true
					}
				}
			}
			dfs(p)
			if borders[Black] && !borders[White] {
				black_territory += len(region)
			} else if borders[White] && !borders[Black] {
				white_territory += len(region)
			}
		}
	}
	return black_territory, white_territory
}

func FinalScore(g *Game, komi float64) (float64, float64) {
	bterr, wterr := CountTerritory(g.Board)
	black_score := float64(bterr + g.Captures[Black])
	white_score := float64(wterr+g.Captures[White]) + komi
	return black_score, white_score
}

func Winner(black, white float64) string {
	if black > white {
		return "Black"
	}

	if white > black {
		return "White"
	}
	return "Draw"
}
