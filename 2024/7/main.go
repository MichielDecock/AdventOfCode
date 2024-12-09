package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc.com/utils"
)

func calc(numbers int, input []string) int {
	local := utils.ToNumber(input[0])
	for i := 1; i < numbers; i++ {
		if input[2*i-1] == "+" {
			local += utils.ToNumber(input[2*i])
		} else {
			local *= utils.ToNumber(input[2*i])
		}
	}
	return local
}

func valid(result int, input []string) bool {
	operators := len(input) / 2
	limit := math.Pow(2, float64(operators))
	for i := 0; i != int(limit); i++ {
		for j := 0; j != operators; j++ {
			if i&int(math.Pow(float64(2), float64(j))) != 0 {
				input[2*j+1] = "*"
			} else {
				input[2*j+1] = "+"
			}
			if result == calc(operators+1, input) {
				return true
			}
		}
	}

	return false
}

func getExpression(input string) []string {
	input = string(input[1:])
	reworked := strings.Split(input, " ")

	limit := len(reworked)
	var out []string
	for i := 0; i < limit; i++ {
		out = append(out, string(reworked[i]))
		if i != limit-1 {
			out = append(out, "+")
		}
	}
	return out
}

func main() {
	input := utils.ReadFile("input")

	total := 0

	for _, line := range input {
		split := strings.Split(line, ":")

		result := utils.ToNumber(split[0])
		expression := getExpression(split[1])

		if valid(result, expression) {
			total += result
		}
	}

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))

}
