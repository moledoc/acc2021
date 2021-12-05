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

func parseNumbers(scanner *bufio.Scanner) (numbers []int) {
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		numbers = make([]int, len(line))
		for i, elem := range line {
			nr, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			}
			numbers[i] = nr
		}
	}
	return numbers
}

func bingo(scanner *bufio.Scanner, numbers []int) bingoInfo {
	board := [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	// for each row/col, keep counter: index 0 = row, index 1 = col
	counter := [5][2]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}
	var unmarkedSum int
	// read in board
	for i := 0; i < 5 && scanner.Scan(); i++ {
		var colInd int
		for _, elem := range strings.Split(scanner.Text(), " ") {
			if elem == "" {
				continue
			}
			nr, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			}
			board[i][colInd] = nr
			colInd++
			unmarkedSum += nr
		}
	}
	// find when board wins
	for i := 0; i < len(numbers); i++ {
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				if numbers[i] == board[row][col] {
					unmarkedSum -= numbers[i]
					counter[row][0]++
					counter[col][1]++
					if counter[row][0] == 5 || counter[col][1] == 5 {
						return bingoInfo{
							lastNr:      numbers[i],
							lastNrInd:   i,
							unmarkedSum: unmarkedSum,
						}
					}
					break
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
	numbers := parseNumbers(scanner)
	// get each board bingo info
	var bingos []bingoInfo
	for scanner.Scan() {
		bingos = append(bingos, bingo(scanner, numbers))
	}
	// get the first and last winner results
	var first int = len(numbers)
	var last int
	var firstWin int
	var lastWin int
	for _, bingo := range bingos {
		if first > bingo.lastNrInd {
			first = bingo.lastNrInd
			firstWin = bingo.lastNr * bingo.unmarkedSum
		}
		if last < bingo.lastNrInd {
			last = bingo.lastNrInd
			lastWin = bingo.lastNr * bingo.unmarkedSum
		}
	}
	return firstWin, lastWin
}

func main() {
	problem1, problem2 := problem()
	fmt.Printf("Problem 1: %v\nProblem 2: %v\n", problem1, problem2)
}
