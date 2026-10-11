package main

import "sort"

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
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

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Del(key string) {
	delete(s.data, key)
}

func (s *Store) Len() int {
	return len(s.data)
}
