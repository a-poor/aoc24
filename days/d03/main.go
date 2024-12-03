package main

import (
	"aoc24"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
  slog.Info("Loading input")
  in := aoc24.GetInput(3)
  
  slog.Info("Finding matches")
  matches := re.FindAllSubmatch(in, -1)
  
  slog.Info("Processing matches")
  var total int
  for i, m := range matches {
    n1, err := strconv.Atoi(string(m[1]))
    if err != nil {
      panic(fmt.Errorf("error converting %q to int from %q in %d-th match: %w",
        string(m[1]),
        string(m[0]),
        i,
        err,
      ))
    }
    n2, err := strconv.Atoi(string(m[2]))
    if err != nil {
      panic(fmt.Errorf("error converting %q to int from %q in %d-th match: %w",
        string(m[1]),
        string(m[0]),
        i,
        err,
      ))
    }
    total += n1 * n2
  }
  slog.Info("Done", "total", total)
}

