package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.com/utils"
)

func uniqueValues(array []int) map[int]int {
	dict := make(map[int]int)

	for _, number := range array {
		dict[number]++
	}

	return dict
}

func main() {
	lines := utils.ReadFile("input")

	var list1 []int
	var list2 []int

	for _, line := range lines {
		fields := strings.Fields(line)

		for index, field := range fields {
			number := utils.ToNumber(field)
			if index % 2 == 0 {
				list1 = append(list1, number)
			} else {
				list2 = append(list2, number)
			}
		}
	}

	unique2 := uniqueValues(list2)

	sum := 0
	for _, number := range list1 {
		sum += number * unique2[number]
	}

	fmt.Println(strconv.FormatFloat(float64(sum), 'f', -1, 64))
}
