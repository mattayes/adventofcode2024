package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	Part2()
}

func Part2() {
	f, err := os.Open("day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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

	antennas := make(map[byte]map[[2]int]struct{})
	for i, row := range grid {
		for j, cell := range row {
			if cell == emptySpace {
				continue
			}

			coords, ok := antennas[cell]
			if !ok {
				coords = make(map[[2]int]struct{})
				antennas[cell] = coords
			}
			coords[[2]int{i, j}] = struct{}{}
		}
	}

	total := make(map[[2]int]struct{})
	for frequency, coords := range antennas {
		for coord1 := range coords {
			for coord2 := range coords {
				if coord1 == coord2 {
					continue
				}
				total[[2]int{coord1[0], coord1[1]}] = struct{}{}

				xChange, yChange := coord2[0]-coord1[0], coord2[1]-coord1[1]

				for i, j := coord2[0]+xChange, coord2[1]+yChange; i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]); i, j = i+xChange, j+yChange {
					cell := grid[i][j]
					if cell == frequency {
						continue
					}
					total[[2]int{i, j}] = struct{}{}
				}

			}
		}
	}

	fmt.Println(len(total))
}

func Part1() {
	f, err := os.Open("day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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

	antennas := make(map[byte]map[[2]int]struct{})
	for i, row := range grid {
		for j, cell := range row {
			if cell == emptySpace {
				continue
			}

			coords, ok := antennas[cell]
			if !ok {
				coords = make(map[[2]int]struct{})
				antennas[cell] = coords
			}
			coords[[2]int{i, j}] = struct{}{}
		}
	}

	total := make(map[[2]int]struct{})
	for frequency, coords := range antennas {
		for coord1 := range coords {
			for coord2 := range coords {
				if coord1 == coord2 {
					continue
				}

				xChange, yChange := coord2[0]-coord1[0], coord2[1]-coord1[1]
				i, j := coord2[0]+xChange, coord2[1]+yChange
				if !(i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])) {
					continue
				}
				cell := grid[i][j]
				if cell == frequency {
					continue
				}
				total[[2]int{i, j}] = struct{}{}
			}
		}
	}

	fmt.Println(len(total))
}

const (
	emptySpace = byte('.')
)
