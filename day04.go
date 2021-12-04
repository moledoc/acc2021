package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// NB This code is not particularly well written nor efficient.
// Unfortunately I don't have too much time, so working solution is the nr 1 priority.

func parseNumbers(scanner *bufio.Scanner) (numbers []int) {
	if scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, elem := range line {
			nr, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, nr)
		}
	}
	return numbers
}

func parseBoard(scanner *bufio.Scanner) map[int][]int {
	board := make(map[int][]int)
	var row int
	for scanner.Scan() {
		lineStr := scanner.Text()
		var line []int
		for _, elem := range strings.Split(lineStr, " ") {
			if elem == "" {
				continue
			}
			elemInt, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, elemInt)
		}
		for i, elem := range line {
			// rows
			board[row] = append(board[row], elem)
			// cols
			board[5+i] = append(board[5+i], elem)
		}
		row++
		if row == 5 {
			break
		}
	}
	return board
}

func findWinning(boards map[int]map[int][]int, numbers []int, findFirst bool) (map[int][]int, int, map[int]int) {
	counter := make(map[int]map[int]int)
	for i := 0; i < len(boards); i++ {
		tmp := make(map[int]int)
		for j := 0; j < len(boards[i]); j++ {
			tmp[j] = 0
		}
		counter[i] = tmp
	}
	ranking := make(map[int]int, len(boards))
	var rank int
	var board map[int][]int
	for last, nr := range numbers {
		for i, board := range boards {
			for j, boardRow := range board {
				for _, elem := range boardRow {
					if elem == nr {
						counter[i][j] += 1
						if counter[i][j] == 5 {
							if findFirst {
								return board, last, ranking
							}
							if _, ok := ranking[i]; !ok {
								ranking[i] = rank
							}
							rank++
							if len(ranking) == len(boards) {
								return board, last, ranking
							}
						}
					}
				}
			}
		}
	}
	return board, -1, ranking
}

func problem1() int {
	file, err := os.Open("input04.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// get numbers
	numbers := parseNumbers(scanner)
	// get boards
	var boardNr int
	boards := make(map[int]map[int][]int)
	for scanner.Scan() {
		boards[boardNr] = parseBoard(scanner)
		boardNr++
	}
	// find winning board
	winningBoard, lastNrInd, _ := findWinning(boards, numbers, true)
	// insert used in a map
	usedNr := make(map[int]bool)
	for i := 0; i <= lastNrInd; i++ {
		usedNr[numbers[i]] = true
	}
	// find not used elems
	var unmarked int
	for i := 0; i < 5; i++ {
		// rows end with index 4
		// since columns contain the same numbers as rows,
		// then summing them would double our result
		for _, elem := range winningBoard[i] {
			if _, ok := usedNr[elem]; ok {
				continue
			}
			unmarked += elem
		}
	}
	return unmarked * numbers[lastNrInd]
}

func problem2() int {
	file, err := os.Open("input04.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// get numbers
	numbers := parseNumbers(scanner)
	// get boards
	var boardNr int
	boards := make(map[int]map[int][]int)
	for scanner.Scan() {
		boards[boardNr] = parseBoard(scanner)
		boardNr++
	}
	// find last winning board
	winningBoard, lastNrInd, ranking := findWinning(boards, numbers, false)
	for i := 0; i < len(boards); i++ {
		if ranking[i] == len(boards)-1 {
			winningBoard = boards[i]
			break
		}
	}
	// insert used in a map
	usedNr := make(map[int]bool)
	for i := 0; i <= lastNrInd; i++ {
		usedNr[numbers[i]] = true
	}
	// find not used elems
	var unmarked int
	for i := 0; i < 5; i++ {
		// rows end with index 4
		// since columns contain the same numbers as rows,
		// then summing them would double our result
		for _, elem := range winningBoard[i] {
			if _, ok := usedNr[elem]; ok {
				continue
			}
			unmarked += elem
		}
	}
	return unmarked * numbers[lastNrInd]
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
