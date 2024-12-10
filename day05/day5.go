package main

import (
	"bufio"
	"encoding/json"
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

func Part2() {
	f, err := os.Open("day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	isRules := true
	leftIsActive := make(map[int64]bool)
	rightToLeft := make(map[int64]map[int64]struct{})
	var ints []int64
	var total int64
	for i := 0; scanner.Scan(); i++ {
		s := scanner.Text()
		if len(s) == 0 {
			// gap
			isRules = false
			continue
		}

		if isRules {
			parts := strings.Split(s, "|")
			left, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			right, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			leftIsActive[left] = false

			lefts := rightToLeft[right]
			if lefts == nil {
				lefts = make(map[int64]struct{})
				rightToLeft[right] = lefts
			}
			lefts[left] = struct{}{}
		} else {
			parts := strings.Split(s, ",")
			if len(parts)%2 == 0 {
				log.Fatalf("didn't see that coming on line %v", i+1)
			}

			for _, s := range parts {
				x, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				ints = append(ints, x)
			}

			isPresent := make(map[int64]struct{}, len(ints))
			for _, x := range ints {
				isPresent[x] = struct{}{}
			}

			var incorrect bool
			for _, x := range ints {
				if _, ok := leftIsActive[x]; ok {
					leftIsActive[x] = true
				}

				for left := range rightToLeft[x] {
					if _, ok := isPresent[left]; !ok {
						continue // not in update
					}

					leftActive, ok := leftIsActive[left]
					if !ok {
						log.Fatal("should not happen")
					}
					if !leftActive {
						incorrect = true
						break
					}
				}
			}

			if incorrect {
				sort.Slice(ints, func(i, j int) bool {
					lefts, ok := rightToLeft[ints[j]]
					if !ok {
						return false
					}
					_, ok = lefts[ints[i]]
					return ok
				})
				middle := ints[len(ints)/2]
				total += middle
			}

			ints = ints[:0]
			for k := range leftIsActive {
				leftIsActive[k] = false
			}
		}
	}

	fmt.Println(total)
}

func Part1() {
	f, err := os.Open("day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	isRules := true
	leftIsActive := make(map[int64]bool)
	rightToLeft := make(map[int64]map[int64]struct{})
	var ints []int64
	var total int64
	for i := 0; scanner.Scan(); i++ {
		s := scanner.Text()
		if len(s) == 0 {
			// gap
			isRules = false

			b, err := json.MarshalIndent(leftIsActive, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))

			b, err = json.MarshalIndent(rightToLeft, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))

			continue
		}

		if isRules {
			parts := strings.Split(s, "|")
			left, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			right, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			leftIsActive[left] = false

			lefts := rightToLeft[right]
			if lefts == nil {
				lefts = make(map[int64]struct{})
				rightToLeft[right] = lefts
			}
			lefts[left] = struct{}{}
		} else {
			parts := strings.Split(s, ",")
			if len(parts)%2 == 0 {
				log.Fatalf("didn't see that coming on line %v", i+1)
			}

			for _, s := range parts {
				x, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				ints = append(ints, x)
			}

			isPresent := make(map[int64]struct{}, len(ints))
			for _, x := range ints {
				isPresent[x] = struct{}{}
			}

			var incorrect bool
			for _, x := range ints {
				if _, ok := leftIsActive[x]; ok {
					leftIsActive[x] = true
				}

				for left := range rightToLeft[x] {
					if _, ok := isPresent[left]; !ok {
						continue // not in update
					}

					leftActive, ok := leftIsActive[left]
					if !ok {
						log.Fatal("should not happen")
					}
					if !leftActive {
						incorrect = true
						break
					}
				}
			}

			log.Printf("row=%d correct=%t", i, !incorrect)
			if !incorrect {
				middle := ints[len(ints)/2]
				total += middle
			}

			ints = ints[:0]
			for k := range leftIsActive {
				leftIsActive[k] = false
			}
		}
	}

	fmt.Println(total)
}
