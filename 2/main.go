package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	filename := "input.txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	lines := strings.Split(string(file), "\n")

	log.Printf("part 1: %d", part_1(lines))
	log.Printf("part 2: %d", part_2(lines))
}

type Game struct {
	Id      int
	Reveals []map[string]int
}

func readGames(input []string) (games []Game) {
	r, _ := regexp.Compile(`Game (\d+): (.*)`)
	for _, line := range input {
		matches := r.FindSubmatch([]byte(line))
		if len(matches) == 0 {
			continue
		}
		id, _ := strconv.Atoi(string(matches[1]))
		game := Game{Id: id}

		reveals := bytes.Split(matches[2], []byte("; "))
		for _, revealStr := range reveals {
			colors := bytes.Split(revealStr, []byte(", "))
			reveal := make(map[string]int)
			for _, color := range colors {
				splitted := strings.Split(string(color), " ")
				amount, _ := strconv.Atoi(splitted[0])
				colorName := splitted[1]
				reveal[colorName] = amount
			}
			game.Reveals = append(game.Reveals, reveal)
		}
		games = append(games, game)
	}
	return
}

func part_1(input []string) int {
	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	games := readGames(input)
	sum := 0
	for _, game := range games {
		possible := true
		for color, amount := range bag {
			for _, reveal := range game.Reveals {
				if reveal[color] > amount {
					possible = false
				}
			}
		}
		if possible {
			sum += game.Id
		}
	}
	return sum
}

func part_2(input []string) int {
	games := readGames(input)
	sum := 0
	for _, game := range games {
		bag := make(map[string]int)
		for _, reveal := range game.Reveals {
			for color, amount := range reveal {
				if bag[color] < amount {
					bag[color] = amount
				}
			}
		}
		values := maps.Values(bag)
		power := values[0]
		for _, v := range values[1:] {
			power *= v
		}
		sum += power
	}
	return sum
}
