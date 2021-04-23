package main

import (
	"fmt"
	"github.com/collinglass/bptree"
)

func main() {
	key := 7
	value := []byte("hello friend")

	t := bptree.NewTree()

	err := t.Insert(key, value)
	if err != nil {
		fmt.Printf("error: %s\n\n", err)
	}
    t.Insert(5, []byte("sea"))
    t.Insert(10, []byte("sky"))
	r, err := t.Find(key, true)
	if err != nil {
		fmt.Printf("error: %s\n\n", err)
	}

	fmt.Printf("%s\n\n", r.Value)

	t.FindAndPrint(key, true)
    t.PrintTree()
    t.PrintLeaves()
}

