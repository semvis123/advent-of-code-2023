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
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

type Block = aoc.Line[int, aoc.Coord3d[int]]

func getFallenBlocks(input []string) (blocks []Block) {
	for _, l := range input {
		block := Block{}
		c1, c2, _ := strings.Cut(l, "~")
		ci1 := aoc.Map(strings.Split(c1, ","), aoc.WrapMust(strconv.Atoi))
		ci2 := aoc.Map(strings.Split(c2, ","), aoc.WrapMust(strconv.Atoi))
		block.Start.X = ci1[0]
		block.Start.Y = ci1[1]
		block.Start.Z = ci1[2]
		block.End.X = ci2[0]
		block.End.Y = ci2[1]
		block.End.Z = ci2[2]
		blocks = append(blocks, block)
	}

	slices.SortFunc(blocks, func(a Block, b Block) int {
		return a.Start.Z - b.Start.Z
	})

	_, blocks = countMoved(Block{}, blocks)

	return blocks
}

func countMoved(removeB Block, blocksO []Block) (int, []Block) {
	heightMap := make(map[aoc.Coord2d[int]]int)
	blocks := slices.Clone(blocksO)
	modified := 0
	for i := range blocks {
		if blocks[i] == removeB {
			continue
		}
		start := aoc.Coord2d[int]{X: blocks[i].Start.X, Y: blocks[i].Start.Y}
		end := aoc.Coord2d[int]{X: blocks[i].End.X, Y: blocks[i].End.Y}
		maxHeight := 1
		for _, step := range start.StepsTo(end) {
			h := heightMap[step]
			maxHeight = aoc.MaxInt(maxHeight, h)
		}
		for _, step := range start.StepsTo(end) {
			heightMap[step] = maxHeight + blocks[i].End.Z - blocks[i].Start.Z + 1
		}
		before := blocks[i].Start.Z
		blocks[i].End.Z = maxHeight + aoc.AbsInt(blocks[i].End.Z-blocks[i].Start.Z)
		blocks[i].Start.Z = maxHeight
		if before != blocks[i].Start.Z {
			modified += 1
		}
	}
	return modified, blocks
}

func part_1(input []string) (s int) {
	blocks := getFallenBlocks(input)

	for _, b := range blocks {
		if moved, _ := countMoved(b, blocks); moved == 0 {
			s++
		}
	}

	return
}

func part_2(input []string) (s int) {
	blocks := getFallenBlocks(input)

	for _, b := range blocks {
		m, _ := countMoved(b, blocks)
		s += m
	}

	return
}
