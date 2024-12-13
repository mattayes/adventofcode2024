package main

import (
	"bytes"
	"container/list"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strconv"
	"sync"
	"sync/atomic"
)

func main() {
	Part2_2()
}

func Part2_2() {
	line, err := os.ReadFile("day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := bytes.Fields(line)
	data := make(map[uint64]int, len(parts))
	for _, b := range parts {
		v, err := strconv.ParseUint(string(b), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		data[v]++
	}

	for i := 0; i < 75; i++ {
		data = transform(data)
	}

	var total int
	for _, v := range data {
		total += v
	}
	fmt.Println(total)
}

func printData2(m map[uint64]int) {
	keys := slices.Sorted(maps.Keys(m))
	for _, k := range keys {
		fmt.Printf("k=%v v=%v\n", k, m[k])
	}
}

func transform(data map[uint64]int) map[uint64]int {
	new := make(map[uint64]int, len(data))
	for v, count := range data {
		if count == 0 {
			continue
		}

		if v == 0 {
			new[1] += count
			continue
		}

		numDigits := 1
		for v := v / 10; v > 0; v /= 10 {
			numDigits++
		}
		if numDigits%2 != 0 {
			new[v*2024] += count
			continue
		}

		half := numDigits / 2

		mod := uint64(10)
		for i := 1; i < half; i++ {
			mod *= 10
		}

		left := v
		for ; left > mod-1; left /= 10 {
		}

		right := v % mod

		new[left] += count
		new[right] += count
	}

	return new
}

func Part2() {
	line, err := os.ReadFile("day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := bytes.Fields(line)
	data := make([]uint64, len(parts))
	for i, b := range parts {
		var err error
		data[i], err = strconv.ParseUint(string(b), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	var total atomic.Uint64
	var wg sync.WaitGroup
	for _, v := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			total.Add(traverse(v, 75))
		}()
	}
	wg.Wait()

	fmt.Println(total.Load())
}

func traverse(v uint64, levelsLeft int) uint64 {
	if levelsLeft == 0 {
		return 1
	}

	if v == 0 {
		return traverse(1, levelsLeft-1)
	}
	numDigits := 1
	for v := v / 10; v > 0; v /= 10 {
		numDigits++
	}
	if numDigits%2 != 0 {
		return traverse(v*2024, levelsLeft-1)
	}

	half := numDigits / 2

	rightMod := uint64(10)
	for i := 1; i < half; i++ {
		rightMod *= 10
	}
	right := v % rightMod

	left := v
	for ; left > rightMod-1; left /= 10 {
	}

	lefts := traverse(left, levelsLeft-1)
	rights := traverse(right, levelsLeft-1)
	return lefts + rights
}

func Part1() {
	line, err := os.ReadFile("day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := list.New()
	for _, b := range bytes.Fields(line) {
		v, err := strconv.ParseUint(string(b), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		data.PushBack(&Element{Value: v})
	}
	printData(data)

	for i := 0; i < 25; i++ {
		for e := data.Front(); e != nil; e = e.Next() {
			v := e.Value.(*Element)

			if v.Value == 0 {
				v.Value = 1
				continue
			}
			numDigits := 1
			for v := v.Value / 10; v > 0; v /= 10 {
				numDigits++
			}
			if numDigits%2 == 0 {
				half := numDigits / 2

				rightMod := uint64(10)
				for i := 1; i < half; i++ {
					rightMod *= 10
				}
				rightV := v.Value % rightMod

				leftV := v.Value
				for ; leftV > rightMod-1; leftV /= 10 {
				}

				right := data.InsertBefore(&Element{Value: rightV}, e)
				data.InsertBefore(&Element{Value: leftV}, right)
				data.Remove(e)

				e = right
				continue
			}

			v.Value *= 2024
		}
	}

	// printData(data)
	fmt.Println(data.Len())
}

type Element struct {
	Value uint64
}

func printData(data *list.List) {
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value.(*Element).Value)
	}
	fmt.Printf("\n")
}
