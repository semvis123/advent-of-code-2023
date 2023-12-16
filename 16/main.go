package main

import (
	"log"
	"semvis123/aoc"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

type Beam struct {
	pos aoc.Coord2d[int]
	dir aoc.Coord2d[int]
}

func (x Beam) withinBounds(m []string) bool {
	return x.pos.Y >= 0 && x.pos.Y < len(m) && x.pos.X >= 0 && x.pos.X < len(m[0])
}

func (x Beam) getChar(m []string) byte {
	return m[x.pos.Y][x.pos.X]
}

func (x Beam) advance() Beam {
	return Beam{x.pos.Add(x.dir).(aoc.Coord2d[int]), x.dir}
}

func getAllowedDirections(x Beam, inp []string) []Beam {
	switch x.getChar(inp) {
	case '.':
		return []Beam{x.advance()}
	case '|':
		if x.dir.Y != 0 {
			return []Beam{x.advance()}
		}
		a := Beam{x.pos, aoc.Coord2d[int]{X: 0, Y: 1}}
		b := Beam{x.pos, aoc.Coord2d[int]{X: 0, Y: -1}}
		return []Beam{a.advance(), b.advance()}
	case '-':
		if x.dir.X != 0 {
			return []Beam{x.advance()}
		}
		a := Beam{x.pos, aoc.Coord2d[int]{X: 1, Y: 0}}
		b := Beam{x.pos, aoc.Coord2d[int]{X: -1, Y: 0}}
		return []Beam{a.advance(), b.advance()}
	case '/':
		var a Beam
		if x.dir.X > 0 {
			a = Beam{x.pos, aoc.Coord2d[int]{X: 0, Y: -1}}
		}
		if x.dir.X < 0 {
			a = Beam{x.pos, aoc.Coord2d[int]{X: 0, Y: 1}}
		}
		if x.dir.Y > 0 {
			a = Beam{x.pos, aoc.Coord2d[int]{X: -1, Y: 0}}
		}
		if x.dir.Y < 0 {
			a = Beam{x.pos, aoc.Coord2d[int]{X: 1, Y: 0}}
		}
		return []Beam{a.advance()}
	case '\\':
		var a Beam
		if x.dir.X > 0 {
			a = Beam{x.pos, aoc.Coord2d[int]{X: 0, Y: 1}}
		}
		if x.dir.X < 0 {
			a = Beam{x.pos, aoc.Coord2d[int]{X: 0, Y: -1}}
		}
		if x.dir.Y > 0 {
			a = Beam{x.pos, aoc.Coord2d[int]{X: 1, Y: 0}}
		}
		if x.dir.Y < 0 {
			a = Beam{x.pos, aoc.Coord2d[int]{X: -1, Y: 0}}
		}
		return []Beam{a.advance()}
	}
	panic("invalid position?")
}

func reflect(input []string, start Beam) (s int) {
	toVisit := []Beam{
		start,
	}
	var curr Beam
	visited := make(map[Beam]struct{})
	for len(toVisit) > 0 {
		curr, toVisit = toVisit[0], toVisit[1:]
		visited[curr] = struct{}{}
		for _, newNode := range getAllowedDirections(curr, input) {
			if !newNode.withinBounds(input) {
				continue
			}
			if _, ok := visited[newNode]; !ok {
				toVisit = append(toVisit, newNode)
			}
		}
	}

	var pos []aoc.Coord2d[int]

	for _, x := range maps.Keys(visited) {
		if !slices.Contains(pos, x.pos) {
			pos = append(pos, x.pos)
		}
	}

	return len(pos)
}

func part_1(input []string) (s int) {
	start := Beam{pos: aoc.Coord2d[int]{X: 0, Y: 0}, dir: aoc.Coord2d[int]{X: 1, Y: 0}}
	return reflect(input, start)
}

func part_2(input []string) (s int) {
	starts := []Beam{}
	for i := range input[0] {
		starts = append(starts, Beam{aoc.Coord2d[int]{X: i, Y: 0}, aoc.Coord2d[int]{X: 0, Y: 1}})
		starts = append(starts, Beam{aoc.Coord2d[int]{X: i, Y: len(input) - 1}, aoc.Coord2d[int]{X: 0, Y: -1}})
	}
	for i := range input {
		starts = append(starts, Beam{aoc.Coord2d[int]{X: 0, Y: i}, aoc.Coord2d[int]{X: 1, Y: 0}})
		starts = append(starts, Beam{aoc.Coord2d[int]{X: len(input[0]) - 1, Y: i}, aoc.Coord2d[int]{X: -1, Y: 0}})
	}
	for _, st := range starts {
		s = aoc.MaxInt(s, reflect(input, st))
	}
	return
}
