package main

import (
	"fmt"

	"aoc.com/utils"
)

func Convert(input string) []int {
	out := Init(input)

	index := 0
	pos := 0

	for i, el := range input {
		if i%2 == 0 {
			for j := 0; j != utils.ToNumber(string(el)); j++ {
				out[pos+j] = index
			}

			index++
		}

		pos += utils.ToNumber(string(el))
	}

	return out
}

func Size(input string) int {
	total := 0
	for _, item := range input {
		total += utils.ToNumber(string(item))
	}
	return total
}

func Init(input string) []int {
	out := make([]int, Size(input))
	for i := range out {
		out[i] = -1
	}
	return out
}

func FirstSpace(input []int) int {
	for i, v := range input {
		if v == -1 {
			return i
		}
	}
	return -1
}

func LastFile(input []int) int {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != -1 {
			return i
		}
	}
	return -1
}

func Move(input []int) []int {
	for {
		space := FirstSpace(input)
		if space == -1 {
			break
		}

		file := LastFile(input)
		if file < space {
			break
		}

		index := input[file]
		input = input[:file]
		input[space] = index
	}
	return input
}

func CheckSum(input []int) uint64 {
	sum := uint64(0)
	for i, v := range input {
		if v == -1 {
			continue
		}
		sum += uint64(i * v)
	}
	return uint64(sum)
}

func main() {
	total := CheckSum(Move(Convert(utils.ReadFile("input")[0])))

	fmt.Println(uint64(total))

}
