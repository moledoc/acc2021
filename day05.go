package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func comp(a, a2, inc int) bool {
	if (inc > 0 && a <= a2) || (inc < 0 && a >= a2) {
		return true
	}
	return false
}

func problem1() int {
	var pointsCounter [1000000]int
	file, err := os.Open("input05.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		fst := strings.Split(points[0], ",")
		snd := strings.Split(points[1], ",")
		x1, err := strconv.Atoi(fst[0])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(fst[1])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(snd[0])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(snd[1])
		if err != nil {
			log.Fatal(err)
		}
		xdiff := abs(x2 - x1)
		ydiff := abs(y2 - y1)
		diff := max(xdiff, ydiff)
		x := (x2 - x1) / diff
		y := (y2 - y1) / diff
		if x1 == x2 || y1 == y2 {
			for i := 0; i <= diff; i++ {
				pointsCounter[(x1+i*x)+1000*(y1+i*y)]++
			}
		}
	}
	var danger int
	for _, elem := range pointsCounter {
		if elem > 1 {
			danger++
		}
	}
	return danger
}

func problem2() int {
	var pointsCounter [1000000]int
	file, err := os.Open("input05.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		fst := strings.Split(points[0], ",")
		snd := strings.Split(points[1], ",")

		x1, err := strconv.Atoi(fst[0])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(fst[1])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(snd[0])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(snd[1])
		if err != nil {
			log.Fatal(err)
		}
		xdiff := abs(x2 - x1)
		ydiff := abs(y2 - y1)
		diff := max(xdiff, ydiff)
		x := (x2 - x1) / diff
		y := (y2 - y1) / diff
		for i := 0; i <= diff; i++ {
			pointsCounter[(x1+i*x)+1000*(y1+i*y)]++
		}
	}
	var danger int
	for _, elem := range pointsCounter {
		if elem > 1 {
			danger++
		}
	}
	return danger
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
