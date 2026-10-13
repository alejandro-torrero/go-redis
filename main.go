package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-redis/store"
	"github.com/go-redis/store/kv"
	"github.com/go-redis/store/ttl"
)

func main() {
	kvStore := kv.NewStore(10)

	val, err := SetEncryptedkey(kvStore, "a", "1")
	fmt.Println("Encoded", val)

	ttlStore := ttl.NewTTLStore(50 * time.Millisecond)

	SetEncryptedkey(ttlStore, "a", "1")
	println(ttlStore.Get("a"))

	time.Sleep(1 * time.Second)

	_, err = ttlStore.Get("a")
	if err != nil {
		fmt.Println(err)
	}
}

func SetEncryptedkey(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))

	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return store.Get(key)
}
