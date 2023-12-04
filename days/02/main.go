package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

	regexColorAndNumber := regexp.MustCompile(`\d{02,4} \w+[,;]`)

	sum := 0

	colors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, line := range lines {
		colorsAndNumbers := regexColorAndNumber.FindAllString(line, -1)

		if isValidColorAndNumber(colorsAndNumbers, colors) {
			sum += getGameId(line)
		}
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	regexColorAndNumber := regexp.MustCompile(`\d+ \w+`)

	sum := 0

	for _, line := range lines {
		colorsAndNumbers := regexColorAndNumber.FindAllString(line, -1)

		colors := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, colorAndNumber := range colorsAndNumbers {
			for color, currentNumber := range colors {
				colors[color] = getHighestValueForColor(colorAndNumber, color, currentNumber)
			}
		}

		multipliedValue := 1
		for _, currentNumber := range colors {
			multipliedValue *= currentNumber
		}

		sum += multipliedValue
	}

	return sum
}

// Part 02
func getHighestValueForColor(s string, color string, currentNumber int) int {
	regex := regexp.MustCompile(`\d+`)
	number, err := strconv.Atoi(regex.FindAllString(s, 1)[0])
	if err != nil {
		panic(err)
	}

	if strings.Contains(s, color) && number > currentNumber {
		return number
	}

	return currentNumber
}

// Part 01
func isValidColorAndNumber(colorsAndNumbers []string, colors map[string]int) bool {
	isAboveLimit := false

	for _, s := range colorsAndNumbers {
		for color, limit := range colors {
			isAboveLimit = isColorAboveLimit(s, color, limit)
			if isAboveLimit {
				return false
			}
		}
	}

	return true
}

func getGameId(line string) int {
	regexGameId := regexp.MustCompile(`\d+`)
	gameId, err := strconv.Atoi(regexGameId.FindAllString(line, 1)[0])
	if err != nil {
		panic(err)
	}

	return gameId
}

func isColorAboveLimit(s string, color string, limit int) bool {
	regex := regexp.MustCompile(`\d+`)

	number, err := strconv.Atoi(regex.FindAllString(s, 1)[0])
	if err != nil {
		panic(err)
	}

	if strings.Contains(s, color) {
		return number > limit
	}

	return false
}
