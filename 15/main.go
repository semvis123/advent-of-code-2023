package main

import (
	"log"
	"semvis123/aoc"
	"strconv"
	"strings"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

func part_1(input []string) (s int) {
	steps := strings.Split(input[0], ",")
	for _, step := range steps {
		t := 0
		for _, c := range step {
			t += int(c)
			t *= 17
			t %= 256
		}
		s += t
	}
	return
}

type lens struct {
	label  string
	length int
}

func part_2(input []string) (s int) {
	steps := strings.Split(input[0], ",")
	boxes := make([][]lens, 256)
	for _, step := range steps {
		t := 0
		loI := len(step) - 2
		loI = aoc.MaxInt(loI, strings.Index(step, "-"))
		for _, c := range step[:loI] {
			t += int(c)
			t *= 17
			t %= 256
		}
		boxI := t
		if step[loI] == '-' {
			tb := []lens{}
			removed := 0
			for _, l := range boxes[boxI] {
				if l.label != step[:loI] {
					tb = append(tb, l)
				} else {
					removed += 1
				}
			}
			if removed > 1 {
				panic("box contained duplicate?")
			}
			boxes[boxI] = tb
		} else if step[loI] == '=' {
			add := true
			for i, le := range boxes[boxI] {
				if le.label == step[:loI] {
					boxes[boxI][i].length, _ = strconv.Atoi(string(step[loI+1]))
					add = false
					break
				}
			}
			if add {
				boxes[boxI] = append(boxes[boxI], lens{step[:loI], aoc.Must(strconv.Atoi(string(step[loI+1])))})
			}
		} else {
			panic("invalid operation")
		}
	}

	for bi, b := range boxes {
		for li, l := range b {
			s += (bi + 1) * (li + 1) * (l.length)
		}
	}
	return
}
