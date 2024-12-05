package main

import (
	"fmt"
	"strconv"
	"regexp"

	"aoc.com/utils"
)

func found(pattern string, order []string) bool {
	for _, orderPattern := range order {
		if (orderPattern == pattern) {
			return true
		}
	}

	return false
}

func move(slice []string, from, to int) []string {
    if from < 0 || from >= len(slice) || to < 0 || to >= len(slice) {
        fmt.Println("Invalid index")
        return slice
    }
    
    item := slice[from]
    slice = append(slice[:from], slice[from+1:]...)
    slice = append(slice[:to], append([]string{item}, slice[to:]...)...)

    return slice
}

func main() {
	order := utils.ReadFile("order")
	updates := utils.ReadFile("updates")
	regex := regexp.MustCompile(`\d\d`)

	total := 0

	for _, update := range updates {
		hits := regex.FindAllString(update, -1)

		valid := true
		repeat := true

		for repeat {
			repeat = false
			stop := false
			for i := 1; i < len(hits) && !stop; i++ {
				for j := 0; j < i && !stop; j++ {
					pattern := hits[j] + "|" + hits[i]
					if !found(pattern, order) {
						stop = true
						valid = false
						repeat = true
						newPattern := hits[i] + "|" + hits[j]
						if !found(newPattern, order) {
							fmt.Println("new pattern should be found!")
						}

						move(hits, i, j)
					}
				}
			}
		}

		if !valid {
			middle := len(hits) / 2
			total += utils.ToNumber(hits[middle])
		}
	}

	fmt.Println(strconv.FormatFloat(float64(total), 'f', -1, 64))
}
