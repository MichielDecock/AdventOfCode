package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func readFile(fileName string) []string {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if (file == nil) {
		return lines
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func uniqueValues(array []int) map[int]int {
	dict := make(map[int]int)

	for _, number := range array {
		dict[number]++
	}

	return dict
}

func main() {
	lines := readFile("input")

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

	unique2 := uniqueValues(list2)

	sum := 0
	for _, number := range list1 {
		sum += number * unique2[number]
	}

	fmt.Println(strconv.FormatFloat(float64(sum), 'f', -1, 64))
}
