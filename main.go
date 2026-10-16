package main

import (
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/store"
	"github.com/go-redis/store/kv"
)

// Mutex 1.14s
// RWMutex 12 ms

func main() {
	s := kv.NewStore(0)

	s.Set("name", "jane")

	const readers = 100
	const readsEach = 10

	start := time.Now()

	var wg sync.WaitGroup
	for range readers {
		wg.Go(func() {
			for range readsEach {
				s.Get("name")
			}
		})
	}

	wg.Wait()

	fmt.Printf("%d readers did %d concurrent reads in %s\n", readers, readsEach, time.Since(start).Round(time.Millisecond))

}

func CrateStorer() store.Storer {
	kvStore := kv.NewStore(10)
	logger := NewLoggerMiddleware(kvStore)
	return logger
}

func SetEncryptedkey(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))

	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return store.Get(key)
}

func PopulateDefaults(s store.Storer) error {
	defaults := map[string]string{
		"env":     "dev",
		"version": "1.0.0",
		"debug":   "true",
	}

	for k, v := range defaults {
		if err := s.Set(k, v); err != nil {
			return fmt.Errorf("PopulatedDefaults: %w", err)
		}
	}
	return nil
}
