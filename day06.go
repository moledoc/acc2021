package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sum(pop []int) (summed int) {
	for _, elem := range pop {
		summed += elem
	}
	return summed
}

func problem(days int) int {
	file, err := os.Open("input06.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	popStr := strings.Split(scanner.Text(), ",")
	population := make([]int, 10)
	for _, elem := range popStr {
		elemInt, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatal(err)
		}
		population[elemInt]++
	}
	populationPrev := population
	for day := 1; day <= days; day++ {
		population[7] += population[0]
		population[9] += population[0]
		population[0] = 0
		for i := 1; i < 10; i++ {
			population[i-1] += population[i]
			population[i] -= populationPrev[i]
		}
		populationPrev = population
	}
	return sum(population)

}

func main() {
	fmt.Printf("Problem 1: %v\n", problem(80))
	fmt.Printf("Problem 2: %v\n", problem(256))
}
