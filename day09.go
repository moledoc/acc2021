package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var maxInt int = int(^uint(0) >> 1)

func addPadding(rowLen int, heightmap [][]int) [][]int {
	padding := make([]int, rowLen)
	for i := 0; i < rowLen; i++ {
		padding[i] = maxInt
	}
	heightmap = append(heightmap, padding)
	return heightmap
}

func makeHeightmap(scanner *bufio.Scanner) [][]int {
	var heightmap [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		if len(heightmap) == 0 {
			heightmap = addPadding(len(line)+2, heightmap)
		}
		row := make([]int, len(line)+2)
		for i, elem := range line {
			val, err := strconv.Atoi(elem)
			check(err)
			row[i+1] = val
		}
		row[0] = maxInt
		row[len(row)-1] = maxInt
		heightmap = append(heightmap, row)
	}
	return addPadding(len(heightmap[0]), heightmap)
}

func problem1() (riskSum int) {
	file, err := os.Open("input09.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	heightmap := makeHeightmap(scanner)
	for row := 1; row < len(heightmap)-1; row++ {
		for col := 1; col < len(heightmap[row])-1; col++ {
			if heightmap[row][col] < heightmap[row][col+1] &&
				heightmap[row][col] < heightmap[row][col-1] &&
				heightmap[row][col] < heightmap[row+1][col] &&
				heightmap[row][col] < heightmap[row-1][col] {
				riskSum += heightmap[row][col] + 1
			}
		}
	}
	return riskSum
}

func makeVisited(row int, col int) [][]bool {
	var visited [][]bool
	for i := 0; i < row; i++ {
		arow := make([]bool, col)
		visited = append(visited, arow)
	}
	return visited
}

// func showVisited(visited [][]bool) {
// 	fmt.Println()
// 	for i := 0; i < len(visited); i++ {
// 		fmt.Println(visited[i])
// 	}
// }

// currently overcounts
func findBasinSize(heightmap [][]int, row int, col int, count int, visited [][]bool) (int, [][]bool) {
	if heightmap[row][col] == 9 {
		visited[row][col] = false
		return count - 1, visited
	}
	nextVals := []int{
		row + 1, col,
		row - 1, col,
		row, col + 1,
		row, col - 1,
	}
	for i := 0; i < len(nextVals); i += 2 {
		nextRow := nextVals[i]
		nextCol := nextVals[i+1]
		if !visited[nextRow][nextCol] && heightmap[nextRow][nextCol] != maxInt && heightmap[row][col] < heightmap[nextRow][nextCol] {
			visited[nextRow][nextCol] = true
			count, visited = findBasinSize(heightmap, nextRow, nextCol, count+1, visited)
		}
	}
	return count, visited
}

func problem2() int {
	file, err := os.Open("input09.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	heightmap := makeHeightmap(scanner)
	var basinSizes []int
	for row := 1; row < len(heightmap)-1; row++ {
		for col := 1; col < len(heightmap[row])-1; col++ {
			if heightmap[row][col] < heightmap[row][col+1] &&
				heightmap[row][col] < heightmap[row][col-1] &&
				heightmap[row][col] < heightmap[row+1][col] &&
				heightmap[row][col] < heightmap[row-1][col] {
				visited := makeVisited(len(heightmap), len(heightmap[0]))
				visited[row][col] = true
				basinSize, visited := findBasinSize(heightmap, row, col, 1, visited)
				// 				showVisited(visited)
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}
	sort.Ints(basinSizes)
	basinsLen := len(basinSizes)
	basins := 1
	for i := 1; i <= 3; i++ {
		basins *= basinSizes[basinsLen-i]
	}
	return basins
}
func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
