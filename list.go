package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Randi")
	data.PushBack("Febriadi")
	data.PushBack("Hebat")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Randi

	next := head.Next() // Febriadi
	fmt.Println(next.Value)

	next = next.Next() // Hebat
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
