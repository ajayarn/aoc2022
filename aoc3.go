package main

import (
	"bufio"
	"fmt"
	"os"
)

func score(letter byte) int {
	if letter >= 'a' && letter <= 'z' {
		return int(letter + 1 - 'a')
	}

	return int(letter + 27 - 'A')
}

func rucksack_reorg() {
	file, err := os.Open("Data\\aoc3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	comparment1 := make(map[byte]bool)

	totalscore := 0
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)

		for k := range comparment1 {
			delete(comparment1, k)
		}

		for i := 0; i < length/2; i++ {
			comparment1[line[i]] = true
		}

		for j := length / 2; j < length; j++ {
			if _, found := comparment1[line[j]]; found {
				totalscore += score(line[j])
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	print("total", ": ", totalscore)
}

func rucksack_reorg2() {
	file, err := os.Open("Data\\aoc3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	founditem := make(map[byte]int)

	totalscore := 0
	groupcount := 0
	for scanner.Scan() {
		line := scanner.Text()
		groupcount++

		if groupcount == 4 {
			groupcount = 1
			for k := range founditem {
				delete(founditem, k)
			}
		}

		length := len(line)

		for j := 0; j < length; j++ {

			if groupcount == 1 {
				founditem[line[j]] = groupcount
				continue
			}

			if _, found := founditem[line[j]]; found {
				if groupcount == 2 {
					founditem[line[j]] = groupcount
					continue
				}

				if groupcount == 3 {
					if founditem[line[j]] == 2 {
						fmt.Println(string(line[j]), "=", score(line[j]))
						totalscore += score(line[j])
						break
					}
				}

				// founditem[line[j]]
				// i need to distinguish between two items in the same person or different person

			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	print("total", ": ", totalscore)
}
