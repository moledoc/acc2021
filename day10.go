package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var errScore map[string]int = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}
var acScore map[string]int = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var mappingClose map[string]string = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var mappingOpen map[string]string = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func problem1() (score int) {
	file, err := os.Open("input10.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var filo []string
		for _, elem := range strings.Split(scanner.Text(), "") {
			errPt, ok := errScore[elem]
			if ok {
				if lastOpened := filo[len(filo)-1]; mappingClose[elem] != lastOpened {
					score += errPt
					break
				} else {
					filo = filo[:len(filo)-1]
					continue
				}
			}
			filo = append(filo, elem)
		}
	}
	return score
}

func problem2() int {
	file, err := os.Open("input10.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var compScores []int
	for scanner.Scan() {
		var syntaxErr bool
		var filo []string
		for _, elem := range strings.Split(scanner.Text(), "") {
			_, ok := errScore[elem]
			if ok {
				if lastOpened := filo[len(filo)-1]; mappingClose[elem] != lastOpened {
					syntaxErr = true
					break
				} else {
					filo = filo[:len(filo)-1]
					continue
				}
			}
			filo = append(filo, elem)
		}
		if syntaxErr {
			continue
		}
		var score int
		for len(filo) > 0 {
			score = score*5 + acScore[mappingOpen[filo[len(filo)-1]]]
			filo = filo[:len(filo)-1]
		}
		compScores = append(compScores, score)
	}
	sort.Ints(compScores)
	return compScores[len(compScores)/2]
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
