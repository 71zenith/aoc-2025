package main

import (
	"aoc/utils"
	"fmt"
)

const DAY = 4

func P(a ...any)            { fmt.Println(a...) }
func Pf(a string, b ...any) { fmt.Printf(a, b...) }

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Part1(input string) int {
	g := utils.ParseGrid(input)
	var papers utils.ArrayList[utils.Point]
	g.FindAll('@', &papers)

	// SOLUTION

	var result int = 0
	for _, p := range papers {
		if g.NeighborCount(p, '@') < 4 {
			result++
		}
	}

	return result
}

func Part2(input string) int {
	g := utils.ParseGrid(input)

	// SOLUTION
	var result int = 0
	var papers utils.ArrayList[utils.Point]
	var legal utils.ArrayList[utils.Point]
	prev := -1

	for legal.Size() > prev {
		for _, p := range legal {
			g.Set(p, '.')
		}
		papers.Clear()
		g.FindAll('@', &papers)
		prev = legal.Size()

		for _, p := range papers {
			if g.NeighborCount(p, '@') < 4 {
				legal.Add(p)
				result++
			}
		}
	}

	return result
}

func main() {
	input = utils.LoadInput(4)
	Pf("Part 1: %d\n", Part1(input))
	Pf("Part 2: %d\n", Part2(input))
}

var input string = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`
