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
// Create some struct with the input data.
// Start with first seed, just check if it is in the range
// If yes, calculate next value to store as input data
// Repeat for each category
// Find lowest value in input data

func main() {
	lines, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(lines)))
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Bytes()
	fmt.Printf("\nline: %s\n", line)

	re := regexp.MustCompile("([0-9]+)")
	seeds := re.FindAll(line, -1)
	fmt.Printf("seeds: %s: %d\n", seeds, len(seeds))

	storage := make([]int, len(seeds))
	storage_set := make([]bool, len(seeds))

	for i := range storage_set {
		storage_set[i] = false
		storage[i], _ = strconv.Atoi(string(seeds[i]))
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Printf("\nline: %s\n", line)

		if bytes.Contains(line, []byte("map")) {
			fmt.Printf("New category, reset!")
			for i := range storage_set {
				storage_set[i] = false
			}
			continue
		}
		if len(line) == 0 {
			fmt.Printf("End of category: ")
			fmt.Printf("storage: %v\n", storage)
			continue
		}
		maps := re.FindAll(line, -1)

		for index, seed_int := range storage {
			if storage_set[index] == false {
				fmt.Printf("Input: %d\n", seed_int)
				dest_start_int, _ := strconv.Atoi(string(maps[0]))
				source_start_int, _ := strconv.Atoi(string(maps[1]))
				range_int, _ := strconv.Atoi(string(maps[2]))

				if seed_int >= source_start_int && seed_int < (source_start_int+range_int) {
					// seed is included
					storage[index] = (dest_start_int - source_start_int) + seed_int
					storage_set[index] = true
					fmt.Printf("Included seed %d [%d]\n", seed_int, storage[index])
				}
			}
		}
		fmt.Printf("storage: %v\n", storage)
	}
	// find lowest
	lowest := storage[0]
	for _, k := range storage {
		if k < lowest {
			lowest = k
		}
	}
	fmt.Printf("Lowest: %d\n", lowest)
}
