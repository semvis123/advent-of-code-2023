package main

import (
	"crypto/sha256"
	"log"
	"semvis123/aoc"

	"golang.org/x/exp/slices"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

func tiltNorth(input [][]byte) [][]byte {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'O' {
				for k := i - 1; k >= 0; k-- {
					if input[k][j] == '.' {
						input[k+1][j], input[k][j] = input[k][j], input[k+1][j]
					} else {
						break
					}
				}
			}
		}
	}
	return input
}

func tiltSouth(input [][]byte) [][]byte {
	for i := len(input) - 1; i >= 0; i-- {
		for j := range input[i] {
			if input[i][j] == 'O' {
				for k := i; k < len(input)-1; k++ {
					if input[k+1][j] == '.' {
						input[k+1][j], input[k][j] = input[k][j], input[k+1][j]
					} else {
						break
					}
				}
			}
		}
	}
	return input
}

func tiltEast(input [][]byte) [][]byte {
	for i := range input {
		for j := len(input[i]) - 1; j >= 0; j-- {
			if input[i][j] == 'O' {
				for k := j; k < len(input[0])-1; k++ {
					if input[i][k+1] == '.' {
						input[i][k+1], input[i][k] = input[i][k], input[i][k+1]
					} else {
						break
					}
				}
			}
		}
	}
	return input
}

func tiltWest(input [][]byte) [][]byte {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'O' {
				for k := j - 1; k >= 0; k-- {
					if input[i][k] == '.' {
						input[i][k+1], input[i][k] = input[i][k], input[i][k+1]
					} else {
						break
					}
				}
			}
		}
	}
	return input
}

func cycle(input [][]byte) [][]byte {
	input = tiltNorth(input)
	input = tiltWest(input)
	input = tiltSouth(input)
	input = tiltEast(input)
	return input
}

func CountSupport(inp [][]byte) (s int) {
	for i := range inp {
		for j := range inp[i] {
			if inp[i][j] == 'O' {
				s += len(inp) - i
			}
		}
	}
	return
}

func part_1(input []string) (s int) {
	m := tiltNorth(aoc.Map(input, func(s string) []byte { return []byte(s) }))
	return CountSupport(m)
}

func hash(inp [][]byte) string {
	s := []byte{}
	for _, x := range inp {
		h := sha256.Sum256(x)
		s = append(s, h[:]...)
	}
	return string(s)
}

var _ = slices.Index([]string{}, "")

func part_2(input []string) (s int) {
	hashes := []string{}
	var m [][]byte = aoc.Map(input, func(s string) []byte { return []byte(s) })
	n := 1_000_000_000
	for i := 0; i < n; i++ {
		m = cycle(m)
		h := hash(m)
		if slices.Contains(hashes, h) {
			hi := slices.Index(hashes, h)
			i += int((n-i)/(i-hi)) * (i - hi)
		}
		hashes = append(hashes, h)
	}
	return CountSupport(m)
}
