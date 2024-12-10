package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	Part2()
}

func Part2() {
	f, err := os.Open("day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]int
	for scanner.Scan() {
		row := scanner.Bytes()
		ints := make([]int, len(row))
		for i, cell := range row {
			ints[i] = int(cell) - 48
		}
		grid = append(grid, ints)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(grid) == 0 {
		log.Fatal("no data")
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	var total int
	for i, row := range grid {
		for j, cell := range row {
			if cell != 0 {
				continue // not a trailhead
			}
			total += traverse2(grid, i, j, 0)
		}
	}

	fmt.Println(total)
}

func traverse2(grid [][]int, i, j, elevation int) int {
	if elevation == 9 {
		return 1
	}

	var total int
	for _, vector := range [][2]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	} {
		i, j := i+vector[0], j+vector[1]
		if !(i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])) {
			continue
		}
		next := grid[i][j]
		if next != elevation+1 {
			continue
		}
		total += traverse2(grid, i, j, next)
	}

	return total
}

func Part1() {
	f, err := os.Open("day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]int
	for scanner.Scan() {
		row := scanner.Bytes()
		ints := make([]int, len(row))
		for i, cell := range row {
			ints[i] = int(cell) - 48
		}
		grid = append(grid, ints)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(grid) == 0 {
		log.Fatal("no data")
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	var total int
	trails := make(map[[2]int]struct{})
	for i, row := range grid {
		for j, cell := range row {
			if cell != 0 {
				continue // not a trailhead
			}
			traverse(grid, trails, i, j, 0)
			total += len(trails)
			clear(trails)
		}
	}

	fmt.Println(total)
}

func traverse(grid [][]int, hits map[[2]int]struct{}, i, j, elevation int) {
	if elevation == 9 {
		hits[[2]int{i, j}] = struct{}{}
	}

	for _, vector := range [][2]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	} {
		i, j := i+vector[0], j+vector[1]
		if !(i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])) {
			continue
		}
		next := grid[i][j]
		if next != elevation+1 {
			continue
		}
		traverse(grid, hits, i, j, next)
	}
}
