package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	Part2()
}

func Part2() {
	f, err := os.Open("day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]byte
	x, y := -1, -1
	for i := 0; scanner.Scan(); i++ {
		row := scanner.Bytes()
		j := bytes.IndexAny(row, directions)
		if j != -1 {
			x, y = i, j
		}
		grid = append(grid, slices.Clone(row))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(grid) == 0 {
		log.Fatal("no data")
	}
	if x == -1 || y == -1 {
		log.Fatalf("x=%d y=%d", x, y)
	}

	route, _, isLoop := solve(grid, x, y)
	if isLoop {
		fmt.Println(1)
		return
	}

	var numLoops int
	for i, row := range route {
		for j, cell := range row {
			if i == x && j == y {
				continue
			}
			if _, ok := directionsIdxs[cell]; !ok {
				continue
			}

			grid[i][j] = obstacle
			if _, _, isLoop := solve(grid, x, y); isLoop {
				numLoops++
			}
			grid[i][j] = space
		}
	}

	fmt.Println(numLoops)
}

func Part1() {
	f, err := os.Open("day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]byte
	x, y := -1, -1
	for i := 0; scanner.Scan(); i++ {
		row := scanner.Bytes()
		j := bytes.IndexAny(row, directions)
		if j != -1 {
			x, y = i, j
		}
		grid = append(grid, slices.Clone(row))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(grid) == 0 {
		log.Fatal("no data")
	}
	if x == -1 || y == -1 {
		log.Fatalf("x=%d y=%d", x, y)
	}

	route, numSeen, _ := solve(grid, x, y)
	printRoute(route)
	fmt.Println(numSeen)
}

const (
	space    = byte('.')
	obstacle = byte('#')
	up       = byte('^')
	right    = byte('>')
	down     = byte('v')
	left     = byte('<')
)

var (
	directions     = string([]byte{up, right, down, left})
	directionsIdxs = map[byte]int{
		up:    0,
		right: 1,
		down:  2,
		left:  3,
	}
	directionToChange = map[byte][2]int{
		up:    {-1, 0},
		right: {0, 1},
		down:  {1, 0},
		left:  {0, -1},
	}
)

func solve(grid [][]byte, x, y int) ([][]byte, int, bool) {
	numRows, numCols := len(grid), len(grid[0])

	route := make([][]byte, numRows)
	for i, row := range grid {
		route[i] = slices.Clone(row)
	}

	seen := make(map[[2]int]map[[2]int]struct{})
	guard := grid[x][y]
	change := directionToChange[guard]

	var isLoop bool
OUTER:
	for {
		k := [2]int{x, y}
		seenChanges, ok := seen[k]
		if !ok {
			seenChanges = make(map[[2]int]struct{})
			seen[k] = seenChanges
		}
		if _, ok := seenChanges[change]; ok {
			// log.Print(k, slices.Collect(maps.Keys(seenChanges)), change)
			isLoop = true
			break
		}
		seenChanges[change] = struct{}{}
		route[x][y] = guard

	INNER:
		for {
			nextX, nextY := x+change[0], y+change[1]
			if nextX < 0 || nextX == numRows || nextY < 0 || nextY == numCols {
				// log.Print(k, slices.Collect(maps.Keys(seenChanges)), change)
				break OUTER
			}

			next := grid[nextX][nextY]
			switch next {
			case space, up: // starting position
				// log.Print(k, slices.Collect(maps.Keys(seenChanges)), change)
				x, y = nextX, nextY
				break INNER
			case obstacle:
				guard = turnRight(guard)
				change = directionToChange[guard]
				continue
			default:
				printRoute(route)
				log.Fatal(string(next))
			}
		}
	}

	return route, len(seen), isLoop
}

func turnRight(direction byte) byte {
	current := directionsIdxs[direction]
	next := (current + 1) % 4
	return directions[next]
}

func printRoute(route [][]byte) {
	for _, row := range route {
		fmt.Println(string(row))
	}
}
