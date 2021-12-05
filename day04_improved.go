package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoInfo struct {
	lastNr, lastNrInd, unmarkedSum int
}

func parseNumbers(scanner *bufio.Scanner) (map[int]int, map[int]int) {
	var numbers map[int]int = make(map[int]int)
	var numbersInd map[int]int = make(map[int]int)
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for i, elem := range line {
			nr, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			}
			numbers[nr] = i
			numbersInd[i] = nr
		}
	}
	return numbers, numbersInd
}

func bingo(scanner *bufio.Scanner, numbers map[int]int, numbersInd map[int]int) bingoInfo {
	board := make(map[int]int, 25)
	// for each row/col, keep counter: index 0 = row, index 1 = col
	counter := map[int][]int{
		0: {0, 0},
		1: {0, 0},
		2: {0, 0},
		3: {0, 0},
		4: {0, 0},
	}
	var boardInd int
	var unmarkedSum int

	for i := 0; i < 5 && scanner.Scan(); i++ {
		for _, elem := range strings.Split(scanner.Text(), " ") {
			if elem == "" {
				continue
			}
			nr, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			}
			board[nr] = boardInd
			boardInd++
			unmarkedSum += nr
		}
	}
	for i := 0; i < len(numbersInd); i++ {
		if ind, ok := board[numbersInd[i]]; ok {
			unmarkedSum -= numbersInd[i]
			// row
			counter[ind/5][0] += 1
			// col
			counter[ind%5][1] += 1
			if counter[ind/5][0] == 5 || counter[ind%5][1] == 5 {
				return bingoInfo{
					lastNr:      numbersInd[i],
					lastNrInd:   i,
					unmarkedSum: unmarkedSum,
				}
			}
		}
	}

	return bingoInfo{}
}

func problem() (int, int) {
	file, err := os.Open("input04.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// get numbers
	numbers, numbersInd := parseNumbers(scanner)
	// get each board bingo info
	bingos := make(map[int]bingoInfo)
	for boardNr := 0; scanner.Scan(); boardNr++ {
		bingos[boardNr] = bingo(scanner, numbers, numbersInd)
	}
	// get the first and last winner indeces
	var first int = len(numbers)
	var firstInd int
	var last int
	var lastInd int
	for i, bingo := range bingos {
		if first > bingo.lastNrInd {
			first = bingo.lastNrInd
			firstInd = i
		}
		if last < bingo.lastNrInd {
			last = bingo.lastNrInd
			lastInd = i
		}
	}
	return bingos[firstInd].unmarkedSum * bingos[firstInd].lastNr, bingos[lastInd].unmarkedSum * bingos[lastInd].lastNr
}

func main() {
	problem1, problem2 := problem()
	fmt.Printf("Problem 1: %v\nProblem 2: %v\n", problem1, problem2)
}
