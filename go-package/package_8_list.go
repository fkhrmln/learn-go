package main

import (
	"container/list"
	"fmt"
)

func main() {
	linkedList := list.New()

	linkedList.PushBack("Fakhri")
	linkedList.PushBack("Rifky")
	linkedList.PushBack("Audry")

	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
