package aoc

import (
	"flag"
	"math"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

var testing bool

func GetInput() []string {
	if !flag.Parsed() {
		flag.BoolVar(&testing, "test", false, "Use test file")
		flag.Parse()
	}
	filename := Iff(testing, "test.txt", "input.txt")
	file := Must(os.ReadFile(filename))
	lines := strings.Split(string(file), "\n")
	return lines
}

func SplitByEmpty(input []string) (output [][]string) {
	var temp []string
	for _, l := range input {
		if len(l) == 0 {
			output = append(output, temp)
			temp = nil
		} else {
			temp = append(temp, l)
		}
	}
	if len(temp) > 0 {
		output = append(output, temp)
	}
	return output
}

func Iff[T any](check bool, ifV T, elseV T) T {
	if check {
		return ifV
	}
	return elseV
}

func Count[T constraints.Ordered](slice []T, el T) int {
	var count int
	for _, s := range slice {
		if s == el {
			count++
		}
	}
	return count
}

func AllV[T constraints.Ordered](slice []T, el T) bool {
	for _, s := range slice {
		if s != el {
			return false
		}
	}
	return true
}

func ReverseCopy[T any](array []T) []T {
	length := len(array)
	result := make([]T, length)
	for i, elem := range array {
		result[length-1-i] = elem
	}
	return result
}

func Transpose[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func TransposeStrs(s []string) []string {
	return Map(
		Transpose(
			Map(s, func(s string) []string {
				return strings.Split(s, "")
			}),
		), func(s []string) string {
			return strings.Join(s, "")
		},
	)
}

func CountFunc[T constraints.Ordered](slice []T, f func(T) bool) int {
	var count int
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
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

func Wrap[A any, B any, C any](a func(A) B, b func(C) A) func(C) B {
	return func(c C) B {
		return a(b(c))
	}
}

func Chunks[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func WrapMust[A any, B any, T func(B) (A, error)](f T) func(B) A {
	return func(param B) A {
		return Must(f(param))
	}
}

func NoErr[T any](v T, err error) T {
	return v
}

func Filter[A any](items []A, f func(A) bool) []A {
	var result []A
	for _, v := range items {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func FilterI[A any](items []A, f func(A, int) bool) []A {
	var result []A
	for i, v := range items {
		if f(v, i) {
			result = append(result, v)
		}
	}

	return result
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func AbsInt(x int) int {
	return int(math.Abs(float64(x)))
}

func MaxInt(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func NotEmpty(v string) bool {
	return len(v) > 0
}

func CastNum[A, B constraints.Integer](x A) B {
	return B(x)
}

func Map[A any, B any](items []A, f func(A) B) []B {
	var result []B
	for _, v := range items {
		result = append(result, f(v))
	}

	return result
}

func Flatten[A any](items [][]A) (result []A) {
	for _, v := range items {
		result = append(result, v...)
	}
	return
}

func FlatMap[A any, B any](items []A, f func(A) []B) []B {
	return Flatten(Map(items, f))
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

type number interface {
	constraints.Integer | constraints.Float
}

func Abs[T number](x T) T {
	return T(math.Abs(float64(x)))
}

func Pow[T number](x T, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Sqrt[T number](x T) T {
	return T(math.Sqrt(float64(x)))
}

func PowSqrt[T number](nums ...T) T {
	var s T
	for _, n := range nums {
		s += Pow(n, 2)
	}
	return Sqrt(s)
}

func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func RemoveIndex[T any](s []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func LCM[T constraints.Integer](a, b T, integers ...T) T {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type Coord[T number] interface {
	Euclidean(Coord[T]) T
}

type Line[T number, C Coord[T]] struct {
	Start, End C
}

func (l Line[T, C]) Length() T {
	return l.Start.Euclidean(l.End)
}

type line1d[T number] struct {
	Line[T, Coord1d[T]]
}

func (a line1d[T]) OverlapAndDifference(b line1d[T]) (overlap line1d[T], difference []line1d[T]) {
	// if b.Start.X < a.End.X && b.End.X > a.Start.X {
	// 	// overlap
	// 	excluded := T(0)
	// 	if a.Start.X < b.End.X {
	// 		// before
	// 		difference = append(difference, line1d[T]{
	// 			Line: Line[T, Coord1d[T]]{
	// 				Coord1d[T]{a.Start.X}, Coord1d[T]{b.Start.X},
	// 			},
	// 		},
	// 		)
	// 		excluded += b.Start.X - a.Start.X
	// 	}
	// 	if b.End.X <= a.End.X {
	// 		// after
	// 		difference = append(difference, []int{b.End.X, a.End.X})
	// 		excluded += end - (m.Src + m.Range)
	// 	}
	// 	changed = []int{m.Dest + aoc.MaxInt(0, start-m.Src), x[1] - excluded}
	// } else {
	// 	difference = [][]int{x}
	// }
	return
}

type Coord1d[T number] struct {
	X T
}

type Coord2d[T number] struct {
	X, Y T
}

type Coord3d[T number] struct {
	X, Y, Z T
}

func (a Coord1d[T]) Add(b Coord[T]) Coord[T] {
	B := b.(Coord1d[T])
	return Coord1d[T]{
		X: a.X + B.X,
	}
}

func (a Coord2d[T]) Add(b Coord[T]) Coord[T] {
	B := b.(Coord2d[T])
	return Coord2d[T]{
		X: a.X + B.X,
		Y: a.Y + B.Y,
	}
}

func (a Coord3d[T]) Add(b Coord[T]) Coord[T] {
	B := b.(Coord3d[T])
	return Coord3d[T]{
		X: a.X + B.X,
		Y: a.Y + B.Y,
		Z: a.Z + B.Z,
	}
}

func (a Coord1d[T]) Substract(b Coord[T]) Coord[T] {
	B := b.(Coord1d[T])
	return Coord1d[T]{
		X: a.X - B.X,
	}
}

func (a Coord2d[T]) Substract(b Coord[T]) Coord[T] {
	B := b.(Coord2d[T])
	return Coord2d[T]{
		X: a.X - B.X,
		Y: a.Y - B.Y,
	}
}

func (a Coord3d[T]) Substract(b Coord[T]) Coord[T] {
	B := b.(Coord3d[T])
	return Coord3d[T]{
		X: a.X - B.X,
		Y: a.Y - B.Y,
		Z: a.Z - B.Z,
	}
}

func (a Coord1d[T]) Euclidean(b Coord[T]) T {
	return PowSqrt(a.X - b.(Coord1d[T]).X)
}

func (a Coord2d[T]) Euclidean(b Coord[T]) T {
	return PowSqrt(a.X-b.(Coord2d[T]).X, a.Y-b.(Coord2d[T]).Y)
}

func (a Coord3d[T]) Euclidean(b Coord[T]) T {
	return PowSqrt(a.X-b.(Coord3d[T]).X, a.Y-b.(Coord3d[T]).Y, a.Z-b.(Coord3d[T]).Z)
}

func (a Coord1d[T]) Manhatten(b Coord[T]) T {
	return Abs(a.X - b.(Coord1d[T]).X)
}

func (a Coord2d[T]) Manhatten(b Coord[T]) T {
	return Abs(a.X-b.(Coord2d[T]).X) + Abs(a.Y-b.(Coord2d[T]).Y)
}

func (a Coord3d[T]) Manhatten(b Coord[T]) T {
	return Abs(a.X-b.(Coord3d[T]).X) + Abs(a.Y-b.(Coord3d[T]).Y) + Abs(a.Z-b.(Coord3d[T]).Z)
}
