package main

import (
	"log"
	"semvis123/aoc"
	"strconv"
	"strings"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines))
	log.Printf("part 2: %d", part_2(lines))
}

func getPyramid(line string) [][]int {
	vals := aoc.Map(strings.Split(line, " "), aoc.WrapMust(strconv.Atoi))
	diffs := [][]int{vals}
	for i := 0; !aoc.AllV(diffs[len(diffs)-1], 0); i++ {
		newL := make([]int, 0, len(vals)-1-i)
		for j := range diffs[i][1:] {
			newL = append(newL, diffs[i][j+1]-diffs[i][j])
		}
		diffs = append(diffs, newL)
	}
	return diffs
}

func part_1(input []string) (s int) {
	for _, l := range input[:len(input)-1] {
		diffs := getPyramid(l)

		n := 0
		for i := len(diffs) - 1; i >= 0; i-- {
			n += diffs[i][len(diffs[i])-1]
		}
		s += n
	}
	return
}

func part_2(input []string) (s int) {
	for _, l := range input[:len(input)-1] {
		diffs := getPyramid(l)

		n := 0
		for i := len(diffs) - 1; i >= 0; i-- {
			n += diffs[i][0]
			n *= -1
		}
		n *= -1
		s += n
	}
	return
}
