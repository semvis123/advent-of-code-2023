package main

import (
	"log"
	"semvis123/aoc"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

type Rule struct {
	name     *string
	operator *byte
	val      *int
	dest     string
}

func (r Rule) apply(part Part) *string {
	if r.name == nil {
		return &r.dest
	}
	if *r.operator == '>' {
		if part[*r.name] > *r.val {
			return &r.dest
		}
	}
	if *r.operator == '<' {
		if part[*r.name] < *r.val {
			return &r.dest
		}
	}
	return nil
}

type workflow struct {
	rules []Rule
}

func ratingNameToIdx(n string) int {
	switch n {
	case "x":
		return 0
	case "m":
		return 1
	case "a":
		return 2
	case "s":
		return 3
	}
	panic("invalid part rating")
}

func combinationsFromRanges(r [][]int) int {
	return (r[0][1] - r[0][0] + 1) *
		(r[1][1] - r[1][0] + 1) *
		(r[2][1] - r[2][0] + 1) *
		(r[3][1] - r[3][0] + 1)
}

func (r Rule) getCombinations(ra [][]int, workflows map[string]workflow) (int, [][]int) {
	ra = slices.Clone(ra)
	nRa := slices.Clone(ra)

	if r.name != nil {
		i := ratingNameToIdx(*r.name)
		if *r.val <= ra[i][0] || *r.val >= ra[i][1] {
			return 0, ra
		}
		if *r.operator == '<' {
			ra[i] = []int{ra[i][0], *r.val - 1}
			nRa[i] = []int{*r.val, nRa[i][1]}
		}
		if *r.operator == '>' {
			ra[i] = []int{*r.val + 1, ra[i][1]}
			nRa[i] = []int{nRa[i][0], *r.val}
		}
	}

	if r.dest == "R" {
		return 0, nRa
	}
	if r.dest == "A" {
		return combinationsFromRanges(ra), nRa
	}

	x := workflows[r.dest].getCombinations(ra, workflows)
	return x, nRa
}

func (w workflow) apply(part Part) string {
	for _, r := range w.rules {
		result := r.apply(part)
		if result != nil {
			return *result
		}
	}
	panic("invalid workflow?")
}

func (w workflow) getCombinations(ra [][]int, workflows map[string]workflow) int {
	s := 0
	for _, x := range w.rules {
		st, nRa := x.getCombinations(ra, workflows)
		ra = nRa
		s += st
	}
	return s
}

type Part = map[string]int

func parse(input []string) (map[string]workflow, []Part) {
	workflowStrs, partStrs := aoc.Unpack2(aoc.SplitByEmpty(input))
	workflows := make(map[string]workflow)

	for _, w := range workflowStrs {
		name, ruleStr, _ := strings.Cut(w, "{")
		rs := strings.Split(ruleStr[:len(ruleStr)-1], ",")
		workflow := workflow{}
		for _, r := range rs {
			cond, dest, ok := strings.Cut(r, ":")
			if !ok {
				workflow.rules = append(workflow.rules, Rule{nil, nil, nil, r})
				continue
			}
			rule := Rule{}
			rule.dest = dest
			if n, val, ok := strings.Cut(cond, "<"); ok {
				valI, _ := strconv.Atoi(val)
				rule.val = &valI
				rule.name = &n
				rule.operator = aoc.Ptr(byte('<'))
			}
			if n, val, ok := strings.Cut(cond, ">"); ok {
				valI, _ := strconv.Atoi(val)
				rule.val = &valI
				rule.name = &n
				rule.operator = aoc.Ptr(byte('>'))
			}
			workflow.rules = append(workflow.rules, rule)
		}
		workflows[name] = workflow
	}

	parts := []Part{}
	for _, p := range partStrs {
		ratings := strings.Split(p[1:len(p)-1], ",")
		part := make(Part)
		for _, r := range ratings {
			part[string(r[0])], _ = strconv.Atoi(r[2:])
		}
		parts = append(parts, part)
	}
	return workflows, parts
}

func part_1(input []string) (s int) {
	workflows, parts := parse(input)
	a := 0
	for _, p := range parts {
		curr := "in"
		for curr != "A" && curr != "R" {
			curr = workflows[curr].apply(p)
		}
		if curr == "A" {
			x := aoc.Sum(maps.Values(p))
			a += x
		}
	}

	return a
}

func part_2(input []string) (s int) {
	workflows, _ := parse(input)

	return workflows["in"].getCombinations([][]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}, workflows)
}
