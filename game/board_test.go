package game

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard(9)
	if b.Size != 9 {
		t.Errorf("expected size 9, got %d", b.Size)
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b.Get(Point{r, c}) != Empty {
				t.Errorf("expected empty at (%d,%d)", r, c)
			}
		}
	}
}

func TestBoardGetSet(t *testing.T) {
	b := NewBoard(9)
	p := Point{3, 4}
	b.Set(p, Black)
	if b.Get(p) != Black {
		t.Errorf("expected Black at %v, got %v", p, b.Get(p))
	}
}

func TestBoardSnapshot(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{0, 0}, Black)
	snap := b.Snapshot()
	b.Set(Point{0, 0}, White)
	if snap[0][0] != Black {
		t.Error("snapshot should be independent of subsequent board changes")
	}
}
