package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	lines := strings.Split(string(file), "\n")

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
		d, _ := strconv.Atoi(string(line[i]))
		return d
	})
}

func wordDigitToInt(line string, i int) int {
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
	if d, _ := strconv.Atoi(string(line[i])); d > 0 {
		return d
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
