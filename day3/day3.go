package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	Part2()
}

func Part2() {
	b, err := os.ReadFile("day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	regex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := regex.FindAll(b, -1)

	var total int64
	enabled := true
	for _, match := range matches {
		switch s := string(match); s {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if !enabled {
				continue
			}

			s = s[4 : len(s)-1]
			parts := strings.Split(s, ",")
			if len(parts) != 2 {
				log.Fatalf("expected 2, got %v", len(parts))
			}

			ints := make([]int64, len(parts))
			for i, s := range parts {
				var err error
				ints[i], err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
			}

			product := int64(1)
			for _, v := range ints {
				product *= v
			}

			total += product
		}
	}

	fmt.Println(total)
}

func Part1() {
	b, err := os.ReadFile("day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	multiplyRegex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := multiplyRegex.FindAll(b, -1)

	var total int64
	for _, match := range matches {
		s := string(match[4 : len(match)-1])
		parts := strings.Split(s, ",")
		if len(parts) != 2 {
			log.Fatalf("expected 2, got %v", len(parts))
		}

		ints := make([]int64, len(parts))
		for i, s := range parts {
			var err error
			ints[i], err = strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
		}

		product := int64(1)
		for _, v := range ints {
			product *= v
		}

		total += product
	}

	fmt.Println(total)
}
