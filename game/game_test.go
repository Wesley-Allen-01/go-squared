package game

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame(9)
	if g.Board == nil {
		t.Error("board should be initialized")
	}
	if g.Current != Black {
		t.Error("Black should move first")
	}
	if g.Captures[Black] != 0 || g.Captures[White] != 0 {
		t.Error("captures should start at 0")
	}
	if g.Over {
		t.Error("game should not be over at start")
	}
}

func TestPlaceStoneValid(t *testing.T) {
	g := NewGame(9)
	err := g.PlaceStone(Point{4, 4})
	if err != nil {
		t.Errorf("valid move should not return error: %v", err)
	}
	if g.Board.Get(Point{4, 4}) != Black {
		t.Error("stone should be placed on the board")
	}
	if g.Current != White {
		t.Error("turn should switch to White after Black plays")
	}
}

func TestPlaceStoneOccupied(t *testing.T) {
	g := NewGame(9)
	g.PlaceStone(Point{4, 4})
	err := g.PlaceStone(Point{4, 4})
	if err == nil {
		t.Error("placing on an occupied position should return an error")
	}
}

func TestPlaceStoneOutOfBounds(t *testing.T) {
	g := NewGame(9)
	err := g.PlaceStone(Point{-1, 0})
	if err == nil {
		t.Error("out-of-bounds move should return an error")
	}
}

func TestPlaceStoneResetsPassCount(t *testing.T) {
	g := NewGame(9)
	g.Pass()
	g.PlaceStone(Point{4, 4})
	if g.Passes != 0 {
		t.Error("placing a stone should reset the pass counter")
	}
}

func TestPass(t *testing.T) {
	g := NewGame(9)
	g.Pass()
	if g.Passes != 1 {
		t.Errorf("expected pass count 1, got %d", g.Passes)
	}
	if g.Current != White {
		t.Error("turn should switch after a pass")
	}
	if g.Over {
		t.Error("one pass should not end the game")
	}
}

func TestTwoConsecutivePassesEndGame(t *testing.T) {
	g := NewGame(9)
	g.Pass()
	g.Pass()
	if !g.Over {
		t.Error("two consecutive passes should end the game")
	}
}

func TestResign(t *testing.T) {
	g := NewGame(9)
	g.Resign()
	if !g.Over {
		t.Error("resigning should end the game")
	}
}

func TestSwitchPlayer(t *testing.T) {
	g := NewGame(9)
	g.SwitchPlayer()
	if g.Current != White {
		t.Errorf("expected White after switching from Black, got %v", g.Current)
	}
	g.SwitchPlayer()
	if g.Current != Black {
		t.Errorf("expected Black after switching from White, got %v", g.Current)
	}
}
