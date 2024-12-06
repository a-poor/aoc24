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
  for _, up := range data.updates {
    // Test against the rules
    var failed bool
    for _, r := range data.rules {
      // Unpack the rule
      rn1, rn2 := r[0], r[1]

      // Do both rule parts exist in the update?
      i1 := slices.Index(up, rn1)
      i2 := slices.Index(up, rn2)

      // If not, skip this rule
      if i1 == -1 || i2 == -1 {
        continue
      }

      // If rn2 comes after rn1, then the rule is followed
      // ...which in this case we can ignore
      if i2 < i1 {
        failed = true
        break
      }
    }

    // Didn't fail? Skip to the next update
    if !failed {
      continue
    }

    // If the update failed one or more rules...
    // - Fix / reorder the update
    // - Find the middle number
    // - Add it to the total
    u2 := reorderUpdate(up, data.rules)
    // slog.Info("FTFY", "before", up, "after", u2)
    total += getMiddleNumber(u2)
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

func reorderUpdate(up []int, rules [][2]int) []int {
  // Create an update by sorting the rule numbers
  //
  // (Reminder that all numbers in the rules must be
  // in the update - but not all numbers in the
  // update must be in the rules)
  matchingRules := filterRules(rules, up)
  sorted := sortRuleNums(matchingRules)

  // Now add in any numbers from the original update
  // that aren't in the new, sorted update
  sorted = addMissingNums(sorted, up)
  return sorted
}

func filterRules(rules [][2]int, up []int) [][2]int {
  // Determine which rules are used
  useRule := make([]bool, len(rules))
  for i, r := range rules {
    n, m := r[0], r[1]
    if slices.Index(up, n) != -1 && slices.Index(up, m) != -1 {
      useRule[i] = true
    }
  }

  // Filter the rules
  var newRules [][2]int
  for i, r := range rules {
    if useRule[i] {
      newRules = append(newRules, r)
    }
  }
  return newRules
}

func insertAt(ns []int, i, n int) []int {
  ns = append(ns, 0)
  copy(ns[i+1:], ns[i:])
  ns[i] = n
  return ns
}

// sortRuleNums creates an update by sorting a list
// of rule numbers -- where a rule [n1, n2] means that
// n1 must come before n2 in the update.
func sortRuleNums(rules [][2]int) []int {
  // A graph of rule from-numbers to to-numbers
  graph := make(map[int][]int)
  for _, r := range rules {
    n1, n2 := r[0], r[1]
    graph[n1] = append(graph[n1], n2)
  }

  // Sort the graph
  var sorted []int // The sorted update list
  for len(graph) > 0 {
    n := findNodeWithZeroIns(graph)
    if n == -1 {
      panic("No node with zero ins found")
    }
    sorted = append(sorted, n)
    delete(graph, n)
  }
 
  // Return the sorted list
  return sorted
}

func reverseGraph(g map[int][]int) map[int][]int {
  rg := make(map[int][]int)
  for nf, nts := range g {
    for _, n := range nts {
      rg[n] = append(rg[n], nf)
    }
  }
  return rg
}

func findNodeWithZeroIns(g map[int][]int) int {
  rg := reverseGraph(g)
  for k := range g {
    if _, ok := rg[k]; !ok {
      return k
    }
  }
  return -1
}

func addMissingNums(some, all []int) []int {
  full := make([]int, 0, len(all))
  for _, n := range some {
    full = append(full, n)
  }
  for _, n := range all {
    if slices.Index(full, n) == -1 {
      full = append(full, n)
    }
  }
  return full
}


