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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func problem1() (fuel int) {
	file, err := os.Open("input07.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	positionsStr := strings.Split(scanner.Text(), ",")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	positions := make([]int, len(positionsStr))
	for i, elem := range positionsStr {
		val, err := strconv.Atoi(elem)
		check(err)
		positions[i] = val
	}
	sort.Ints(positions)
	var median int
	length := len(positions)
	switch {
	case length%2 == 0:
		midpoint := length / 2
		median = (positions[midpoint-1] + positions[midpoint]) / 2
	case length%2 != 0:
		median = positions[length/2+1]
	default:
		log.Fatal("Unreachable")
	}
	for _, elem := range positions {
		check(err)
		fuel += abs(elem - median)
	}
	return fuel
}

// idea/hint from here: https://www.reddit.com/r/adventofcode/comments/rawxad/2021_day_7_part_2_i_wrote_a_paper_on_todays/
func problem2() (fuel int) {
	file, err := os.Open("input07.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	positionsStr := strings.Split(scanner.Text(), ",")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	positions := make([]int, len(positionsStr))
	var mean int
	for i, elem := range positionsStr {
		val, err := strconv.Atoi(elem)
		check(err)
		positions[i] = val
		mean += val
	}
	mean = mean / len(positions)
	for i := mean - 1; i <= mean+1; i++ {
		var fuelTmp int
		for _, elem := range positions {
			step := abs(elem - i)
			fuelTmp += step * (1 + step) / 2 // arihmetic progression
		}
		if fuel == 0 || fuelTmp < fuel {
			fuel = fuelTmp
		}
	}
	return fuel
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
