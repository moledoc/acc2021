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

// kinda brute forcing it, not deducing numbers
func problem1() (unique int) {
	file, err := os.Open("input08.txt")
	defer file.Close()
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		io := strings.Split(scanner.Text(), " | ")
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

func sortStr(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func mapLetters() {

}

// Could use some refactoring, but currently will not do it.
func problem2() (sum int) {
	file, err := os.Open("input08.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		io := strings.Split(scanner.Text(), " | ")
		input := io[0]

		// find the freq letters to identify the letter positions (from 0 to 6)
		// count 6 identifies position 1
		// count 4 identifies position 4
		// count 9 identifies position 5
		// count 8 identifies positions 0 and 2
		// // by comparing numbers 1 (length = 2) and 7 (lenth = 3),
		// // then 7 contains position 0, but 1 don't,
		// // so we can identify position 0 by that.
		// // Position 2 can be inferred from position 0 and count 8.
		// count 7 identifies positions 3 and 6
		// // Number 4 contains position 3 and not position 6.
		// // Position 6 can be inferred from position 3 and count 7

		elemCounter := make(map[string]int, 7)
		uniques := make(map[int]string, 4)
		position := make([]string, 7)
		// find freq of the letters and identify uniques with length 3 and 4
		for _, inp := range strings.Split(input, " ") {
			length := len(inp)
			for _, elem := range strings.Split(inp, "") {
				if length == 2 || length == 4 {
					uniques[length] = inp
				}
				elemCounter[elem]++
			}
		}
		// map letter with position in display
		for key, count := range elemCounter {
			var pos int
			switch count {
			case 6:
				pos = 1
			case 4:
				pos = 4
			case 9:
				pos = 5
			case 8:
				if strings.Contains(uniques[2], key) {
					pos = 2
				} else {
					pos = 0
				}
			case 7:
				if strings.Contains(uniques[4], key) {
					pos = 3
				} else {
					pos = 6
				}
			default:
				log.Fatal("Unreachable")
			}
			position[pos] = key
		}

		displayMap := map[string]string{
			position[0] + position[1] + position[2] + position[4] + position[5] + position[6]: "0",
			position[2] + position[5]: "1",
			position[0] + position[2] + position[3] + position[4] + position[6]:                             "2",
			position[0] + position[2] + position[3] + position[5] + position[6]:                             "3",
			position[1] + position[2] + position[3] + position[5]:                                           "4",
			position[0] + position[1] + position[3] + position[5] + position[6]:                             "5",
			position[0] + position[1] + position[3] + position[4] + position[5] + position[6]:               "6",
			position[0] + position[2] + position[5]:                                                         "7",
			position[0] + position[1] + position[2] + position[3] + position[4] + position[5] + position[6]: "8",
			position[0] + position[1] + position[2] + position[3] + position[5] + position[6]:               "9",
		}
		// make the keys sorted
		for key, elem := range displayMap {
			newKey := sortStr(key)
			displayMap[sortStr(key)] = elem
			if newKey != key {
				delete(displayMap, key)
			}
		}

		// decode output
		output := io[1]
		var decoded string
		for _, coded := range strings.Split(output, " ") {
			decoded += displayMap[sortStr(coded)]
		}
		decodedInt, err := strconv.Atoi(decoded)
		if err != nil {
			log.Fatal(err)
		}
		sum += decodedInt
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
