package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ToNumber(value string) int {
	number, error := strconv.Atoi(value)
	if error != nil {
		fmt.Println("Error converting string to int:", error)
		return -1
	}

	return number
}

func ReadFile(fileName string) []string {
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
