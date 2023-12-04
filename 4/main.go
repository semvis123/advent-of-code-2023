package main

import (
	"log"
	"regexp"
	"semvis123/aoc"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines))
	log.Printf("part 2: %d", part_2(lines))
}

type Card struct {
	Id      int
	Winning []int
	Present []int
}

func ReadCards(input []string) []Card {
	r, _ := regexp.Compile(`Card\s+(\d+): (.*)`)
	rw, _ := regexp.Compile(`(\s+)`)
	var cards []Card
	for _, cardStr := range input {
		matches := r.FindStringSubmatch(cardStr)
		if len(matches) == 0 {
			continue
		}
		id := aoc.Must(strconv.Atoi(matches[1]))
		parts := strings.Split(matches[2], "|")
		winning := aoc.Map(aoc.Filter(rw.Split(parts[0], -1), aoc.NotEmpty), aoc.WrapMust(strconv.Atoi))
		present := aoc.Map(aoc.Filter(rw.Split(parts[1], -1), aoc.NotEmpty), aoc.WrapMust(strconv.Atoi))
		card := Card{
			Id:      id,
			Winning: winning,
			Present: present,
		}
		cards = append(cards, card)
	}

	return cards
}

func part_1(input []string) int {
	cards := ReadCards(input)
	values := aoc.Map(cards, func(c Card) int {
		nums := aoc.Filter(c.Present, func(n int) bool {
			return slices.Contains(c.Winning, n)
		})
		return aoc.Iff(len(nums) == 0, 0, aoc.PowInt(2, len(nums)-1))
	})
	return aoc.Sum(values)
}

func part_2(input []string) int {
	cards := ReadCards(input)
	values := aoc.Map(cards, func(c Card) int {
		nums := aoc.Filter(c.Present, func(n int) bool {
			return slices.Contains(c.Winning, n)
		})
		return len(nums)
	})

	amount := make(map[int]int)
	for i := range cards {
		amount[i] = 1
	}

	for i, n := range values {
		for j := 0; j < n; j++ {
			amount[i+j+1] = amount[i+j+1] + amount[i]
		}
	}

	return aoc.Sum(maps.Values(amount))
}
