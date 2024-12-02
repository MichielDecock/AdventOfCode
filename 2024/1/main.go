package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func toNumber(value string) int {
	number, error := strconv.Atoi(value)
	if error != nil {
		fmt.Println("Error converting string to int:", error)
		return -1
	}

	return number
}

func main() {
	var lines []string

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var list1 []int
	var list2 []int

	for _, line := range lines {
		fields := strings.Fields(line)

		for index, field := range fields {
			number := toNumber(field)
			if index % 2 == 0 {
				list1 = append(list1, number)
			} else {
				list2 = append(list2, number)
			}
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	totalDifference := float64(0)
	for i := 0; i < len(list1); i++ {
		totalDifference += math.Abs(float64(list1[i] - list2[i]))
	}

	fmt.Println(strconv.FormatFloat(totalDifference, 'f', -1, 64))
}
