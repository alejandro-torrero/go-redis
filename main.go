package main

import (
	"encoding/base64"
	"fmt"

	"github.com/go-redis/store"
	"github.com/go-redis/store/kv"
)

func main() {
	s := kv.NewStore(0)

	s.Set("pageViews", "0")

	const hits = 1000

	cmds := make([]Command, hits)
	for i := range cmds {
		cmds[i] = Command{Op: "INCR", Key: "pageViews"}
	}

	RestoreOnBoot(s, cmds)

	got, _ := s.Get("pageViews")

	fmt.Println("expected: ", hits)
	fmt.Println("got: ", got)

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
