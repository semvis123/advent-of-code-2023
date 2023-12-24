package main

import (
	"log"
	"semvis123/aoc"

	"fmt"
)

func main() {
	lines := aoc.GetInput()
	log.Printf("part 1: %d", part_1(lines[:len(lines)-1]))
	log.Printf("part 2: %d", part_2(lines[:len(lines)-1]))
}

type Hail struct {
	posX, posY, posZ, velX, velY, velZ float64
}

func part_1(input []string) (s int) {
	var hails []Hail
	for _, l := range input {
		hail := Hail{}
		fmt.Sscanf(l, "%f, %f, %f @ %f, %f, %f", &hail.posX, &hail.posY, &hail.posZ, &hail.velX, &hail.velY, &hail.velZ)
		log.Printf("hail: %v", hail)
		hails = append(hails, hail)
	}

	for i, a := range hails {
		for _, b := range hails[i:] {
			if a == b {
				continue
			}

			// <math>
			if a.velY*b.velX-a.velX*b.velY == 0 || a.velX*b.velX == 0 {
				continue
			}
			x := (-a.posY*a.velX*b.velX + a.posX*a.velY*b.velX + a.velX*b.posY*b.velX - a.velX*b.posX*b.velY) / (a.velY*b.velX - a.velX*b.velY)
			y := a.posY + (a.velY/a.velX)*(x-a.posX)

			past := false
			if (x-a.posX)/a.velX < 0 || (y-a.posY)/a.velY < 0 {
				past = true
			}
			if (x-b.posX)/b.velX < 0 || (y-b.posY)/b.velY < 0 {
				past = true
			}
			// </math>

			lowerB := float64(200000000000000)
			upperB := float64(400000000000000)
			inside := x >= lowerB && x <= upperB && y >= lowerB && y <= upperB
			if inside && !past {
				log.Printf("%v %v %v %v", x, y, a, b)
				s++
			}
		}
	}

	return
}

func part_2(input []string) (s int) {

	return
}
