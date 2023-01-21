package main

import (
	"bufio"
	"fmt"
	"os"
)

func nospaceleft() {

	file, err := os.Open("Data\\aoc7.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sizes := make(map[string]int)
	parents := make(map[string]string)

	currentdir := ""

	linenr := 1

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(linenr, line)
		linenr++

		if line[0] == '$' {
			var (
				cmd, param string
			)

			fmt.Sscanf(line, "$ %s %s", &cmd, &param)

			// do something with this:
			// add a key in a map and let my people go!
			// fmt.Println(cmd, param)
			if cmd == "cd" {

				if param == ".." {
					currentdir = parents[currentdir]
				} else {
					currentdir = currentdir + "-" + param
				}
			}

			continue
		}

		if line[:3] == "dir" {

			// dir - add to hash

			var dirname string
			fmt.Sscanf(line, "dir %s", &dirname)
			parents[currentdir+"-"+dirname] = currentdir
			continue
		}

		var (
			size     int
			filename string
		)

		fmt.Sscanf(line, "%d %s", &size, &filename)

		sizes[currentdir] = sizes[currentdir] + size

		p := currentdir

		for parents[p] != "" {
			p = parents[p]
			sizes[p] += size
		}

		// fmt.Println(sizes)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// part 1:
	//---------
	// totsize := 0
	// for dirname := range sizes {
	// 	if sizes[dirname] <= 100000 {
	// 		totsize += sizes[dirname]
	// 	}
	// }
	// fmt.Println("totsize=", totsize)

	// part 2:
	//---------
	remainingSpace := 70000000 - sizes["-/"]
	if remainingSpace >= 30000000 {
		fmt.Println("No need to delete anything")
		os.Exit(0)
	}
	neededspace := 30000000 - remainingSpace
	mindir := "-/"
	minsize := sizes[mindir]
	for dirname := range sizes {
		s := sizes[dirname]
		if s >= neededspace && s < minsize {
			minsize = s
			mindir = dirname
		}
	}
	fmt.Println(mindir, minsize)
}
