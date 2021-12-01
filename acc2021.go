package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func main() {
	// read input
	input := readFile("input.txt")
	// 01
	fmt.Println(problem01(input))
}
