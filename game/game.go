package game

import "fmt"

type Game struct {
	Board     *Board
	Current   Color
	Passes    int
	History [][][]Color
	Over      bool
}

func NewGame(size int) *Game {
	board := NewBoard(size)
	return &Game{
		Board:     board,
		Current:   Black,
		Passes:    0,
		History:   nil,
		Over:      false,
	}
}

func (g *Game) PlaceStone(p Point) error {
	if !InBounds(g.Board, p) {
		return fmt.Errorf("invalid move: point out of bounds")
	}

	if !IsEmpty(g.Board, p) {
		return fmt.Errorf("invalid move: point is occupied")
	}
	previous := g.Board.Snapshot()

	g.Board.Set(p, g.Current)

	opponent := White
	if g.Current == White {
		opponent = Black
	}

	CheckCaptures(g.Board, p, opponent)

	if IsCaptured(g.Board, FindGroup(g.Board, p)) {
		g.Board.Cells = previous
		return nil
	}

	if CheckKo(g, g.Board.Snapshot()) {
		g.Board.Cells = previous
		return nil
	}
	g.History = append(g.History, previous)
	g.Passes = 0
	g.SwitchPlayer()
	return nil
}

func (g *Game) Pass() {
	g.Passes += 1
	if g.Passes == 2 {
		g.Over = true
		return
	}
	g.SwitchPlayer()

}

func (g *Game) Resign() {
	g.Over = true
	//NewGame(9)
}

func (g *Game) SwitchPlayer() {
	if g.Current == Black {
		g.Current = White
	} else {
		g.Current = Black
	}
}
