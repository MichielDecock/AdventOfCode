package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"aoc.com/utils"
)

type Maze [][]string
type Pos struct {
	Row int
	Col int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Reindeer struct {
	path      []Pos
	distance  int
	direction Direction
}

type Visited [][]int

func initVisited(rows int, cols int) Visited {
	visited := make([][]int, rows)
	for row := range rows {
		visited[row] = make([]int, cols)
		for col := range cols {
			visited[row][col] = math.MaxInt
		}
	}
	return visited
}

func parse(fileName string) Maze {
	lines := utils.ReadFile(fileName)

	maze := make(Maze, len(lines))
	for i, line := range lines {
		maze[i] = make([]string, len(line))
		for j, el := range line {
			maze[i][j] = string(el)
		}
	}

	return maze
}

func rotate(direction Direction, target Direction) int {
	count := 0
	for ; direction%4 != target; direction++ {
		count++
	}
	return count % 2
}

func beenHereBefore(row int, col int, reindeer Reindeer) bool {
	for _, cell := range reindeer.path {
		if cell.Row == row && cell.Col == col {
			return true
		}
	}
	return false
}

func neighbors(reindeer Reindeer, maze Maze) []Reindeer {
	limits := Pos{len(maze), len(maze[0])}

	var neighbors []Reindeer

	pos := currentPosition(reindeer)

	for row := int(math.Max(0, float64(pos.Row-1))); row < int(math.Min(float64(limits.Row), float64(pos.Row+2))); row++ {
		for col := int(math.Max(0, float64(pos.Col-1))); col < int(math.Min(float64(limits.Col), float64(pos.Col+2))); col++ {
			if (row != pos.Row && col != pos.Col) || (row == pos.Row && col == pos.Col) {
				continue
			}

			if maze[row][col] == "#" {
				continue
			}

			if beenHereBefore(row, col, reindeer) {
				continue
			}

			distance := reindeer.distance + 1
			direction := reindeer.direction

			if row < pos.Row {
				distance += rotate(reindeer.direction, North) * 1000
				direction = North
			} else if row > pos.Row {
				distance += rotate(reindeer.direction, South) * 1000
				direction = South
			} else if col < pos.Col {
				distance += rotate(reindeer.direction, West) * 1000
				direction = West
			} else if col > pos.Col {
				distance += rotate(reindeer.direction, East) * 1000
				direction = East
			}

			path := make([]Pos, len(reindeer.path))
			copy(path, reindeer.path)

			path = append(path, Pos{row, col})

			neighbors = append(neighbors, Reindeer{path, distance, direction})
		}
	}

	sort.Slice(neighbors, func(i, j int) bool {
		neighbor1 := neighbors[i]
		neighbor2 := neighbors[j]

		posI := currentPosition(neighbor1)
		posJ := currentPosition(neighbor2)

		if posI.Row > posI.Col {
			if posI.Row == posJ.Row {
				return posI.Col < posJ.Col
			}

			return posI.Row > posJ.Row

		} else {
			if posI.Col == posJ.Col {
				return posI.Row > posJ.Row
			}

			return posI.Col < posJ.Col
		}
	})

	return neighbors
}

func pop_back(reindeers *[]Reindeer) Reindeer {
	if len(*reindeers) == 0 {
		fmt.Println("Trying to pop an empty reindeer list")
		return Reindeer{}
	}

	back := (*reindeers)[len(*reindeers)-1]
	*reindeers = (*reindeers)[:len(*reindeers)-1]
	return back
}

func shortestPath(target Pos, current Pos) int {
	return 1000 + (target.Col - current.Col) + (target.Row - current.Row)
}

func currentPosition(reindeer Reindeer) Pos {
	return reindeer.path[len(reindeer.path)-1]
}

func walk(maze Maze) int {
	var reindeers []Reindeer

	reindeers = append(reindeers, findReindeer(maze))

	target := findTarget(maze)

	minDistance := math.MaxInt

	visited := initVisited(len(maze), len(maze[0]))

	for {
		if len(reindeers) == 0 {
			break
		}

		reindeer := pop_back(&reindeers)

		pos := currentPosition(reindeer)

		if reindeer.distance >= visited[pos.Row][pos.Col] {
			continue
		}

		visited[pos.Row][pos.Col] = reindeer.distance

		if reindeer.distance >= minDistance-shortestPath(currentPosition(target), pos) {
			continue
		}

		if maze[pos.Row][pos.Col] == "E" {
			minDistance = int(math.Min(float64(minDistance), float64(reindeer.distance)))
			continue
		}

		reindeers = append(reindeers, neighbors(reindeer, maze)...)
	}

	return minDistance
}

func findReindeer(maze Maze) Reindeer {
	for row, line := range maze {
		for col, el := range line {
			if el == "S" {
				return Reindeer{[]Pos{{row, col}}, 0, East}
			}
		}
	}

	fmt.Println("The reindeer wasn't found!")
	return Reindeer{[]Pos{}, 0, East}
}

func findTarget(maze Maze) Reindeer {
	for row, line := range maze {
		for col, el := range line {
			if el == "E" {
				return Reindeer{[]Pos{{row, col}}, 0, East}
			}
		}
	}

	fmt.Println("The target wasn't found!")
	return Reindeer{[]Pos{}, 0, East}
}

func main() {
	maze := parse("input")

	total := walk(maze)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
