package main

import (
	"aoc24"
	"bytes"
	"fmt"
	"log/slog"
)

func main() {
  slog.Info("Loading input") 
  in := aoc24.GetInput(4)

  slog.Info("Parsing input as a grid")
  grid := gridify(in)
  

  slog.Info("Counting words")
  var total int
  p := findNextStartingPoint(grid, nil)
  for p != nil {
    total += findWordCounts(grid, *p)
    p = findNextStartingPoint(grid, p)
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

// findNextStartingPoint returns the next possible starting point
// for the target word `xmas` in the grid, after the given point.
func findNextStartingPoint(grid [][]rune, after *point) *point {
  // The size of the grid
  gw, gh := len(grid[0]), len(grid)

  // Pick a starting x/y
  xstart, ystart := 0, 0
  if after != nil {
    xstart = after.x + 1
    ystart = after.y
  }

  // Check if we've looped around L/R
  if xstart >= gw {
    xstart = 0
    ystart++
  }

  // Look for the next possible starting point
  for i := ystart; i < gh; i++ {
    xs := xstart
    if i > ystart {
      xs = 0
    }
    for j := xs; j < gw; j++ {
      r := grid[i][j]
      if r == 'x' || r == 'X' {
        return &point{y: i, x: j}
      }
    }
  }
  return nil
}

// findWordCounts returns the number of times the target word
// `xmas` appears in the grid starting from the given startingPoint.
func findWordCounts(grid [][]rune, startingPoint point) int {
  var n int
  if findRight(grid, startingPoint) {
    n++
  }
  if findLeft(grid, startingPoint) {
    n++
  }
  if findUp(grid, startingPoint) {
    n++
  }
  if findDown(grid, startingPoint) {
    n++
  }
  if findUpRight(grid, startingPoint) {
    n++
  }
  if findUpLeft(grid, startingPoint) {
    n++
  }
  if findDownRight(grid, startingPoint) {
    n++
  }
  if findDownLeft(grid, startingPoint) {
    n++
  }
  return n
}

func findRight(grid [][]rune, p point) bool {
  var word []rune
  w := len(grid[0])
  for i := 0; i < len(`xmas`) && p.x+i < w; i++ {
    word = append(word, grid[p.y][p.x+i])
  }
  s := string(word)
  return s == "xmas" || s == "XMAS"
}

func findLeft(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`xmas`) && p.x-i >= 0; i++ {
    word = append(word, grid[p.y][p.x-i])
  }
  s := string(word)
  return s == "xmas" || s == "XMAS"
}

func findUp(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`xmas`) && p.y-i >= 0; i++ {
    word = append(word, grid[p.y-i][p.x])
  }
  s := string(word)
  return s == "xmas" || s == "XMAS"
}

func findDown(grid [][]rune, p point) bool {
  var word []rune
  h := len(grid)
  for i := 0; i < len(`xmas`) && p.y+i < h; i++ {
    word = append(word, grid[p.y+i][p.x])
  }
  s := string(word)
  return s == "xmas" || s == "XMAS"
}

func findUpRight(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`xmas`) && p.y-i >= 0 && p.x+i < len(grid[0]); i++ {
    word = append(word, grid[p.y-i][p.x+i])
  }
  s := string(word)
  return s == "xmas" || s == "XMAS"
}

func findUpLeft(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`xmas`) && p.y-i >= 0 && p.x-i >= 0; i++ {
    word = append(word, grid[p.y-i][p.x-i])
  }
  s := string(word)
  return s == "xmas" || s == "XMAS"
}

func findDownRight(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`xmas`) && p.y+i < len(grid) && p.x+i < len(grid[0]); i++ {
    word = append(word, grid[p.y+i][p.x+i])
  }
  s := string(word)
  return s == "xmas" || s == "XMAS"
}

func findDownLeft(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`xmas`) && p.y+i < len(grid) && p.x-i >= 0; i++ {
    word = append(word, grid[p.y+i][p.x-i])
  }
  s := string(word)
  return s == "xmas" || s == "XMAS"
}


