package kv

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// Sentinel errors
var ErrKeyDoesNotExist = errors.New("key does not exist")
var ErrEmptyStringKey = errors.New("An empty string is not a valid key")
var ErrFullStore = errors.New("Store is full, no more keys can be stored")

const DEFAULT_SIZE = 10

type Store struct {
	mu      sync.Mutex
	data    map[string]string
	maxSize int
}

func NewStore(maxSize int) *Store {

	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize, // 0 means unlimited
	}
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (s *Store) Get(key string) (string, error) {

	if key == "" {
		return "", ErrEmptyStringKey
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.data[key]

	if !ok {
		return "", ErrKeyDoesNotExist
	}

	return val, nil
}

func (s *Store) Set(key, value string) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[key]

	if s.IsFull() && exists {
		return ErrFullStore
	}

	s.data[key] = value

	time.Sleep(1 * time.Second)
	return nil
}

func (s *Store) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}

func (s *Store) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.data)
}

func (s *Store) IsFull() bool {
	return s.maxSize != 0 && len(s.data) == s.maxSize
}

func (s *Store) KeyValue() map[string]string {
	return s.data
}

func (s *Store) Clone() *Store {

	s.mu.Lock()
	defer s.mu.Unlock()

	cp := &Store{
		data:    make(map[string]string, len(s.data)),
		maxSize: s.maxSize,
	}

	for k, v := range s.data {
		cp.data[k] = v
	}

	return cp
}

func (s *Store) Incr(key string) (int, error) {
	if key == "" {
		return 0, ErrEmptyStringKey
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	current := 0

	if val, ok := s.data[key]; ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0, fmt.Errorf("INCR %q: %w", key, err)
		}
		current = v
	}

	current++
	time.Sleep(time.Millisecond)

	s.data[key] = strconv.Itoa(current)

	return current, nil
}
