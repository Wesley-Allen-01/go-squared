package game

import "testing"

func TestInBounds(t *testing.T) {
	b := NewBoard(9)
	cases := []struct {
		p       Point
		want    bool
	}{
		{Point{0, 0}, true},
		{Point{8, 8}, true},
		{Point{4, 4}, true},
		{Point{-1, 0}, false},
		{Point{0, -1}, false},
		{Point{9, 0}, false},
		{Point{0, 9}, false},
	}
	for _, c := range cases {
		if InBounds(b, c.p) != c.want {
			t.Errorf("InBounds(%v) = %v, want %v", c.p, !c.want, c.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	b := NewBoard(9)
	p := Point{4, 4}
	if !IsEmpty(b, p) {
		t.Error("new board position should be empty")
	}
	b.Set(p, Black)
	if IsEmpty(b, p) {
		t.Error("occupied position should not be empty")
	}
}

func TestCheckSuicideFalse(t *testing.T) {
	b := NewBoard(9)
	if CheckSuicide(b, Point{4, 4}, Black) {
		t.Error("placing in open center should not be suicide")
	}
}

func TestCheckSuicideTrue(t *testing.T) {
	b := NewBoard(9)
	// corner surrounded on both sides by white
	b.Set(Point{0, 1}, White)
	b.Set(Point{1, 0}, White)
	if !CheckSuicide(b, Point{0, 0}, Black) {
		t.Error("placing into fully surrounded corner should be suicide")
	}
}

func TestCheckSuicideCapturesEnemy(t *testing.T) {
	b := NewBoard(9)
	// white stone in corner, surrounded by black on two sides
	// placing black at {1,1} captures white at {0,0}, so it is not suicide
	b.Set(Point{0, 0}, White)
	b.Set(Point{0, 1}, Black)
	b.Set(Point{1, 0}, Black)
	if CheckSuicide(b, Point{1, 1}, Black) {
		t.Error("move that captures an enemy group should not be flagged as suicide")
	}
}

func TestCheckKoFalse(t *testing.T) {
	g := NewGame(9)
	// no previous board state — ko never triggers
	snap := g.Board.Snapshot()
	if CheckKo(g, snap) {
		t.Error("ko should not trigger when there is no previous board state")
	}
}

func TestCheckKoTrue(t *testing.T) {
	g := NewGame(9)
	g.Board.Set(Point{4, 4}, Black)
	snap := g.Board.Snapshot()
	g.PrevBoard = snap
	if !CheckKo(g, snap) {
		t.Error("proposed board state matching PrevBoard should trigger ko")
	}
}
