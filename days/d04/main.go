package main

import (
	"aoc24"
	"bytes"
	"fmt"
	"log/slog"
	"os"
)

const extraLogs = true

type Direction int

const (
  Up Direction = 1 << iota
  Down
  Left
  Right
  
  UpRight = Up | Right
  UpLeft = Up | Left
  DownRight = Down | Right
  DownLeft = Down | Left
)

var gridSize point

func init() {
  lvl := slog.LevelInfo
  if extraLogs {
    lvl = slog.LevelDebug
  }
  slog.SetDefault(slog.New(slog.NewTextHandler(
    os.Stdout,
    &slog.HandlerOptions{
      Level: lvl,
    },
  )))
}

func main() {
  slog.Info("Loading input") 
  in := aoc24.GetInput(4, false)

  slog.Info("Parsing input as a grid")
  grid := gridify(in)
  gridSize = point{y: len(grid), x: len(grid[0])}

  slog.Info("Finding all 'A's in the grid")
  ps := findMidPoints(grid)
 
  slog.Info("Looking for crosses around each 'A'")
  var total int
  for _, p := range ps { 
    total += findMatches(grid, p)
  }

  slog.Info("Done", "total", total)
}

func gridify(in []byte) [][]rune {
  ls := bytes.Split(in, []byte("\n"))
  grid := make([][]rune, 0)
  for i, l := range ls {
    if len(l) == 0 {
      continue
    }
    grid = append(grid, make([]rune, len(l)))
    for j, c := range l {
      grid[i][j] = rune(c)
    }
  }
  return grid
}

type point struct { x, y int }

func (p point) String() string {
  return fmt.Sprintf("[y=%d,x=%d]", p.y, p.x)
}

func (p point) Right() point {
  return point{y: p.y, x: p.x + 1}
}

func (p point) Left() point {
  return point{y: p.y, x: p.x - 1}
}

func (p point) Up() point {
  return point{y: p.y - 1, x: p.x}
}

func (p point) Down() point {
  return point{y: p.y + 1, x: p.x}
}

func (p point) UpRight() point {
  return p.Up().Right()
}

func (p point) UpLeft() point {
  return p.Up().Left()
}

func (p point) DownRight() point {
  return p.Down().Right()
}

func (p point) DownLeft() point {
  return p.Down().Left()
}

func (p point) IsPos() bool {
  return p.x >= 0 && p.y >= 0
}

// In checks if the point is within the bounds
// of the given other point.
func (p point) In(o point) bool {
  return p.IsPos() && p.x < o.x && p.y < o.y
}

func findMidPoints(grid [][]rune) []point {
  var mps []point
  for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid[0]); j++ {
      if grid[i][j] == 'A' {
        mps = append(mps, point{y: i, x: j})
      }
    }
  }
  return mps
}

func getWord(grid [][]rune, ps ...point) string {
  var word []rune
  for _, p := range ps {
    if !p.In(gridSize) {
      continue
    }
    word = append(word, grid[p.y][p.x])
  }
  return string(word)
}

func findMatches(grid [][]rune, p point) int {
  matches := make([]Direction, 0)
  
  // Check: L  => R
  if getWord(grid, p.Left(), p, p.Right()) == "MAS" {
    matches = append(matches, Left)
  }

  // Check: R  => L
  if getWord(grid, p.Right(), p, p.Left()) == "MAS" {
    matches = append(matches, Right)
  }

  // Check: U  => D
  if getWord(grid, p.Up(), p, p.Down()) == "MAS" {
    matches = append(matches, Up)
  }

  // Check: D  => U
  if getWord(grid, p.Down(), p, p.Up()) == "MAS" {
    matches = append(matches, Down)
  }

  // Check: UL => DR
  if getWord(grid, p.UpLeft(), p, p.DownRight()) == "MAS" {
    matches = append(matches, UpLeft)
  }

  // Check: DR => UL
  if getWord(grid, p.DownRight(), p, p.UpLeft()) == "MAS" {
    matches = append(matches, DownRight)
  }

  // Check: DL => UR
  if getWord(grid, p.DownLeft(), p, p.UpRight()) == "MAS" {
    matches = append(matches, DownLeft)
  }

  // Check: UR => DL
  if getWord(grid, p.UpRight(), p, p.DownLeft()) == "MAS" {
    matches = append(matches, UpRight)
  }

  var count int
  for i, a := range matches {
    for _, b := range matches[i+1:] {
      if a.IsPerp(b) {
        count++
      }
    }
  }
  return count
}

func (d Direction) IsPerp(o Direction) bool {
  switch d {
  case Up, Down:
    return o == Left || o == Right
  case Left, Right:
    return o == Up || o == Down
  case UpRight, DownLeft:
    return o == UpLeft || o == DownRight
  case UpLeft, DownRight:
    return o == UpRight || o == DownLeft
  }
  panic(fmt.Errorf("Unknown direction: %d", d))
}

