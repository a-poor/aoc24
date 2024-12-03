package main

import (
	"fmt"
	"testing"
)

func Test_isLevelSafe(t *testing.T) {
	cases := []struct {
		level []int
		want  bool
	}{
		{level: []int{7, 6, 4, 2, 1}, want: true},
		{level: []int{1, 2, 7, 8, 9}, want: false},
		{level: []int{9, 7, 6, 2, 1}, want: false},
		{level: []int{1, 3, 2, 4, 5}, want: false},
		{level: []int{8, 6, 4, 4, 1}, want: false},
		{level: []int{1, 3, 6, 7, 9}, want: true},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isLevelSafe(c.level)
			if got != c.want {
				t.Errorf("isLevelSafe(%v) == %v, want %v", c.level, got, c.want)
			}
		})
	}
}

func Test_isLevelSafeWithRemoval(t *testing.T) {
	cases := []struct {
		level []int
		want  bool
	}{
		{level: []int{7, 6, 4, 2, 1}, want: true},
		{level: []int{1, 2, 7, 8, 9}, want: false},
		{level: []int{9, 7, 6, 2, 1}, want: false},
		{level: []int{1, 3, 2, 4, 5}, want: true},
		{level: []int{8, 6, 4, 4, 1}, want: true},
		{level: []int{1, 3, 6, 7, 9}, want: true},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isLevelSafeWithRemoval(c.level)
			if got != c.want {
				t.Errorf("isLevelSafe(%v) == %v, want %v", c.level, got, c.want)
			}
		})
	}
}
