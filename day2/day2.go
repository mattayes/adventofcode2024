package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	Part2_2()
}

func Part1() {
	f, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var numSafe int
	for ints := make([]int64, 0); scanner.Scan(); ints = ints[:0] {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 1 {
			continue
		}

		for _, s := range parts {
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, v)
		}

		if _, ok := isSafe(ints); ok {
			numSafe++
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(numSafe)
}

func isSafe(ints []int64) (int, bool) {
	var isInc *bool
	validate := func(a, b int64) bool {
		diff := b - a
		if !validateDiffMagnitude(diff) {
			return false
		}

		isGT0 := diff > 0
		if isInc == nil {
			isInc = &isGT0
		} else if *isInc != isGT0 {
			return false
		}

		return true
	}

	for i := 1; i < len(ints); i++ {
		a, b := ints[i-1], ints[i]
		if !validate(a, b) {
			return i, false
		}
	}

	return -1, true
}

func Part2_2() {
	f, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var numSafe int

SCAN:
	for ints := make([]int64, 0); scanner.Scan(); ints = ints[:0] {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 1 {
			numSafe++
			continue
		}

		for _, s := range parts {
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, v)
		}

		if _, ok := isSafe(ints); ok {
			numSafe++
			continue
		}

		otherInts := make([]int64, len(ints)-1)
		for i := range ints {
			copy(otherInts[:i], ints[:i])
			copy(otherInts[i:], ints[i+1:])
			if _, ok := isSafe(otherInts); ok {
				numSafe++
				continue SCAN
			}
		}
	}

	fmt.Println(numSafe)
}

func Part2() {
	f, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var numSafe int
SCAN:
	for ints := make([]int64, 0); scanner.Scan(); ints = ints[:0] {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 1 {
			numSafe++
			continue
		}

		for _, s := range parts {
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, v)
		}
		if len(ints) == 2 && !validateDiffMagnitude(ints[1]-ints[0]) {
			numSafe++
			continue
		}

		validate := func(a, b int64, isInc *bool) (*bool, bool) {
			diff := b - a
			if !validateDiffMagnitude(diff) {
				return isInc, false
			}

			isGT0 := diff > 0
			if isInc == nil {
				return &isGT0, true
			} else if *isInc != isGT0 {
				return isInc, false
			}

			return isInc, true
		}

		var isInc *bool
		var numSkipped int
		for i := 2; i < len(ints); i++ {
			a, b, c := ints[i-2], ints[i-1], ints[i]

			abIsInc, abOK := validate(a, b, isInc)
			bcIsInc, bcOK := validate(b, c, abIsInc)
			if abOK && bcOK {
				isInc = bcIsInc
				continue
			}
			if numSkipped == 1 {
				continue SCAN
			}

			numSkipped++
			if abOK {
				isInc = abIsInc
				continue
			}

			acIsInc, acOK := validate(a, c, isInc)
			if !acOK {
				// can't skip two
				continue SCAN
			}

			isInc = acIsInc
		}

		numSafe++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(numSafe)
}

func validateDiffMagnitude(diff int64) bool {
	if diff == 0 {
		return false
	}
	if diff > 3 || diff < -3 {
		return false
	}
	return true
}
