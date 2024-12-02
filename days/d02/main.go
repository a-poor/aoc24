package main

import (
	"aoc24"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  in := aoc24.GetInput(2)

}

func isLevelSafe(level []int) bool {
  for i, n := range level {
    if i == 0 {
      continue
    }
    m := level[i-1]
    d := dist(n, m)
  }
  return true
}

func dist(a, b int) int {
  d := b - a
  if d < 0 {
    return -d
  }
  return d
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

