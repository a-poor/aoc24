package main

import (
	"aoc24"
	"cmp"
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
  // Load the input file
  slog.Info("Loading input")
  in := aoc24.GetInput(1)

  // Split it into two lists of integers
  slog.Info("Parsing input")
  a1, a2 := parseFile(in)

  // Sort the two slices
  slog.Info("Sorting lists")
  slices.SortFunc(a1, cmpNumber)
  slices.SortFunc(a2, cmpNumber)

  // Find the distances
  slog.Info("Calculating distances")
  var d int
  for i := 0; i < len(a1); i++ {
    d += dist(a1[i].pos, a2[i].pos)
  }
  slog.Info("Done", "answer", d)
}

func dist(a, b int) int {
  d := b - a
  if d < 0 {
    return -d
  }
  return d
}

type number struct {
  n, pos int
}

func cmpNumber(a, b number) int {
  return cmp.Compare(a.n, b.n)
}

func parseFile(in []byte) ([]number, []number) {
  // The slices to be returned 
  var arr1 []number
  var arr2 []number

  // Regex for matching the lines
  re := regexp.MustCompile(`(\d+)\s+(\d+)`)
 
  // Loop over the lines
  for i, line := range strings.Split(string(in), "\n") {
    if line == "" {
      continue
    }

    // Split on strings
    m := re.FindStringSubmatch(line)
    if m == nil {
      panic(fmt.Errorf("Invalid line #%d=%q no matches found", i, line))
    }
    if len(m) != 3 {
      panic(fmt.Errorf("Invalid line #%d=%q matches=%#v", i, line, m))
    }
   
    // Get the matches
    s1 := m[1]
    n1, err := strconv.Atoi(s1)
    if err != nil {
      panic(err)
    }
    s2 := m[2]
    n2, err := strconv.Atoi(s2)
    if err != nil {
      panic(err)
    }

    // Append to the return value
    arr1 = append(arr1, number{n: n1, pos: i})
    arr2 = append(arr2, number{n: n2, pos: i})
  }

  // Return the slices
  return arr1, arr2
}

