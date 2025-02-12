package main

import (
	"fmt"
	"regexp"
	"strconv"

	"aoc.com/utils"
)

type Direction int

const (
	None  = 0
	Up    = 1
	Down  = 2
	Left  = 4
	Right = 8
)

type Visited struct {
	visited   bool
	direction Direction
}

func direction(direction string) Direction {
	if direction == "<" {
		return Left
	}
	if direction == ">" {
		return Right
	}
	if direction == "^" {
		return Up
	}
	if direction == "v" {
		return Down
	}

	return None
}

func initVisited(plan []string) [][]Visited {
	visited := make([][]Visited, len(plan))

	for i := range visited {
		visited[i] = make([]Visited, len(plan[0]))
	}

	return visited
}

func initObstacles(plan []string) [][]bool {
	obstacles := make([][]bool, len(plan))

	for i := range obstacles {
		obstacles[i] = make([]bool, len(plan[0]))
	}

	return obstacles
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

func getDistance(visited [][]Visited) int {
	distance := 0

	for _, i := range visited {
		for _, j := range i {
			if j.visited {
				distance++
			}
		}
	}

	return distance
}

func getObstacles(obstacles [][]bool) int {
	total := 0

	for _, i := range obstacles {
		for _, j := range i {
			if j {
				total++
			}
		}
	}

	return total
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

func visit(guard Guard, visited *[][]Visited, plan *[]string) bool {
	direction := direction(string((*plan)[guard.Row][guard.Col]))
	if (*visited)[guard.Row][guard.Col].direction&direction != 0 {
		return false
	}

	(*visited)[guard.Row][guard.Col].direction |= direction
	(*visited)[guard.Row][guard.Col].visited = true
	return true
}

func distance() int {
	plan := utils.ReadFile("test")
	visited := initVisited(plan)

	guard := findGuard(plan)

	for {
		visit(guard, &visited, &plan)
		if !up(&plan, &guard) && !down(&plan, &guard) && !right(&plan, &guard) && !left(&plan, &guard) {
			break
		}
	}

	return getDistance(visited)
}

func obstacles() int {
	plan := utils.ReadFile("input")
	visited := initVisited(plan)
	obstacles := initObstacles(plan)

	guard := findGuard(plan)

	for {
		visit(guard, &visited, &plan)
		placeObstacle(guard, &visited, &plan, &obstacles)

		if !up(&plan, &guard) && !down(&plan, &guard) && !right(&plan, &guard) && !left(&plan, &guard) {
			break
		}
	}

	return getObstacles(obstacles)
}

func placeObstacle(guard Guard, visited *[][]Visited, plan *[]string, obstacles *[][]bool) {
	if placeUp(plan, guard, visited, obstacles) {
		return
	}

	if placeDown(plan, guard, visited, obstacles) {
		return
	}

	if placeLeft(plan, guard, visited, obstacles) {
		return
	}

	if placeRight(plan, guard, visited, obstacles) {
		return
	}
}

func copyPlan(plan *[]string, row int, col int) []string {
	planCopy := make([]string, len(*plan))
	copy(planCopy, *plan)

	str := planCopy[row]
	if col < len((*plan)[0])-1 {
		str = str[:col] + "#" + str[col+1:]
	} else {
		str = str[:col] + "#"
	}

	planCopy[row] = str

	return planCopy
}

func rotate(plan *[]string, guard Guard) {
	if (*plan)[guard.Row][guard.Col] == 'v' {
		replace(&(*plan)[guard.Row], guard.Col, '<')
	} else if (*plan)[guard.Row][guard.Col] == '<' {
		replace(&(*plan)[guard.Row], guard.Col, '^')
	} else if (*plan)[guard.Row][guard.Col] == '^' {
		replace(&(*plan)[guard.Row], guard.Col, '>')
	} else if (*plan)[guard.Row][guard.Col] == '>' {
		replace(&(*plan)[guard.Row], guard.Col, 'v')
	}
}

func deepCopy(visited *[][]Visited) [][]Visited {
	visitedCopy := make([][]Visited, len(*visited))

	for row, u := range *visited {
		line := make([]Visited, len((*visited)[0]))
		for col, v := range u {
			line[col].visited = v.visited
			line[col].direction = v.direction
		}
		visitedCopy[row] = line
	}

	return visitedCopy
}

func copyAll(plan *[]string, guard Guard, visited *[][]Visited, next int, horizontal bool) ([]string, Guard, [][]Visited) {
	var planCopy []string
	if horizontal {
		planCopy = copyPlan(plan, guard.Row, next)
	} else {
		planCopy = copyPlan(plan, next, guard.Col)
	}

	visitedCopy := deepCopy(visited)

	guardCopy := guard

	return planCopy, guardCopy, visitedCopy
}

func place(plan *[]string, guard Guard, visited *[][]Visited, obstacles *[][]bool, next int, horizontal bool) {
	planCopy, guardCopy, visitedCopy := copyAll(plan, guard, visited, next, horizontal)
	rotate(&planCopy, guardCopy)
	visit(guardCopy, &visitedCopy, &planCopy)

	for {
		if !up(&planCopy, &guardCopy) && !down(&planCopy, &guardCopy) && !right(&planCopy, &guardCopy) && !left(&planCopy, &guardCopy) {
			break
		}

		if !visit(guardCopy, &visitedCopy, &planCopy) {
			if horizontal {
				(*obstacles)[guard.Row][next] = true
			} else {
				(*obstacles)[next][guard.Col] = true
			}
			break
		}
	}
}

func placeUp(plan *[]string, guard Guard, visited *[][]Visited, obstacles *[][]bool) bool {
	if (*plan)[guard.Row][guard.Col] != '^' {
		return false
	}

	if guard.Row == 0 {
		return true
	}

	next := guard.Row - 1
	if (*plan)[next][guard.Col] == '#' {
		return true
	}

	if !(*visited)[next][guard.Col].visited {
		place(plan, guard, visited, obstacles, next, false)
	}

	return true
}

func placeDown(plan *[]string, guard Guard, visited *[][]Visited, obstacles *[][]bool) bool {
	if (*plan)[guard.Row][guard.Col] != 'v' {
		return false
	}

	if guard.Row == len(*plan)-1 {
		return true
	}

	next := guard.Row + 1
	if (*plan)[next][guard.Col] == '#' {
		return true
	}

	if !(*visited)[next][guard.Col].visited {
		place(plan, guard, visited, obstacles, next, false)
	}

	return true
}

func placeRight(plan *[]string, guard Guard, visited *[][]Visited, obstacles *[][]bool) bool {
	if (*plan)[guard.Row][guard.Col] != '>' {
		return false
	}

	if guard.Col == len((*plan)[0])-1 {
		return true
	}

	next := guard.Col + 1
	if (*plan)[guard.Row][next] == '#' {
		return true
	}

	if !(*visited)[guard.Row][next].visited {
		place(plan, guard, visited, obstacles, next, true)
	}

	return true
}

func placeLeft(plan *[]string, guard Guard, visited *[][]Visited, obstacles *[][]bool) bool {
	if (*plan)[guard.Row][guard.Col] != '<' {
		return false
	}

	if guard.Col == 0 {
		return true
	}

	next := guard.Col - 1
	if (*plan)[guard.Row][next] == '#' {
		return true
	}

	if !(*visited)[guard.Row][next].visited {
		place(plan, guard, visited, obstacles, next, true)
	}

	return true
}

func main() {
	total := obstacles()
	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
