package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func vel(p1, p2 float64) int {
	// 	fmt.Println("NEW", p1, p2)
	var vel int = -1
	for p := p1; p < p2; p++ {
		quadSol := (-1 + math.Sqrt(4*(p*2)+1)) / 2
		if quadSol == math.Floor(quadSol) {
			vel = int(quadSol)
			break
		}
	}
	return vel
}

func problem1() int {
	//file, err := os.Open("sample.txt")
	file, err := os.Open("input17.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	scanner.Scan()
	elems := strings.Split(scanner.Text(), "..")
	// 	x1, err := strconv.Atoi(strings.Split(elems[0], "=")[1])
	// 	check(err)
	// 	x2, err := strconv.Atoi(strings.Split(elems[1], ", ")[0])
	// 	check(err)
	y1, err := strconv.Atoi(strings.Split(elems[1], "=")[1])
	check(err)
	y2, err := strconv.Atoi(elems[2])
	check(err)
	// 	fmt.Println(x1, x2, y1, y2)
	//
	// 	var xvel int = vel(float64(x1), float64(x2))
	//
	var haveHitT bool
	var haveNotHitCtr int // allow to overshoot fixed amout of times after hitting target the first time
	var height int
	// 	var yvel int
	for i := 0; ; i++ {
		h := i * (i + 1) / 2
		tvel := vel(float64(h-y2), float64(h-y1))
		if tvel == -1 && haveHitT && haveNotHitCtr > 100 {
			break
		}
		if tvel == -1 {
			if haveHitT {
				haveNotHitCtr++
			}
			continue
		}
		haveHitT = true
		if height < tvel*(tvel+1)/2 {
			height = tvel * (tvel + 1) / 2
			// 			yvel = tvel
		}
	}
	// 	fmt.Println(xvel, yvel, height)
	return height
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func problem2() int {
	file, err := os.Open("input17.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	defer checkScanner(scanner)
	scanner.Scan()
	elems := strings.Split(scanner.Text(), "..")
	x1, err := strconv.Atoi(strings.Split(elems[0], "=")[1])
	check(err)
	x2, err := strconv.Atoi(strings.Split(elems[1], ", ")[0])
	check(err)
	y1, err := strconv.Atoi(strings.Split(elems[1], "=")[1])
	check(err)
	y2, err := strconv.Atoi(elems[2])
	check(err)
	var yStart int = -max(abs(y1), abs(y2))
	var yEnd int = max(abs(y1), abs(y2))
	var velos int
	for x := 0; x <= x2; x++ {
		for y := yStart; y <= yEnd; y++ {
			yvel := y
			xvel := x
			var xpos int
			var ypos int
			for i := 0; i < 10000; i++ {
				xpos += xvel
				ypos += yvel
				yvel -= 1
				xvel += boolToInt(xvel < 0) - boolToInt(xvel > 0)
				if xpos >= x1 && xpos <= x2 && ypos >= y1 && ypos <= y2 {
					velos++
					break
				}
			}
		}
	}
	return velos
}

func main() {
	fmt.Printf("Problem 1: %v\n", problem1())
	fmt.Printf("Problem 2: %v\n", problem2())
}
