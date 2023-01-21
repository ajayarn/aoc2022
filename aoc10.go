package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func cathoderay() {

	file, err := os.Open("Data\\aoc10.txt")

	// file, err := os.Open("Data\\test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cycle := 0
	X := 1

	// targetcycle := 20
	targetcycle := 40 // for the new line
	// totalvalue := 0

	var cmd string
	var val int

	var dark_linebuffer = [40]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}

	linebuffer := dark_linebuffer

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Sscanf(line, "%s %d", &cmd, &val)

		c := cycle % 40

		if math.Abs(float64(X-c)) < 2 {
			linebuffer[c] = '#'
		}

		cycle++

		if cycle == targetcycle {
			// totalvalue += X * targetcycle
			fmt.Println(string(linebuffer[:]))
			linebuffer = dark_linebuffer
			targetcycle += 40
		}

		if cmd == "noop" {
			continue
		}

		c = cycle % 40

		if math.Abs(float64(X-c)) < 2 {
			linebuffer[c] = '#'
		}

		cycle++

		if cycle == targetcycle {
			// totalvalue += X * targetcycle

			fmt.Println(string(linebuffer[:]))
			linebuffer = dark_linebuffer
			targetcycle += 40
		}
		X += val
	}

	// fmt.Println("totalvalue = ", totalvalue)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
