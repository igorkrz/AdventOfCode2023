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

var regex = regexp.MustCompile(`-?\d+`)

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
	lines := strings.Split(input, "\n")

	result := 0

	for _, line := range lines {
		var edges []int
		convertToInt := convertToIntArray(regex.FindAllString(line, -1))

		if part == 1 {
			edges = []int{convertToInt[len(convertToInt)-1]}
		} else {
			edges = []int{convertToInt[0]}
		}

		for !isAllZero(convertToInt) {
			convertToInt = reduce(convertToInt)

			if part == 1 {
				edges = append(edges, convertToInt[len(convertToInt)-1])
			} else {
				edges = append([]int{convertToInt[0]}, edges...)
			}
		}
		result += sumArray(edges, part)
	}

	return result
}

func convertToIntArray(arrayToConvert []string) []int {
	var result []int

	for _, convert := range arrayToConvert {
		j, err := strconv.Atoi(convert)
		if err != nil {
			panic(err)
		}
		result = append(result, j)
	}

	return result
}

func reduce(arr []int) []int {
	var result []int

	for i := 0; i < len(arr)-1; i++ {
		result = append(result, arr[i+1]-arr[i])
	}

	return result
}

func isAllZero(arr []int) bool {
	for _, element := range arr {
		if element != 0 {
			return false
		}
	}

	return true
}

func sumArray(arr []int, part int) int {
	result := 0

	if part == 1 {
		for _, element := range arr {
			result += element
		}

		return result
	}

	for i := 1; i < len(arr); i++ {
		result = arr[i] - result
	}

	return result
}
