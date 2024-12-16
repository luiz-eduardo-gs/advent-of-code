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
	//input := [][]int{
	//	{7, 6, 4, 2, 1},
	//	{1, 2, 7, 8, 9},
	//	{9, 7, 6, 2, 1},
	//	{1, 3, 2, 4, 5},
	//	{8, 6, 4, 4, 1},
	//	{1, 3, 6, 7, 9},
	//}
	input := readFromFile()

	var result []string

	for _, report := range input {
		valid := validate(report)

		if !valid {
			nowValid := removeAndValidate(report)

			if nowValid {
				result = append(result, Safe)
			} else {
				result = append(result, Unsafe)
			}
		} else {
			result = append(result, Safe)
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

func readFromFile() [][]int {
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

	return input
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func removeAndValidate(report []int) bool {
	tempReport := make([]int, len(report))

	copy(tempReport, report)

	for i := 0; i < len(report); i++ {
		tempReport = remove(tempReport, i)
		valid := validate(tempReport)
		if !valid {
			tempReport = make([]int, len(report))
			copy(tempReport, report)
			continue
		}

		return true
	}

	return false
}

func validate(report []int) bool {
	g := growth(report[0], report[1])

	if g == Same {
		return false
	}

	if !validDegree(report[0], report[1]) {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		if i == len(report)-1 {
			return true
		}

		tempG := growth(report[i], report[i+1])

		if tempG != g {
			return false
		}

		if !validDegree(report[i], report[i+1]) {
			return false
		}
	}

	return true
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
