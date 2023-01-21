package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func movestacks() {

	var stacks = map[int]string{
		1: "FCJPHTW",
		2: "GRVFZJBH",
		3: "HPTR",
		4: "ZSNPHT",
		5: "NVFZHJCD",
		6: "PMGFWDZ",
		7: "MVZWSJDP",
		8: "NDS",
		9: "DZSFM",
	}

	file, err := os.Open("Data\\aoc5.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		n, stackfrom, stackto int
	)
	fmt.Println(stacks)
	for scanner.Scan() {
		line := scanner.Text()

		// fmt.Println("\n" + line)
		fmt.Sscanf(line, "move %d from %d to %d", &n, &stackfrom, &stackto)
		// stacks[stackto] = stacks[stackto] + reverse(stacks[stackfrom][len(stacks[stackfrom])-n:])
		stacks[stackto] = stacks[stackto] + stacks[stackfrom][len(stacks[stackfrom])-n:]
		stacks[stackfrom] = stacks[stackfrom][:len(stacks[stackfrom])-n]
		// fmt.Println(stacks)
	}
	var result string
	for k := 1; k <= len(stacks); k++ {
		result = result + string(stacks[k][len(stacks[k])-1])
	}

	print("Finally: ", result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
