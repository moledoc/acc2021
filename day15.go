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
var gridSize int
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

func dumpPath(grid [][]int, prev []node) {
	fmt.Println()
	for i, row := range grid {
		for j := range row {
			inPath := false
			for _, n := range prev {
				if n.row == i && n.col == j {
					fmt.Printf("%v]", grid[i][j])
					inPath = true
					break
				}
			}
			if inPath {
				continue
			}
			fmt.Printf("%v ", grid[i][j])
		}
		fmt.Println()
	}
}

func minDist(dist []int, queue []node) (ind int) {
	unreachable := node{-1, -1, -1}
	var min int = inf
	for i, elem := range dist {
		if elem < min && queue[i] != unreachable {
			min = elem
			ind = i
		}
	}
	return ind
}

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
// It needs better variable names and more optimal structres.
func dijkstra(grid [][]int) int {
	unreachable := node{-1, -1, -1}
	source := node{0, 0, 0}
	queue := make([]node, gridSize*gridSize)
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
	queue[0] = source
	qCtr := len(queue)
	var u node
	for qCtr > 0 {
		safest := minDist(dist, queue)
		point := queue[safest]
		queue[safest] = unreachable
		qCtr--
		if safest == gridSize*gridSize-1 {
			u = point
			break
		}
		for i := 0; i < len(dirs); i += 2 {
			newRow := point.row + dirs[i]
			newCol := point.col + dirs[i+1]
			next := newRow*gridSize + newCol
			if newRow < 0 ||
				newCol < 0 ||
				newRow >= gridSize ||
				newCol >= gridSize ||
				queue[next] == unreachable {
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
	// 	var prevs []node
	for {
		// 		prevs = append(prevs, u)
		i := u.row*gridSize + u.col
		if !(prev[i] != unreachable || prev[i] == source) {
			break
		}
		riskSum += u.dist
		u = prev[i]
	}
	// 	dumpPath(grid, prevs)
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
		if j == 0 {
			gridSize = len(elems)
		}
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

func problem2() int {
	file, err := os.Open(filename)
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	var grid [][]int
	for j := 0; scanner.Scan(); j++ {
		elems := strings.Split(scanner.Text(), "")
		elen := len(elems)
		if j == 0 {
			gridSize = elen * 5
			grid = make([][]int, gridSize)
		}
		risks := make([]int, elen*5)
		for i, elem := range elems {
			val, err := strconv.Atoi(elem)
			check(err)
			for k := 0; k < 5; k++ {
				risks[k*elen+i] = val + k
				if risks[k*elen+i] > 9 {
					risks[k*elen+i] = risks[k*elen+i] - 9
				}
			}
		}
		grid[j] = risks
	}
	gridSizeOrig := gridSize / 5
	for i := gridSizeOrig; i < gridSize; i++ {
		row := make([]int, gridSize)
		for j, elem := range grid[i-gridSizeOrig] {
			incr := elem + 1
			if elem+1 > 9 {
				incr = 1
			}
			row[j] = incr
		}
		grid[i] = row
	}
	// 	dump(grid)
	return dijkstra(grid)
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
