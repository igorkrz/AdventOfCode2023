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

	defer utils.Timer("day1 " + "part" + strconv.Itoa(part))()

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

	regex := regexp.MustCompile(`\d`)

	sum := 0

	for _, s := range lines {
		numerics := regex.FindAllString(s, -1)
		result, err := strconv.Atoi(numerics[0] + numerics[len(numerics)-1])
		if err != nil {
			panic(err)
		}

		sum += result
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	pattern := `(\d|zero|one|two|three|four|five|six|seven|eight|nine)`
	patternFromEnd := `(\d|orez|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`
	regex := regexp.MustCompile(pattern)
	regexFromEnd := regexp.MustCompile(patternFromEnd)

	sum := 0

	for _, s := range lines {
		numerics := regex.FindAllString(s, 1)

		numericsFromEnd := regexFromEnd.FindAllString(Reverse(s), 1)
		numerics = append(numerics, numericsFromEnd...)

		result, err := strconv.Atoi(getStringDigit(numerics[0]) + getStringDigit(numerics[len(numerics)-1]))
		if err != nil {
			panic(err)
		}

		sum += result
	}

	return sum
}

func getStringDigit(numeric string) string {
	switch numeric {
	case "0", "zero", "orez":
		return "0"
	case "1", "one", "eno":
		return "1"
	case "2", "two", "owt":
		return "2"
	case "3", "three", "eerht":
		return "3"
	case "4", "four", "ruof":
		return "4"
	case "5", "five", "evif":
		return "5"
	case "6", "six", "xis":
		return "6"
	case "7", "seven", "neves":
		return "7"
	case "8", "eight", "thgie":
		return "8"
	case "9", "nine", "enin":
		return "9"
	default:
		panic(numeric)
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
