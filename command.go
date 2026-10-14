package main

import (
	"log"
	"sync"

	"github.com/go-redis/store"
)

type Command struct {
	Op    string
	Key   string
	Value string
}

func dispatch(s store.Storer, c Command) {
	switch c.Op {
	case "SET":
		log.Printf("set KEY %q", c.Key)
		s.Set(c.Key, c.Value)
	case "DEL":
		s.Del(c.Key)
	case "GET":
		s.Get(c.Key)
	case "INCR":
		s.Incr(c.Key)
	}
}

func RestoreOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup

	for _, c := range cmds {
		wg.Go(func() {
			dispatch(s, c)
		})

		// Older version
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	dispatch(s, c)
		// }()
	}
	wg.Wait()
}
