package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//in := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	in := readInput("../input.txt")

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches := re.FindAllString(in, -1)

	result := strings.Join(matches, "\n")

	total := 0

	for _, s := range strings.Split(result, "\n") {
		replaced := strings.Replace(s, "mul(", "", 1)
		replaced = strings.Replace(replaced, ")", "", 1)

		x := strings.Split(replaced, ",")
		a, _ := strconv.Atoi(x[0])
		b, _ := strconv.Atoi(x[1])

		total += a * b
	}

	fmt.Println(total)
}

func readInput(file string) string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var buffer bytes.Buffer

	for s.Scan() {
		buffer.WriteString(s.Text())
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return buffer.String()
}
