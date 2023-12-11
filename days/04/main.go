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

	defer utils.Timer("day4 " + "part" + strconv.Itoa(part))()

	if part == 1 {
		ans := part1()
		fmt.Println("Output:", ans)
	} else {
		ans := part2([]int{1})
		fmt.Println("Output:", ans)
	}
}

var lines = strings.Split(input, "\n")
var sum = 0

func part1() int {
	for _, line := range lines {
		// Get winning and guessed numbers
		cardSeparator := strings.Index(line, ": ")
		numberSeparator := strings.Index(line, "|")
		winningNumbers := regex.FindAllString(line[cardSeparator:numberSeparator], -1)
		guessedNumbers := regex.FindAllString(line[numberSeparator:], -1)

		// Add score if guessed number is in winning numbers
		score := 0
		for _, guess := range guessedNumbers {
			if containsElement(winningNumbers, guess) {
				score += 1
			}
		}

		// Calculate sum
		sum += int(math.Pow(2, float64(score-1)))
	}

	return sum
}

var copies = map[int]int{}

func part2(queue []int) int {
	if queue[0] == len(lines)+1 {
		// Add last original
		copies[queue[0]-1] += 1
		// And sum everything
		for _, number := range copies {
			sum += number
		}

		return sum
	}

	index := queue[0]

	for _, line := range lines[index-1 : index] {
		// Get winning and guessed numbers
		cardSeparator := strings.Index(line, ": ")
		numberSeparator := strings.Index(line, "|")
		winningNumbers := regex.FindAllString(line[cardSeparator:numberSeparator], -1)
		guessedNumbers := regex.FindAllString(line[numberSeparator:], -1)

		// Get number of copies
		score := 0
		for _, guess := range guessedNumbers {
			if containsElement(winningNumbers, guess) {
				score += 1
			}
		}

		// Add original
		if index-1 > 0 {
			copies[index-1] += 1
		}

		// Populate copies of tickets
		for i := 1; i < score+1; i++ {
			copyNumber, ok := copies[index]
			// If the key exists
			if ok {
				copies[index+i] += 1 + copyNumber
			} else {
				copies[index+i] += 1
			}
		}
	}

	// Do the same for the next line
	queue[0] += 1
	part2(queue)

	return sum
}

func containsElement(arr []string, target string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}
