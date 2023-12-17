package main

import (
	"log"
	"semvis123/aoc"
	"strconv"

	"github.com/oleiade/lane/v2"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

type Node struct {
	pos  aoc.Coord2d[int]
	dir  aoc.Coord2d[int]
	dirN int
	cost int
}

type NodeV struct {
	pos  aoc.Coord2d[int]
	dir  aoc.Coord2d[int]
	dirN int
}

type NodeTV struct {
	pos  aoc.Coord2d[int]
	dir  aoc.Coord2d[int]
	dirN int
	cost int
}

func NewNodeTV(node Node) NodeTV {
	return NodeTV(node)
}

func NewNodeV(node Node) NodeV {
	return NodeV{node.pos, node.dir, node.dirN}
}

func (x Node) withinBounds(m []string) bool {
	return x.pos.Y >= 0 && x.pos.Y < len(m) && x.pos.X >= 0 && x.pos.X < len(m[0])
}

func (x Node) getChar(m []string) byte {
	return m[x.pos.Y][x.pos.X]
}

func (x Node) advance() Node {
	return Node{x.pos.Add(x.dir).(aoc.Coord2d[int]), x.dir, x.dirN + 1, x.cost}
}

func findPath(input []string, canKeepDir func(int) bool, mustKeepDir func(int) bool) (s int) {
	toVisit := lane.NewMinPriorityQueue[Node, int]()
	toVisitEntries := make(map[NodeTV]struct{})
	toVisit.Push(Node{pos: aoc.Coord2d[int]{X: 0, Y: 0}, dir: aoc.Coord2d[int]{X: 1, Y: 0}}, 0)
	toVisit.Push(Node{pos: aoc.Coord2d[int]{X: 0, Y: 0}, dir: aoc.Coord2d[int]{X: 0, Y: 1}}, 0)

	end := aoc.Coord2d[int]{X: len(input[0]) - 1, Y: len(input) - 1}
	var curr Node
	visited := make(map[NodeV]int)
	for toVisit.Size() > 0 {
		curr, _, _ = toVisit.Pop()
		delete(toVisitEntries, NewNodeTV(curr))
		currV := NewNodeV(curr)
		visited[currV] = curr.cost
		for _, dir := range aoc.Directions(false, false) {
			d := aoc.Coord2d[int]{X: dir[0], Y: dir[1]}
			if curr.dir.Add(d).(aoc.Coord2d[int]).IsZero() {
				continue
			}
			if curr.dir != d && mustKeepDir(curr.dirN) {
				continue
			}
			newNode := Node{
				pos:  curr.pos,
				dir:  d,
				dirN: aoc.Iff(d == curr.dir, curr.dirN, 0),
				cost: curr.cost,
			}
			newNode = newNode.advance()

			if !newNode.withinBounds(input) || !canKeepDir(newNode.dirN) {
				continue
			}
			newNode.cost += aoc.Must(strconv.Atoi(string(newNode.getChar(input))))
			newNodeV := NewNodeV(newNode)

			if newNode.pos == end {
				return newNode.cost
			}

			if c, ok := visited[newNodeV]; !ok || c > newNode.cost {
				newNodeTV := NewNodeTV(newNode)
				if _, ok := toVisitEntries[newNodeTV]; !ok {
					toVisitEntries[newNodeTV] = struct{}{}
					toVisit.Push(newNode, newNode.cost+newNode.pos.Manhatten(end))
				}
			}
		}
	}

	return
}

func part_1(input []string) (s int) {
	return findPath(input, func(x int) bool { return x <= 3 }, func(x int) bool { return false })
}

func part_2(input []string) (s int) {
	return findPath(input, func(x int) bool { return x <= 10 }, func(x int) bool { return x < 4 })
}
