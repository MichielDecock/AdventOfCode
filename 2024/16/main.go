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

func initBoolMatrix(rows int, cols int) [][]bool {
	visited := make([][]bool, rows)
	for row := range rows {
		visited[row] = make([]bool, cols)
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

func currentPosition(reindeer Reindeer) Pos {
	return reindeer.path[len(reindeer.path)-1]
}

func printMaze(maze Maze) {
	for _, i := range maze {
		for _, j := range i {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func routes(minimalPaths []Reindeer, rows int, cols int, distance int) [][]bool {
	routes := initBoolMatrix(rows, cols)

	for _, path := range minimalPaths {
		if path.distance != distance {
			continue
		}

		for _, cell := range path.path {
			routes[cell.Row][cell.Col] = true
		}
	}

	return routes
}

func seats(routes [][]bool) int {
	total := 0
	for _, i := range routes {
		for _, j := range i {
			if j {
				total++
			}
		}
	}

	return total
}

func toMaze(routes [][]bool, original Maze) Maze {
	maze := make([][]string, len(original))
	for r := range maze {
		maze[r] = make([]string, len(original[0]))
		for c := range len(original[0]) {
			maze[r][c] = original[r][c]
		}
	}

	for r, i := range routes {
		for c, j := range i {
			if j {
				maze[r][c] = "o"
			}
		}
	}

	return maze
}

func walk(maze Maze, minDistance int) int {
	var reindeers []Reindeer

	reindeers = append(reindeers, findReindeer(maze))

	update := (minDistance == math.MaxInt)

	visited := initVisited(len(maze), len(maze[0]))

	var minimalPaths []Reindeer

	for {
		if len(reindeers) == 0 {
			break
		}

		reindeer := pop_back(&reindeers)

		pos := currentPosition(reindeer)

		if visited[pos.Row][pos.Col] != math.MaxInt {
			if update {
				if reindeer.distance > (visited[pos.Row][pos.Col]) {
					continue
				}
			} else if (reindeer.distance != (visited[pos.Row][pos.Col])+1000) && (reindeer.distance != (visited[pos.Row][pos.Col])+2000) {
				if reindeer.distance > (visited[pos.Row][pos.Col]) {
					continue
				}
			}
		}

		visited[pos.Row][pos.Col] = reindeer.distance

		if reindeer.distance > minDistance {
			continue
		}

		if maze[pos.Row][pos.Col] == "E" {
			if update {
				minDistance = int(math.Min(float64(minDistance), float64(reindeer.distance)))
			}
			minimalPaths = append(minimalPaths, reindeer)
			continue
		}

		reindeers = append(reindeers, neighbors(reindeer, maze)...)
	}

	shortest := routes(minimalPaths, len(maze), len(maze[0]), minDistance)

	printMaze(toMaze(shortest, maze))

	if update {
		return minDistance
	}

	return seats(shortest)
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

func main() {
	maze := parse("input")

	total := walk(maze, math.MaxInt)
	fmt.Println("first run")
	total = walk(maze, total)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
