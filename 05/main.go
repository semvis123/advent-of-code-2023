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

type Mapping struct {
	Src   int
	Dest  int
	Range int
}

func (m Mapping) Apply(x int) (int, bool) {
	d := x - m.Src
	if d >= 0 && d < m.Range {
		return m.Dest + d, true
	}
	return x, false
}

func (m Mapping) ApplyRange(x []int) (unchanged [][]int, changed []int) {
	start, end := x[0], x[0]+x[1]-1
	if m.Src < end && m.Src+m.Range > start {
		// overlap
		excluded := 0
		if start < m.Src {
			// before
			unchanged = append(unchanged, []int{start, m.Src - start})
			excluded += m.Src - start
		}
		if m.Src+m.Range <= end {
			// after
			unchanged = append(unchanged, []int{m.Src + m.Range, end - (m.Src + m.Range)})
			excluded += end - (m.Src + m.Range)
		}
		changed = []int{m.Dest + aoc.MaxInt(0, start-m.Src), x[1] - excluded}
	} else {
		unchanged = [][]int{x}
	}
	return
}

type MappingGroup struct {
	Name     string
	Mappings []Mapping
}

func readMappings(input []string) (seeds []int, mappings []MappingGroup) {
	seeds = aoc.Map(strings.Split(input[0][7:], " "), aoc.WrapMust(strconv.Atoi))
	for i := 2; i < len(input); i++ {
		mg := MappingGroup{Name: input[i]}
		for i++; input[i] != ""; i++ {
			var Src, Dest, Rang int
			aoc.Unpack(aoc.Map(strings.Split(input[i], " "), aoc.WrapMust(strconv.Atoi)), &Dest, &Src, &Rang)
			mg.Mappings = append(mg.Mappings, Mapping{Src, Dest, Rang})
		}
		mappings = append(mappings, mg)
	}
	return
}

func part_1(input []string) int {
	seeds, mappings := readMappings(input)
	for _, mg := range mappings {
		applied := make([]bool, len(input))
		for _, m := range mg.Mappings {
			for i := range seeds {
				if applied[i] {
					continue
				}
				seeds[i], applied[i] = m.Apply(seeds[i])
			}
		}
	}
	return slices.Min(seeds)
}

func part_2(input []string) int {
	seedsA, mappings := readMappings(input)
	seeds := aoc.Chunks(seedsA, 2)
	for _, mg := range mappings {
		toChange := seeds
		changed := [][]int{}
		for _, m := range mg.Mappings {
			notChanged := [][]int{}
			for _, seedRange := range toChange {
				orig, newSeeds := m.ApplyRange(seedRange)
				notChanged = append(notChanged, orig...)
				if len(newSeeds) != 0 {
					changed = append(changed, newSeeds)
				}
			}
			toChange = notChanged
		}
		seeds = append(toChange, changed...)
	}
	return slices.Min(aoc.Map(seeds, func(x []int) int { return x[0] }))
}
