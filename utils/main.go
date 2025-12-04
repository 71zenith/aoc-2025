package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func LoadInput(day int) string {
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "../")
	path := filepath.Join(projectRoot, fmt.Sprintf("input/day%d/input", day))
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func ToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func ExtractInts(input string) []int {
	fields := strings.Fields(input)
	array := make([]int, len(fields))
	for i, field := range fields {
		array[i] = ToInt(field)
	}
	return array
}

func ParseIntLines(input string) []int {
	lines := strings.Split(input, "\n")
	array := make([]int, len(lines))
	for i, line := range lines {
		array[i] = ToInt(strings.TrimSpace(line))
	}
	return array
}

type Grid [][]byte

func ParseGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return grid
}

type Point struct {
	X int
	Y int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

var (
	N      = Point{-1, 0}
	S      = Point{1, 0}
	W      = Point{0, -1}
	E      = Point{0, 1}
	Direc4 = []Point{N, S, W, E}
	NW     = N.Add(W)
	NE     = N.Add(E)
	SW     = S.Add(W)
	SE     = S.Add(E)
	Direc8 = []Point{NW, NE, SW, SE, N, S, W, E}
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (g Grid) InBounds(p Point) bool {
	return p.X >= 0 && p.X < len(g) && p.Y >= 0 && p.Y < len(g[p.X])
}

func (g Grid) At(p Point) byte {
	return g[p.X][p.Y]
}

func (g Grid) Set(p Point, val byte) {
	g[p.X][p.Y] = val
}

func (g Grid) Find(symbol byte) Point {
	for x := range(g) {
		for y := range(g[x]) {
			if g.At(Point{x,y}) == symbol {
				return Point{x,y}
			}
		}
	}
	return Point{-1,-1}
}

func (g Grid) FindAll(symbol byte, list *ArrayList[Point]) {
	for x := range(g) {
		for y := range(g[x]) {
			if g.At(Point{x,y}) == symbol {
				list.Add(Point{x,y})
			}
		}
	}
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

func (g Grid) NeighborCount(p Point, sym byte) int {
	count := 0
	for _, dir := range Direc8 {
		neighbor := p.Add(dir)
		if g.InBounds(neighbor) && g.At(neighbor) == sym {
			count += 1
		}
	}
	return count
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

func (q Queue[T]) Len() int {
	return len(q)
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

func (s Stack[T]) Len() int {
	return len(s)
}

type HashSet[T comparable] map[T]bool

func (s HashSet[T]) Add(v T) {
	s[v] = true
}

func (s HashSet[T]) Contains(v T) bool {
	_, exists := s[v]
	return exists
}

func (s HashSet[T]) Len() int {
	return len(s)
}

func (s HashSet[T]) GetorAdd(v T) bool {
	_, exists := s[v]
	if !exists {
		s.Add(v)
		return false
	}
	return exists
}

func (s HashSet[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s HashSet[T]) Remove(v T) {
	delete(s, v)
}

type HashMap[T comparable, V any] map[T]V

func (s HashMap[T, V]) Add(k T, v V) {
	s[k] = v
}

func (s HashMap[T, V]) Contains(k T) bool {
	_, exists := s[k]
	return exists
}

func (s HashMap[T, V]) Len() int {
	return len(s)
}

func (s HashMap[T, V]) Get(k T) V {
	return s[k]
}

func (s HashMap[T, V]) GetorAdd(k T, v V) bool {
	_, exists := s[k]
	if !exists {
		s.Add(k, v)
		return false
	}
	return exists
}

func (s HashMap[T, V]) Clear() {
	for k := range s {
		delete(s, k)
	}
}
func (s HashMap[T, V]) Remove(k T) {
	delete(s, k)
}

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

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
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

func FreqTable(nums []int) map[int]int {
	table := make(map[int]int, len(nums))
	for i := range nums {
		table[nums[i]]++
	}
	return table
}

type ArrayList[T any] []T

func (arr *ArrayList[T]) Add(item T) {
	*arr = append(*arr, item)
}

func (arr *ArrayList[T]) AddAt(index int, item T) {
	if index < 0 || index > len(*arr) {
		panic("out of bounds")
	}
	*arr = append(*arr, item)
	copy((*arr)[index+1:], (*arr)[index:])
	(*arr)[index] = item
}
func (arr ArrayList[T]) Get(index int) T {
	if index < 0 || index >= len(arr) {
		panic("out of bounds")
	}
	return arr[index]
}

func (arr *ArrayList[T]) Remove(index int) T {
	if index < 0 || index >= len(*arr) {
		panic("index out of bounds")
	}
	item := (*arr)[index]
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
	return item
}

func (arr ArrayList[T]) Size() int {
	return len(arr)
}

func (arr *ArrayList[T]) Clear() {
	*arr = (*arr)[:0]
}

func (arr *ArrayList[T]) Pop() T {
	if len(*arr) == 0 {
		panic("empty list")
	}
	lastIndex := len(*arr) - 1
	item := (*arr)[lastIndex]
	*arr = (*arr)[:lastIndex]
	return item
}

func BFS(grid Grid, start, end Point, block byte) int {
	queue := Queue[Point]{}
	visited := HashSet[Point]{}
	distance := HashMap[Point, int]{}
	queue.Enqueue(start)
	visited.Add(start)
	distance.Add(start, 0)
	for queue.Len() > 0 {
		current := queue.Dequeue()
		currentDist := distance.Get(current)
		if current == end {
			return currentDist
		}
		for _, neighbor := range grid.Neighbors4(current) {
			if !visited.Contains(neighbor) && grid.At(neighbor) != block {
				queue.Enqueue(neighbor)
				visited.Add(neighbor)
				distance.Add(neighbor, currentDist+1)
			}
		}
	}
	return -1
}

func Mod(a, b int) int {
    m := a % b
    if m < 0 {
        m += b
    }
    return m
}


func Floor(a float64) int {
	if a >= 0 {
		return int(a)
	}
	i := int(a)
	if float64(i) == a {
		return i
	}
	return i - 1
}
