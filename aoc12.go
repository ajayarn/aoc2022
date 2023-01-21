package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	r int
	c int
}

var distanced = make(map[coord]int)

func printpath(visited map[coord]coord, topo [][]byte) {

	rows := len(topo)
	columns := len(topo[0])

	for r := 0; r < rows; r++ {
		visitmap := []byte{}
		for c := 0; c < columns; c++ {
			rc := coord{r: r, c: c}
			if _, f := visited[rc]; !f {
				visitmap = append(visitmap, '.')
				continue
			}

			next := visited[rc]
			if next.r > r {
				visitmap = append(visitmap, 'V')
				continue
			}
			if next.r < r {
				visitmap = append(visitmap, '^')
				continue
			}
			if next.c > c {
				visitmap = append(visitmap, '>')
				continue
			}
			if next.c < c {
				visitmap = append(visitmap, '<')
				continue
			}
		}
		fmt.Println(string(visitmap))
	}

	fmt.Println("-----")
}

func dynamicsearch(start, end coord, topo [][]byte, visited map[coord]coord) int {

	dbg := false
	dbg = start.r == 2 && start.c == 1

	if start == end {
		// fmt.Println("Found path:", visited)
		// printpath(visited, topo)
		return 0
	}

	if distance, found := distanced[start]; found {
		return distance
	}

	// fmt.Println(start, "start")

	minsteps := len(topo)*len(topo[0]) + 1

	rightcoord := coord{r: start.r, c: start.c + 1}
	visited[start] = rightcoord
	rightdist, err := trysearch(visited, rightcoord, start, end, topo, dbg)

	if !err && minsteps > rightdist {
		minsteps = rightdist
		if dbg {
			fmt.Println("right: ", minsteps)
		}
	}

	leftcoord := coord{r: start.r, c: start.c - 1}
	visited[start] = leftcoord
	leftdist, err := trysearch(visited, leftcoord, start, end, topo, dbg)

	if !err && minsteps > leftdist {
		minsteps = leftdist
		if dbg {
			fmt.Println("left: ", minsteps)
		}
	}

	upcoord := coord{r: start.r - 1, c: start.c}
	visited[start] = upcoord
	updist, err := trysearch(visited, upcoord, start, end, topo, dbg)

	if !err && minsteps > updist {
		minsteps = updist
		if dbg {
			fmt.Println("up: ", minsteps)
		}
	}

	downcoord := coord{r: start.r + 1, c: start.c}
	visited[start] = downcoord
	downdist, err := trysearch(visited, downcoord, start, end, topo, dbg)
	if dbg {
		fmt.Println(downdist, err)
	}
	if !err && minsteps > downdist {
		minsteps = downdist
		if dbg {
			fmt.Println("down: ", minsteps)
		}
	}

	if dbg {
		printpath(visited, topo)
	}

	distanced[start] = minsteps
	delete(visited, start)
	return minsteps
}

func trysearch(
	visited map[coord]coord,
	test coord,
	start coord,
	end coord,
	topo [][]byte, dbg bool) (int, bool) {

	if _, found := visited[test]; found {
		if dbg {
			fmt.Println("rejected due to visited")
		}
		return 0, true
	}

	if test.r < 0 || test.r >= len(topo) || test.c < 0 || test.c >= len(topo[0]) {
		if dbg {
			fmt.Println("rejected due to border")
		}
		return 0, true
	}

	p := int(topo[test.r][test.c])
	q := int(topo[start.r][start.c])
	if p-q > 1 {
		if dbg {
			fmt.Println("rejected due to height", string(byte(p)), string(byte(q)))
		}
		return 0, true
	}

	dist := 1 + dynamicsearch(test, end, topo, visited)

	return dist, false
}

func hillclimbing() {
	file, err := os.Open("Data\\test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var topo [][]byte = [][]byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		topo = append(topo, []byte(line))
	}

	fmt.Println("-----")

	var start, end coord

	for r := 0; r < len(topo); r++ {
		for c := 0; c < len(topo[r]); c++ {
			if topo[r][c] == 'S' {
				start = coord{r: r, c: c}
				topo[r][c] = 'a'
			}

			if topo[r][c] == 'E' {
				end = coord{r: r, c: c}
				topo[r][c] = 'z'
			}
		}
	}

	visited := make(map[coord]coord)
	dist := dynamicsearch(start, end, topo, visited)
	fmt.Println("dist=", dist)
	printDistanced(topo)

	fmt.Println(visited)
	// fmt.Println(distanced)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func printDistanced(topo [][]byte) {
	for r := 0; r < len(topo); r++ {
		for c := 0; c < len(topo[r]); c++ {
			if d, f := distanced[coord{r: r, c: c}]; f {
				fmt.Printf("%02d ", d)
			} else {
				fmt.Print(" x ")
			}
		}
		fmt.Println()
	}
}
