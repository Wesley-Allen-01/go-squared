package game

import "testing"

func TestRemoveGroup(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{4, 4}, Black)
	b.Set(Point{4, 5}, Black)
	RemoveGroup(b, FindGroup(b, Point{4, 4}))
	if b.Get(Point{4, 4}) != Empty || b.Get(Point{4, 5}) != Empty {
		t.Error("all stones in removed group should be set to Empty")
	}
}

func TestCheckCapturesNone(t *testing.T) {
	b := NewBoard(9)
	b.Set(Point{4, 4}, Black)
	captured := CheckCaptures(b, Point{4, 4}, Black)
	if captured != 0 {
		t.Errorf("expected 0 captures, got %d", captured)
	}
}

func TestCheckCapturesOne(t *testing.T) {
	b := NewBoard(9)
	// white in corner, black places last stone to surround it
	b.Set(Point{0, 0}, White)
	b.Set(Point{0, 1}, Black)
	b.Set(Point{1, 0}, Black)
	captured := CheckCaptures(b, Point{1, 0}, Black)
	if captured != 1 {
		t.Errorf("expected 1 capture, got %d", captured)
	}
	if b.Get(Point{0, 0}) != Empty {
		t.Error("captured stone should be removed from board")
	}
}

func TestCheckCapturesGroup(t *testing.T) {
	b := NewBoard(9)
	// white group of two in corner surrounded by black
	b.Set(Point{0, 0}, White)
	b.Set(Point{0, 1}, White)
	b.Set(Point{1, 0}, Black)
	b.Set(Point{1, 1}, Black)
	b.Set(Point{0, 2}, Black)
	captured := CheckCaptures(b, Point{0, 2}, Black)
	if captured != 2 {
		t.Errorf("expected 2 captures, got %d", captured)
	}
	if b.Get(Point{0, 0}) != Empty || b.Get(Point{0, 1}) != Empty {
		t.Error("both captured stones should be removed from board")
	}
}
