package main

import (
	"fmt"
	"github.com/tzdanows/advent-of-code/2024-advent/Day1"
	"github.com/tzdanows/advent-of-code/2024-advent/Day2"
	"github.com/tzdanows/advent-of-code/utils"
	"log"
	"math/rand"
)

func main() {
	fmt.Println(getFestiveQuote())

	// set the current day to run (EX: "Day1, Day30, Day1337, Day2147 etc.."
	currDay := "Day1"

	// mapping of days to solutions
	days := map[string]struct {
		inputPath  string
		solvePart1 func(string) int
		solvePart2 func(string) int
	}{
		"Day1": {
			inputPath:  "2024-advent/Day1/input.txt",
			solvePart1: Day1.SolvePart1,
			solvePart2: Day1.SolvePart2,
		},
		"Day2": {
			inputPath:  "2024-advent/Day2/input.txt",
			solvePart1: Day2.SolvePart1,
			solvePart2: Day2.SolvePart2,
		},
	}

	// check if the current day exists in the map
	day, exists := days[currDay]
	if !exists {
		log.Fatalf("Day %s is not defined in the days map", currDay)
	}

	// NOTE: some questions may differ in input, we need to address this as it comes
	input := utils.ReadInput(day.inputPath)

	part1Result := day.solvePart1(input)
	fmt.Printf("%s, Part 1: %d\n", currDay, part1Result)

	part2Result := day.solvePart2(input)
	fmt.Printf("%s, Part 2: %d\n", currDay, part2Result)
}

func getFestiveQuote() string {
	quotes := []string{
		"Churning the advent calendar...",
		"‘Tis the season...",
		"I'm hoping I finish more days this time around!",
		"Santa’s workshop isn’t the only place running on magic...",
		"Building festive solutions...",
		"Let it snow, let it sno, let it snowww!",
		"Jingle bells, Jingle bells, Jingle all the way...",
		"It’s beginning to look a lot like runtime... everywhere in dismay...",
		"Dashing through the loops...",
	}
	return quotes[rand.Intn(len(quotes))]
}
