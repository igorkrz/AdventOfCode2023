package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/igorkrz/AdventOfCode2023/utils"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bet   int
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
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	return solve(1)
}

func part2(input string) int {
	return solve(2)
}

func solve(part int) int {
	lines := strings.Split(input, "\n")
	var groupedHands = make([][]Hand, 8)

	for _, line := range lines {
		split := strings.Split(line, " ")
		bet, _ := strconv.Atoi(split[1])
		hand := Hand{split[0], bet}
		group := groupCards(hand.cards, part)

		groupedHands[group] = append(groupedHands[group], hand)
	}

	for _, hand := range groupedHands {
		sortArray(hand, part)
	}

	result := 0
	rank := 1

	for _, hands := range groupedHands {
		for _, hand := range hands {
			result += rank * hand.bet
			rank += 1
		}
	}

	return result
}

func groupCards(cards string, part int) int {
	charCount := make(map[rune]int)

	if part == 1 {
		for _, char := range cards {
			charCount[char]++
		}

		return calculateScore(charCount)
	}

	maxCharLen := 0
	var maxChar int32
	jokerLen := 0

	for _, char := range cards {
		if string(char) == "J" {
			jokerLen += 1
			continue
		}

		charCount[char] += 1

		if charCount[char] > maxCharLen {
			maxCharLen = charCount[char]
			maxChar = char
		}
	}

	charCount[maxChar] += jokerLen

	return calculateScore(charCount)
}

func calculateScore(charCount map[rune]int) int {
	pairs := 0
	isFourOfKind := false
	for _, occurrence := range charCount {
		if occurrence == 2 {
			pairs += 1
		}
		if occurrence == 4 {
			isFourOfKind = true
		}
	}

	switch len(charCount) {
	case 1:
		return 7
	case 2:
		if isFourOfKind {
			return 6
		}
		return 5
	case 3:
		if pairs == 2 {
			return 3
		}
		return 4
	case 4:
		return 2
	default:
		return 1
	}
}

func sortArray(arr []Hand, part int) {
	sort.Slice(arr, func(i, j int) bool {
		cardsLeft := arr[i].cards
		cardsRight := arr[j].cards

		for i := 0; i < len(cardsLeft); i++ {
			lhs := string(cardsLeft[i])
			rhs := string(cardsRight[i])

			if lhs == rhs {
				continue
			}

			return !shouldReorder(lhs, rhs, part)
		}

		return true
	})
}

func shouldReorder(lhs string, rhs string, part int) bool {
	if part == 2 && rhs == "J" {
		return true
	}
	if lhs == "A" {
		return true
	}
	if lhs == "K" {
		if rhs == "A" {
			return false
		}
		return true
	}
	if lhs == "Q" {
		if strings.Contains("AK", rhs) {
			return false
		}
		return true
	}
	if lhs == "J" {
		if part == 2 {
			return false
		}
		if strings.Contains("AKQ", rhs) {
			return false
		}
		return true
	}
	if lhs == "T" {
		if strings.Contains("AKQJ", rhs) {
			return false
		}
		return true
	}

	return lhs > rhs
}
