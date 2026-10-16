package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/store"
)

type Command struct {
	Op    string
	Key   string
	Value string
}

func dispatch(s store.Storer, c Command) error {
	switch c.Op {
	case "SET":
		log.Printf("set KEY %q", c.Key)
		return s.Set(c.Key, c.Value)
	case "DEL":
		s.Del(c.Key)
		return nil
	case "GET":
		_, err := s.Get(c.Key)
		return err
	case "INCR":
		_, err := s.Incr(c.Key)
		return err
	default:
		return fmt.Errorf("Unexpected operation %q", c.Op)
	}
}

func RestoreOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup

	errs := make(chan error, len(cmds))

	for _, c := range cmds {
		wg.Go(func() {
			errs <- dispatch(s, c)
		})

		// Older version
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	dispatch(s, c)
		// }()
	}

	// Close the channel at the end
	go func() {
		wg.Wait()
		close(errs)
	}()

	// log errors

	for err := range errs {
		if err != nil {
			fmt.Printf("Error %v\n", err)
		}
	}
}
