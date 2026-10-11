package main

import (
	"fmt"
)

func main() {
	store := NewStore(1)

	store.Set("a", "10")
	err := store.Set("b", "20")
	fmt.Println(err)

	aVal, _ := store.Get("a")
	bVal, _ := store.Get("b")

	fmt.Println(aVal)
	fmt.Println(bVal)

	store.Del("a")
	store.Set("c", "30")
	fmt.Println(store.Keys())

	store.Set("c", "40")
	fmt.Println(store.KeyValue())
}
