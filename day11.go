package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var mapSize int = 10

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

func makeFlashed() [][]bool {
	flashed := make([][]bool, mapSize)
	for i := 0; i < mapSize; i++ {
		row := make([]bool, mapSize)
		flashed[i] = row
	}
	return flashed
}

func flashSimu(octomap [][]int, row int, col int, flashed [][]bool, count int) ([][]int, int, [][]bool) {
	octomap[row][col] = 0
	var sur []int = []int{
		row + 1, col,
		row + 1, col + 1,
		row + 1, col - 1,
		row - 1, col,
		row - 1, col + 1,
		row - 1, col - 1,
		row, col + 1,
		row, col - 1,
	}
	for i := 0; i < len(sur); i += 2 {
		if sur[i] < 0 || sur[i+1] < 0 || sur[i] > 9 || sur[i+1] > 9 {
			continue
		}
		if !flashed[sur[i]][sur[i+1]] {
			octomap[sur[i]][sur[i+1]] += 1
			if octomap[sur[i]][sur[i+1]] > 9 {
				flashed[sur[i]][sur[i+1]] = true
				octomap, count, flashed = flashSimu(octomap, sur[i], sur[i+1], flashed, count+1)
			}
		}
	}

	return octomap, count, flashed
}

func step(octomap [][]int, count int) ([][]int, int) {
	flashed := makeFlashed()
	for row := 0; row < mapSize; row++ {
		for col := 0; col < mapSize; col++ {
			if !flashed[row][col] {
				octomap[row][col] += 1
			}
			if octomap[row][col] > 9 {
				flashed[row][col] = true
				octomap, count, flashed = flashSimu(octomap, row, col, flashed, count+1)
			}
		}
	}
	return octomap, count
}

func dumpMap(octomap [][]int) {
	for i := 0; i < mapSize; i++ {
		fmt.Println(octomap[i])
	}
	fmt.Println()
}

func problem1() (count int) {
	file, err := os.Open("input11.txt")
	check(err)
	defer file.Close()
	octomap := make([][]int, mapSize)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	for i := 0; scanner.Scan(); i++ {
		row := make([]int, mapSize)
		for i, elem := range strings.Split(scanner.Text(), "") {
			val, err := strconv.Atoi(elem)
			check(err)
			row[i] = val
		}
		octomap[i] = row
	}
	steps := 100
	//dumpMap(octomap)
	for i := 0; i < steps; i++ {
		octomap, count = step(octomap, count)
		// 		dumpMap(octomap)
	}
	return count
}

func isSynced(octomap [][]int) bool {
	for row := 0; row < mapSize; row++ {
		for col := 0; col < mapSize; col++ {
			if octomap[row][col] != 0 {
				return false
			}
		}
	}
	return true
}

func problem2() int {
	file, err := os.Open("input11.txt")
	check(err)
	defer file.Close()
	octomap := make([][]int, mapSize)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	for i := 0; scanner.Scan(); i++ {
		row := make([]int, mapSize)
		for i, elem := range strings.Split(scanner.Text(), "") {
			val, err := strconv.Atoi(elem)
			check(err)
			row[i] = val
		}
		octomap[i] = row
	}

	var count int
	var curStep int
	// 	dumpMap(octomap)
	for ; !isSynced(octomap); curStep++ {
		octomap, count = step(octomap, count)
		// 		dumpMap(octomap)
	}
	return curStep
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
