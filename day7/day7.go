package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	Part2()
}

func Part2() {
	f, err := os.Open("day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int64
	var possibleTotal int64
	for i := 0; scanner.Scan(); i++ {
		b := scanner.Bytes()
		before, after, ok := bytes.Cut(b, []byte{':', ' '})
		if !ok {
			log.Fatal(i)
		}

		trueTotal, err := strconv.ParseInt(string(before), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		possibleTotal += trueTotal

		parts := bytes.Fields(after)
		ints := make([]int64, len(parts))
		for i, b := range parts {
			var err error
			ints[i], err = strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(ints) == 0 {
			log.Fatal("no numbers")
		}

		ops := []opFunc{
			multiplication,
			combination,
			addition,
		}
		if check(ops, trueTotal, 0, ints) {
			total += trueTotal
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total, "not", possibleTotal)
}

func Part1() {
	f, err := os.Open("day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int64
	var possibleTotal int64
	for i := 0; scanner.Scan(); i++ {
		b := scanner.Bytes()
		before, after, ok := bytes.Cut(b, []byte{':', ' '})
		if !ok {
			log.Fatal(i)
		}

		trueTotal, err := strconv.ParseInt(string(before), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		possibleTotal += trueTotal

		parts := bytes.Fields(after)
		ints := make([]int64, len(parts))
		for i, b := range parts {
			var err error
			ints[i], err = strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
		if len(ints) == 0 {
			log.Fatal("no numbers")
		}

		ops := []opFunc{
			multiplication,
			addition,
		}
		if check(ops, trueTotal, 0, ints) {
			total += trueTotal
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total, "not", possibleTotal)
}

func check(ops []opFunc, trueTotal, total int64, remaining []int64) bool {
	if len(remaining) == 0 {
		return trueTotal == total
	}

	for i, v := range remaining {
		for _, opFunc := range ops {
			total := opFunc(total, v)
			if check(ops, trueTotal, total, remaining[i+1:]) {
				return true
			}
		}
	}

	return false
}

type opFunc func(int64, int64) int64

func addition(a, b int64) int64 {
	return a + b
}

func multiplication(a, b int64) int64 {
	return a * b
}

func combination(a, b int64) int64 {
	multiplier := int64(10)
	for b := b / 10; b > 0; multiplier *= 10 {
		b /= 10
	}

	return a*multiplier + b
}
