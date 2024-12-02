package main

import (
	"aoc24"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  slog.Info("Loading input")
  in := aoc24.GetInput(2)

  slog.Info("Parsing input")
  var count int
  for _, level := range strings.Split(string(in), "\n") {
    if level == "" {
      continue
    }
    if isLevelSafe(parseLevel([]byte(level))) {
      count++
    }
  }
  slog.Info("Done", "answer", count)
}

func isLevelSafe(level []int) bool {
  var dir int
  for i, n := range level {
    if i == 0 {
      continue
    }
    m := level[i-1]
    d, r := distAndDirection(n, m)
    
    // No change is not safe
    if d == 0 {
      return false
    }

    // Different direction is not safe
    if dir != 0 && dir != r {
      return false
    }
    dir = r

    // Too big a step is not safe
    if d > 3 {
      return false
    }
  }
  return true
}

func distAndDirection(a, b int) (int, int) {
  d := b - a
  if d < 0 {
    return -d, -1
  }
  return d, 1
}

func parseLevel(in []byte) []int {
  trimmed := strings.TrimSpace(string(in))
  parts := regexp.MustCompile(`\s+`).Split(trimmed, -1)
  res := make([]int, len(parts))
  for i, p := range parts {
    n, err := strconv.Atoi(p)
    if err != nil {
      panic(err)
    }
    res[i] = n
  }
  return res
}

