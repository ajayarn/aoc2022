package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	r int
	c int
}

type Val struct {
	letter byte
	dist   int
}

func hillclimbing() {
	file, err := os.Open("data/aoc12.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var topo [][]byte = [][]byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		topo = append(topo, []byte(line))
	}

	var start, end Coord
	visited := make(map[Coord]Val)
	unvisited := make(map[Coord]Val)

	// find the start and end
	R := len(topo)
	C := len(topo[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {

			coord := Coord{r: r, c: c}
			d := R * C

			if topo[r][c] == 'S' {
				topo[r][c] = 'a'
				start = coord
				d = 0
			}

			if topo[r][c] == 'E' {
				topo[r][c] = 'z'
				end = coord
			}

			val := Val{letter: topo[r][c], dist: d}
			unvisited[coord] = val
		}
	}

	fmt.Println("start: ", start, "end: ", end)

	node := start

	prevnode := node

	for node != end {
		val := unvisited[node]

		// fmt.Println("Current node: ", node, val)

		top := Coord{node.r - 1, node.c}
		bottom := Coord{node.r + 1, node.c}
		right := Coord{node.r, node.c + 1}
		left := Coord{node.r, node.c - 1}

		neighbors := []Coord{top, bottom, left, right}

		for _, neighbor := range neighbors {
			if nv, found := unvisited[neighbor]; found {
				if val.letter >= nv.letter || val.letter == nv.letter-1 {
					d := val.dist + 1
					if nv.dist > d {
						unvisited[neighbor] = Val{nv.letter, d}
					}
				}
			}
		}

		visited[node] = val

		delete(unvisited, node)

		M := R * C
		for k := range unvisited {
			if unvisited[k].dist < M {
				M = unvisited[k].dist
				node = k
			}
		}

		if node == prevnode {
			fmt.Println("Stuck in a loop?", node, prevnode)
			// M := R * C
			// for k := range unvisited {
			// 	fmt.Println(k, unvisited[k])
			// 	if unvisited[k].dist < M {
			// 		fmt.Println(k, unvisited[k])
			// 		M = unvisited[k].dist
			// 		node = k
			// 	}
			// }

			break
		}

		prevnode = node
	}

	fmt.Println("Result:", unvisited[end])

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
