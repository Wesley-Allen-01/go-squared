package game

import "testing"

func TestFindGroupSingleStone(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{4, 4}, Black)
	group := FindGroup(b, Point{4, 4})
	if len(group) != 1 {
		t.Errorf("expected group of 1, got %d", len(group))
	}
}

func TestFindGroupConnected(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{4, 4}, Black)
	b.Set(Point{4, 5}, Black)
	b.Set(Point{4, 6}, Black)
	group := FindGroup(b, Point{4, 4})
	if len(group) != 3 {
		t.Errorf("expected group of 3, got %d", len(group))
	}
}

func TestFindGroupDoesNotIncludeEnemy(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{4, 4}, Black)
	b.Set(Point{4, 5}, White)
	group := FindGroup(b, Point{4, 4})
	if len(group) != 1 {
		t.Errorf("enemy stone should not be included in group, got size %d", len(group))
	}
}

func TestCountLibertiesCenter(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{4, 4}, Black)
	libs := CountLiberties(b, FindGroup(b, Point{4, 4}))
	if libs != 4 {
		t.Errorf("expected 4 liberties in center, got %d", libs)
	}
}

func TestCountLibertiesCorner(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{0, 0}, Black)
	libs := CountLiberties(b, FindGroup(b, Point{0, 0}))
	if libs != 2 {
		t.Errorf("expected 2 liberties in corner, got %d", libs)
	}
}

func TestCountLibertiesEdge(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{0, 4}, Black)
	libs := CountLiberties(b, FindGroup(b, Point{0, 4}))
	if libs != 3 {
		t.Errorf("expected 3 liberties on edge, got %d", libs)
	}
}

func TestIsCapturedFalse(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{4, 4}, Black)
	if IsCaptured(b, FindGroup(b, Point{4, 4})) {
		t.Error("stone with liberties should not be captured")
	}
}

func TestIsCapturedTrue(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{0, 0}, Black)
	b.Set(Point{0, 1}, White)
	b.Set(Point{1, 0}, White)
	if !IsCaptured(b, FindGroup(b, Point{0, 0})) {
		t.Error("stone with no liberties should be captured")
	}
}
