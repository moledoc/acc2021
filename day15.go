package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	row, col, dist int
}

var inf int = int(^uint(0) >> 1)
var gridSize int = 100
var filename string = "input15.txt"
var dirs []int = []int{
	1, 0, // up
	-1, 0, //down
	0, 1, // right
	0, -1, // left
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func checkScanner(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func dump(grid [][]int) {
	fmt.Println()
	for _, row := range grid {
		fmt.Println(row)
	}
}

func makeVisited() [][]int {
	var visited [][]int
	for i := 0; i < gridSize; i++ {
		row := make([]int, gridSize)
		visited = append(visited, row)
	}
	return visited
}

func minDist(dist []int, queue map[int]node) (ind int) {
	var min int = inf
	for i, elem := range dist {
		if _, ok := queue[i]; elem < min && ok {
			min = elem
			ind = i
		}
	}
	return ind
}

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
// It needs better variable names and more optimal structres.
func dijkstra(grid [][]int) int {
	queue := make(map[int]node)
	dist := make([]int, gridSize*gridSize)
	prev := make([]node, gridSize*gridSize)
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			dist[row*gridSize+col] = inf
			prev[row*gridSize+col] = node{-1, -1, -1}
			queue[row*gridSize+col] = node{row, col, grid[row][col]}
		}
	}
	dist[0] = 0
	queue[0] = node{0, 0, 0}
	var u node
	for len(queue) > 0 {
		safest := minDist(dist, queue)
		point := queue[safest]
		delete(queue, safest)
		if safest == gridSize*gridSize-1 {
			u = point
			break
		}
		for i := 0; i < len(dirs); i += 2 {
			newRow := point.row + dirs[i]
			newCol := point.col + dirs[i+1]
			next := newRow*gridSize + newCol
			if _, ok := queue[next]; newRow < 0 ||
				newCol < 0 ||
				newRow >= gridSize ||
				newCol >= gridSize ||
				!ok {
				continue
			}
			alt := dist[safest] + queue[next].dist
			if alt < dist[next] {
				dist[next] = alt
				prev[next] = point
			}
		}
	}
	var riskSum int
	unreachable := node{-1, -1, -1}
	source := node{0, 0, 0}
	for {
		i := u.row*gridSize + u.col
		if !(prev[i] != unreachable || prev[i] == source) {
			break
		}
		riskSum += u.dist
		u = prev[i]
	}
	return riskSum
}

func problem1() int {
	file, err := os.Open(filename)
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	var grid [][]int
	for j := 0; scanner.Scan(); j++ {
		elems := strings.Split(scanner.Text(), "")
		risks := make([]int, len(elems))
		for i, elem := range elems {
			val, err := strconv.Atoi(elem)
			check(err)
			risks[i] = val
		}
		grid = append(grid, risks)
	}
	return dijkstra(grid)
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	//     fmt.Printf("Problem 2: %v\n",problem2())
}
