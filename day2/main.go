package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Get the sum of IDs that is possible with 12 red cubes, 13 green
// cubes, and 14 blue cubes.

// Get each line, regex out and check if they are lower than amount of color.

var max_red = 12
var max_green = 13
var max_blue = 14

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	scanner.Split(bufio.ScanLines)

	re_game_id := regexp.MustCompile("[0-9]+") // Expect first number to be game id.
	total := 0
	total_cube := 0
	for scanner.Scan() {
		possible := true
		line := scanner.Bytes()
		fmt.Printf("\n-> %s\n", line)
		// Find ID.
		gameIdInt, _ := strconv.Atoi(string(re_game_id.Find(line)))
		// Find all colors and numbers.

		var least_cubes = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for k, v := range cubes {
			re := regexp.MustCompile("([0-9]+) " + k) // match color, make a group of the number

			output := re.FindAllSubmatch(line, 10)
			fmt.Printf("%q\n", output)

			for i := 0; i < len(output); i++ {
				output_int, _ := strconv.Atoi(string(output[i][1]))
				if output_int > v {
					fmt.Printf("IMPOSSIBLE!\n")
					possible = false
				}
				if output_int > least_cubes[k] {
					least_cubes[k] = output_int
				}

			}
		}
		if possible {
			total += gameIdInt
		}
		total_cube += (least_cubes["red"] * least_cubes["green"] * least_cubes["blue"])
	}
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Total cubes: %d\n", total_cube)
}
