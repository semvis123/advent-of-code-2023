package main

import (
	"log"
	"regexp"
	"semvis123/aoc"
	"strconv"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines))
	log.Printf("part 2: %d", part_2(lines))
}

type Race struct {
	Time     int
	Distance int
}

func (r Race) Calc(n int) int {
	return n * (r.Time - n)
}

func readRaces(input []string) (races []Race) {
	r := regexp.MustCompile(`\s+`)
	times := aoc.Map(r.Split(input[0], -1)[1:], aoc.WrapMust(strconv.Atoi))
	distances := aoc.Map(r.Split(input[1], -1)[1:], aoc.WrapMust(strconv.Atoi))

	for i := range distances {
		races = append(races, Race{
			Time:     times[i],
			Distance: distances[i],
		})
	}

	return
}

func readRace(input []string) Race {
	r := regexp.MustCompile(`\s+`)
	time := r.ReplaceAllString(input[0][11:], "")
	distance := r.ReplaceAllString(input[1][11:], "")

	return Race{
		Time:     aoc.Must(strconv.Atoi(time)),
		Distance: aoc.Must(strconv.Atoi(distance)),
	}

}

func part_1(input []string) (s int) {
	races := readRaces(input)
	var possibilities []int
	for _, r := range races {
		n := 0
		for i := 0; i < r.Time; i++ {
			n += aoc.Iff(r.Calc(i) > r.Distance, 1, 0)
		}
		possibilities = append(possibilities, n)
	}
	return aoc.Multiply(possibilities)
}

func part_2(input []string) int {
	r := readRace(input)
	n := 0
	for i := 0; i < r.Time; i++ {
		n += aoc.Iff(r.Calc(i) > r.Distance, 1, 0)
	}
	return n
}
