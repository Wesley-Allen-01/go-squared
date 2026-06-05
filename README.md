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
