package main

import (
	"aoc24"
	"log/slog"
	"strconv"
	"strings"
)

func main() {
  slog.Info("Loading input")
  in := aoc24.GetInput(7, false)

  slog.Info("Parsing input")
  eqs := parseInput(in)

  slog.Info("Testing equations")
  var total int
  for _, eq := range eqs {
    if test(eq) {
      total += eq.result
    }
  }
  slog.Info("Done", "total", total)
}

type equation struct {
  result int
  parts []int
}

func parseInput(in []byte) []equation {
  sin := strings.TrimSpace(string(in))
  lines := strings.Split(sin, "\n")

  eqs := make([]equation, len(lines))
  for i, line := range lines {
    sides := strings.Split(line, ": ")
    n, err := strconv.Atoi(sides[0])
    if err != nil {
      panic(err)
    }
    eqs[i].result = n

    parts := strings.Split(strings.TrimSpace(sides[1]), " ")
    eqs[i].parts = make([]int, len(parts))
    for j, part := range parts {
      n, err := strconv.Atoi(part)
      if err != nil {
        panic(err)
      }
      eqs[i].parts[j] = n
    }
  }
  return eqs
}

func test(eq equation) bool {
  return _test(eq.result, eq.parts[0], eq.parts[1:])
}

func _test(target, current int, rest []int) bool {
  // Exactly right?
  if target == current {
    return true
  }

  // No more to try?
  if len(rest) == 0 {
    return false
  }

  // Overshot?
  if current > target {
    return false
  }

  // Get the next number
  next, rest := rest[0], rest[1:]

  // Try adding
  if _test(target, current + next, rest) {
    return true
  }

  // Try multiplying
  return _test(target, current * next, rest)
}














