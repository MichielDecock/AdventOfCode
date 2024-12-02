package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.com/utils"
)

func main() {
	const minDiff = 1
	const maxDiff = 3

	lines := utils.ReadFile("input")

	validReports := 0

	for _, line := range lines {
		fields := strings.Fields(line)

		var increasing bool
		valid := true

		for i := 0; i != len(fields) - 1; i++ {
			number1 := utils.ToNumber(fields[i])
			number2 := utils.ToNumber(fields[i + 1])

			diff := number1 - number2

			if (i == 0) {
				increasing = (diff >= 0)
			}

			if (increasing) {
				if !(diff >= minDiff && diff <= maxDiff) {
					valid = false
				}
			} else {
				if !(diff <= -minDiff && diff >= -maxDiff) {
					valid = false
				}
			}
		}

		if valid {
			validReports++
		}
	}

	// sum := 0
	// for _, number := range list1 {
	// 	sum += number * unique2[number]
	// }

	fmt.Println(strconv.FormatFloat(float64(validReports), 'f', -1, 64))
}
