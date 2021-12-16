package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var verSum int
var hexToBin map[string]string = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

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

func version(binary string, ptr int) int {
	verStr := binary[ptr : ptr+3]
	ptr += 3
	ver64, err := strconv.ParseInt(verStr, 2, 64)
	check(err)
	// 	fmt.Println("VER", ver64)
	verSum += int(ver64)
	return ptr
}

func typeId(binary string, ptr int) (int, int) {
	idStr := binary[ptr : ptr+3]
	ptr += 3
	id64, err := strconv.ParseInt(idStr, 2, 64)
	check(err)
	return int(id64), ptr
}
func lenTypeId(binary string, ptr int) (int, int) {
	idStr := binary[ptr : ptr+1]
	ptr += 1
	id64, err := strconv.ParseInt(idStr, 2, 64)
	check(err)
	return int(id64), ptr
}

func packet(binary string, ptr int) int {
	var id int
	var lenType int
	ptr = version(binary, ptr)
	id, ptr = typeId(binary, ptr)
	// 	fmt.Println("HERE", ptr, id)
	if id == 4 {
		_, ptr = literal(binary, ptr)
	}
	if id != 4 {
		lenType, ptr = lenTypeId(binary, ptr)
		// 		fmt.Println("HERE2", lenType)
		// 		if id == 6 && lenType == 0 {
		if lenType == 0 {
			ptr = operator0(binary, ptr)
		}
		// 		if id == 3 && lenType == 1 {
		if lenType == 1 {
			ptr = operator1(binary, ptr)
		}
	}
	return ptr
}

// assume that version and typeId are already run
func literal(binary string, ptr int) (int, int) {
	var stop bool
	var literalStr string
	for !stop {
		part := binary[ptr : ptr+5]
		ptr += 5
		if string(part[0]) == "0" {
			stop = true
		}
		literalStr += part[1:5]
	}
	// 	ptr += 3 // skip last 3 bits, because they are for hex representation
	lit64, err := strconv.ParseInt(literalStr, 2, 64)
	// 	fmt.Println("LIT", lit64)
	check(err)
	return int(lit64), ptr
}

func operator0(binary string, ptr int) int {
	subPackLen, err := strconv.ParseInt(binary[ptr:ptr+15], 2, 64)
	ptr += 15
	check(err)
	for i := ptr; ptr-i < int(subPackLen); {
		ptr = packet(binary, ptr)
		// 		fmt.Println("HERERER", i, ptr, ptr-i, subPackLen)
	}
	//ptr += len(binary) - ptr // ??
	return ptr
}

func operator1(binary string, ptr int) int {
	subPackCnt, err := strconv.ParseInt(binary[ptr:ptr+11], 2, 64)
	ptr += 11
	check(err)
	for i := 0; i < int(subPackCnt); i++ {
		ptr = packet(binary, ptr)
	}
	// 	ptr += len(binary) - ptr // ??
	return ptr
}

func problem1() int {
	//file, err := os.Open("sample.txt")
	file, err := os.Open("input16.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	var binary string
	scanner.Scan()
	elems := strings.Split(scanner.Text(), "")
	for _, elem := range elems {
		binary += hexToBin[elem]
	}
	// 	fmt.Println(binary)
	_ = packet(binary, 0)
	return verSum
}

func problem2() int {
	file, err := os.Open("sample.txt")
	//file, err := os.Open("input16.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	var binary string
	scanner.Scan()
	elems := strings.Split(scanner.Text(), "")
	for _, elem := range elems {
		binary += hexToBin[elem]
	}
	// 	fmt.Println(binary)
	_ = packet(binary, 0)
	return verSum
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	//     fmt.Printf("Problem 2: %v\n",problem2())
}
