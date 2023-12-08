package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/igorkrz/AdventOfCode2023/utils"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	time, distance int
}

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
	flag.IntVar(&part, "part", 1, "part 01 or 02")
	flag.Parse()
	fmt.Println("Running part", part)

	defer utils.Timer("day6 " + "part" + strconv.Itoa(part))()

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
	times := convertToIntArray(regex.FindAllString(lines[0], -1))
	distances := convertToIntArray(regex.FindAllString(lines[1], -1))

	result := 1

	for i := 0; i < len(times); i++ {
		pair := Pair{times[i], distances[i]}
		result *= calculateResult(pair)
	}

	return result
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	time, _ := strconv.Atoi(strings.Join(regex.FindAllString(lines[0], -1), ""))
	distance, _ := strconv.Atoi(strings.Join(regex.FindAllString(lines[1], -1), ""))

	pair := Pair{time, distance}

	return calculateResult(pair)
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

// Quadratic equation
// x=(time±(sqrRoot(timeSquared−4*distance)) / 2
func calculateResult(pair Pair) int {
	timeSquared := math.Pow(float64(pair.time), 2)
	sqrRoot := math.Sqrt(timeSquared - float64(4*pair.distance))
	lowerBound := int((float64(pair.time) - sqrRoot) / 2)
	upperBound := int((float64(pair.time) + sqrRoot) / 2)

	result := upperBound - lowerBound

	if result == lowerBound {
		return result - 1
	}

	return result
}
