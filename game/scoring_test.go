package game

import "testing"

func TestCountTerritoryEmptyBoard(t *testing.T) {
	b := NewBoard(9)
	black, white := CountTerritory(b)
	if black != 0 || white != 0 {
		t.Errorf("empty board should have 0 territory, got black=%d white=%d", black, white)
	}
}

func TestCountTerritoryBlack(t *testing.T) {
	b := NewBoard(9)
	// black wall along row 2 and up column 8 encloses the top-left 2x8 region (16 points)
	for i := 0; i < 9; i++ {
		b.Set(Point{2, i}, Black)
	}
	for i := 0; i < 2; i++ {
		b.Set(Point{i, 8}, Black)
	}
	black, _ := CountTerritory(b)
	if black != 16 {
		t.Errorf("expected 16 black territory points, got %d", black)
	}
}

func TestFinalScore(t *testing.T) {
	g := NewGame(9)
	g.Captures[Black] = 3
	g.Captures[White] = 1
	black, white := FinalScore(g, 6.5)
	// empty board: territory is 0 for both; scores are captures + komi
	if black != 3.0 {
		t.Errorf("expected black score 3.0, got %.1f", black)
	}
	if white != 7.5 {
		t.Errorf("expected white score 7.5 (1 capture + 6.5 komi), got %.1f", white)
	}
}

func TestWinnerBlack(t *testing.T) {
	if Winner(50, 43.5) != "Black" {
		t.Error("expected Black to win with score 50 vs 43.5")
	}
}

func TestWinnerWhite(t *testing.T) {
	if Winner(40, 49.5) != "White" {
		t.Error("expected White to win with score 40 vs 49.5")
	}
}
