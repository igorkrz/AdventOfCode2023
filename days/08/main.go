package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/igorkrz/AdventOfCode2023/utils"
	"strconv"
	"strings"
)

type Instruction struct {
	value string
	left  string
	right string
}

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 01 or 02")
	flag.Parse()
	fmt.Println("Running part", part)

	defer utils.Timer("day7 " + "part" + strconv.Itoa(part))()

	if part == 1 {
		ans := p1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func p1(input string) int {
	lines := strings.Split(input, "\n\n")
	instructions := strings.Trim(lines[0], " ")
	network := strings.Split(lines[1], "\n")

	steps := 0
	lastStep := "AAA"
	networkSet := make(map[string]Instruction)

	for _, nodesString := range network {
		instructionsSplit := strings.Split(nodesString, "=")
		currentNode := strings.Trim(instructionsSplit[0], " ")
		leftAndRight := strings.Split(instructionsSplit[1], ",")
		left := strings.Trim(leftAndRight[0], " (")
		right := strings.Trim(leftAndRight[1], " )")

		networkSet[currentNode] = Instruction{currentNode, left, right}
	}

	l := 'L'
	r := 'R'

	for lastStep != "ZZZ" {
		for _, instruction := range instructions {
			if instruction == l {
				lastStep = networkSet[lastStep].left
			} else if instruction == r {
				lastStep = networkSet[lastStep].right
			} else {
				panic(l)
			}
			steps += 1
		}
	}

	return steps
}

func part2(input string) int {
	lines := strings.Split(input, "\n\n")
	instructions := strings.Trim(lines[0], " ")
	network := strings.Split(lines[1], "\n")

	steps := 0
	lastStep := ""
	networkSet := make(map[string]Instruction)

	var networkEndsWithA []Instruction

	for _, nodesString := range network {
		instructionsSplit := strings.Split(nodesString, "=")
		currentNode := strings.Trim(instructionsSplit[0], " ")
		leftAndRight := strings.Split(instructionsSplit[1], ",")
		left := strings.Trim(leftAndRight[0], " (")
		right := strings.Trim(leftAndRight[1], " )")

		networkSet[currentNode] = Instruction{currentNode, left, right}

		if strings.HasSuffix(currentNode, "A") {
			networkEndsWithA = append(networkEndsWithA, Instruction{currentNode, left, right})
		}
	}

	l := 'L'
	r := 'R'

	var results []int
	for _, n := range networkEndsWithA {
		lastStep = n.value
		for !strings.HasSuffix(lastStep, "Z") {
			for _, instruction := range instructions {
				if instruction == l {
					lastStep = networkSet[lastStep].left
				} else if instruction == r {
					lastStep = networkSet[lastStep].right
				} else {
					panic(l)
				}
				steps += 1
			}
		}
		results = append(results, steps)
		steps = 0
	}

	return lcmOfArray(results)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcmOfArray(numbers []int) int {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
