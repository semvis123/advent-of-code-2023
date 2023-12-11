package main

import (
	"log"
	"semvis123/aoc"
	"strings"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

func expand(input []string) []string {
	newMap := make([]string, len(input))
	copy(newMap, input)
row:
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if newMap[i][j] != '.' {
				continue row
			}
		}
		newMap = append(newMap[:i], append([]string{strings.Repeat(".", len(newMap[0]))}, newMap[i:]...)...)
		i++
	}
col:
	for i := 0; i < len(newMap[0]); i++ {
		for j := 0; j < len(newMap); j++ {
			if newMap[j][i] != '.' {
				continue col
			}
		}
		for j := 0; j < len(newMap); j++ {
			newMap[j] = newMap[j][:i] + "." + newMap[j][i:]
		}
		i++
	}
	return newMap
}

func pathLen(start, end aoc.Coord2d[int], input []string) int {
	return aoc.AbsInt(start.X-end.X) + aoc.AbsInt(start.Y-end.Y)
}

func sumPaths(n []string) (s int) {
	coords := []aoc.Coord2d[int]{}
	for i := range n {
		for j := range n[i] {
			if n[i][j] == '#' {
				coords = append(coords, aoc.Coord2d[int]{X: j, Y: i})
			}
		}
	}

	for i := range coords {
		for j := range coords[i+1:] {
			s += pathLen(coords[i], coords[i+1+j], n)
		}
	}
	return
}

func part_1(input []string) (s int) {
	n := expand(input)
	return sumPaths(n)
}

func part_2(input []string) (s int) {
	n := expand(input)
	x := sumPaths(input)
	return (sumPaths(n)-x)*999_999 + x
}
