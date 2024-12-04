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
	f, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var data [][]byte
	for scanner.Scan() {
		b := scanner.Bytes()
		b = slices.Clone(b)
		data = append(data, b)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(data) == 0 {
		log.Fatal("no data")
	}

	numRows, numCols := len(data), len(data[0])
	var totalMatches int
	for i, line := range data {
		for j, b := range line {
			if b != 'A' {
				continue
			}

			// look in a X shape
			coords := [2][2][2]int{
				{{-1, 1}, {1, -1}},
				{{1, 1}, {-1, -1}},
			}
			var hits int
			for _, spot := range coords {
				i0, j0 := i+spot[0][0], j+spot[0][1]
				if i0 < 0 || i0 >= numRows || j0 < 0 || j0 >= numCols {
					continue
				}

				i1, j1 := i+spot[1][0], j+spot[1][1]
				if i1 < 0 || i1 >= numRows || j1 < 0 || j1 >= numCols {
					continue
				}

				b0, b1 := data[i0][j0], data[i1][j1]
				if (b0 != 'M' && b0 != 'S') || (b1 != 'M' && b1 != 'S') {
					continue
				}

				isM := b0 == 'M'
				isS := b1 == 'S'
				if isM == isS {
					hits++
				}
			}
			if hits == len(coords) {
				totalMatches++
			}
		}
	}

	fmt.Println(totalMatches)
}

func Part1() {
	f, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var data [][]byte
	for scanner.Scan() {
		b := scanner.Bytes()
		b = slices.Clone(b)
		data = append(data, b)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(data) == 0 {
		log.Fatal("no data")
	}

	numRows, numCols := len(data), len(data[0])
	letters := []byte{'M', 'A', 'S'}
	var totalMatches int
	for i, line := range data {
		for j, b := range line {
			if b != 'X' {
				continue
			}

			// look in a cirle
			for _, spot := range [][2]int{
				{-1, 0},
				{-1, 1},
				{0, 1},
				{1, 1},
				{1, 0},
				{1, -1},
				{0, -1},
				{-1, -1},
			} {
				i, j := i, j // shadow for modification below
				var hit int
				for _, letter := range letters {
					i, j = i+spot[0], j+spot[1]
					if i < 0 || i >= numRows || j < 0 || j >= numCols {
						continue
					}

					b := data[i][j]
					if b == letter {
						hit++
					}
				}

				if hit == len(letters) {
					totalMatches++
				}
			}
		}
	}

	fmt.Println(totalMatches)
}
