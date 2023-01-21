package main

import "fmt"

type calc func(int) int

type monkey struct {
	items      []int
	fcalc      calc
	test       int
	throwtrue  int
	throwfalse int
}

const maxrounds int = 10000
const nummonkeys int = 8

func monkeyinthemiddle() {

	monkeys := getmonkeys_big()

	supermod := 1

	for _, m := range monkeys {
		supermod *= m.test
	}

	counts := [nummonkeys]int{}

	for round := 0; round < maxrounds; round++ {
		for m := 0; m < nummonkeys; m++ {

			counts[m] += len(monkeys[m].items)

			for i := 0; i < len(monkeys[m].items); i++ {
				oldval := monkeys[m].items[i]

				newval := monkeys[m].fcalc(oldval)

				newval = newval % supermod

				to := 0
				if newval%monkeys[m].test == 0 {
					to = monkeys[m].throwtrue
				} else {
					to = monkeys[m].throwfalse
				}

				monkeys[to].items = append(monkeys[to].items, newval)

				// fmt.Println(m, "throws", newval, "to", to)

			}

			monkeys[m].items = []int{}

		}

	}

	fmt.Println("Finally:")
	fmt.Println(counts)

	// fmt.Println(counts)

}

func getmonkeys_big() []monkey {

	var monkeys [nummonkeys]monkey

	monkeys[0] = monkey{
		items:      []int{85, 79, 63, 72},
		fcalc:      func(old int) int { return old * 17 },
		test:       2,
		throwtrue:  2,
		throwfalse: 6,
	}
	monkeys[1] = monkey{
		items:      []int{53, 94, 65, 81, 93, 73, 57, 92},
		fcalc:      func(old int) int { return old * old },
		test:       7,
		throwtrue:  0,
		throwfalse: 2,
	}
	monkeys[2] = monkey{
		items:      []int{62, 63},
		fcalc:      func(old int) int { return old + 7 },
		test:       13,
		throwtrue:  7,
		throwfalse: 6,
	}
	monkeys[3] = monkey{
		items:      []int{57, 92, 56},
		fcalc:      func(old int) int { return old + 4 },
		test:       5,
		throwtrue:  4,
		throwfalse: 5,
	}
	monkeys[4] = monkey{
		items:      []int{67},
		fcalc:      func(old int) int { return old + 5 },
		test:       3,
		throwtrue:  1,
		throwfalse: 5,
	}
	monkeys[5] = monkey{
		items:      []int{85, 56, 66, 72, 57, 99},
		fcalc:      func(old int) int { return old + 6 },
		test:       19,
		throwtrue:  1,
		throwfalse: 0,
	}
	monkeys[6] = monkey{
		items:      []int{86, 65, 98, 97, 69},
		fcalc:      func(old int) int { return old * 13 },
		test:       11,
		throwtrue:  3,
		throwfalse: 7,
	}
	monkeys[7] = monkey{
		items:      []int{87, 68, 92, 66, 91, 50, 68},
		fcalc:      func(old int) int { return old + 2 },
		test:       17,
		throwtrue:  4,
		throwfalse: 3,
	}
	return monkeys[:]
}

// func getmonkeys() []monkey {

// 	// const nummonkeys int = 4
// 	var monkeys [nummonkeys]monkey

// 	monkeys[0] = monkey{
// 		items:      []int{79, 98},
// 		fcalc:      func(old int) int { return old * 19 },
// 		test:       23,
// 		throwtrue:  2,
// 		throwfalse: 3,
// 	}
// 	monkeys[1] = monkey{
// 		items:      []int{54, 65, 75, 74},
// 		fcalc:      func(old int) int { return old + 6 },
// 		test:       19,
// 		throwtrue:  2,
// 		throwfalse: 0,
// 	}
// 	monkeys[2] = monkey{
// 		items:      []int{79, 60, 97},
// 		fcalc:      func(old int) int { return old * old },
// 		test:       13,
// 		throwtrue:  1,
// 		throwfalse: 3,
// 	}
// 	monkeys[3] = monkey{
// 		items:      []int{74},
// 		fcalc:      func(old int) int { return old + 3 },
// 		test:       17,
// 		throwtrue:  0,
// 		throwfalse: 1,
// 	}
// 	return monkeys[:]

// }
