package main

import (
	"aoc24"
	"fmt"
	"log/slog"
	"regexp"
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
  slog.Info("Counting list 2")
  cs := toCounts(a2)

  // Find the distances
  slog.Info("Calculating similarity score")
  var ts int
  for _, n := range a1 {
    ts += n * cs[n]
  }
  slog.Info("Done", "answer", ts)
}

func dist(a, b int) int {
  d := b - a
  if d < 0 {
    return -d
  }
  return d
}

func parseFile(in []byte) ([]int, []int) {
  // The slices to be returned 
  var arr1 []int
  var arr2 []int

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
    arr1 = append(arr1, n1)
    arr2 = append(arr2, n2)
  }

  // Return the slices
  return arr1, arr2
}

func toCounts(ns []int) map[int]int {
  m := make(map[int]int)
  for _, n := range ns {
    m[n]++
  }
  return m
}

