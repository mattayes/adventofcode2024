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
	Part2()
}

const (
	shortExample = `AAAA
BBCD
BBCC
EEEC`
	longerExample = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
)

// Part2 does not work
func Part2() {
	// f, err := os.Open("day12.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	f := bytes.NewBufferString(shortExample)

	scanner := bufio.NewScanner(f)

	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, slices.Clone(scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(grid) == 0 {
		fmt.Println(0)
		return
	}

	visited := make(map[[2]int]struct{})
	possibleCorners := make(map[[2]int]struct{})
	var total int
	for i, row := range grid {
		for j, cell := range row {
			area := traverse2(grid, visited, possibleCorners, i, j, cell)
			if area == 0 {
				continue
			}

			var numSides int
			for coord := range possibleCorners {
				matches := map[[2]int]bool{
					{-1, 0}:  false,
					{0, 0}:   false,
					{0, -1}:  false,
					{-1, -1}: false,
				}
				for vector := range matches {
					i, j := coord[0]+vector[0], coord[1]+vector[1]
					if !isInGrid(grid, i, j) {
						continue
					}
					next := grid[i][j]
					if next == cell {
						matches[vector] = true
					}
				}

				var numMatching int
				for _, b := range matches {
					if b {
						numMatching++
					}
				}
				if numMatching == 1 || numMatching == 3 {
					numSides++
					continue
				}

				for _, diagonals := range [2][2][2]int{
					{{-1, -1}, {0, 0}},
					{{-1, 0}, {0, -1}},
				} {
					if matches[diagonals[0]] && matches[diagonals[1]] {
						numSides += 2
						break
					}
				}
			}

			total += area * numSides
			clear(possibleCorners)
		}
	}

	fmt.Println(total)
}

func traverse2(grid [][]byte, visited map[[2]int]struct{}, possibleCorners map[[2]int]struct{}, i, j int, cell byte) int {
	k := [2]int{i, j}
	if _, ok := visited[k]; ok {
		return 0
	}
	visited[k] = struct{}{}

	for _, vector := range [4][2]int{
		{0, 0},
		{0, 1},
		{1, 1},
		{1, 0},
	} {
		possibleCorners[[2]int{i + vector[0], j + vector[1]}] = struct{}{}
	}

	area := 1
	for _, vector := range [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	} {
		i, j := i+vector[0], j+vector[1]
		if !isInGrid(grid, i, j) {
			continue
		}
		next := grid[i][j]
		if next != cell {
			continue
		}

		moreArea := traverse2(grid, visited, possibleCorners, i, j, next)
		area += moreArea
	}

	return area
}

func isInGrid(grid [][]byte, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func Part1() {
	f, err := os.Open("day12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// 	f := bytes.NewBufferString(`RRRRIICCFF
	// RRRRIICCCF
	// VVRRRCCFFF
	// VVRCCCJFFF
	// VVVVCJJCFE
	// VVIVCCJJEE
	// VVIIICJJEE
	// MIIIIIJJEE
	// MIIISIJEEE
	// MMMISSJEEE`)

	scanner := bufio.NewScanner(f)

	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, slices.Clone(scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(grid) == 0 {
		fmt.Println(0)
		return
	}

	visited := make(map[[2]int]struct{})
	var total int
	for i, row := range grid {
		for j, cell := range row {
			area, perimeter := traverse(grid, visited, i, j, cell)
			if area != 0 {
				total += area * perimeter
			}
		}
	}

	fmt.Println(total)
}

func traverse(grid [][]byte, visited map[[2]int]struct{}, i, j int, cell byte) (int, int) {
	k := [2]int{i, j}
	if _, ok := visited[k]; ok {
		return 0, 0
	}
	visited[k] = struct{}{}

	area, perimeter := 1, 0
	for _, vector := range [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	} {
		i, j := i+vector[0], j+vector[1]
		if !(i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])) {
			perimeter++
			continue
		}

		next := grid[i][j]
		if cell != next {
			perimeter++
			continue
		}

		moreArea, morePerimeter := traverse(grid, visited, i, j, next)
		area += moreArea
		perimeter += morePerimeter
	}

	return area, perimeter
}
