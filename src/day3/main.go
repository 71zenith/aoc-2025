package main

import (
	"aoc/utils"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

const DAY = 3

func P(a ...any)            { fmt.Println(a...) }
func Pf(a string, b ...any) { fmt.Printf(a, b...) }

func Int(b byte) int {
	return utils.ToInt(string(b))
}

func Max(s string) int {
	max := 0
	for i := 1; i < len(s); i++ {
		if Int(s[i]) > Int(s[max]) {
			max = i
		}
	}
	return max
}

func Part1(input string) int32 {
	lines := strings.SplitSeq(input, "\n")

	// SOLUTION
	var result int32 = 0
	var wg sync.WaitGroup

	for line := range lines {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			f := Max(line[:len(line)-1])
			s := f + 1 + Max(line[f+1:])
			atomic.AddInt32(&result, int32(Int(line[f])*10 + Int(line[s])))
		}(line)
	}
	wg.Wait()
	return result
}

func Part2(input string) int64 {
	lines := strings.SplitSeq(input, "\n")

	// SOLUTION
	var result int64 = 0
	var wg sync.WaitGroup

	for line := range lines {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			local := 0
			bat := 12
			max := -1
			for i := range bat {
				if (len(line) - bat + 1 + i) > len(line) {
					break
				}
				max = max + 1 + Max(line[max+1:len(line)-bat+i+1])
				local = local*10 + Int(line[max])
			}
			atomic.AddInt64(&result, int64(local))
		}(line)
	}
	wg.Wait()
	return result
}

func main() {
	input = utils.LoadInput(3)
	Pf("Part 1: %d\n", Part1(input))
	Pf("Part 2: %d\n", Part2(input))
}

var input string = `987654321111111
811111111111119
234234234234278
818181911112111`
