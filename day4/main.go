package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Idea
// Split in winning numbers and cards
// Use regex with groups, check if each winning card exist in my
// cards

func main() {

	lines, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(lines)))
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile("([0-9]+)")

	total := 0
	for scanner.Scan() {
		round := 0
		line := scanner.Bytes()
		fmt.Printf("line: %s\n", line)
		split := bytes.Split(line, []byte("|"))

		winning := re.FindAll(split[0], -1)
		cards := re.FindAll(split[1], -1)
		winning = winning[1:]
		fmt.Printf("%s\n", winning)
		fmt.Printf("%s\n", cards)

		for _, k := range winning {
			for _, y := range cards {
				if bytes.Equal(k, y) {
					if round == 0 {
						round = 1
					} else {
						round *= 2
					}
				}
			}
		}
		total += round
		fmt.Printf("total: %v\n", total)
	}
}
