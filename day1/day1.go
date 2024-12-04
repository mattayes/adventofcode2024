package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Part2()
}

func Part1() {
	f, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lefts, rights []int64
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("Expected 2 parts, got %v", len(parts))
		}
		left, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		lefts = append(lefts, left)

		right, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		rights = append(rights, right)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})
	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})

	var total int64
	for i, left := range lefts {
		right := rights[i]
		if left < right {
			// abs
			left, right = right, left
		}
		diff := left - right
		total += diff
	}
}

func Part2() {
	f, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lefts []int64
	rights := make(map[int64]int64)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("Expected 2 parts, got %v", len(parts))
		}
		left, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		lefts = append(lefts, left)

		right, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		rights[right]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var similarityScore int64
	for _, left := range lefts {
		right := rights[left]
		score := left * right
		similarityScore += score
	}

	fmt.Println(similarityScore)
}
