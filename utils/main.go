package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"path/filepath"
	"runtime"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func LoadInput(day int) string {
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "../")
	path := filepath.Join(projectRoot, fmt.Sprintf("input/day%d/input", day))
	input, err := os.ReadFile(path)
	check(err)
	return string(input)
}

func ExtractInts(input string) []int {
	fields := strings.Fields(input)
	array := make([]int, len(fields))
	for i, field := range fields {
		array[i], _ = strconv.Atoi(field)
	}
	return array
}

func ParseIntLines(input string) []int {
	lines := strings.Split(input, "\n")
	array := make([]int, len(lines))
	for i, line := range lines {
		array[i], _ = strconv.Atoi(strings.TrimSpace(line))
	}
	return array
}

type Grid [][]byte

func ParseGrid(input string) [][]byte {
	lines := strings.Split(input, "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return grid
}

type Point struct {
	x int
	y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

var (
	U      = Point{0, -1}
	D      = Point{0, 1}
	L      = Point{-1, 0}
	R      = Point{1, 0}
	Direc4 = []Point{U, D, L, R}
	UL     = U.Add(L)
	UR     = U.Add(R)
	DL     = D.Add(L)
	DR     = D.Add(R)
	Direc8 = []Point{UL, UR, DL, DR, U, D , L , R}
)

func (g Grid) InBounds(p Point) bool {
	return p.y >= 0 && p.y < len(g) && p.x >= 0 && p.x < len(g[p.y])
}

func (g Grid) At(p Point) byte {
	return g[p.x][p.y]
}

func (g Grid) Set(p Point, val byte) {
	g[p.x][p.y] = val
}

func (g Grid) Neighbors4(p Point) []Point {
	var neighbors []Point
	for _, dir := range Direc4 {
		neighbor := p.Add(dir)
		if g.InBounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (g Grid) Neighbors8(p Point) []Point {
	var neighbors []Point
	for _, dir := range Direc8 {
		neighbor := p.Add(dir)
		if g.InBounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

type Queue[T any] []T

func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Dequeue() T {
	if len(*q) == 0 {
		panic("queue is empty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

type HashSet[T comparable] map[T]bool

type Item struct {
	Value    any
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func Min(nums ...int) int {
    min := nums[0]
    for _, n := range nums {
        if n < min {
            min = n
        }
    }
    return min
}

func Max(nums ...int) int {
    max := nums[0]
    for _, n := range nums {
        if n > max {
            max = n
        }
    }
    return max
}

