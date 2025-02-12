package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func getInput() map[int]string {
	input := make(map[int]string)
	input[0] =  "MMMSXXMASM"
	input[1] =  "MSAMXMSMSA"
	input[2] =  "AMXSXMAAMM"
	input[3] =  "MSAMASMSMX"
	input[4] =  "XMASAMXAMM"
	input[5] =  "XXAMMXXAMA"
	input[6] =  "SMSMSASXSS"
	input[7] =  "SAXAMASAAA"
	input[8] =  "MAMMMXMMMM"
	input[9] =  "MXMXAXMASX"
	return input
}

func TestFindHorizontal(t *testing.T) {
	input := getInput()
	assert.Equal(t, 5, FindInMap(input));
}

func TestFindVertical(t *testing.T) {
	input := getInput()
	assert.Equal(t, 3, FindInMap(Flip(input)));
}

func TestFindDiagonal(t *testing.T) {
	input := getInput()
	assert.Equal(t, 10, FindDiagonal(input));
}
