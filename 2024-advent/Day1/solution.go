package Day1

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1

// subproblems:
// pair smallest numbers in each list together
// find the difference between the two
// add all differences

// sampledata: 11
// input: 1970720

func SolvePart1(input string) int {
	var col1, col2 []int
	line := strings.Split(input, "\n")

	for _, l := range line {
		locations := strings.Fields(l)

		num1, err1 := strconv.Atoi(locations[0])
		num2, err2 := strconv.Atoi(locations[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Failed to convert strings to integers")
		}

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}
	sort.Ints(col1)
	sort.Ints(col2)

	finalDist := 0
	for i := 0; i < len(col1); i++ {
		diff := col1[i] - col2[i]
		if diff < 0 {
			diff = -diff
		}
		finalDist += diff
	}
	return finalDist
}

// subproblems:
// modifying Part1 since it's reusable up after turning the columns into arrays
// fully loop the right list and store entries in a dictionary/hashmap
// loop the left list while checking if they exist in dictionary, multiply by the count
// add all values

// sampledata: 31
// input: 17191599

func SolvePart2(input string) int {
	var col1, col2 []int
	line := strings.Split(input, "\n")
	numCount := map[int]int{}

	for _, l := range line {
		locations := strings.Fields(l)

		num1, err1 := strconv.Atoi(locations[0])
		num2, err2 := strconv.Atoi(locations[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Failed to convert strings to integers")
		}

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	for _, num := range col2 {
		numCount[num]++
	}

	similarityScore := 0
	for i := 0; i < len(col1); i++ {
		if count, exists := numCount[col1[i]]; exists {
			similarityScore += col1[i] * count
		}
	}
	return similarityScore
}
