package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func checkScanner(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func printPretty(paper [][]int) {
	fmt.Println()
	var pretty [][]string
	for row := 0; row < len(paper); row++ {
		prettyRow := make([]string, len(paper[0]))
		for col := 0; col < len(paper[0]); col++ {
			symb := " "
			if paper[row][col] > 0 {
				symb = "."
			}
			prettyRow[col] = symb
		}
		pretty = append(pretty, prettyRow)
	}
	for row := 0; row < len(pretty); row++ {
		fmt.Println(pretty[row])
	}
}

func printer(arr [][]int, part int) (result int) {
	var xMax int
	var yMax int
	for i := 0; i < len(arr[0]); i++ {
		if xMax < arr[0][i] {
			xMax = arr[0][i]
		}
		if yMax < arr[1][i] {
			yMax = arr[1][i]
		}
	}
	xMax++
	yMax++
	pretty := make([][]int, yMax)
	for i := 0; i < yMax; i++ {
		tmp := make([]int, xMax)
		pretty[i] = tmp
	}
	for i := 0; i < len(arr[0]); i++ {
		pretty[arr[1][i]][arr[0][i]] = 1
	}
	for row := 0; row < len(pretty); row++ {
		for col := 0; col < len(pretty[0]); col++ {
			result += pretty[row][col]
		}
	}
	if part == 2 {
		printPretty(pretty)
	}
	return result
}

func updCoord(arr [][]int, ind int, fold int) [][]int {
	for i, elem := range arr[ind] {
		if fold < elem {
			arr[ind][i] = 2*fold - elem
		}
	}
	return arr
}

func problem(part int) int {
	file, err := os.Open("input13.txt")
	//file, err := os.Open("sample.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	arr := make([][]int, 2)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		check(err)
		y, err := strconv.Atoi(coords[1])
		check(err)
		arr[0] = append(arr[0], x)
		arr[1] = append(arr[1], y)
	}
	for scanner.Scan() {
		instr := strings.Split(scanner.Text(), "=")
		fold, err := strconv.Atoi(instr[1])
		check(err)
		if strings.Contains(instr[0], "x") {
			arr = updCoord(arr, 0, fold)
		}
		if strings.Contains(instr[0], "y") {
			arr = updCoord(arr, 1, fold)
		}
		if part == 1 {
			break
		}
	}
	return printer(arr, part)
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem(1))
	fmt.Printf("Problem 2: %v\n", problem(2))
	// fmt.Printf("Problem 1: %v\n", problem1("sample.txt"))
	//fmt.Printf("Problem 1: %v\n", problem1("input13.txt"))
	//     fmt.Printf("Problem 2: %v\n",problem2())
}
