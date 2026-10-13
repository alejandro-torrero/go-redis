package ttl

import (
	"testing"
	"time"
)

func TestTTLStore_GetBeforeExpiry(t *testing.T) {
	s := NewTTLStore(1 * time.Second)
	_ = s.Set("k", "v")

	got, err := s.Get("k")

	if got != "v" || err != nil {
		t.Errorf("expected v, got %q %v", got, err)
	}
}

func TestTTLStore_GetAfterExpiry(t *testing.T) {
	s := NewTTLStore(50 * time.Millisecond)

	_ = s.Set("k", "v")

	time.Sleep(1 * time.Second)

	val, err := s.Get("k")

	if err == nil {
		t.Errorf("Key didnt expired as expected. %q", val)
	}
}

func TestTTLStore_ExpiredKeyEvicted(t *testing.T) {
	s := NewTTLStore(50 * time.Millisecond)
	_ = s.Set("k", "v")

	time.Sleep(100 * time.Millisecond)

	s.Get("k")

	if s.Len() != 0 {
		t.Errorf("The key didnt got deleted on time")
	}

}
