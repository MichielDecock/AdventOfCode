package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	input := "233"
	out := Convert(input)
	assert.Equal(t, []int{0, 0, -1, -1, -1, 1, 1, 1}, out)
}

func TestSize(t *testing.T) {
	input := "233"
	assert.Equal(t, 8, Size(input))
}

func TestInit(t *testing.T) {
	input := "233"
	assert.Equal(t, []int{-1, -1, -1, -1, -1, -1, -1, -1}, Init(input))
}

func TestMove(t *testing.T) {
	input := []int{0, 0, -1, -1, -1, 1, 1, 1}
	assert.Equal(t, []int{0, 0, 1, 1, 1, -1, -1, -1}, Move(input))
}

func TestFirstSpace(t *testing.T) {
	input := []int{0, 0, -1, -1, -1, 1, 1, 1}
	assert.Equal(t, 2, FirstSpace(input))
}

func TestLastFile(t *testing.T) {
	input := []int{0, 0, -1, -1, -1, 1, 1, 1}
	assert.Equal(t, 7, LastFile(input))
}

func TestCheckSum(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, -1, -1, -1}
	assert.Equal(t, uint64(9), CheckSum(input))
}
