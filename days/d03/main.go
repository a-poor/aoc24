package main

import (
	"aoc24"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

var (
	reMul  = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	reCond = regexp.MustCompile(`(do\(\)|don't\(\))`)
)

func main() {
	slog.Info("Loading input")
	in := aoc24.GetInput(3, false)

	slog.Info("Processing input")
	var total int
	on := true // Tracks if counting is turned on/off

	// Loop through the whole input
loop:
	for len(in) > 0 {
		// Find the next match - either a mul or a cond
		mMul := reMul.FindIndex(in)
		mCond := reCond.FindIndex(in)

		switch {
		case mMul == nil:
			// There aren't any more multiplications, so we're done
			break loop

		case mCond == nil || mMul[0] < mCond[0]:
			// The next match is a multiplication
			m := reMul.FindSubmatch(in[mMul[0]:mMul[1]])
			n1, err := strconv.Atoi(string(m[1]))
			if err != nil {
				slog.Error("Failed to convert to int", "err", err, "match", string(m[1]))
				os.Exit(1)
			}
			n2, err := strconv.Atoi(string(m[2]))
			if err != nil {
				slog.Error("Failed to convert to int", "err", err, "match", string(m[1]))
				os.Exit(1)
			}

			// Only add the multiplication if counting is turned on
			if on {
				total += n1 * n2
			}

			// Now trim the input to the end of the match
			in = in[mMul[1]:]

		case mCond[0] < mMul[0]:
			// The next match is a condition

			// - Get the match
			switch m := in[mCond[0]:mCond[1]]; string(m) {
			case "do()":
				on = true

			case "don't()":
				on = false
			}

			// Now trim the input to the end of the match
			in = in[mCond[1]:]
		}
	}

	slog.Info("Done", "total", total)
}
