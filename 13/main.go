package main

import (
	"log"
	"semvis123/aoc"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

func mirrors(p []string, smudges int) (int, bool) {
	for i := range p {
		if i == 0 {
			continue
		}
		x := len(p) - (2 * i)
		st := 0
		en := len(p)
		if x < 0 {
			st = -x
		} else {
			en = len(p) - x
		}
		a, b := p[st:i], aoc.ReverseCopy(p[i:en])

		dd := 0
		for i := range a {
			d := 0
			for j := range a[0] {
				if a[i][j] != b[i][j] {
					d++
				}
			}
			dd += d
		}
		if dd == smudges {
			return i, true
		}
	}
	return 0, false
}

func calcMirrors(input []string, smudges int) (s int) {
	patterns := aoc.SplitByEmpty(input)

	for _, p := range patterns {
		x, ok := mirrors(p, smudges)
		s += x * 100

		if !ok {
			p = aoc.TransposeStrs(p)
			x, ok = mirrors(p, smudges)
			if !ok {
				panic("pattern doesn't mirror?")
			}
			s += x
		}
	}
	return
}

func part_1(input []string) (s int) {
	return calcMirrors(input, 0)
}

func part_2(input []string) (s int) {
	return calcMirrors(input, 1)
}
