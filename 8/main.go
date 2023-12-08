package main

import (
	"log"
	"semvis123/aoc"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	lines := aoc.GetInput()
	// log.Printf("part 1: %d", part_1(lines))
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
	type state struct {
		instr int
		steps int
		node  string
	}
	log.Printf("%v", curr)

	offsets := []int{}
	distances := []int{}
	distancesToZ := []int{}
	for i := range curr {
		seen := []state{{0, 0, curr[i]}}

		steps := int(0)
		stepsZ := int(0)
		for ; (!slices.ContainsFunc(seen, func(x state) bool {
			return x.instr == (steps)%int(len(instructions)) && x.node == curr[i]
		})) || steps == 0; steps++ {
			name := curr[i]
			n := network[name]
			if instructions[steps%int(len(instructions))] {
				curr[i] = n.left
			} else {
				curr[i] = n.right
			}
			seen = append(seen, state{steps, (steps + 1) % int(len(instructions)), name})
			if name[2] == 'Z' {
				if stepsZ == 0 {
					stepsZ = steps
				}
			}
		}

		var looping *state = nil
		for _, x := range seen {
			if x.node == curr[i] && x.instr == steps%int(len(instructions)) {
				looping = &x
			}
		}
		if looping == nil {
			panic("hmmmm")
		}

		offsets = append(offsets, looping.steps)
		distances = append(distances, steps-looping.steps)
		distancesToZ = append(distancesToZ, stepsZ-looping.steps)
	}
	log.Printf("%v", offsets)
	log.Printf("%v", distances)
	log.Printf("%v", distancesToZ)

	// something's off here, not sure what
	steps := 0
	for i := range curr {
		d := aoc.RemoveIndex(distances, i)
		if steps < offsets[i] {
			steps += aoc.LCM(1, offsets[i], d...)
		}
		if (steps-offsets[i])%distances[i] != distancesToZ[i] {
			steps += aoc.LCM(1, distancesToZ[i], d...)
		}
	}

	return steps
}
