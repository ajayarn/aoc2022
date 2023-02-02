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
			//	 fmt.Println(k, unvisited[k])
			//	 if unvisited[k].dist < M {
			//		 fmt.Println(k, unvisited[k])
			//		 M = unvisited[k].dist
			//		 node = k
			//	 }
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

func printtopo(
	visited, unvisited map[Coord]Val,
	R, C int,
) {
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			coord := Coord{r, c}
			p := 0
			if val, found := visited[coord]; found {
				p = val.dist
			} else if val, found := unvisited[coord]; found {
				p = val.dist
			}
			if p == R*C {
				p = -1
			}

			fmt.Printf("%03d ", p)
		}

		fmt.Println()
	}
}

func hillclimbing2() {
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

	var start Coord
	visited := make(map[Coord]Val)
	unvisited := make(map[Coord]Val)

	// find the start and end
	R := len(topo)
	C := len(topo[0])

	setofas := make(map[Coord]bool)
	distofas := make(map[Coord]int)

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {

			coord := Coord{r: r, c: c}
			d := R * C

			if topo[r][c] == 'S' {
				topo[r][c] = 'a'
			}

			if topo[r][c] == 'E' {
				topo[r][c] = 'z'
				start = coord
				d = 0
			}

			if topo[r][c] == 'a' {
				setofas[coord] = true
				distofas[coord] = d
			}

			val := Val{letter: topo[r][c], dist: d}
			unvisited[coord] = val
		}
	}

	node := start

	prevnode := node

	for len(setofas) > 0 && len(unvisited) > 0 {

		val := unvisited[node]

		fmt.Println("Current node: ", node, val)

		top := Coord{node.r - 1, node.c}
		bottom := Coord{node.r + 1, node.c}
		right := Coord{node.r, node.c + 1}
		left := Coord{node.r, node.c - 1}

		neighbors := []Coord{top, bottom, left, right}
		for _, neighbor := range neighbors {
			if nv, found := unvisited[neighbor]; found {
				if val.letter <= nv.letter+1 {
					d := val.dist + 1
					if nv.dist > d {
						unvisited[neighbor] = Val{nv.letter, d}
					}
				}
			}
		}

		visited[node] = val

		delete(unvisited, node)

		if _, found := setofas[node]; found {
			delete(setofas, node)
			distofas[node] = val.dist
		}

		M := R * C
		for k, v := range unvisited {
			if v.dist < M {
				M = v.dist
				node = k
			}
		}

		if node == prevnode {
			fmt.Println("Stuck in a loop?", node, prevnode)
			fmt.Println(unvisited)
			// M := R * C
			// for k := range unvisited {
			//	 fmt.Println(k, unvisited[k])
			//	 if unvisited[k].dist < M {
			//		 fmt.Println(k, unvisited[k])
			//		 M = unvisited[k].dist
			//		 node = k
			//	 }
			// }

			break
		}

		prevnode = node

		printtopo(visited, unvisited, R, C)
	}

	M := R * C
	coord := Coord{}
	for k, v := range distofas {
		if v < M {
			M = v
		}
		coord = k
	}

	fmt.Printf("Best dist: %d at %d, %d\n", M, coord.r, coord.c)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
