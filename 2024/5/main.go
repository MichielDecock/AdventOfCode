package main

import (
	"fmt"
	"strconv"
	"regexp"

	"aoc.com/utils"
)

func found(pattern string, order []string) bool {
	for _, orderPattern := range order {
		if (orderPattern == pattern) {
			return true
		}
	}

	return false
}

func main() {
	order := utils.ReadFile("order")
	updates := utils.ReadFile("updates")
	regex := regexp.MustCompile(`\d\d`)

	total := 0

	for _, update := range updates {
		hits := regex.FindAllString(update, -1)

		valid := true

		for i := 1; i < len(hits) && valid; i++ {
			for j := 0; j < i && valid; j++ {
				pattern := hits[j] + "|" + hits[i]
				if !found(pattern, order) {
					valid = false
				}
			}
		}

		if valid {
			middle := len(hits) / 2
			total += utils.ToNumber(hits[middle])
		}
	}

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
