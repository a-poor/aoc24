package main

import (
	"aoc24"
	"cmp"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

func main() {
  slog.Info("Loading input data")
  in := aoc24.GetInput(6, false)

  slog.Info("Parsing input data")
  h, w, g, obs := parseInput(in)
  slog.Info("Parsed", "height", h, "width", w, "guard", g, "obstacles", len(obs))

  slog.Info("Running guard's inital walkthrough")
  _, spots, _ := walkTheGuard(g, obs, h, w)
 
  slog.Info("Looking for obstacle positions to create loops")
  var ls int
  for p := range spots {
    // Can't put an obstacle on the guard's 
    // current position.
    if p == g {
      continue
    }
    // slog.Info("Checking obstacle", "spot", p)

    // Add the spot to the list of obstacles
    newObs := binaryInsert(obs, p)

    // Run the guard's walkthrough with that new
    // obstacle added to the list
    _, _, loop := walkTheGuard(g, newObs, h, w)

    // If the guard is in a loop, count it
    if loop {
      // slog.Info("Found loop", "spot", p)
      ls++
    }
  }
  slog.Info("Done", "loopableSpots", ls)
}

func binaryInsert(ps []point, p point) []point {
  // Copy the ps slice
  p2 := make([]point, len(ps))
  copy(p2, ps)

  // Then find the insertion point (w/ binary search)
  i, _ := slices.BinarySearchFunc(p2, p, comparePoint)

  // Then insert into p2
  p2 = slices.Insert(p2, i, p)

  // Then return the copy
  return p2
}

func walkTheGuard(g point, obs []point, h, w int) (point, map[point]direction, bool) {
  d := Up // Starts up
  spots := map[point]direction{
    g: d,
  }
  for {
    // slog.Info("Move", "guard", g, "direction", d)

    // Find the next move. Start by turning until
    // there is an un-obstructed next step
    var turns int
    for {
      // If the guard has turned 360 degrees without
      // finding an un-obstructed next step, then panic
      if turns >= 4 {
        slog.Error("No un-obstructed next step", "guard", g, "direction", d)
        panic("unreachable")
      }

      // Pick a next step in the current direction
      next := move(g, d)

      // Check if that position is in the list of obstacles
      _, found := slices.BinarySearchFunc(obs, next, comparePoint)
      if !found {
        // Not an obstacle! Choose this move
        g = next
        break
      }

      // Otherwise, try turning
      d = d.turn()
      turns++
    }
    
    // Is the guard out of bounds?
    // Then stop the loop
    if !inBounds(g, h, w) {
      // slog.Info("Guard out of bounds", "guard", g)
      break
    }

    // Is the guard in a spot it has already visited,
    // moving in the same direction? (aka a loop)
    if spots[g] & d != 0 {
      // slog.Info("Guard in loop", "guard", g, "direction", d)
      break
    }

    // Otherwise, count the move and keep going
    spots[g] |= d
  }
  loop := spots[g] & d != 0
  return g, spots, loop
}

type direction int

const (
  Up direction = 1 << iota
  Right
  Down
  Left
)

func (d direction) turn() direction {
  switch d {
  case Up, Right, Down:
    return d << 1
  case Left:
    return Up
  }
  panic("unreachable")
}

func (d direction) String() string {
  switch d {
  case Up:
    return "U"
  case Right:
    return "R"
  case Down:
    return "D"
  case Left:
    return "L"
  }
  panic("unreachable")
}

type point struct {
  x, y int
}

func (p point) String() string {
  return fmt.Sprintf("[%d,%d]", p.y, p.x)
}

func comparePoint(a, b point) int {
  switch c := cmp.Compare(a.y, b.y); c {
  case 0:
    return cmp.Compare(a.x, b.x)
  default:
    return c
  }
}

func inBounds(p point, h, w int) bool {
  if p.y < 0 || p.x < 0 {
    return false
  }
  if p.y >= h || p.x >= w {
    return false
  }
  return true
}

func move(p point, d direction) point {
  switch d {
  case Up:
    return point{y: p.y - 1, x: p.x}
  case Right:
    return point{y: p.y, x: p.x + 1}
  case Down:
    return point{y: p.y + 1, x: p.x}
  case Left:
    return point{y: p.y, x: p.x - 1}
  }
  panic("unreachable")
}

func parseInput(in []byte) (int, int, point, []point) {
  lines := strings.Split(string(in), "\n")

  var h, w int
  var p point
  var obs []point
  for i, line := range lines {
    if len(line) == 0 {
      continue
    }
    cols := []rune(line)
    w = len(cols)
    h = i + 1
    for j, c := range cols {
      switch c {
      case '#':
        obs = append(obs, point{y: i, x: j})
      case '^':
        p = point{y: i, x: j}
      }
    }
  }
  return h, w, p, obs
}

