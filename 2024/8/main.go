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

func nodeTypes() []string {
	var types []string
	for i := range 10 {
		types = append(types, string('0'+i))
	}

	for i := range 26 {
		types = append(types, string('a'+i))
		types = append(types, string('A'+i))
	}

	return types
}

func nodes(nodeType string, lines []string) []Pos {
	var nodes []Pos

	for row, line := range lines {
		for col, char := range line {
			if string(char) == nodeType {
				nodes = append(nodes, Pos{row, col})
			}
		}
	}

	return nodes
}

func contains(positions []Pos, position Pos) bool {
	for _, pos := range positions {
		if pos.Row == position.Row && pos.Col == position.Col {
			return true
		}
	}
	return false
}

func antiNodes(nodeType string, lines []string, antiNodes *[]Pos) {
	start := 0
	end := Pos{len(lines), len(lines[0])}

	nodes := nodes(nodeType, lines)
	for i := 0; i != len(nodes); i++ {
		for j := 0; j != len(nodes); j++ {
			if i == j {
				continue
			}

			span := Pos{nodes[i].Row - nodes[j].Row, nodes[i].Col - nodes[j].Col}

			n := []Pos{{nodes[i].Row + span.Row, nodes[i].Col + span.Col}, {nodes[j].Row - span.Row, nodes[j].Col - span.Col}}
			for _, node := range n {
				if node.Row < end.Row && node.Row >= start && node.Col < end.Col && node.Col >= start {
					if !contains(*antiNodes, node) {
						*antiNodes = append(*antiNodes, node)
					}
				}
			}
		}
	}
}

func total(lines []string) int {
	total := 0

	types := nodeTypes()

	anti := make([]Pos, 0)

	for _, nodeType := range types {
		antiNodes(nodeType, lines, &anti)
	}

	total += len(anti)

	return total
}

func main() {
	lines := utils.ReadFile("input")

	total := total(lines)

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
