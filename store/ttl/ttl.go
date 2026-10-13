package ttl

import (
	"errors"
	"sort"
	"time"
)

var ErrEmptyKey = errors.New("Invalid empty key")
var ErrMissingKey = errors.New("Key doesnt exists")

type TTLStore struct {
	data map[string]TTLEntry
	ttl  time.Duration
}

type TTLEntry struct {
	value     string
	expiresAt time.Time
}

func NewTTLStore(ttl time.Duration) *TTLStore {
	return &TTLStore{
		data: make(map[string]TTLEntry),
		ttl:  ttl,
	}
}

func (t *TTLStore) Set(key, val string) error {
	if key == "" {
		return ErrEmptyKey
	}

	t.data[key] = TTLEntry{
		value:     val,
		expiresAt: time.Now().Add(t.ttl),
	}

	return nil
}

func (t *TTLStore) Get(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	entry, ok := t.data[key]

	if !ok || time.Now().After(entry.expiresAt) {
		delete(t.data, key)
		return "", errors.New("Expired value")
	}
	return entry.value, nil
}

func (t *TTLStore) Len() int {
	return len(t.data)
}

func (t *TTLStore) Del(key string) {
	delete(t.data, key)
}

func (s *TTLStore) Keys() []string {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}
