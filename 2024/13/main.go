package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"aoc.com/utils"
)

type Pair struct {
	X int
	Y int
}

type Claws = [][]Pair

func parse(input []string) Claws {
	regex := regexp.MustCompile(`(\d+\b)`)

	var claws Claws
	pairs := []Pair{}
	for _, line := range input {
		if len(line) == 0 {
			claws = append(claws, pairs)
			pairs = []Pair{}
			continue
		}
		numbers := regex.FindAllString(line, -1)
		pairs = append(pairs, Pair{utils.ToNumber(numbers[0]), utils.ToNumber(numbers[1])})
	}

	return claws
}

func tokens(claw []Pair) int {
	goal := claw[2]
	a := claw[0]
	b := claw[1]

	A := -1
	B := -1
	T := math.MaxInt

	for u := 0; u != 101; u++ {
		for v := 0; v != 101; v++ {
			if u*a.X+v*b.X != goal.X {
				continue
			}

			if u*a.Y+v*b.Y != goal.Y {
				continue
			}

			cost := 3*u + v
			if cost < T {
				T = cost
				A = u
				B = v
			}
		}
	}

	if A != -1 && B != -1 {
		return T
	}

	return 0
}

func total(claws Claws) int {
	total := 0

	for _, claw := range claws {
		total += tokens(claw)
	}

	return total
}

func main() {
	claws := parse(utils.ReadFile("input"))

	total := total(claws)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
