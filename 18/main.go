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

type dig struct {
	dir   byte
	steps int
	color string
}

func (d dig) getDir() aoc.Coord2d[int] {
	switch d.dir {
	case 'R', '0':
		return aoc.Coord2d[int]{X: 1, Y: 0}
	case 'L', '2':
		return aoc.Coord2d[int]{X: -1, Y: 0}
	case 'U', '3':
		return aoc.Coord2d[int]{X: 0, Y: -1}
	case 'D', '1':
		return aoc.Coord2d[int]{X: 0, Y: 1}
	}
	panic("invalid dir")
}

func calcArea(digs []dig) int {
	coords := []aoc.Coord2d[int]{{}}
	curr := aoc.Coord2d[int]{}
	for _, d := range digs {
		curr = curr.Add(d.getDir().Scale(d.steps)).(aoc.Coord2d[int])
		coords = append(coords, curr)
	}

	t := 0
	for i := 0; i < len(coords)-1; i++ {
		t += coords[i].X * coords[i+1].Y
		t -= coords[i].Y * coords[i+1].X
	}

	return t/2 + aoc.Sum(aoc.Map(digs, func(d dig) int { return d.steps }))/2 + 1
}

func part_1(input []string) (s int) {
	digs := []dig{}
	for _, l := range input {
		parts := strings.Split(l, " ")
		digs = append(digs, dig{
			dir:   parts[0][0],
			steps: aoc.Must(strconv.Atoi(parts[1])),
			color: parts[2],
		})
	}
	return calcArea(digs)
}

func part_2(input []string) (s int) {
	digs := []dig{}
	for _, l := range input {
		parts := strings.Split(l, " ")
		digs = append(digs, dig{
			dir:   parts[2][7],
			steps: int(aoc.Must(strconv.ParseInt(parts[2][2:7], 16, 64))),
			color: parts[2][2:8],
		})
	}

	return calcArea(digs)
}
