package aoc24

import (
	"fmt"
	"os"
)

func GetInput(d uint8, test bool) []byte {
	var sfx string
	if test {
		sfx = ".test"
	}
	fn := fmt.Sprintf("./inputs/d%02d%s.txt", d, sfx)
	b, err := os.ReadFile(fn)
	if err != nil {
		err = fmt.Errorf("Error reading file %q: %w", fn, err)
		panic(err)
	}
	return b
}
