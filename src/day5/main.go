package main

import (
	"fmt"
	"strings"
	"aoc/utils"
	"sort"
)

const DAY = 5

func P(a ...any) { fmt.Println(a...) }
func Pf(a string, b ...any) { fmt.Printf(a, b...) }

type Range struct {
	L int
	R int
}

func Part1(input string) int {
	lines := strings.Split(input, "\n\n")
	var ranges utils.ArrayList[Range]
	for line := range(strings.SplitSeq(lines[0], "\n")) {
		L,R := utils.ToInt(strings.Split(line, "-")[0]), utils.ToInt(strings.Split(line, "-")[1])
		ranges.Add(Range{L: L, R: R})
	}
	// SOLUTION

	var result int = 0
	for line := range(strings.SplitSeq(lines[1], "\n")) {
		N := utils.ToInt(line)
		for _,r := range(ranges) {
			if N >= r.L && N <= r.R {
				result++
				break
			}
		}
	}
	
	return result
}

type SRange []Range
func (s SRange) Len() int {return(len(s))}
func (s SRange) Swap(i,j int) {s[i], s[j] = s[j], s[i]}
func (s SRange) Less(i, j int) bool { return s[i].L < s[j].L }

func Part2(input string) int {
	lines := strings.Split(input, "\n\n")
	var ranges utils.ArrayList[Range]
	for line := range(strings.SplitSeq(lines[0], "\n")) {
		L,R := utils.ToInt(strings.Split(line, "-")[0]), utils.ToInt(strings.Split(line, "-")[1])
		ranges.Add(Range{L: L, R: R})
	}
	var result int = 0
	sort.Sort(SRange(ranges))
	var L = ranges[0].L
	var R = ranges[0].R
	for _,r := range(ranges[1:]) {
		if r.L > R {
			result += R - L + 1
			L = r.L
			R = r.R
		} else if r.R > R {
			R = r.R
		}
	}
	result += R - L + 1

	return result
}

func main() {
	input = utils.LoadInput(5)
	Pf("Part 1: %d\n", Part1(input))
	Pf("Part 2: %d\n", Part2(input))
}

var input string = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
