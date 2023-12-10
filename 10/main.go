package main

import (
	"bytes"
	"fmt"
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

type Node struct {
	pos   aoc.Coord2d[int]
	steps int
}

func (n Node) getChar(m []string) byte {
	return m[n.pos.Y][n.pos.X]
}

func (n Node) withinBounds(m []string) bool {
	return n.pos.Y >= 0 && n.pos.Y < len(m) && n.pos.X >= 0 && n.pos.X < len(m[0])
}

func getAllowedDirections(char byte, pos aoc.Coord2d[int], input []string) [][]int {
	switch char {
	case '|':
		return [][]int{{0, -1}, {0, 1}}
	case '-':
		return [][]int{{1, 0}, {-1, 0}}
	case 'L':
		return [][]int{{1, 0}, {0, -1}}
	case 'J':
		return [][]int{{-1, 0}, {0, -1}}
	case '7':
		return [][]int{{-1, 0}, {0, 1}}
	case 'F':
		return [][]int{{1, 0}, {0, 1}}
	case '.':
		return nil
	case 'S':
		return aoc.Filter(aoc.Directions(false, false), func(d []int) bool {
			newPos := pos.Add(aoc.Coord2d[int]{X: d[0], Y: d[1]}).(aoc.Coord2d[int])
			newNode := Node{newPos, 0}
			if !newNode.withinBounds(input) {
				return false
			}
			return slices.ContainsFunc(getAllowedDirections(newNode.getChar(input), newPos, input), func(x []int) bool {
				return d[0] == -x[0] && d[1] == -x[1]
			})
		})
	}
	return nil
}

func printGrid(input []string, visited map[aoc.Coord2d[int]]int) {
	for i := range input {
		for j := range input[i] {
			x, ok := visited[aoc.Coord2d[int]{X: j, Y: i}]
			if ok {
				fmt.Print(x % 10)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part_1(input []string) (s int) {
	var start aoc.Coord2d[int]
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'S' {
				start = aoc.Coord2d[int]{X: j, Y: i}
			}
		}
	}

	toVisit := []Node{
		{pos: start, steps: 0},
	}
	var curr Node
	visited := make(map[aoc.Coord2d[int]]int)
	for len(toVisit) > 0 {
		curr, toVisit = toVisit[0], toVisit[1:]
		visited[curr.pos] = curr.steps
		for _, d := range getAllowedDirections(curr.getChar(input), curr.pos, input) {
			newPos := curr.pos.Add(aoc.Coord2d[int]{X: d[0], Y: d[1]}).(aoc.Coord2d[int])
			newNode := Node{newPos, curr.steps + 1}
			if !newNode.withinBounds(input) {
				continue
			}
			if x, ok := visited[newNode.pos]; !ok || x != newNode.steps-2 {
				toVisit = append(toVisit, newNode)
			}
		}
	}

	printGrid(input, visited)

	return slices.Max(maps.Values(visited))
}

func part_2(input []string) (s int) {
	var start aoc.Coord2d[int]
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'S' {
				start = aoc.Coord2d[int]{X: j, Y: i}
			}
		}
	}

	toVisit := []Node{
		{pos: start, steps: 0},
	}
	var curr Node
	visited := make(map[aoc.Coord2d[int]]int)
	for len(toVisit) > 0 {
		curr, toVisit = toVisit[0], toVisit[1:]
		visited[curr.pos] = curr.steps
		for _, d := range getAllowedDirections(curr.getChar(input), curr.pos, input) {
			newPos := curr.pos.Add(aoc.Coord2d[int]{X: d[0], Y: d[1]}).(aoc.Coord2d[int])
			newNode := Node{newPos, curr.steps + 1}
			if !newNode.withinBounds(input) {
				continue
			}
			if x, ok := visited[newNode.pos]; !ok || x != newNode.steps-2 {
				toVisit = append(toVisit, newNode)
			}
		}
	}

	emptyTiles := []aoc.Coord2d[int]{}
	for i := range input {
		for j := range input[i] {
			c := aoc.Coord2d[int]{X: j, Y: i}
			_, ok := visited[c]
			if !ok {
				emptyTiles = append(emptyTiles, c)
			}
		}
	}

	inside := 0
	for _, t := range emptyTiles {
		left := []byte{}
		for x := 0; x <= t.X; x++ {
			_, inLoop := visited[aoc.Coord2d[int]{X: x, Y: t.Y}]
			if inLoop {
				left = append(left, input[t.Y][x])
			}
		}
		left = bytes.ReplaceAll(left, []byte{'-'}, nil)
		left = bytes.ReplaceAll(left, []byte{'F', '7'}, nil)
		left = bytes.ReplaceAll(left, []byte{'L', 'J'}, nil)
		left = bytes.ReplaceAll(left, []byte{'S', '7'}, nil)
		left = bytes.ReplaceAll(left, []byte{'S', 'J'}, nil)
		left = bytes.ReplaceAll(left, []byte{'F', 'J'}, []byte{'|'})
		left = bytes.ReplaceAll(left, []byte{'F', 'S'}, []byte{'|'})
		left = bytes.ReplaceAll(left, []byte{'S', 'J'}, []byte{'|'})
		left = bytes.ReplaceAll(left, []byte{'S', '7'}, []byte{'|'})
		left = bytes.ReplaceAll(left, []byte{'L', '7'}, []byte{'|'})

		if len(left)%2 != 0 {
			inside++
		}
	}
	return inside
}
