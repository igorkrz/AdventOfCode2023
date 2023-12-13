package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/igorkrz/AdventOfCode2023/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

type Grid struct {
	populated    []Point
	emptyRows    map[int]bool
	emptyColumns map[int]bool
}

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

	defer utils.Timer("day11 " + "part" + strconv.Itoa(part))()

	if part == 1 {
		ans := part1()
		fmt.Println("Output:", ans)
	} else {
		ans := part2()
		fmt.Println("Output:", ans)
	}
}

func part1() int {
	return solve(1)
}

func part2() int {
	return solve(2)
}

func solve(part int) int {
	grid := createGrid()

	if part == 1 {
		expandUniverse(grid, 1)
	} else {
		expandUniverse(grid, 1000000-1)
	}

	result := 0

	for _, galaxy := range grid.populated {
		for _, point := range grid.populated {
			result += manhattanDistance(galaxy, point)
		}
	}

	return result / 2
}

func createGrid() Grid {
	var grid [][]string
	var populated []Point
	emptyColumns := make(map[int]bool)
	emptyRows := make(map[int]bool)

	lines := strings.Split(input, "\n")

	for x, line := range lines {
		var row []string

		for y, char := range line {
			if char == '#' {
				emptyRows[x] = true
				emptyColumns[y] = true
				populated = append(populated, Point{x, y})
			}

			row = append(row, string(char))

			if !emptyColumns[y] {
				emptyColumns[y] = false
			}
		}

		if !emptyRows[x] {
			emptyRows[x] = false
		}

		grid = append(grid, row)
	}

	for i, b := range emptyColumns {
		if b {
			delete(emptyColumns, i)
		}
	}

	for i, b := range emptyRows {
		if b {
			delete(emptyRows, i)
		}
	}

	return Grid{populated, emptyRows, emptyColumns}
}

func expandUniverse(grid Grid, howMuch int) {
	for in, galaxy := range grid.populated {
		for f, _ := range grid.emptyRows {
			if galaxy.x > f {
				grid.populated[in].x += howMuch
			}
		}

		for i, _ := range grid.emptyColumns {
			if galaxy.y > i {
				grid.populated[in].y += howMuch
			}
		}
	}
}

func manhattanDistance(lhs Point, rhs Point) int {
	return abs(lhs.x, rhs.x) + abs(lhs.y, rhs.y)
}

func abs(lhs int, rhs int) int {
	result := lhs - rhs

	if result < 0 {
		return result * -1
	}

	return result
}
