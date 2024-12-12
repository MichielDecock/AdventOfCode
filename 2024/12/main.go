package main

import (
	"fmt"
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
			perimeter := perimeter(member)
			total += area * perimeter
		}
	}
	return total
}

func main() {
	regions := convert(utils.ReadFile("input"))

	total := total(regions)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
