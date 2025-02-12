package main

import (
	"fmt"
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
	goal.X += 10000000000000
	goal.Y += 10000000000000

	a := claw[0]
	b := claw[1]

	n1 := goal.X*a.Y - goal.Y*a.X
	n2 := b.X*a.Y - b.Y*a.X

	if n2 == 0 {
		fmt.Println("Division by zero!")
		return 0
	}

	if n1%n2 != 0 {
		return 0
	}

	v := n1 / n2
	if (goal.X-v*b.X)%a.X != 0 {
		return 0
	}

	u := (goal.X - v*b.X) / a.X

	return 3*u + v
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
