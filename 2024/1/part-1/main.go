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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lList []int
	var rList []int

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")
		l, _ := strconv.Atoi(values[0])
		r, _ := strconv.Atoi(values[1])
		lList = append(lList, l)
		rList = append(rList, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	length := len(lList)

	sortedLeftList := quickSort(lList, 0, length-1)
	sortedRightList := quickSort(rList, 0, length-1)
	var result []int

	for i := 0; i <= length-1; i++ {
		result = append(result, int(math.Abs(float64(sortedLeftList[i]-sortedRightList[i]))))
	}

	total := 0
	for i := 0; i < len(result); i++ {
		total += result[i]
	}

	fmt.Println(total)
}

func swap(values []int, i int, j int) {
	values[i], values[j] = values[j], values[i]
}

func partition(values []int, l, r int) int {
	pivot := values[l]
	j := l

	for i := l + 1; i <= r; i++ {
		if values[i] <= pivot {
			j++
			swap(values, j, i)
		}
	}

	swap(values, l, j)

	return j
}

func quickSort(values []int, l, r int) []int {
	if l < r {
		pivotIdx := partition(values, l, r)
		quickSort(values, l, pivotIdx-1)
		quickSort(values, pivotIdx+1, r)
	}

	return values
}
