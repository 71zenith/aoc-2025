package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

const DAY = 1

func P(a ...any) { fmt.Println(a...) }
func Pf(a string, b ...any) { fmt.Printf(a, b...) }

func check(err error) {
	if (err != nil) {
		panic(err)
	}
}

type Op struct {
	Dir byte
	Angle int
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	var Ops utils.ArrayList[Op]

	for _, line := range(lines) {
		Ops.Add(Op{Dir: line[0], Angle: utils.ToInt(line[1:])})
	}
	// SOLUTION

	var angle int = 9999950
	var result int = 0

	for _, op := range(Ops) {
		if op.Dir == 'R' {
			angle += op.Angle
		}
		if op.Dir == 'L' {
			angle -= op.Angle
		}
		if angle % 100 == 0 {
			result++
		}
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	var Ops utils.ArrayList[Op]

	for _, line := range(lines) {
		Ops.Add(Op{Dir: line[0], Angle: utils.ToInt(line[1:])})
	}
	
	// SOLUTION

	var angle int = 50
	var result int = 0
	var prev int = 0
	var diff int = 0

	for _, op := range(Ops) {
		prev = angle
		if op.Dir == 'R' {
			angle += op.Angle
			diff = op.Angle
		}
		if op.Dir == 'L' {
			angle -= op.Angle
			diff = -op.Angle
		}
		angle = utils.Mod(angle, 100)
		if angle % 100 == 0 {
			result++
		}
		result += utils.Floor((float64)(utils.Abs(diff))/ 100.0)
		if (prev + (diff % 100) > 100) || ((prev + (diff % 100) < 0) && prev != 0) {
			result++
		}
	}
	
	return result
}

func main() {
	input = utils.LoadInput(1)
	Pf("Part 1: %d\n", Part1(input))
	Pf("Part 2: %d\n", Part2(input))
}

var input string = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
