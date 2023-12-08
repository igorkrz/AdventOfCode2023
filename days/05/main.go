package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/igorkrz/AdventOfCode2023/utils"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	start, end int
}

type Triple struct {
	destination, source, rangeLength int
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

	defer utils.Timer("day5 " + "part" + strconv.Itoa(part))()

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	var locations []int
	var seeds []int
	var stepsToLocation [][]Triple

	lines := strings.Split(input, "\n\n")

	// Get seeds from the first line
	seeds = convertToIntArray(regex.FindAllString(lines[0], -1))

	// Set other lines without seeds
	lines = lines[1:]

	// Get all the steps - source, destination and rangeLength and group them by steps
	for _, line := range lines {
		convertToInt := convertToIntArray(regex.FindAllString(line, -1))
		var currentStep []Triple

		for i := 0; i < len(convertToInt); i += 3 {
			triple := Triple{
				destination: convertToInt[i],
				source:      convertToInt[i+1],
				rangeLength: convertToInt[i+2],
			}
			currentStep = append(currentStep, triple)
		}
		stepsToLocation = append(stepsToLocation, currentStep)
	}

	// Iterate through each seed
	for _, seed := range seeds {
		currentSeed := seed
		// Iterate through the process from seed to location
		for _, steps := range stepsToLocation {
			for _, step := range steps {
				sourceEnd := step.source + step.rangeLength
				// If seed is gte than source and lt addition of source and rangeLength set current seed to diff of
				// destination and source, it seems that only one mapping can be used. For example:
				// [seed 14->soil 14->fertilizer] 53 fertilizer to water 49 53 8 is a match and immediately step after
				// 0 11 42 is also a match, so if we do not use break our seed would be 14-14-53-49-32, but if we use
				// break it stays 49
				if currentSeed >= step.source && currentSeed < sourceEnd {
					currentSeed += step.destination - step.source
					break
				}
			}
		}
		locations = append(locations, currentSeed)
	}

	return slices.Min(locations)
}

func part2(input string) int {
	var seeds []Pair
	var stepsToLocation [][]Triple

	lines := strings.Split(input, "\n\n")

	// Get seeds from the first line
	initialSeeds := convertToIntArray(regex.FindAllString(lines[0], -1))

	// Set other lines without seeds
	lines = lines[1:]

	for i := 0; i < len(initialSeeds)-1; i++ {
		if i%2 == 0 {
			seed := initialSeeds[i]
			seedRange := initialSeeds[i+1]
			seeds = append(seeds, Pair{seed, seed + seedRange})
		}
	}

	// Get all the steps - source, destination and rangeLength and group them by steps
	for _, line := range lines {
		convertToInt := convertToIntArray(regex.FindAllString(line, -1))
		var currentStep []Triple

		for i := 0; i < len(convertToInt); i += 3 {
			triple := Triple{
				destination: convertToInt[i],
				source:      convertToInt[i+1],
				rangeLength: convertToInt[i+2],
			}
			currentStep = append(currentStep, triple)
		}
		stepsToLocation = append(stepsToLocation, currentStep)
	}

	// Some interval magic
	var viablePairs []Pair
	var viableRanges []Pair
	for _, seed := range seeds {
		viablePairs = []Pair{seed}
		for _, steps := range stepsToLocation {
			viablePairs = findViablePairs(viablePairs, steps)
		}
		viableRanges = append(viableRanges, viablePairs...)
	}

	// Find the lowest pair
	lowestPair := Pair{math.MaxInt32, math.MaxInt32}

	for _, pair := range viableRanges {
		if pair.start < lowestPair.start {
			lowestPair = pair
		}
	}

	return min(lowestPair.start, lowestPair.end)
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

func findViablePairs(viablePairs []Pair, steps []Triple) []Pair {
	var interPairs []Pair

	for _, step := range steps {
		var beforeAndAfterPairs []Pair
		sourceEnd := step.source + step.rangeLength

		for len(viablePairs) > 0 {
			currentSeed := viablePairs[0]
			viablePairs = viablePairs[1:]
			start := currentSeed.start
			end := currentSeed.end
			before := Pair{start, min(end, step.source)}
			inter := Pair{max(start, step.source), min(sourceEnd, end)}
			after := Pair{max(sourceEnd, start), end}

			if before.end > before.start {
				beforeAndAfterPairs = append(beforeAndAfterPairs, before)
			}

			if inter.end > inter.start {
				interPair := Pair{inter.start - step.source + step.destination, inter.end - step.source + step.destination}
				interPairs = append(interPairs, interPair)
			}

			if after.end > after.start {
				beforeAndAfterPairs = append(beforeAndAfterPairs, after)
			}
		}
		viablePairs = beforeAndAfterPairs
	}

	return append(viablePairs, interPairs...)
}
