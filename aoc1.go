package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calories() {
	file, err := os.Open("Data\\aoc1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf := 0
	elfcalories := 0

	topthree := map[int]int{1: 0, 2: 0, 3: 0}
	for scanner.Scan() {
		line := scanner.Text()

		c, err := strconv.Atoi(line)
		if err != nil {
			for elfcalories > topthree[3] {
				fmt.Println(topthree)
				if elfcalories > topthree[1] {
					topthree[2] = topthree[1]
					topthree[3] = topthree[2]
					topthree[1] = elfcalories
					break
				}

				if elfcalories > topthree[2] {
					topthree[3] = topthree[2]
					topthree[2] = elfcalories
					break
				}

				topthree[3] = elfcalories
				break
			}

			elfcalories = 0
			elf++
			continue

		}

		elfcalories += c
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	print("total", ": ", topthree[3]+topthree[1]+topthree[2])
}
