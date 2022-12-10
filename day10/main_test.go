package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", inputFile, 13760},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"actual", inputFile, "⬜⬜⬜⬛⬛⬜⬜⬜⬜⬛⬜⬛⬛⬜⬛⬜⬜⬜⬜⬛⬛⬜⬜⬛⬛⬜⬜⬜⬛⬛⬜⬜⬜⬜⬛⬜⬜⬜⬜⬛\n⬜⬛⬛⬜⬛⬜⬛⬛⬛⬛⬜⬛⬜⬛⬛⬛⬛⬛⬜⬛⬜⬛⬛⬜⬛⬜⬛⬛⬜⬛⬜⬛⬛⬛⬛⬜⬛⬛⬛⬛\n⬜⬛⬛⬜⬛⬜⬜⬜⬛⬛⬜⬜⬛⬛⬛⬛⬛⬜⬛⬛⬜⬛⬛⬛⬛⬜⬛⬛⬜⬛⬜⬜⬜⬛⬛⬜⬜⬜⬛⬛\n⬜⬜⬜⬛⬛⬜⬛⬛⬛⬛⬜⬛⬜⬛⬛⬛⬜⬛⬛⬛⬜⬛⬛⬛⬛⬜⬜⬜⬛⬛⬜⬛⬛⬛⬛⬜⬛⬛⬛⬛\n⬜⬛⬜⬛⬛⬜⬛⬛⬛⬛⬜⬛⬜⬛⬛⬜⬛⬛⬛⬛⬜⬛⬛⬜⬛⬜⬛⬛⬛⬛⬜⬛⬛⬛⬛⬜⬛⬛⬛⬛\n⬜⬛⬛⬜⬛⬜⬛⬛⬛⬛⬜⬛⬛⬜⬛⬜⬜⬜⬜⬛⬛⬜⬜⬛⬛⬜⬛⬛⬛⬛⬜⬜⬜⬜⬛⬜⬛⬛⬛⬛\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
