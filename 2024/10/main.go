package main

import (
	"fmt"
	"math"
	"strconv"

	"aoc.com/utils"
)

type Pos struct {
	Row int
	Col int
}

func findPath(number int, pos Pos, plan [][]int, found *[][]bool) int {
	if plan[pos.Row][pos.Col] != number {
		return 0
	}

	if plan[pos.Row][pos.Col] == 9 {
		return 1
	}

	paths := 0

	for row := int(math.Max(0, float64(pos.Row-1))); row != int(math.Min(float64(pos.Row+2), float64(len(plan)))); row++ {
		for col := int(math.Max(0, float64(pos.Col-1))); col != int(math.Min(float64(pos.Col+2), float64(len(plan[0])))); col++ {
			if row == pos.Row && col == pos.Col {
				continue
			}
			if math.Abs(float64(row-pos.Row)) == 1 && math.Abs(float64(col-pos.Col)) == 1 {
				continue
			}
			paths += findPath(number+1, Pos{Row: row, Col: col}, plan, found)
		}
	}

	return paths
}

func convert(plan []string) [][]int {
	var matrix [][]int
	for _, str := range plan {
		var row []int

		for _, elem := range str {
			if elem == '.' {
				row = append(row, -1)
			} else {
				row = append(row, utils.ToNumber(string(elem)))
			}
		}

		matrix = append(matrix, row)
	}
	return matrix
}

func initFound(plan [][]int) [][]bool {
	found := make([][]bool, len(plan))
	for i := range found {
		found[i] = make([]bool, len(plan[0]))
	}
	return found
}

func main() {
	input := convert(utils.ReadFile("input"))

	total := 0

	for row, line := range input {
		for col := range line {
			found := initFound(input)
			res := findPath(0, Pos{Row: row, Col: col}, input, &found)
			total += res
		}
	}

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))

}
