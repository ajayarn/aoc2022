package main

import (
	"bufio"
	"fmt"
	"os"
)

const R int = 99
const C int = 99

// const R int = 5
// const C int = 5

func treetops() {
	file, err := os.Open("Data\\aoc8.txt")

	// file, err := os.Open("Data\\test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var trees [R][C]int

	r := 0
	for scanner.Scan() {
		line := scanner.Text()

		for c := 0; c < C; c++ {
			trees[r][c] = int(line[c] - '0')
		}
		r++
	}

	//---
	// findviewable(trees)
	findbesttreehouse(trees)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func findviewable(trees [R][C]int) {

	var viewable [R][C]bool

	for r := 0; r < R; r++ {
		M := trees[r][0]
		viewable[r][0] = true
		for c := 0; c < C; c++ {
			viewable[r][c] = viewable[r][c] || M < trees[r][c]
			if M < trees[r][c] {
				M = trees[r][c]
			}
		}

		M = trees[r][C-1]
		viewable[r][C-1] = true
		for c := C - 1; c >= 0; c-- {
			viewable[r][c] = viewable[r][c] || M < trees[r][c]
			if M < trees[r][c] {
				M = trees[r][c]
			}
		}
	}

	for c := 0; c < C; c++ {
		M := trees[0][c]
		viewable[0][c] = true
		for r := 0; r < R; r++ {
			viewable[r][c] = viewable[r][c] || M < trees[r][c]
			if M < trees[r][c] {
				M = trees[r][c]
			}
		}

		M = trees[R-1][c]
		viewable[R-1][c] = true
		for r := R - 1; r >= 0; r-- {
			viewable[r][c] = viewable[r][c] || M < trees[r][c]
			if M < trees[r][c] {
				M = trees[r][c]
			}
		}
	}

	total := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if viewable[r][c] {
				total++
			}
		}
	}

	fmt.Println(viewable)

	fmt.Println("total = ", total)
}

func findbesttreehouse(trees [R][C]int) {
	var scores [R][C]int

	bestscore := 0

	for r := 1; r < R-1; r++ {
		for c := 1; c < C-1; c++ {
			dbg := false // r == 3 && c == 2

			up := 1
			s := r - 1
			for s > 0 {

				if trees[s][c] >= trees[r][c] {
					break
				}
				up++
				s--
			}

			if dbg {
				fmt.Println("up=", up)
			}

			down := 1
			s = r + 1
			for s < R-1 {

				if trees[s][c] >= trees[r][c] {
					break
				}
				down++
				s++
			}

			if dbg {
				fmt.Println("down=", down)
			}

			left := 1
			s = c - 1
			for s > 0 {

				if trees[r][s] >= trees[r][c] {
					break
				}
				left++
				s--
			}

			if dbg {
				fmt.Println("left=", left)
			}

			right := 1
			s = c + 1
			for s < C-1 {

				if trees[r][s] >= trees[r][c] {
					break
				}
				s++
				right++
			}
			if dbg {
				fmt.Println("right=", right)
			}
			score := up * down * right * left

			scores[r][c] = score
			if bestscore < score {
				bestscore = score
			}
		}
	}

	fmt.Println("best score", bestscore)
	// fmt.Println(trees)
	// fmt.Println(scores)
}
