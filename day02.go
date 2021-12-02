package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func problem01() int {
	file, err := os.Open("input02.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	var forward int
	var depth int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), " ")
		direction := instructions[0]
		valStr := instructions[1]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "up":
			depth -= val
		case "down":
			depth += val
		case "forward":
			forward += val
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return forward * depth
}

func problem02() int {
	file, err := os.Open("input02.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	var forward int
	var depth int
	var aim int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), " ")
		direction := instructions[0]
		valStr := instructions[1]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "up":
			aim -= val
		case "down":
			aim += val
		case "forward":
			forward += val
			depth += aim * val
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return forward * depth
}

func main() {
	// 01
	fmt.Printf("Problem 01: %v\n", problem01())
	// 02
	fmt.Printf("Problem 02: %v\n", problem02())
}
