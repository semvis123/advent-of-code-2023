package aoc

import (
	"flag"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

func GetInput() []string {
	testFlag := flag.Bool("test", false, "Use test file")
	flag.Parse()
	filename := Iff(*testFlag, "test.txt", "input.txt")
	file := Must(os.ReadFile(filename))
	lines := strings.Split(string(file), "\n")
	return lines
}

func Iff[T any](check bool, ifV T, elseV T) T {
	if check {
		return ifV
	}
	return elseV
}

type summable interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Sum[T summable](slice []T) T {
	var total T
	for _, x := range slice {
		total += x
	}
	return total
}

func Unpack[T any](s []T, vars ...*T) {
	for i, str := range s {
		*vars[i] = str
	}
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func NoErr[T any](v T, err error) T {
	return v
}

func Map[A any, B any](items []A, f func(A) B) []B {
	var result []B
	for _, v := range items {
		result = append(result, f(v))
	}

	return result
}

func Reduce[T any](items []T, f func(curr T, acc T) T) T {
	result := items[0]
	for _, x := range items[1:] {
		result = f(x, result)
	}
	return result
}

func Multiply[T summable](items []T) T {
	return Reduce(items, func(x T, acc T) T {
		return acc * x
	})
}

func Any(items []bool) bool {
	result := false
	for _, x := range items {
		result = result || x
	}
	return result
}

func All(items []bool) bool {
	result := true
	for _, x := range items {
		result = result && x
	}
	return result
}

func Directions(diagonal bool, center bool) [][]int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	if diagonal {
		dirs = append(dirs, [][]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}...)
	}
	if center {
		dirs = append(dirs, []int{0, 0})
	}
	return dirs
}
