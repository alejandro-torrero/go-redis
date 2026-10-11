package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore()
	store.Set("a", "10")
	store.Set("b", "20")

	got := store.Keys()
	want := []string{"a", "b"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("The keys are different. Expected %v, and got %v", want, got)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore()

	// Dev comment for testing
	// store.Set("test", "1")

	got := store.Keys()

	if len(got) > 0 {
		t.Errorf("Test failed, emtpy store didnt return empty keys")
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore()

	store.Set("test", "right-value")

	val, _ := store.Get("test")

	if val != "right-value" {
		t.Errorf("Failed to retrieve key from store")
	}

	// store.Set("missing-key", "missing")
	val, _ = store.Get("missing-key")

	if val != "" {
		t.Error("A value was retrieved from a missing key, failed")
	}

}
