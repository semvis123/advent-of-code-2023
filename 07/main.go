package main

import (
	"log"
	"semvis123/aoc"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines))
	log.Printf("part 2: %d", part_2(lines))
}

type Hand struct {
	Cards   []int
	Bidding int
	Str     string
	Type    int
}

var order1 = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
var order2 = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

func getHandTypeOfCards(cardsN int, maxCount int) int {
	t := 0
	switch maxCount {
	case 5:
		t = 0
	case 4:
		t = 1
	case 3:
		if cardsN == 2 {
			t = 2
		} else {
			t = 3
		}
	case 2:
		if cardsN == 3 {
			t = 4
		} else {
			t = 5
		}
	case 1:
		t = 6
	}
	return t
}

func getHandTypeF(withJoker bool) func(cards []int) int {
	return func(cards []int) int {
		var counts []int
		var used []int
		joker := 0
		for _, c := range cards {
			if c == 1 && withJoker {
				joker++
				continue
			}
			if !slices.Contains(used, c) {
				count := aoc.Count(cards, c)
				counts = append(counts, count)
				used = append(used, c)
			}
		}

		maxCount := slices.Max(aoc.Iff(len(counts) == 0, []int{0}, counts))
		if withJoker {
			maxCount += joker
		}
		return getHandTypeOfCards(len(counts), maxCount)
	}
}

func readHands(input []string, order []string, handF func([]int) int) (hands []Hand) {
	for _, line := range input {
		if len(line) == 0 {
			continue
		}

		lineSplitted := strings.Split(line, " ")

		cards := aoc.Map(strings.Split(lineSplitted[0], ""), func(s string) int {
			return len(order) - slices.Index(order, s)
		})
		t := handF(cards)

		hands = append(hands, Hand{
			Cards:   cards,
			Bidding: aoc.Must(strconv.Atoi(lineSplitted[1])),
			Str:     lineSplitted[0],
			Type:    t,
		})
	}

	return
}

func orderHands(hands []Hand) []Hand {
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Type > b.Type {
			return 1
		} else if a.Type < b.Type {
			return -1
		}
		for i := range a.Cards {
			if a.Cards[i] < b.Cards[i] {
				return 1
			} else if a.Cards[i] > b.Cards[i] {
				return -1
			}
		}
		return 0
	})
	slices.Reverse(hands)
	return hands
}

func total(hands []Hand) int {
	var sum int
	for i, h := range hands {
		sum += h.Bidding * (i + 1)
	}
	return sum
}

func part_1(input []string) (s int) {
	hands := readHands(input, order1, getHandTypeF(false))
	hands = orderHands(hands)

	return total(hands)
}

func part_2(input []string) int {
	hands := readHands(input, order2, getHandTypeF(true))
	hands = orderHands(hands)

	return total(hands)
}
