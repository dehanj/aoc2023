package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Idea
// Read all input
// Read a line with scanner.
// append all numbers found in a integer slice
// Take first and last number and sum

var part2 = true
var digits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getDigit(input []byte) []int {
	var digit []int

	for i := 0; i < len(input); i++ {
		if unicode.IsDigit(rune(input[i])) {
			out, _ := strconv.Atoi(string(input[i]))
			digit = append(digit, out)
		}
		if part2 { // part 2

			for k := 0; k < len(digits); k++ {
				if bytes.HasPrefix(input[i:], []byte(digits[k])) {
					digit = append(digit, k+1)
				}
			}
		}
	}
	return digit
}

func main() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanLines)

	i := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		res := getDigit(line)

		digit := (res[0]*10 + res[len(res)-1])
		sum += digit

		// debug
		fmt.Printf("%s\n", line)
		fmt.Printf("%d\n", digit)
		i++
	}
	fmt.Printf("Number of rows: %d\n", i)
	fmt.Printf("Sum: %d\n", sum)
}
