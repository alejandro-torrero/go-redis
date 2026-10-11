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

	want := []string{}
	got := store.Keys()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Test failed, emtpy store didnt return empty keys")
	}
}
