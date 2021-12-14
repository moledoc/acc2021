package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func problem(steps int) int {
	file, err := os.Open("input14.txt")
	//file, err := os.Open("sample.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	scanner.Scan()
	formula := make(map[string]int)
	template := scanner.Text()
	for i := 1; i < len(template); i++ {
		formula[string(template[i-1])+string(template[i])] += 1
	}
	mapping := make(map[string]string)
	for scanner.Scan() {
		elems := strings.Split(scanner.Text(), " -> ")
		if len(elems) == 1 {
			continue
		}
		mapping[elems[0]] = elems[1]
	}

	for step := 0; step < steps; step++ {
		tmp := make(map[string]int)
		for key, elem := range formula {
			e := mapping[key]
			a := string(key[0]) + e
			b := e + string(key[1])
			// NOTE: adding e count twice, so I have to divide by 2 in the end
			tmp[a] += elem
			tmp[b] += elem
		}
		formula = tmp
	}
	counter := make(map[string]int)
	for key, elem := range formula {
		counter[string(key[0])] += elem
		counter[string(key[1])] += elem
	}
	var least int = int(^uint(0) >> 1)
	var greatest int
	for _, elem := range counter {
		if elem > greatest {
			greatest = elem
			continue
		}
		if elem < least {
			least = elem
			continue
		}
	}
	return ((greatest + greatest%2) / 2) - ((least + least%2) / 2)
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem(10))
	fmt.Printf("Problem 2: %v\n", problem(40))
}
