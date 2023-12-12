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

type Node struct {
	value string
	point Point
	level int
}

type Solution struct {
	nodes   []Node
	visited map[string]bool
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

	defer utils.Timer("day9 " + "part" + strconv.Itoa(part))()

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	var grid [][]string
	var start Point

	// Create grid and find starting point
	for x, line := range lines {
		var row []string
		for y, char := range line {
			if char == 'S' {
				start = Point{x, y}
			}
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	// Width and height are the same
	width := len(lines[0])

	node := Node{"S", start, 0}
	solution := getSolution(node, grid, width)

	// Function returns last element found, but actually its neighbour is the answer, so we add +1
	return solution.nodes[len(solution.nodes)-1].level + 1
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	var grid [][]string
	var start Point

	for x, line := range lines {
		var row []string
		for y, char := range line {
			if char == 'S' {
				start = Point{x, y}
			}
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	width := len(lines[0])

	node := Node{"S", start, 0}
	solution := getSolution(node, grid, width)
	visited := solution.visited
	result := 0

	// Check if character is in a shape by taking its starting position and then going diagonally to see how many times
	// it crosses visited nodes, if the number is odd then the character is inside a shape
	for x, line := range grid {
		for y, char := range line {
			key := strconv.Itoa(x) + char + strconv.Itoa(y)
			if visited[key] {
				continue
			}

			crosses := 0
			x1 := x
			y1 := y

			for x1 < width && y1 < width {
				ch := grid[x1][y1]
				k := strconv.Itoa(x1) + ch + strconv.Itoa(y1)

				if visited[k] && ch != "L" && ch != "7" {
					crosses += 1
				}

				x1 += 1
				y1 += 1
			}

			if crosses%2 == 1 {
				result += 1
			}
		}
	}

	return result
}

// This function does two things:
// 1) Creates a node array which keeps track of its value(character), point(x,y) and level(step number)
// 2) Creates visitedPoints map which is used for efficiency in part two
func getSolution(node Node, grid [][]string, width int) Solution {
	visitedPoints := make(map[string]bool)
	var nodes []Node
	queue := []Node{node}

	// Start by adding the starting point(S) to the queue
	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]

		// Get adjacent points for the character and iterate through them
		for _, point := range getAdjacent(node.value) {
			x := node.point.x + point.x
			y := node.point.y + point.y

			if x >= 0 && x < width && y >= 0 && y < width {
				value := grid[x][y]
				if value == "." {
					continue
				}

				// Create new node and if it does not already exist - add it to visitedPoints and next in queue
				adjacentPoint := Point{x, y}
				newNode := Node{value, adjacentPoint, node.level + 1}

				key := strconv.Itoa(x) + value + strconv.Itoa(y)
				_, ok := visitedPoints[key]
				if !ok {
					visitedPoints[key] = true
					queue = append(queue, newNode)
				}
			}
		}
		nodes = append(nodes, node)
	}

	return Solution{nodes, visitedPoints}
}

func getAdjacent(value string) []Point {
	up := Point{-1, 0}
	left := Point{0, -1}
	down := Point{1, 0}
	right := Point{0, 1}

	switch value {
	case "|":
		return []Point{up, down}
	case "-":
		return []Point{left, right}
	case "L":
		return []Point{up, right}
	case "J":
		return []Point{up, left}
	case "7":
		return []Point{down, left}
	case "F":
		return []Point{down, right}
	case "S":
		return []Point{up, left, down, right}
	default:
		return nil
	}
}
