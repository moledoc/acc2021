package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// kinda brute forcing it, not deducing numbers
func problem1() (unique int) {
	// file, err := os.Open("sample.txt")
	file, err := os.Open("input08.txt")
	defer file.Close()
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		io := strings.Split(scanner.Text(), " | ")
		// input := io[0]
		output := io[1]
		for _, elem := range strings.Split(output, " ") {
			length := len(elem)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				unique += 1
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return unique
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	//     fmt.Printf("Problem 2: %v\n",problem2())
}
