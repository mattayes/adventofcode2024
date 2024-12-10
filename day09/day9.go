package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
)

func main() {
	Part2()
}

func Part2() {
	// line := []byte(`2333133121414131402`)
	line, err := os.ReadFile("day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := list.New()
	for i, b := range line {
		size := int(b) - 48
		if size == 0 {
			continue
		}
		e := Element{
			Index: i,
			Size:  size,
		}
		if i%2 == 0 {
			// file
			id := i / 2
			e.FileID = &id
		}
		data.PushBack(&e)
	}
	if data.Len() == 0 {
		log.Fatal("no data")
	}

	back := data.Back()
	for back != nil {
		backV := back.Value.(*Element)
		if backV.FileID == nil || backV.Index < 0 {
			// space or already moved
			back = back.Prev()
			continue
		}

		prev := back.Prev()
	FRONTLOOP:
		for front := data.Front(); front != nil; front = front.Next() {
			frontV := front.Value.(*Element)
			if frontV.Index < 0 {
				continue // moved file
			}
			if !(frontV.Index < backV.Index) {
				break // too far
			}
			if frontV.FileID != nil {
				continue // file
			}

			switch {
			case frontV.Size == backV.Size:
				if prev == front {
					prev = prev.Prev() // skip
				}
				space := Element{
					Index: -backV.Index,
					Size:  backV.Size,
				}
				data.InsertAfter(&space, back)

				backV.Index = -backV.Index
				data.MoveBefore(back, front)

				data.Remove(front)
				break FRONTLOOP
			case frontV.Size > backV.Size:
				space := Element{
					Index: -backV.Index,
					Size:  backV.Size,
				}
				data.InsertAfter(&space, back)

				backV.Index = -backV.Index
				data.MoveBefore(back, front)

				frontV.Size -= backV.Size
				break FRONTLOOP
			default:
				continue // not enough room
			}
		}

		back = prev
	}

	var total int
	var i int
	for e := data.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Element)
		for j := 0; j < v.Size; j++ {
			if v.FileID != nil {
				total += i * *v.FileID
			}
			i++
		}
	}

	fmt.Println(total)
}

func Part1() {
	line, err := os.ReadFile("day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := list.New()
	for i, b := range line {
		size := int(b) - 48
		if size == 0 {
			continue
		}
		e := Element{
			Index: i,
			Size:  size,
		}
		if i%2 == 0 {
			// file
			id := i / 2
			e.FileID = &id
		}
		data.PushBack(&e)
	}
	if data.Len() == 0 {
		log.Fatal("no data")
	}

	front := data.Front()
	frontV := front.Value.(*Element)

	back := data.Back()
	backV := back.Value.(*Element)
OUTER:
	for frontV.Index < backV.Index {
		if frontV.FileID != nil {
			front = front.Next()
			frontV = front.Value.(*Element)
			continue
		}

		for frontV.Size > 0 {
			for backV.FileID == nil {
				prev := back.Prev()
				data.Remove(back)
				back = prev
				backV = back.Value.(*Element)
				if !(frontV.Index < backV.Index) {
					break OUTER
				}
			}

			switch {
			case frontV.Size == backV.Size:
				prev := back.Prev()
				data.MoveBefore(back, front)
				frontV.Size = 0

				back = prev
				backV = back.Value.(*Element)
				if !(frontV.Index < backV.Index) {
					break OUTER
				}
			case frontV.Size > backV.Size:
				prev := back.Prev()
				data.MoveBefore(back, front)
				frontV.Size -= backV.Size

				back = prev
				backV = back.Value.(*Element)
				if !(frontV.Index < backV.Index) {
					break OUTER
				}

			default:
				// frontV.Size < backV.Size
				fileID := *backV.FileID
				e := Element{
					Index:  -1,
					Size:   frontV.Size,
					FileID: &fileID,
				}
				data.InsertBefore(&e, front)
				backV.Size -= frontV.Size
				frontV.Size = 0
			}
		}

		next := front.Next()
		data.Remove(front)
		front = next
		if front == nil {
			break
		}
		frontV = front.Value.(*Element)
	}

	var total int
	var i int
	for e := data.Front(); e != nil && e.Value.(*Element).FileID != nil; e = e.Next() {
		v := e.Value.(*Element)
		for j := 0; j < v.Size; j++ {
			total += i * *v.FileID
			i++
		}
	}

	fmt.Println(total)
}

type Element struct {
	Index  int
	Size   int
	FileID *int
}
