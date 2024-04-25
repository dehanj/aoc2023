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

// Idea
// Split in winning numbers and cards
// Use regex with groups, check if each winning card exist in my
// cards, sum.

// Part 2
// add a map with values of orignals + copies.
// Sum the map in the end.

var copies = map[int]int{}

func main() {

	lines, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(lines)))
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile("([0-9]+)")
	var gameId int
	total := 0

	for scanner.Scan() {
		round := 0
		line := scanner.Bytes()
		fmt.Printf("\nline: %s\n", line)
		split := bytes.Split(line, []byte("|"))

		winning := re.FindAll(split[0], -1)
		cards := re.FindAll(split[1], -1)
		gameId, _ = strconv.Atoi(string(winning[0]))
		winning = winning[1:]
		fmt.Printf("%s\n", winning)
		fmt.Printf("%s\n", cards)

		for _, k := range winning {
			for _, y := range cards {
				if bytes.Equal(k, y) {
					round++

					// Part 1
					// 	if round == 0 {
					// 		round = 1
					// 	} else {
					// 		round *= 2
					// 	}
				}
			}
		}

		fmt.Printf("round: %d\n", round)
		// Part 1
		// total += round
		// fmt.Printf("total: %v\n", total)

		copies[gameId] = copies[gameId] + 1 // Add the original one

		// Add copies to list
		for i := 1; i <= round; i++ {
			copies[gameId+i] = copies[gameId+i] + copies[gameId]
		}

		fmt.Printf("copies: %v\n", copies)

	}
	for l := range copies {
		total += copies[l]
	}
	fmt.Printf("total: %v\n", total)
}
