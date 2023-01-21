package main

import (
	"bufio"
	"fmt"
	"os"
)

func rockpaperscissors() {
	file, err := os.Open("Data\\aoc2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// winning1 := map[string]int{
	// 	// score: sign + winning
	// 	"A X": 1 + 3,
	// 	"A Y": 2 + 6,
	// 	"A Z": 3 + 0,
	// 	"B X": 1 + 0,
	// 	"B Y": 2 + 3,
	// 	"B Z": 3 + 6,
	// 	"C X": 1 + 6,
	// 	"C Y": 2 + 0,
	// 	"C Z": 3 + 3,
	// }

	winning2 := map[string]int{
		// score: sign + winning
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}

	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		score += winning2[line]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	print("total", ": ", score)
}
