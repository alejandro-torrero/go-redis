package main

import (
	"encoding/base64"
	"fmt"

	"github.com/go-redis/store"
	"github.com/go-redis/store/kv"
)

// Mutex 1.14s
// RWMutex 12 ms

func main() {
	s := kv.NewStore(0)

	s.Set("PageViews", "0")

	const hits = 1000
	cmds := make([]Command, hits)

	for i := range hits {
		cmds[i] = Command{Op: "INCR", Key: "PageViews"}
	}

	cmds = append(cmds, Command{Op: "INCR", Key: ""}, Command{Op: "WAT", Key: "PageViews"})

	RestoreOnBoot(s, cmds)

	got, _ := s.Get("PageViews")
	fmt.Printf("expected %d\n", hits)
	fmt.Printf("got %s\n", got)

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
