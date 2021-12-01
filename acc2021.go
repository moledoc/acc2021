package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sample map[int]int = map[int]int{
	0: 199,
	1: 200,
	2: 208,
	3: 210,
	4: 200,
	5: 207,
	6: 240,
	7: 269,
	8: 260,
	9: 263,
}

func readFile(fname string) map[int]int {
	input := make(map[int]int)
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input[i] = val
		i++
	}
	return input
}

// 01
func problem01(input map[int]int) (increased int) {
	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			increased++
		}
	}
	return increased
}

// 02
func problem02(input map[int]int) (increased int) {
	var prev int = input[0] + input[1] + input[2]
	for i := 3; i < len(input); i++ {
		next := prev - input[i-3] + input[i]
		if prev < next {
			increased++
		}
		prev = next
	}
	return increased
}

func main() {
	_ = sample
	// read input
	input := readFile("input.txt")
	// 01
	fmt.Printf("Problem 01: %v\n", problem01(input))
	// 02
	fmt.Printf("Problem 02: %v\n", problem02(input))
}
