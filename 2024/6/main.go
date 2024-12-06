package main

import (
	"fmt"
	"regexp"
	"strconv"

	"aoc.com/utils"
)

func initVisited(plan []string) [][]bool {
	visited := make([][]bool, len(plan))

	for i := range visited {
		visited[i] = make([]bool, len(plan[0]))
	}

	return visited
}

type Guard struct {
	Row int
	Col int
}

func findGuard(plan []string) Guard {
	regex := regexp.MustCompile(`[><^v]`)

	for row, line := range plan {
		index := regex.FindStringIndex(line)
		if index != nil {
			return Guard{Row: row, Col: index[0]}
		}
	}

	fmt.Println("Could not find a guard!")
	return Guard{-1, -1}
}

func getDistance(visited [][]bool) int {
	distance := 0

	for _, i := range visited {
		for _, j := range i {
			if j {
				distance++
			}
		}
	}

	return distance
}

func replace(line *string, pos int, char rune) {
	newLine := []rune(*line)
	newLine[pos] = char
	*line = string(newLine)
}

func up(plan *[]string, guard *Guard) bool {
	if (*plan)[guard.Row][guard.Col] != '^' {
		return false
	}

	if guard.Row == 0 {
		return false
	}

	next := guard.Row - 1
	if (*plan)[next][guard.Col] == '#' {
		replace(&(*plan)[guard.Row], guard.Col, '>')
		return true
	}

	replace(&(*plan)[guard.Row], guard.Col, '.')
	(*guard).Row = next
	replace(&(*plan)[guard.Row], guard.Col, '^')
	return true
}

func down(plan *[]string, guard *Guard) bool {
	if (*plan)[guard.Row][guard.Col] != 'v' {
		return false
	}

	if guard.Row == len(*plan)-1 {
		return false
	}

	next := guard.Row + 1
	if (*plan)[next][guard.Col] == '#' {
		replace(&(*plan)[guard.Row], guard.Col, '<')
		return true
	}

	replace(&(*plan)[guard.Row], guard.Col, '.')
	(*guard).Row = next
	replace(&(*plan)[guard.Row], guard.Col, 'v')
	return true
}

func right(plan *[]string, guard *Guard) bool {
	if (*plan)[guard.Row][guard.Col] != '>' {
		return false
	}

	if guard.Col == len((*plan)[0])-1 {
		return false
	}

	next := guard.Col + 1
	if (*plan)[guard.Row][next] == '#' {
		replace(&(*plan)[guard.Row], guard.Col, 'v')
		return true
	}

	replace(&(*plan)[guard.Row], guard.Col, '.')
	(*guard).Col = next
	replace(&(*plan)[guard.Row], guard.Col, '>')
	return true
}

func left(plan *[]string, guard *Guard) bool {
	if (*plan)[guard.Row][guard.Col] != '<' {
		return false
	}

	if guard.Col == 0 {
		return false
	}

	next := guard.Col - 1
	if (*plan)[guard.Row][next] == '#' {
		replace(&(*plan)[guard.Row], guard.Col, '^')
		return true
	}

	replace(&(*plan)[guard.Row], guard.Col, '.')
	(*guard).Col = next
	replace(&(*plan)[guard.Row], guard.Col, '<')
	return true
}

func distance() int {
	plan := utils.ReadFile("input")
	visited := initVisited(plan)

	guard := findGuard(plan)
	visited[guard.Row][guard.Col] = true

	stop := false
	for !stop {
		if !up(&plan, &guard) && !down(&plan, &guard) && !right(&plan, &guard) && !left(&plan, &guard) {
			stop = true
		}

		visited[guard.Row][guard.Col] = true
	}

	return getDistance(visited)
}

func main() {
	total := distance()
	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
