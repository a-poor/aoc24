package main

import (
	"aoc24"
	"cmp"
	"fmt"
	"log/slog"
	"strings"
)

func main() {
  slog.Info("Loading input data")
  in := aoc24.GetInput(8, false)

  slog.Info("Parsing input data")
  h, w, as := parseInput(in)
  slog.Info("Parsed", "height", h, "width", w, "len(antennas)", len(as))

  slog.Info("Grouping antennas by type")
  ag := groupAntennas(as)
  slog.Info("Grouped", "len(groups)", len(ag))

  slog.Info("Looking for convergence points")
  spots := map[point]bool{}
  for i := 0; i < h; i++ {
    for j := 0; j < w; j++ {
      // Make the point of the current position
      p := point{y: i, x: j}

      // Look through the antennas
      for ai, a1 := range as {
        // How far is it from the current position
        // to this antenna?
        distFromPToA1 := p.diff(a1.p)

        // Now look through the remaining antennas
        for _, a2 := range as[ai+1:] {
          // If it isn't the same type, skip it
          if a1.r != a2.r {
            continue
          }

          // How far is it from the current position?
          distFromPToA2 := p.diff(a2.p)

          // Get the doubled positions
          doublePToA1 := distFromPToA1.mul(2)
          doublePToA2 := distFromPToA2.mul(2)

          // If the distance isn't either double or
          // half, then skip it
          if !(distFromPToA1 == doublePToA2 || doublePToA1 == distFromPToA2) {
            continue
          }

          // If we made it this far, then we have a
          // convergence point
          // slog.Info("Found convergence point", "r", string(a1.r), "point", p, "antenna1", a1, "antenna2", a2)
          spots[p] = true
        }
      }
    }
  }
  slog.Info("Done", "count", len(spots))
}

type antenna struct {
  p point
  r rune
}

type point struct {
  x, y int
}

func (p point) String() string {
  return fmt.Sprintf("[%d,%d]", p.y, p.x)
}

func (p point) diff(q point) point {
  return point{
    x: p.x - q.x,
    y: p.y - q.y,
  }
}

func (p point) mul(n int) point {
  return point{
    x: p.x * n,
    y: p.y * n,
  }
}

func parseInput(in []byte) (int, int, []antenna) {
  // Capture the size of the grid
  var h, w int
 
  // Split the input into lines
  ins := strings.TrimSpace(string(in))
  lines := strings.Split(ins, "\n")
  h = len(lines)

  // Define the antennas
  var as []antenna
  for i, line := range lines {
    // Skip empty lines
    if len(line) == 0 {
      continue
    }

    // Trim the line (to be safe) and get the columns
    line = strings.TrimSpace(line)
    cols := []rune(line)
    w = len(cols)

    // Look for antennas
    for j, r := range cols {
      if r == '.' {
        continue
      }
      as = append(as, antenna{
        p: point{x: j, y: i},
        r: r,
      })
    }
  }
  return h, w, as
}

func groupAntennas(as []antenna) map[rune][]point {
  // Group the antennas by type
  groups := make(map[rune][]point)
  for _, a := range as {
    groups[a.r] = append(groups[a.r], a.p)
  }
  return groups
}

func comparePoint(a, b point) int {
  switch c := cmp.Compare(a.y, b.y); c {
  case 0:
    return cmp.Compare(a.x, b.x)
  default:
    return c
  }
}

func getAntennaType(as []antenna, p point) rune {
  for _, a := range as {
    if a.p == p {
      return a.r
    }
  }
  return '.'
}
