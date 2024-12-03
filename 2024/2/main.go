package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.com/utils"
)

func check(subfield[]string) bool {
	const minDiff = 1
	const maxDiff = 3

	var increasing bool

	for i := 0; i != len(subfield) - 1; i++ {
		number1 := utils.ToNumber(subfield[i])
		number2 := utils.ToNumber(subfield[i + 1])

		diff := number1 - number2

		if (i == 0) {
			increasing = (diff >= 0)
		}

		if (increasing) {
			if !(diff >= minDiff && diff <= maxDiff) {
				return false
			}
		} else {
			if !(diff <= -minDiff && diff >= -maxDiff) {
				return false
			}
		}
	}

	return true
}

func getSubfields(fields []string)  map[int][]string {
	subfields := make(map[int][]string)
	for i := 0; i != len(fields); i++ {
		var subfield []string
		for j := 0; j != len(fields); j++ {
			if j != i {
				subfield = append(subfield, fields[j])
			}
		}
		subfields[i] = subfield
	}

	return subfields
}

func main() {
	lines := utils.ReadFile("input")
	validReports := 0

	for _, line := range lines {
		subfields := getSubfields(strings.Fields(line))

		for i := 0; i != len(subfields); i++ {
			if check(subfields[i]) {
				validReports++
				break
			}
		}
	}

	fmt.Println(strconv.FormatFloat(float64(validReports), 'f', -1, 64))
}
