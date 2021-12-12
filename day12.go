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

func findPaths(nodes map[string][]string, node string, paths []string, path string, visited string, allowTwice bool) []string {
	if node == "end" {
		return append(paths, path)
	}
	if strings.ToLower(node) == node {
		visited += "," + node
	}
	if strings.Count(visited, node) == 2 {
		allowTwice = false
	}
	for _, elem := range nodes[node] {
		if elem == "start" || (strings.Contains(visited, elem) && !allowTwice) {
			continue
		}
		paths = findPaths(nodes, elem, paths, path+","+elem, visited, allowTwice)
	}
	return paths
}

func readInput(filename string) map[string][]string {
	file, err := os.Open(filename)
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	nodes := make(map[string][]string)
	for scanner.Scan() {
		vertex := strings.Split(scanner.Text(), "-")
		nodes[vertex[0]] = append(nodes[vertex[0]], vertex[1])
		nodes[vertex[1]] = append(nodes[vertex[1]], vertex[0])
	}
	return nodes
}

func showPaths(paths []string) {
	for _, path := range paths {
		fmt.Println(path)
	}
}

func problem(allowTwice bool) int {
	nodes := readInput("input12.txt")
	var paths []string
	paths = findPaths(nodes, "start", paths, "start", "", allowTwice)
	//showPaths(paths)
	return len(paths)
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem(false))
	fmt.Printf("Problem 2: %v\n", problem(true))
}
