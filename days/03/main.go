package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/igorkrz/AdventOfCode2023/utils"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var regex = regexp.MustCompile(`\d+`)

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	defer utils.Timer("day3 " + "part" + strconv.Itoa(part))()

	if part == 1 {
		fmt.Println("Output:", part1())
	} else {
		fmt.Println("Output:", part2())
	}
}

func part1() int {
	return solve(1)
}

func part2() int {
	return solve(2)
}

func solve(part int) int {
	adjacent := getAdjacentCoords()

	// Create grid and star character coordinates
	var grid [][]string
	var specialCharCoords [][]int
	for x, line := range strings.Split(input, "\n") {
		var row []string
		for y, char := range line {
			character := string(char)
			row = append(row, character)
			if part == 1 && isSpecialCharacter(character) {
				specialCharCoords = append(specialCharCoords, []int{x, y})
			} else if part == 2 && character == "*" {
				specialCharCoords = append(specialCharCoords, []int{x, y})
			}
		}
		grid = append(grid, row)
	}

	// Loop through special character coordinates and its adjacent elements
	sum := 0
	for _, coords := range specialCharCoords {
		var visited []int
		for _, adj := range adjacent {
			x := coords[0] + adj[0]
			y := coords[1] + adj[1]
			startingY := y
			result := ""

			// Do not get any characters from either left or right since we encountered a dot
			if grid[x][y] == "." {
				continue
			}

			// Get characters on the left side
			for isInBounds(x, y, grid) && isDigit(grid[x][y]) {
				result = grid[x][y] + result
				y -= 1
			}

			// Adjust y
			y = startingY + 1

			// Get characters on the right side
			for isInBounds(x, y, grid) && isDigit(grid[x][y]) {
				result += grid[x][y]
				y += 1
			}

			// Sum result if it is not an empty string
			if result != "" {
				num, _ := strconv.Atoi(result)
				// Do not add duplicates from the same row
				if !containsElement(visited, num) {
					visited = append(visited, num)
					if part == 1 {
						sum += num
					} else if part == 2 && len(visited) == 2 {
						sum += visited[0] * visited[1]
					}
				}
			}
		}
	}

	return sum
}

func getAdjacentCoords() [][]int {
	return [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
}

func isSpecialCharacter(char string) bool {
	return char != "." && len(regex.FindAllString(char, -1)) == 0
}

func isDigit(char string) bool {
	return len(regex.FindAllString(char, -1)) > 0
}

func isInBounds(x int, y int, grid [][]string) bool {
	return (x >= 0 && x < len(grid)) && (y >= 0 && y < len(grid))
}

func containsElement(arr []int, target int) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}
