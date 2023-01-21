package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	r int
	c int
}

func ropebridge() {

	file, err := os.Open("Data\\aoc9.txt")

	// file, err := os.Open("Data\\test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// H := Point{0, 0}
	// T := Point{0, 0}
	const numknots int = 10

	var knots [numknots]Point

	Tvisited := make(map[Point]bool)

	Tvisited[knots[numknots-1]] = true

	for scanner.Scan() {
		line := scanner.Text()

		var direction string
		m := 0
		fmt.Sscanf(line, "%s %d", &direction, &m)

		fmt.Println(line)
		for i := 0; i < m; i++ {

			switch direction {
			case "L":
				knots[0].c--
			case "R":
				knots[0].c++
			case "U":
				knots[0].r--
			case "D":
				knots[0].r++
			}

			for k := 1; k < numknots; k++ {

				cdiff := knots[k].c - knots[k-1].c
				rdiff := knots[k].r - knots[k-1].r

				if math.Abs(float64(cdiff)) <= 1 && math.Abs(float64(rdiff)) <= 1 {
					break
				}

				// have to move diagnonally - inc or dec both c and r
				if cdiff > 0 {
					knots[k].c--
				} else if cdiff < 0 {
					knots[k].c++
				}

				if rdiff > 0 {
					knots[k].r--
				} else if rdiff < 0 {
					knots[k].r++
				}

				// fmt.Println(T, Tvisited)
			}

			Tvisited[knots[numknots-1]] = true
			// fmt.Println("H=", H, "T=", T)
		}

		for i := 0; i < numknots; i++ {
			fmt.Println(i, knots[i])
		}

	}

	fmt.Println("Nr coords: ", len(Tvisited))
	// fmt.Println(Tvisited)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
