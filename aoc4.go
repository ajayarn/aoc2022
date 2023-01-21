package main

import (
	"bufio"
	"fmt"
	"os"
)

func iscontained(a, b, c, d int) bool {
	iscont := c <= b && a <= d

	return iscont
}

func fullycontained() {
	file, err := os.Open("Data\\aoc4.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalcount := 0
	for scanner.Scan() {
		line := scanner.Text()
		m1, M1, m2, M2 := 0, 0, 0, 0
		fmt.Sscanf(line, "%d-%d,%d-%d", &m1, &M1, &m2, &M2)
		if iscontained(m1, M1, m2, M2) {
			totalcount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	print("total", ": ", totalcount)
}
