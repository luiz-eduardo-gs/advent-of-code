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

const (
	Increasing = 0
	Decreasing = 1
	Same       = -1

	Safe   = "safe"
	Unsafe = "unsafe"
)

func main() {
	// The unusual data (your puzzle input) consists of many reports, one report per line.
	// Each report is a list of numbers called levels that are separated by spaces.
	//input := [][]int{
	//	{7, 6, 4, 2, 1},
	//	{1, 2, 7, 8, 9},
	//	{9, 7, 6, 2, 1},
	//	{1, 3, 2, 4, 5},
	//	{8, 6, 4, 4, 1},
	//	{1, 3, 6, 7, 9},
	//}
	var input [][]int

	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "\n")
		var vSlice []int

		for _, v := range values {
			for _, s := range strings.Split(v, " ") {
				n, _ := strconv.Atoi(s)
				vSlice = append(vSlice, n)
			}
		}
		input = append(input, vSlice)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// The levels are either all increasing or all decreasing.
	// Any two adjacent levels differ by at least 1 and at most 3.
	var result []string

	for i := 0; i <= len(input)-1; i++ {
		g := growth(input[i][0], input[i][1])

		if g == Same {
			result = append(result, Unsafe)
			continue
		}

		if !validDegree(input[i][0], input[i][1]) {
			result = append(result, Unsafe)
			continue
		}

		for j := 1; j <= len(input[i])-1; j++ {
			if j == len(input[i])-1 {
				result = append(result, Safe)
				break
			}

			tempG := growth(input[i][j], input[i][j+1])

			if tempG != g {
				result = append(result, Unsafe)
				break
			}

			if !validDegree(input[i][j], input[i][j+1]) {
				result = append(result, Unsafe)
				break
			}
		}
	}

	total := 0
	for _, v := range result {
		if v == Safe {
			total++
		}
	}

	fmt.Println(total)
}

func validDegree(x, y int) bool {
	abs := math.Abs(float64(x - y))

	return abs >= 1 && abs <= 3
}

func growth(x, y int) int {
	if x-y < 0 {
		return Increasing
	}

	if x-y > 0 {
		return Decreasing
	}

	return Same
}
