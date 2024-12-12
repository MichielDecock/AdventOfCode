package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.com/utils"
)

func convert(input string) map[int]int {
	out := make(map[int]int)
	fields := strings.Fields(input)
	for _, field := range fields {
		add(&out, utils.ToNumber(field), 1)
	}
	return out
}

func add(out *map[int]int, stone int, count int) {
	_, ok := (*out)[stone]
	if !ok {
		(*out)[stone] = count
	} else {
		(*out)[stone] += count
	}
}

func blink(stones map[int]int) map[int]int {
	out := make(map[int]int)

	for stone, count := range stones {
		if stone == 0 {
			add(&out, 1, count)
			continue
		}

		str := strconv.Itoa(stone)
		len := len(str)

		if len%2 == 0 {
			half := len / 2
			add(&out, utils.ToNumber(str[:half]), count)
			add(&out, utils.ToNumber(str[half:]), count)
			continue
		}

		add(&out, 2024*stone, count)
	}

	return out
}

func total(stones map[int]int) int {
	total := 0
	for _, count := range stones {
		total += count
	}
	return total
}

func main() {
	stones := convert(utils.ReadFile("input")[0])

	for count := 0; count != 75; count++ {
		stones = blink(stones)
	}

	total := total(stones)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
