package main

import (
	"semvis123/aoc"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	lines := aoc.GetInput()

	for i := 0; i < b.N; i++ {
		part_1(lines[:len(lines)-1])
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := aoc.GetInput()

	for i := 0; i < b.N; i++ {
		part_2(lines[:len(lines)-1])
	}
}
