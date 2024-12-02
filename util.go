package aoc24

import (
	"fmt"
	"os"
)

func GetInput(d uint8) []byte {
  fn := fmt.Sprintf("./inputs/d%02d.txt", d)
  b, err := os.ReadFile(fn)
  if err != nil {
    err = fmt.Errorf("Error reading file %q: %w", fn, err)
    panic(err)
  }
  return b
}

