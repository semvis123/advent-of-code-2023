package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

func main() {
	filename := "input.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	lines := strings.Split(string(file), "\n")

	log.Printf("part 1: %d", part_1(lines))
	log.Printf("part 2: %d", part_2(lines))
}

func part_1(input []string) int {
	var numbers []int
	for i, line := range input {
		num := ""
		adjecend := false

		for j := range line {
			if unicode.IsDigit(rune(input[i][j])) {
				if !adjecend {
					for di := -1; di <= 1; di++ {
						for dj := -1; dj <= 1; dj++ {
							if i+di > 0 && i+di < len(input) && j+dj > 0 && j+dj < len(input[i+di]) {
								if input[i+di][j+dj] != '.' && !unicode.IsDigit(rune(input[i+di][j+dj])) {
									adjecend = true
								}
							}
						}
					}
				}
				num += string(input[i][j])
				if (j+1 >= len(input[i]) || !unicode.IsDigit(rune(input[i][j+1]))) && adjecend {
					n, _ := strconv.Atoi(num)
					numbers = append(numbers, n)
					continue
				}
				if j+1 < len(input[i]) && unicode.IsDigit(rune(input[i][j+1])) {
					continue
				}
			}
			num = ""
			adjecend = false
		}
	}
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}

func part_2(input []string) int {
	numbers := make(map[int][]int)
	for i, line := range input {
		var num string
		var adjecend []int

		for j := range line {
			if unicode.IsDigit(rune(input[i][j])) {
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if i+di > 0 && i+di < len(input) && j+dj > 0 && j+dj < len(input[i+di]) {
							if input[i+di][j+dj] == '*' {
								a := (i+di)*len(input[0]) + (j + dj)
								if !slices.Contains(adjecend, a) {
									adjecend = append(adjecend, a)
								}
							}
						}
					}
				}
				num += string(input[i][j])
				if (j+1 >= len(input[i]) || !unicode.IsDigit(rune(input[i][j+1]))) && adjecend != nil {
					n, _ := strconv.Atoi(num)
					for _, a := range adjecend {
						numbers[a] = append(numbers[a], n)
					}
					continue
				}
				if j+1 < len(input[i]) && unicode.IsDigit(rune(input[i][j+1])) {
					continue
				}
			}
			num = ""
			adjecend = nil
		}
	}
	s := 0
	for _, nums := range numbers {
		if len(nums) == 2 {
			s += nums[0] * nums[1]
		}
	}
	return s
}
