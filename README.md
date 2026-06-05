# go-squared

A terminal implementation of the board game Go, written in Go.

## Project Structure

```
go-squared/
├── main.go              # Entry point and game loop
├── game/
│   ├── board.go         # Board struct, types, Get/Set/Snapshot
│   ├── group.go         # FindGroup, CountLiberties, IsCaptured
│   ├── rules.go         # InBounds, IsEmpty, CheckSuicide, CheckKo
│   ├── capture.go       # RemoveGroup, CheckCaptures
│   ├── scoring.go       # CountTerritory, FinalScore, Winner
│   └── game.go          # Game struct, PlaceStone, Pass, Resign, SwitchPlayer
└── ui/
    └── graphics.go      # Terminal rendering
```

## File Descriptions

- **main.go** — Entry point. Runs the game loop, reads player input from the terminal, and calls into the game package.
- **game/board.go** — Defines the core types (`Color`, `Point`, `Board`). Handles creating the board, reading and writing cell values, and snapshotting board state.
- **game/group.go** — Logic for finding connected groups of stones and counting their liberties. Used by capture and rules logic.
- **game/rules.go** — Move validation. Checks whether a move is in bounds, targets an empty cell, violates the ko rule, or is a suicide move.
- **game/capture.go** — Removes captured stone groups from the board and counts how many stones were taken.
- **game/scoring.go** — Counts territory, tallies final scores (including komi), and determines the winner.
- **game/game.go** — Top-level game state. Orchestrates turns, placing stones, passing, resigning, and switching the active player.

## Ownership

| File | Owner |
|---|---|
| `game/scoring.go` | Jaden |
| `game/game.go` | Jaden |
| `game/board.go` | Wesley |
| `game/group.go` | Wesley |
| `game/rules.go` | Wesley |
| `game/capture.go` | Wesley |

## Requirements

- Go 1.23+

## Running the Game

```
go run main.go
```

## Testing

Run all tests:
```
go test ./game/...
```

Run tests with verbose output:
```
go test ./game/... -v
```

Run a single test by name:
```
go test ./game/... -run TestNewBoard
```

Run a group of tests by prefix:
```
go test ./game/... -run TestBoard
```

Run tests matching a pattern (regex):
```
go test ./game/... -run "TestPlace|TestPass"
```

Run tests in a specific file (by matching the names of tests in that file):
```
go test ./game/... -run "TestNewBoard|TestBoardGetSet|TestBoardSnapshot"
```
