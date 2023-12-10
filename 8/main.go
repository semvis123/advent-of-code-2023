package main

import (
	"log"
	"semvis123/aoc"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines))
	log.Printf("part 2: %d", part_2(lines))
}

type Node struct {
	left  string
	right string
	toZ   int
}

func readNodes(input []string) (instructions []bool, network map[string]Node) {
	instructions = aoc.Map(strings.Split(input[0], ""), func(s string) bool { return s == "L" })

	network = make(map[string]Node)
	for _, l := range input[2 : len(input)-1] {
		parts := strings.Split(l, " ")
		name := parts[0]
		left := parts[2][1:4]
		right := parts[3][:3]
		network[name] = Node{left, right, -1}
	}
	return
}

func part_1(input []string) (s int) {
	instructions, network := readNodes(input)
	curr := "AAA"
	steps := 0
	for ; curr != "ZZZ"; steps++ {
		n := network[curr]
		if instructions[steps%len(instructions)] {
			curr = n.left
		} else {
			curr = n.right
		}
	}
	return steps
}

func part_2(input []string) int {
	instructions, network := readNodes(input)

	curr := aoc.Filter(maps.Keys(network), func(k string) bool { return k[2] == 'A' })

	stepsToZ := []int{}
	for i := range curr {
		steps := 0
		for ; ; steps++ {
			n := network[curr[i]]
			if curr[i][2] == 'Z' {
				break
			}
			if instructions[steps%int(len(instructions))] {
				curr[i] = n.left
			} else {
				curr[i] = n.right
			}
		}

		stepsToZ = append(stepsToZ, steps)
	}

	return aoc.LCM(1, 1, stepsToZ...)
}
