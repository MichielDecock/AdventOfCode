package main

import (
	"fmt"
	"strconv"
	"regexp"

	"aoc.com/utils"
)

func Print(input map[int]string) {
	for row := 0; row != len(input); row++ {
		fmt.Println(input[row])
	}
}

func Find(input string) int {
	regex1 := regexp.MustCompile(`XMAS`)
	regex2 := regexp.MustCompile(`SAMX`)
	return len(regex1.FindAllString(input, -1)) + len(regex2.FindAllString(input, -1));
}

func FindInMap(input map[int]string) int {
	total := 0
	for i := 0; i != len(input); i++ {
		total += Find(input[i])
	}

	return total
}

func Flip(input map[int]string) map[int]string {
	out := make(map[int]string)

	for row := 0; row != len(input); row++ {
		line := input[row]
		for col :=0; col != len(line); col++ {
			out[col] += string(line[col])
		}
	}

	return out
}

func FindDiagonal(input map[int]string) int {
	total := 0

	for row := 0; row != len(input) - 3; row++ {
		part := make(map[int]string)
		
		for partRow := 0; partRow != 4; partRow++ {
			part[partRow] = input[row + partRow]
		}

		total += FindInMap(Flip(ShiftLeft(part)))
		total += FindInMap(Flip(ShiftLeft(inverse(part))))
	}

	return total
}

func ShiftLeft(input map[int]string) map[int]string {
	out := make(map[int]string)

	for row := 0; row != len(input); row++ {
		line := input[row]
		out[row] = line[row:]
	}

	return out
}

func inverse(input map[int]string) map[int]string {
	out := make(map[int]string)

	for row := 0; row != len(input); row++ {
		out[len(input) - row - 1] = input[row]
	}

	return out
}

func main() {
	lines := utils.ReadFile("input")

	input := make(map[int]string)

	for index, line := range lines {
		input[index] = line
	}

	total := FindInMap(Flip(input)) + FindInMap(input) + FindDiagonal(input)


	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
