package main

import (
	"fmt"
	"strconv"
	"regexp"

	"aoc.com/utils"
)

func multiply(input string) int {
	regex := regexp.MustCompile(`\d{1,3}`)
	numbers := regex.FindAllString(input, -1);
	return utils.ToNumber(numbers[0]) * utils.ToNumber(numbers[1])
}

func main() {
	lines := utils.ReadFile("input")

	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	sum := 0

	for _, line := range lines {
		hits := regex.FindAllString(line, -1);

		for _, hit := range hits {
			sum += multiply(hit)
		}
	}

	fmt.Println(strconv.FormatFloat(float64(sum), 'f', -1, 64))
}
