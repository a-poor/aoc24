package main

import (
	"aoc24"
	"log/slog"
	"strings"
)

func main() {
  slog.Info("Loading input data")
  in := []rune(strings.TrimSpace(string(aoc24.GetInput(9, true))))

  slog.Info("Unpacking input data")
  var unpacked []rune
  for i := 0; i < len(in); i += 2 {
    // Get the next two runes
    j := i + 1
    a, b := in[i], in[j]

    // Add the file blocks
  }
}

func r2n(r rune) int {
  return int(r - '0')
}


