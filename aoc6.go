package main

import (
	"bufio"
	"fmt"
	"os"
)

func isalldifferent(s string) (bool, int) {

	tracking := make(map[byte]int)
	for c := 0; c < len(s); c++ {
		i, f := tracking[s[c]]
		if f {
			return false, i
		}
		tracking[s[c]] = c
	}
	return true, -1
}

func tuningtrouble() {

	file, err := os.Open("Data\\aoc6.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		start := 0
		L := 14
		for start+L < len(line) {
			fmt.Println(line[start : start+L])
			isdiff, nextchar := isalldifferent(line[start : start+L])

			if !isdiff {
				start = start + nextchar + 1
			} else {
				fmt.Println(line[start:start+L], start+L)
				break
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
