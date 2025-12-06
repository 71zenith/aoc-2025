package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

const DAY = 6

func P(a ...any)            { fmt.Println(a...) }
func Pf(a string, b ...any) { fmt.Printf(a, b...) }

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	var nums utils.ArrayList[[]int]
	for _, l := range lines[:len(lines)-1] {
		x := utils.ExtractInts(l)
		nums.Add(x)
	}
	var oprs utils.ArrayList[string]
	fields := strings.FieldsSeq(lines[len(lines)-1])
	for field := range fields {
		oprs.Add(field)
	}

	// SOLUTION
	var result int = 0

	for r := range len(oprs) {
		var local int = 0
		switch oprs[r] {
		case "*":
			local = 1
			var len = nums.Size()
			for x := range len {
				local *= nums.Get(x)[r]
			}
		case "+":
			local = 0
			var len = nums.Size()
			for x := range len {
				local += nums.Get(x)[r]
			}
		}
		result += local
	}

	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	var trans string
	for l := range len(lines[0]) {
		var local string
		for r := range len(lines) - 1 {
			local += string(lines[r][l])
		}
		if strings.TrimSpace(local) == "" {
			trans += "\n"
		} else {
			trans += local + " "
		}
	}
	splits := strings.Split(trans, "\n")

	var oprs utils.ArrayList[string]
	fields := strings.FieldsSeq(lines[len(lines)-1])
	for field := range fields {
		oprs.Add(field)
	}

	// SOLUTION
	var result int = 0

	for i := range(oprs) {
		var local int
		var nums = utils.ExtractInts(splits[i])
		switch oprs[i] {
		case "*":
			local = 1
			for x := range len(nums) {
				local *= nums[x]
			}
		case "+":
			local = 0
			for x := range len(nums) {
				local += nums[x]
			}
		}
		result += local
	}
	return result
}

func main() {
	input = utils.LoadInput(6)
	Pf("Part 1: %d\n", Part1(input))
	Pf("Part 2: %d\n", Part2(input))
}

var input string = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `
