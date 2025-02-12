package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"aoc.com/utils"
)

type Registers struct {
	A int
	B int
	C int
}

type Instructions []int

type Data struct {
	registers    Registers
	instructions Instructions
}

func parse(fileName string) Data {
	lines := utils.ReadFile(fileName)

	if len(lines) != 5 {
		fmt.Println("Invalid input!")
	}

	regex := regexp.MustCompile(`\d+`)

	var registers Registers
	registers.A = utils.ToNumber(regex.FindString(lines[0]))
	registers.B = utils.ToNumber(regex.FindString(lines[1]))
	registers.C = utils.ToNumber(regex.FindString(lines[2]))

	var instructions Instructions
	input := regex.FindAllString(lines[4], -1)
	for _, v := range input {
		instructions = append(instructions, utils.ToNumber(v))
	}

	return Data{registers, instructions}
}

func combo(operand int, data Data) int {
	if operand < 4 {
		return operand
	}

	if operand >= 7 {
		fmt.Println("Invalid operand!")
		return 0
	}

	if operand == 4 {
		return data.registers.A
	}

	if operand == 5 {
		return data.registers.B
	}

	if operand == 6 {
		return data.registers.C
	}

	return 0
}

func Adv(data Data, operand int) int {
	return data.registers.A / int(math.Pow(2, float64(combo(operand, data))))
}

func Bxl(data Data, operand int) int {
	return data.registers.B ^ operand
}

func Bst(data Data, operand int) int {
	return combo(operand, data) % 8
}

func Jnz(data Data, operand int) (bool, int) {
	if data.registers.A == 0 {
		return false, operand
	}

	return true, operand
}

func Bxc(data Data) int {
	return data.registers.B ^ data.registers.C
}

func Out(data Data, operand int) string {
	return strconv.Itoa(combo(operand, data)%8) + ","
}

func Bdv(data Data, operand int) int {
	return Adv(data, operand)
}

func Cdv(data Data, operand int) int {
	return Adv(data, operand)
}

func process(data Data) string {
	pos := 0

	var out string

	for pos < len(data.instructions) {
		instruction := data.instructions[pos]

		var operand int

		if instruction != 4 {
			if pos == len(data.instructions)-1 {
				fmt.Println("Invalid instructions sequence!")
				break
			}

			operand = data.instructions[pos+1]
		}

		switch instruction {
		case 0:
			data.registers.A = Adv(data, operand)
		case 1:
			data.registers.B = Bxl(data, operand)
		case 2:
			data.registers.B = Bst(data, operand)
		case 3:
			jump, newPos := Jnz(data, operand)
			if jump {
				pos = newPos
				continue
			}
		case 4:
			data.registers.B = Bxc(data)
		case 5:
			out += Out(data, operand)
		case 6:
			data.registers.B = Bdv(data, operand)
		case 7:
			data.registers.C = Cdv(data, operand)
		}

		pos += 2
	}

	return out[:len(out)-1]
}

func main() {
	fmt.Println(process(parse("input")))
}
