package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.com/utils"
)

type Pos struct {
	Row int
	Col int
}

func convert(input string) []int {
	var out []int
	fields := strings.Fields(input)
	for _, field := range fields {
		out = append(out, utils.ToNumber(field))
	}
	return out
}

func blink(stones []int) []int {
	var out []int

	for _, stone := range stones {
		if stone == 0 {
			out = append(out, 1)
			continue
		}

		str := strconv.Itoa(stone)
		len := len(str)

		if len%2 == 0 {
			half := len / 2
			out = append(out, utils.ToNumber(str[:half]))
			out = append(out, utils.ToNumber(str[half:]))
			continue
		}

		out = append(out, stone*2024)
	}

	return out
}

func main() {
	stones := convert(utils.ReadFile("input")[0])

	for count := 0; count != 25; count++ {
		stones = blink(stones)
	}

	total := len(stones)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
