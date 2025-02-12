package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func initData() Data {
	var data Data
	data.registers.A = 729
	data.registers.B = 9

	return data
}

func TestAdv(t *testing.T) {
	data := initData()
	assert.Equal(t, 729, Adv(data, 0))
	assert.Equal(t, 364, Adv(data, 1))
	assert.Equal(t, 182, Adv(data, 2))
	assert.Equal(t, 91, Adv(data, 3))
	assert.Equal(t, 0, Adv(data, 4))
	assert.Equal(t, 1, Adv(data, 5))
	assert.Equal(t, 729, Adv(data, 6))
}

func TestBxl(t *testing.T) {
	data := initData()
	assert.Equal(t, 9, Bxl(data, 0))
	assert.Equal(t, 8, Bxl(data, 1))
	assert.Equal(t, 11, Bxl(data, 2))
	assert.Equal(t, 10, Bxl(data, 3))
	assert.Equal(t, 13, Bxl(data, 4))
	assert.Equal(t, 12, Bxl(data, 5))
	assert.Equal(t, 15, Bxl(data, 6))
}
