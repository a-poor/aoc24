package main

import (
	"aoc24"
	"bytes"
	"fmt"
	"log/slog"
	"os"
)

const extraLogs = false

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
  

  slog.Info("Counting words")
  
  // Keep track of "MAS" centers (aka "A"s) to see if they
  // come up again later.
  //
  // Note: I'm splitting out diagonals and horizontals(aka UDLR)
  // so I only count perpindicular matches, not oblique matches.
  dcenters := make(map[point]bool)
  hcenters := make(map[point]bool)
  var total int

  p := findNextStartingPoint(grid, nil)
  for p != nil {
    // Look through the new centers
    hcs, dcs := findCentersOfMas(grid, *p)
    for _, p := range hcs {
      // If the center is already in the map,
      // it's a cross - increment and remove
      if hcenters[p] {
        slog.Debug("Found a cross", "point", p)
        total++
      } else {
        hcenters[p] = true
      }
    }
    for _, p := range dcs {
      // If the center is already in the map,
      // it's a cross - increment and remove
      if dcenters[p] {
        total++
        slog.Debug("Found a cross", "point", p)
      } else {
        dcenters[p] = true
      }
    }

    // Move the starting point
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
  return point{y: p.y - 1, x: p.x + 1}
}

func (p point) UpLeft() point {
  return point{y: p.y - 1, x: p.x - 1}
}

func (p point) DownRight() point {
  return point{y: p.y + 1, x: p.x + 1}
}

func (p point) DownLeft() point {
  return point{y: p.y + 1, x: p.x - 1}
}

// findNextStartingPoint returns the next possible starting point
// for the target word `mas` in the grid, after the given point.
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
      if r == 'm' || r == 'M' {
        return &point{y: i, x: j}
      }
    }
  }
  return nil
}

// findCentersOfMas returns the center points of the target word
// `mas` grid, starting from the given startingPoint character.
func findCentersOfMas(grid [][]rune, startingPoint point) ([]point, []point) {
  var dcs []point
  var hcs []point
  if findRight(grid, startingPoint) {
    hcs = append(hcs, startingPoint.Right())
  }
  if findLeft(grid, startingPoint) {
    hcs = append(hcs, startingPoint.Left())
  }
  if findUp(grid, startingPoint) {
    hcs = append(hcs, startingPoint.Up())
  }
  if findDown(grid, startingPoint) {
    hcs = append(hcs, startingPoint.Down())
  }
  if findUpRight(grid, startingPoint) {
    dcs = append(dcs, startingPoint.UpRight())
  }
  if findUpLeft(grid, startingPoint) {
    dcs = append(dcs, startingPoint.UpLeft())
  }
  if findDownRight(grid, startingPoint) {
    dcs = append(dcs, startingPoint.DownRight())
  }
  if findDownLeft(grid, startingPoint) {
    dcs = append(dcs, startingPoint.DownLeft())
  }
  return hcs, dcs
}

func findRight(grid [][]rune, p point) bool {
  var word []rune
  w := len(grid[0])
  for i := 0; i < len(`mas`) && p.x+i < w; i++ {
    word = append(word, grid[p.y][p.x+i])
  }
  s := string(word)
  return s == "mas" || s == "MAS"
}

func findLeft(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`mas`) && p.x-i >= 0; i++ {
    word = append(word, grid[p.y][p.x-i])
  }
  s := string(word)
  return s == "mas" || s == "MAS"
}

func findUp(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`mas`) && p.y-i >= 0; i++ {
    word = append(word, grid[p.y-i][p.x])
  }
  s := string(word)
  return s == "mas" || s == "MAS"
}

func findDown(grid [][]rune, p point) bool {
  var word []rune
  h := len(grid)
  for i := 0; i < len(`mas`) && p.y+i < h; i++ {
    word = append(word, grid[p.y+i][p.x])
  }
  s := string(word)
  return s == "mas" || s == "MAS"
}

func findUpRight(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`mas`) && p.y-i >= 0 && p.x+i < len(grid[0]); i++ {
    word = append(word, grid[p.y-i][p.x+i])
  }
  s := string(word)
  return s == "mas" || s == "MAS"
}

func findUpLeft(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`mas`) && p.y-i >= 0 && p.x-i >= 0; i++ {
    word = append(word, grid[p.y-i][p.x-i])
  }
  s := string(word)
  return s == "mas" || s == "MAS"
}

func findDownRight(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`mas`) && p.y+i < len(grid) && p.x+i < len(grid[0]); i++ {
    word = append(word, grid[p.y+i][p.x+i])
  }
  s := string(word)
  return s == "mas" || s == "MAS"
}

func findDownLeft(grid [][]rune, p point) bool {
  var word []rune
  for i := 0; i < len(`mas`) && p.y+i < len(grid) && p.x-i >= 0; i++ {
    word = append(word, grid[p.y+i][p.x-i])
  }
  s := string(word)
  return s == "mas" || s == "MAS"
}


