package main

import (
	"aoc24"
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
  reRule = regexp.MustCompile(`(\d+)\|(\d+)`)
)

func main() {
  slog.Info("Loading the input") 
  in := aoc24.GetInput(5, false)

  slog.Info("Parsing the input")
  data := parseInput(in)

  slog.Info("Walking through the updates")
  var total int
updateLoop:
  for _, up := range data.updates {
    // Test against the rules
  ruleLoop:
    for _, r := range data.rules {
      // Unpack the rule
      rn1, rn2 := r[0], r[1]

      // Does rn1 exist in the update?
      i1 := slices.Index(up, rn1)
      i2 := slices.Index(up, rn2)

      // Are they both in the update?
      // If not, skip this rule
      if i1 == -1 || i2 == -1 {
        continue ruleLoop
      }

      // If rn2 comes after rn1, then we're good
      if i2 < i1 {
        continue updateLoop
      }
    }

    // If the update passed all the rules, add
    // the middle number to the total
    total += getMiddleNumber(up)
  }
  slog.Info("Done", "total", total)
}

type inputData struct {
  rules [][2]int
  updates [][]int
}

func parseInput(input []byte) inputData {
  // Split into lines
  lines := strings.Split(string(input), "\n")

  // Split into two sections
  var ruleLines []string
  var updateLines []string
  toRules := true 
  for _, line := range lines { 
    switch lineEmpty := line == ""; {
    case lineEmpty && toRules:
      toRules = false

    case lineEmpty && !toRules:
      // Skip this line

    case toRules && !lineEmpty:
      ruleLines = append(ruleLines, line)

    case !toRules && !lineEmpty:
      updateLines = append(updateLines, line)
    }
  }

  // Parse the rules section
  rules := parseRulesSection(ruleLines)

  // Parse the updates section
  updates := parseUpdatesSection(updateLines)

  // Combine and return
  return inputData{
    rules: rules,
    updates: updates,
  }
}

func parseRulesSection(lines []string) [][2]int {
  rules := make([][2]int, len(lines))
  for i, line := range lines {
    // Parse the matches from the line
    match := reRule.FindStringSubmatch(line)
    if len(match) != 3 {
      panic(fmt.Errorf("Invalid rule line=%q - no match", line))
    }

    // Convert the submatches to numbers
    n1, err := strconv.Atoi(match[1])
    if err != nil {
      panic(fmt.Errorf("Invalid rule line=%q - failed converting 1: %w", line, err))
    }
    n2, err := strconv.Atoi(match[2])
    if err != nil {
      panic(fmt.Errorf("Invalid rule line=%q - failed converting 1: %w", line, err))
    }

    // Add the rule to the list
    rules[i] = [2]int{n1, n2}
  }
  return rules
}

func parseUpdatesSection(lines []string) [][]int {
  updates := make([][]int, len(lines))
  for i, line := range lines {
    // Split the line into numbers
    parts := strings.Split(line, ",")

    // Parse the strings as numbers
    update := make([]int, len(parts))
    for j, part := range parts {
      n, err := strconv.Atoi(part)
      if err != nil {
        panic(fmt.Errorf("Invalid update %dth line=%q - failed converting %d: %w", i, line, j, err))
      }
      update[j] = n
    }
    updates[i] = update
  }
  return updates
}

func getMiddleNumber(ns []int) int {
  return ns[len(ns) / 2]
}

