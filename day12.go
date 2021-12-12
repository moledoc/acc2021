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

func findPaths(nodes map[string][]string, node string, paths []string, path string, visited string) []string {
	if node == "end" {
		return append(paths, path)
	}
	if strings.ToLower(node) == node {
		visited += "," + node
	}
	for _, elem := range nodes[node] {
		if elem == "start" || strings.Contains(visited, elem) {
			continue
		}
		paths = findPaths(nodes, elem, paths, path+","+elem, visited)
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

func problem() int {
	//nodes := readInput("input12.txt")
	nodes := readInput("sample.txt")
	//fmt.Println(nodes)
	var paths []string
	paths = findPaths(nodes, "start", paths, "start", "")
	// 	fmt.Println(paths)
	// 	for i, elem := range paths {
	// 		fmt.Println(i, elem)
	// 	}
	return len(paths)
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem())
	//     fmt.Printf("Problem 2: %v\n",problem2())
}
