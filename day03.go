package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func problem01() int64 {
	// define counters
	counter0 := make(map[int]int)
	counter1 := make(map[int]int)
	// open file
	file, err := os.Open("input03.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// increment the corresponding counter
		for i := 0; i < len(line); i++ {
			switch string(line[i]) {
			case "0":
				counter0[i] += 1
			case "1":
				counter1[i] += 1
			default:
				log.Fatalf("Unexpected value: %v\n", string(line[i]))
			}
		}
	}
	// get gamma and epsilon binary representations.
	var gammaStr string
	var epsilonStr string
	for i := 0; i < len(counter0); i++ {
		switch {
		case counter0[i] < counter1[i]:
			gammaStr += "1"
			epsilonStr += "0"
		case counter0[i] > counter1[i]:
			gammaStr += "0"
			epsilonStr += "1"
		default:
			log.Fatal("Unreachable")
		}
	}
	// convert binary representation to int
	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return gamma * epsilon
}

func findRating(data []string, offset int) string {
	var bitNr int
	strLen := len(data[0])
	for i := 0; i < strLen; i++ {
		counter := make(map[int][]string)
		for _, elem := range data {
			switch string(elem[bitNr]) {
			case "0":
				counter[0] = append(counter[0], elem)
			case "1":
				counter[1] = append(counter[1], elem)
			default:
				log.Fatal("Unreachable")
			}
		}
		switch len(counter[0]) <= len(counter[1]) {
		case true:
			data = counter[(1+offset)%2]
		case false:
			data = counter[(0+offset)%2]
		}
		bitNr++
		if len(data) == 1 {
			return data[0]
		}
	}
	if len(data) != 1 {
		log.Fatal("Did not find the correct rating value, since there are more than 1 rating value in 'data'")
	}
	return data[0]
}

func problem02() int64 {
	file, err := os.Open("input03.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// read data to array, so we could define a function to find the rating for both O2 and CO2.
	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	o2Str := findRating(data, 0)
	co2Str := findRating(data, 1)
	o2, err := strconv.ParseInt(o2Str, 2, 64)
	co2, err := strconv.ParseInt(co2Str, 2, 64)
	return o2 * co2
}

func main() {
	// 01
	fmt.Printf("Problem 01: %v\n", problem01())
	// 02
	fmt.Printf("Problem 02: %v\n", problem02())
}
