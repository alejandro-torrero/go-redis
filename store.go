package main

import (
	"errors"
	"sort"
)

var ErrKeyDoesNotExist = errors.New("key does not exist")
var ErrEmptyStringKey = errors.New("An empty string is not a valid key")
var ErrFullStore = errors.New("Store is full, no more keys can be stored")

const DEFAULT_SIZE = 10

type Store struct {
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

	val, ok := s.data[key]

	if !ok {
		return "", ErrKeyDoesNotExist
	}

	return val, nil
}

func (s *Store) Set(key, value string) error {

	_, exists := s.Get(key)

	if s.IsFull() && exists != nil {
		return ErrFullStore
	}

	s.data[key] = value

	return nil
}

func (s *Store) Del(key string) {
	delete(s.data, key)
}

func (s *Store) Len() int {
	return len(s.data)
}

func (s *Store) IsFull() bool {
	return s.maxSize != 0 && s.Len() == s.maxSize
}

func (s *Store) KeyValue() map[string]string {
	return s.data
}
