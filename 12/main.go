package main

import (
	"crypto/sha256"
	"log"
	"semvis123/aoc"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

func hash(x []int) string {
	sum := sha256.Sum256([]byte(strings.Join(aoc.Map(x, strconv.Itoa), ",")))
	return string(sum[:])
}

type CacheK = struct {
	s string
	g string
}

type CacheV = struct {
	x  int
	ok bool
}

var cache = expirable.NewLRU[CacheK, CacheV](50000, nil, time.Second*100)

func isPossible(s string, g int) bool {
	if len(s) < g {
		return false
	}
	for i := 0; i < g; i++ {
		if s[i] != '?' && s[i] != '#' {
			return false
		}
	}
	return len(s) == g || s[g] == '.' || s[g] == '?'
}

func arrangements(s string, g []int) (int, bool) {
	if len(s) == 0 && len(g) == 0 {
		return 0, true
	}

	if len(g) == 0 && strings.Count(s, "#") > 0 {
		return 0, false
	}

	if len(g) == 0 {
		return 0, true
	}

	if len(s) < g[0] {
		return 0, false
	}

	a := 0
	ap := false
	if isPossible(s, g[0]) {
		if len(s[g[0]:]) >= 1 {
			var x int
			key := CacheK{
				s: s[g[0]+1:],
				g: hash(g[1:]),
			}
			cacheVal, ok := cache.Get(key)
			if !ok {
				x, ap = arrangements(key.s, g[1:])
				cache.Add(key, CacheV{x: x, ok: ap})
			} else {
				x = cacheVal.x
				ap = cacheVal.ok
			}
			if ap {
				a = x
			}
		} else {
			a = 0
			ap = len(g) == 1
		}
	}
	b := 0
	bp := false
	if s[0] == '.' || s[0] == '?' {
		var bx int
		key := CacheK{
			s: s[1:],
			g: hash(g),
		}
		cacheVal, ok := cache.Get(key)
		if !ok {
			bx, bp = arrangements(key.s, g)
			cache.Add(key, CacheV{x: bx, ok: bp})
		} else {
			bx = cacheVal.x
			bp = cacheVal.ok
		}
		if bp {
			b += bx
		}
	}

	return a + b + aoc.Iff(ap && bp, 1, 0), ap || bp
}

func part_1(input []string) (s int) {
	for _, l := range input {
		springs, groups, _ := strings.Cut(l, " ")
		g := aoc.Map(strings.Split(groups, ","), aoc.WrapMust(strconv.Atoi))
		x, p := arrangements(springs, g)
		if !p {
			panic("not possible?")
		}
		s += x + 1
	}
	return
}

func part_2(input []string) (s int) {
	for _, l := range input {
		springs, groups, _ := strings.Cut(l, " ")
		springs = strings.Repeat(springs+"?", 5)
		groups = strings.Repeat(groups+",", 5)
		springs = springs[:len(springs)-1]
		groups = groups[:len(groups)-1]
		g := aoc.Map(strings.Split(groups, ","), aoc.WrapMust(strconv.Atoi))
		x, p := arrangements(springs, g)
		if !p {
			panic("not possible?")
		}
		s += x + 1
	}
	return
}
