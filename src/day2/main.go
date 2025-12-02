package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
	"sync"
	"sync/atomic"
)

const DAY = 2

func P(a ...any)            { fmt.Println(a...) }
func Pf(a string, b ...any) { fmt.Printf(a, b...) }

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Range struct {
	L int
	R int
}

var factors = [11][]int{{0}, {0}, {1}, {1}, {2, 1}, {1}, {3, 2, 1}, {1}, {4, 2, 1}, {3, 1}, {5, 2, 1}}

func length(x int) int {
	return int(math.Log10(float64(x))) + 1
}
func subset(x int, diff int) bool {
	var y utils.ArrayList[int]
	var p10 = int(math.Pow10(diff))
	for x > 0 {
		y.Add(x % p10)
		x = x / p10
	}
	var count = y.Get(0)
	for _, k := range y {
		if k != count {
			return false
		}
	}
	return true
}

func Part1(input string) int {
	var Ranges utils.ArrayList[Range]
	for i := range strings.SplitSeq(input, ",") {
		var lr = strings.Split(i, "-")
		Ranges.Add(Range{L: utils.ToInt(lr[0]), R: utils.ToInt(lr[1])})
	}

	// SOLUTION
	var result int64 = 0
	var wg sync.WaitGroup

	for _, i := range Ranges {
		wg.Add(1)
		go func(i Range) {
			defer wg.Done()
			local := 0
			for j := i.L; j < i.R+1; j++ {
				var l = length(j)
				if l%2 != 0 {
					continue
				}
				if subset(j, l/2) {
					local += j
				}
			}
			atomic.AddInt64(&result, int64(local))
		}(i)
	}
	wg.Wait()

	return int(result)
}

func Part2(input string) int {
	var Ranges utils.ArrayList[Range]
	for i := range strings.SplitSeq(input, ",") {
		var lr = strings.Split(i, "-")
		Ranges.Add(Range{L: utils.ToInt(lr[0]), R: utils.ToInt(lr[1])})
	}

	// SOLUTION

	var result int64 = 0
	var wg sync.WaitGroup

	for _, i := range Ranges {
		wg.Add(1)
		go func(i Range) {
			defer wg.Done()
			local := 0
			for j := i.L; j < i.R+1; j++ {
				var l = length(j)
				if l < 2 {
					continue
				}
				var factor = factors[l]
				for _, k := range factor {
					if subset(j, k) {
						local += j
						break
					}
				}
			}
			atomic.AddInt64(&result, int64(local))
		}(i)
	}
	wg.Wait()

	return int(result)
}

func main() {
	input = utils.LoadInput(2)
	Pf("Part 1: %d\n", Part1(input))
	Pf("Part 2: %d\n", Part2(input))
}

var input string = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
