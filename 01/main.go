package main

import (
	"fmt"
	"log"
	"semvis123/aoc"
	"strconv"
	"strings"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %s", part_1(lines))
	log.Printf("part 2: %s", part_2(lines))
}

func sumFirstAndLast(lines []string, digitFunc func(string, int) int) string {
	sum := 0
	for _, line := range lines {
		first := 0
		last := 0

		for i := 0; i < len(line); i++ {
			if d := digitFunc(line, i); d > 0 {
				if first == 0 {
					first = d
				}
				last = d
			}
		}
		sum += first*10 + last

	}

	return fmt.Sprint(sum)
}

func part_1(input []string) string {
	return sumFirstAndLast(input, func(line string, i int) int {
		return aoc.NoErr(strconv.Atoi(string(line[i])))
	})
}

func wordDigitToInt(line string, i int) int {
	if d, _ := strconv.Atoi(string(line[i])); d > 0 {
		return d
	}

	digits := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for di, digit := range digits {
		if strings.Index(line[i:], digit) == 0 {
			return di + 1
		}
	}

	return 0
}

func part_2(input []string) string {
	return sumFirstAndLast(input, wordDigitToInt)
}
