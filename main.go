package main

import (
	"encoding/base64"
	"fmt"

	"github.com/go-redis/store"
	"github.com/go-redis/store/kv"
)

func main() {

	// s := CrateStorer()
	// mm := NewMetricsMiddleware(s)
	// val, err := SetEncryptedkey(mm, "a", "1")
	// fmt.Println("Encoded", val)

	// mm.Get("d")

	// ttlStore := ttl.NewTTLStore(50 * time.Millisecond)

	// SetEncryptedkey(ttlStore, "a", "1")
	// println(ttlStore.Get("a"))

	// time.Sleep(1 * time.Second)

	// _, err = ttlStore.Get("a")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// mm.Report()

	// kv := kv.NewStore(10)

	// PopulateDefaults(kv)

	// nKv := kv.Clone()
	// nKv.Set("env", "prod")

	// fmt.Println(kv)
	// fmt.Println(nKv)

	cmds := []Command{
		{Op: "SET", Key: "env", Value: "production"},
		{Op: "SET", Key: "version", Value: "0.0.1"},
		{Op: "SET", Key: "debug", Value: "true"},
		{Op: "GET", Key: "env"},
		{Op: "SET", Key: "region", Value: "us-west-1"},
		{Op: "GET", Key: "version"},
	}

	s := kv.NewStore(10)

	RestoreOnBoot(s, cmds)

	fmt.Println("Restored, store size: ", s.Len())
	fmt.Println("Keys", s.Keys())
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
