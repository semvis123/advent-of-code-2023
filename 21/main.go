package main

import (
	"golang.org/x/exp/slices"
	"log"
	"semvis123/aoc"
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
	visited := make(map[Node]struct{})
	endPositions := make(map[aoc.Coord2d[int]]struct{})
	for len(toVisit) > 0 {
		curr, toVisit = toVisit[0], toVisit[1:]
		visited[curr] = struct{}{}
		for _, d := range aoc.Directions(false, false) {
			newPos := curr.pos.Add(aoc.Coord2d[int]{X: d[0], Y: d[1]}).(aoc.Coord2d[int])
			newNode := Node{newPos, curr.steps + 1}
			if !newNode.withinBounds(input) {
				continue
			}
			if newNode.getChar(input) == '#' {
				continue
			}
			if newNode.steps == 64 {
				endPositions[newNode.pos] = struct{}{}
			}
			if newNode.steps > 64 {
				continue
			}
			if _, ok := visited[newNode]; !ok && !slices.Contains(toVisit, newNode) {
				toVisit = append(toVisit, newNode)
			}
		}
	}
	return len(endPositions)
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
	visited := make(map[Node]struct{})
	endPositions := make(map[aoc.Coord2d[int]]struct{})
	for len(toVisit) > 0 {
		curr, toVisit = toVisit[0], toVisit[1:]
		visited[curr] = struct{}{}
		for _, d := range aoc.Directions(false, false) {
			newPos := curr.pos.Add(aoc.Coord2d[int]{X: d[0], Y: d[1]}).(aoc.Coord2d[int])
			newNode := Node{newPos, curr.steps + 1}
			if !newNode.withinBounds(input) {
				continue
			}
			if newNode.getChar(input) == '#' {
				continue
			}
			if newNode.steps == 1000 {
				endPositions[newNode.pos] = struct{}{}
			}
			if newNode.steps > 1000 {
				continue
			}
			if _, ok := visited[newNode]; !ok && !slices.Contains(toVisit, newNode) {
				toVisit = append(toVisit, newNode)
			}
		}
	}
	return len(endPositions)
}
