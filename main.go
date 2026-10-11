package main

import (
	"fmt"
)

func main() {
	store := NewStore()

	store.Set("b", "20")
	store.Set("a", "10")

	bVal, _ := store.Get("b")
	aVal, _ := store.Get("a")

	fmt.Println(aVal)
	fmt.Println(bVal)
	fmt.Println(store.Keys())
}
