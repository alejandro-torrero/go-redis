package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore(10)
	store.Set("a", "10")
	store.Set("b", "20")

	got := store.Keys()
	want := []string{"a", "b"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("The keys are different. Expected %v, and got %v", want, got)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore(10)

	// Dev comment for testing
	// store.Set("test", "1")

	got := store.Keys()

	if len(got) > 0 {
		t.Errorf("Test failed, emtpy store didnt return empty keys")
	}
}

func TestKeys_EmptyKey(t *testing.T) {
	store := NewStore(10)

	_, err := store.Get("")

	if err == nil {
		t.Errorf("Store shouldnt accept empty strings as keys")
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore(10)

	store.Set("test", "right-value")

	val, err := store.Get("test")

	if val != "right-value" || err != nil {
		t.Errorf("Failed to retrieve key from store")
	}

	// store.Set("missing-key", "missing")
	val, err = store.Get("missing-key")

	if val != "" || err == nil {
		t.Error("A value was retrieved from a missing key, failed")
	}
}

func TestDelete(t *testing.T) {
	// Table driven test

	tests := []struct {
		name      string
		setup     map[string]string
		deleteKey string
		wantLen   int
	}{
		{"deletes existing key", map[string]string{"a": "1", "b": "2", "c": "3"}, "a", 2},
		{"no-op on missing key", map[string]string{"a": "1"}, "x", 1},
		{"no-op on empty store", map[string]string{}, "x", 0},
		{"deletes last remaining key", map[string]string{"only": "value"}, "only", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create store
			store := NewStore(10)

			// Fill store
			for k, v := range tc.setup {
				store.Set(k, v)
			}

			store.Del(tc.deleteKey)

			if got := store.Len(); got != tc.wantLen {
				t.Errorf("After delete(%q): Len() = %d, want %d", tc.deleteKey, got, tc.wantLen)
			}

			if _, ok := store.Get(tc.deleteKey); ok == nil {
				t.Errorf("After delete(%q): Key still present", tc.deleteKey)
			}

		})
	}
}

func TestCapacity(t *testing.T) {
	store := NewStore(1)
	store.Set("key", "val")

	err := store.Set("key2", "val2")

	if err == nil {
		t.Errorf("Store allowed more capacity than intended")
	}

	err = store.Set("key", "updated value")

	if err != nil {
		t.Errorf("Store didnt allow to update a key when full")
	}
}
