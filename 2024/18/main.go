package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"aoc.com/utils"
)

type Maze [][]string

type DistanceMaze [][]int
type Pos struct {
	Row int
	Col int
}

type Reindeer struct {
	pos      Pos
	distance int
}

func initMaze(size int, fall []Pos, bytes int) Maze {
	maze := make(Maze, size)
	for i := range maze {
		maze[i] = make([]string, size)
		for j := range size {
			maze[i][j] = "."
		}
	}

	for b := 0; b != bytes; b++ {
		el := fall[b]
		maze[el.Row][el.Col] = "#"
	}

	maze[0][0] = "S"
	maze[size-1][size-1] = "E"

	return maze
}

func initDistanceMaze(m Maze) DistanceMaze {
	size := len(m)
	maze := make(DistanceMaze, size)
	for i := range maze {
		maze[i] = make([]int, size)
		for j := range size {
			if m[i][j] == "#" {
				maze[i][j] = -1
			}
		}
	}

	return maze
}

func parse(fileName string) ([]Pos, int) {
	lines := utils.ReadFile(fileName)

	regex := regexp.MustCompile(`\d+`)

	var fall []Pos

	size := 0

	for _, line := range lines {
		hits := regex.FindAllString(line, -1)
		row := utils.ToNumber(hits[1])
		col := utils.ToNumber(hits[0])
		fall = append(fall, Pos{row, col})

		size = int(math.Max(float64(row), float64(size)))
		size = int(math.Max(float64(col), float64(size)))
	}

	return fall, size + 1
}

func neighbors(reindeer Reindeer, maze Maze, distances DistanceMaze) []Reindeer {
	var neighbors []Reindeer

	pos := reindeer.pos
	size := len(maze)

	for row := int(math.Max(0, float64(pos.Row-1))); row < int(math.Min(float64(size), float64(pos.Row+2))); row++ {
		for col := int(math.Max(0, float64(pos.Col-1))); col < int(math.Min(float64(size), float64(pos.Col+2))); col++ {
			if (row != pos.Row && col != pos.Col) || (row == pos.Row && col == pos.Col) {
				continue
			}

			if maze[row][col] == "#" {
				continue
			}

			if distances[row][col] != 0 {
				continue
			}

			distances[row][col] = reindeer.distance + 1

			neighbors = append(neighbors, Reindeer{Pos{row, col}, reindeer.distance + 1})
		}
	}

	return neighbors
}

func pop(reindeers *[]Reindeer) Reindeer {
	if len(*reindeers) == 0 {
		fmt.Println("Trying to pop an empty reindeer list")
		return Reindeer{Pos{}, -1}
	}

	front := (*reindeers)[0]
	*reindeers = (*reindeers)[1:len(*reindeers)]
	return front
}

func printMaze(maze Maze) {
	for _, i := range maze {
		for _, j := range i {
			fmt.Print(j)
		}
		fmt.Println()
	}
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

func walk(maze Maze) int {
	reindeers := []Reindeer{findReindeer(maze)}

	distances := initDistanceMaze(maze)

	for {
		reindeer := pop(&reindeers)

		if reindeer.distance == -1 {
			break
		}

		pos := reindeer.pos
		distances[pos.Row][pos.Col] = reindeer.distance

		if maze[pos.Row][pos.Col] == "E" {
			fmt.Println("reached E")
			return distances[pos.Row][pos.Col]
		}

		reindeers = append(reindeers, neighbors(reindeer, maze, distances)...)
	}

	return -1
}

func findReindeer(maze Maze) Reindeer {
	for row, line := range maze {
		for col, el := range line {
			if el == "S" {
				return Reindeer{Pos{row, col}, 0}
			}
		}
	}

	fmt.Println("The reindeer wasn't found!")
	return Reindeer{Pos{}, 0}
}

func main() {
	fall, size := parse("input")

	maze := initMaze(size, fall, 1024)

	// printMaze(maze)

	total := walk(maze)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
