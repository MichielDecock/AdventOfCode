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

type Region struct {
	Members [][]Pos
}

type Regions = map[string]Region
type Lines = []string

func contains(positions []Pos, position Pos) bool {
	for _, pos := range positions {
		if pos.Row == position.Row && pos.Col == position.Col {
			return true
		}
	}
	return false
}

func delete(members [][]Pos, index int) [][]Pos {
	if index < len(members)-1 {
		return append(members[:index], members[index+1:]...)
	}

	return members[:index]
}

func addToRegion(symbol string, pos Pos, regions *Regions) {
	newRow := pos.Row - 1
	newCol := pos.Col - 1

	region, ok := (*regions)[symbol]
	if ok {
		left := -1
		up := -1
		members := (*regions)[symbol].Members
		for i := range members {
			if left == -1 && newCol >= 0 && contains(members[i], Pos{pos.Row, newCol}) {
				left = i

				if up != -1 {
					members[i] = append(members[i], members[up]...)
					members = delete(members, up)
					break

				} else {
					members[i] = append(members[i], pos)
				}

				continue
			}

			if up == -1 && newRow >= 0 && contains(members[i], Pos{newRow, pos.Col}) {
				up = i

				if left != -1 {
					members[i] = append(members[i], members[left]...)
					members = delete(members, left)
					break
				} else {
					members[i] = append(members[i], pos)
				}
			}
		}

		if left == -1 && up == -1 {
			members = append(members, []Pos{pos})
		}

		region.Members = members
		(*regions)[symbol] = region
	} else {
		(*regions)[symbol] = Region{[][]Pos{{pos}}}
	}
}

func convert(lines Lines) Regions {
	regions := make(Regions)
	for row, line := range lines {
		for col, symbol := range line {
			addToRegion(string(symbol), Pos{row, col}, &regions)
		}
	}
	return regions
}

func perimeter(positions []Pos) int {
	perimeter := 0

	for _, tile := range positions {
		neighbors := []Pos{{tile.Row - 1, tile.Col}, {tile.Row + 1, tile.Col}, {tile.Row, tile.Col - 1}, {tile.Row, tile.Col + 1}}
		for _, neighbor := range neighbors {
			if !contains(positions, neighbor) {
				perimeter++
			}
		}
	}

	return perimeter
}

func total(regions Regions) int {
	total := 0
	for _, region := range regions {
		for _, member := range region.Members {
			area := len(member)
			sides := sides(member)
			total += area * sides
		}
	}
	return total
}

func boundingBox(positions []Pos) (Pos, Pos) {
	min := Pos{math.MaxInt, math.MaxInt}
	max := Pos{0, 0}

	for _, pos := range positions {
		min.Row = int(math.Min(float64(min.Row), float64(pos.Row)))
		max.Row = int(math.Max(float64(max.Row), float64(pos.Row)))
		min.Col = int(math.Min(float64(min.Col), float64(pos.Col)))
		max.Col = int(math.Max(float64(max.Col), float64(pos.Col)))
	}

	return min, max
}

func initSides(rows int, cols int) [][]bool {
	sides := make([][]bool, rows)
	for i := range rows {
		sides[i] = make([]bool, cols)
	}
	return sides
}

func distinctHorizontal(fences [][]bool) int {
	sides := 0

	for _, row := range fences {
		onFence := false

		for _, fence := range row {
			if fence {
				if !onFence {
					onFence = true
					sides++
				}
			} else {
				onFence = false
			}
		}
	}

	return sides
}

func distinctVertical(fences [][]bool) int {
	sides := 0

	for col := range fences[0] {
		onFence := false

		for row := range fences {
			if fences[row][col] {
				if !onFence {
					onFence = true
					sides++
				}
			} else {
				onFence = false
			}
		}
	}

	return sides
}

func sides(positions []Pos) int {
	min, max := boundingBox(positions)

	spanCol := max.Col - min.Col + 1
	spanRow := max.Row - min.Row + 1

	up := initSides(spanRow, spanCol)
	down := initSides(spanRow, spanCol)
	left := initSides(spanRow, spanCol)
	right := initSides(spanRow, spanCol)

	for _, tile := range positions {
		row := tile.Row - min.Row
		col := tile.Col - min.Col

		if !contains(positions, Pos{tile.Row - 1, tile.Col}) {
			up[row][col] = true
		}
		if !contains(positions, Pos{tile.Row + 1, tile.Col}) {
			down[row][col] = true
		}
		if !contains(positions, Pos{tile.Row, tile.Col - 1}) {
			left[row][col] = true
		}
		if !contains(positions, Pos{tile.Row, tile.Col + 1}) {
			right[row][col] = true
		}
	}

	return distinctHorizontal(up) + distinctHorizontal(down) + distinctVertical(left) + distinctVertical(right)
}

func main() {
	regions := convert(utils.ReadFile("input"))

	total := total(regions)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
