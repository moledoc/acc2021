package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
var verSum int
var polish []string

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
	if id == 4 {
		ptr = literal(binary, ptr)
	}
	if id != 4 {
		lenType, ptr = lenTypeId(binary, ptr)
		var op string
		switch id {
		case 0:
			op = "+"
		case 1:
			op = "*"
		case 2:
			op = "min"
		case 3:
			op = "max"
		case 5:
			op = ">"
		case 6:
			op = "<"
		case 7:
			op = "="
		}
		polish = append(polish, op)
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
func literal(binary string, ptr int) int {
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
	lit64, err := strconv.ParseInt(literalStr, 2, 64)
	check(err)
	polish = append(polish, fmt.Sprint(lit64))
	return ptr
}

func operator0(binary string, ptr int) int {
	subPackLen, err := strconv.ParseInt(binary[ptr:ptr+15], 2, 64)
	ptr += 15
	check(err)
	for i := ptr; ptr-i < int(subPackLen); {
		ptr = packet(binary, ptr)
	}
	polish = append(polish, "|")
	return ptr
}

func operator1(binary string, ptr int) int {
	subPackCnt, err := strconv.ParseInt(binary[ptr:ptr+11], 2, 64)
	ptr += 11
	check(err)
	for i := 0; i < int(subPackCnt); i++ {
		ptr = packet(binary, ptr)
	}
	polish = append(polish, "|")
	return ptr
}

func problem1() int {
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

var ops *regexp.Regexp = regexp.MustCompile("\\+|\\*|min|max|>|<|=")

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func parsePolish(pol []string) int {
	var result int
	for len(pol) != 1 {
		elem := pol[0]
		if ops.MatchString(pol[1]) || elem == "|" {
			pol = append(pol, elem)
			pol = append(pol[:0], pol[1:]...)
			continue
		}
		var done bool
		if elem == "+" {
			var j int = 1
			var sum int
			for ; j < len(pol) && !ops.MatchString(pol[j]); j++ {
				if pol[j] == "|" {
					done = true
					j++
					break
				}
				a, err := strconv.Atoi(pol[j])
				check(err)
				sum += a
			}
			pol = append(pol[:0], pol[j:]...)
			if !done {
				pol = append(pol, "+")
			}
			pol = append(pol, fmt.Sprint(sum))
			continue
		}
		if elem == "*" {
			var j int = 1
			var prod int = 1
			for ; j < len(pol) && !ops.MatchString(pol[j]); j++ {
				if pol[j] == "|" {
					done = true
					j++
					break
				}
				a, err := strconv.Atoi(pol[j])
				check(err)
				prod *= a
			}
			pol = append(pol[:0], pol[j:]...)
			if !done {
				pol = append(pol, "*")
			}
			pol = append(pol, fmt.Sprint(prod))
			continue
		}
		if elem == "min" {
			var min int = int(^uint(0) >> 1)
			var j int = 1
			for ; j < len(pol) && !ops.MatchString(pol[j]); j++ {
				if pol[j] == "|" {
					done = true
					j++
					break
				}
				a, err := strconv.Atoi(pol[j])
				check(err)
				if a < min {
					min = a
				}
			}
			pol = append(pol[:0], pol[j:]...)
			if !done {
				pol = append(pol, "min")
			}
			pol = append(pol, fmt.Sprint(min))
			continue
		}
		if elem == "max" {
			var max int
			var j int = 1
			for ; j < len(pol) && !ops.MatchString(pol[j]); j++ {
				if pol[j] == "|" {
					done = true
					j++
					break
				}
				a, err := strconv.Atoi(pol[j])
				check(err)
				if a > max {
					max = a
				}
			}
			pol = append(pol[:0], pol[j:]...)
			if !done {
				pol = append(pol, "max")
			}
			pol = append(pol, fmt.Sprint(max))
			continue
		}
		if elem == "<" {
			a, err := strconv.Atoi(pol[1])
			check(err)
			b, err := strconv.Atoi(pol[2])
			check(err)
			pol = append(pol[:0], pol[3+1:]...) // the +1 for removing |
			pol = append(pol, fmt.Sprint(boolToInt(a < b)))
			continue
		}
		if elem == ">" {
			a, err := strconv.Atoi(pol[1])
			check(err)
			b, err := strconv.Atoi(pol[2])
			check(err)
			pol = append(pol[:0], pol[3+1:]...) // the +1 for removing |
			pol = append(pol, fmt.Sprint(boolToInt(a > b)))
			continue
		}
		if elem == "=" {
			a, err := strconv.Atoi(pol[1])
			check(err)
			b, err := strconv.Atoi(pol[2])
			check(err)
			pol = append(pol[:0], pol[3+1:]...) // the +1 for removing |
			pol = append(pol, fmt.Sprint(boolToInt(a == b)))
			continue
		}
		pol = append(pol[:0], pol[1:]...)
		pol = append(pol, elem)
	}
	result, err := strconv.Atoi(pol[0])
	check(err)
	return result
}

func problem2() int {
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

	var tmp []string
	polish = tmp
	_ = packet(binary, 0)
	// 	fmt.Println(polish)
	return parsePolish(polish)
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
