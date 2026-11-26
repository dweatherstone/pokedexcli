package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		value    []byte
		initial  []byte
		expected []byte
	}{{
		name:     "add new entry",
		key:      "goodbye",
		value:    []byte("everyone"),
		expected: []byte("everyone"),
	}, {
		name:     "update existing entry",
		key:      "hello",
		initial:  []byte("world"),
		value:    []byte("someone"),
		expected: []byte("someone"),
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(time.Hour)
			if c.initial != nil {
				cache.Add(c.key, c.initial)
			}

			cache.Add(c.key, c.value)

			actual, found := cache.Get(c.key)
			if !found {
				t.Fatal("expected to find key in cache")
			}
			if !bytes.Equal(actual, c.expected) {
				t.Errorf("Expected: %q, got: %q", c.expected, actual)
			}
		})
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		name          string
		key           string
		expectedFound bool
		expectedValue []byte
	}{{
		name:          "retrieve existing key",
		key:           "hello",
		expectedFound: true,
		expectedValue: []byte("world"),
	}, {
		name:          "retrieve missing key",
		key:           "world",
		expectedFound: false,
		expectedValue: nil,
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(time.Hour)
			cache.Add("hello", []byte("world"))

			actualValue, actualFound := cache.Get(c.key)
			if actualFound != c.expectedFound {
				t.Errorf("expected found to be: %t, got: %t", c.expectedFound, actualFound)
			}
			if !bytes.Equal(actualValue, c.expectedValue) {
				t.Errorf("expected value: %q, got: %q", c.expectedValue, actualValue)
			}
		})
	}
}

func TestDifferentKeys(t *testing.T) {
	cache := NewCache(time.Hour)

	cache.Add("hello", []byte("world"))
	cache.Add("goodbye", []byte("everyone"))

	actual, found := cache.Get("hello")
	if !found {
		t.Fatal("expected to find key 'hello'")
	}
	if !bytes.Equal(actual, []byte("world")) {
		t.Errorf("expected %q, got %q", []byte("world"), actual)
	}

	actual, found = cache.Get("goodbye")
	if !found {
		t.Fatal("expected to find key 'goodbye'")
	}
	if !bytes.Equal(actual, []byte("everyone")) {
		t.Errorf("expected %q, got %q", []byte("everyone"), actual)
	}
}

func TestFreshEntryRemains(t *testing.T) {
	cache := NewCache(100 * time.Millisecond)

	cache.Add("hello", []byte("world"))

	time.Sleep(50 * time.Millisecond)

	actual, found := cache.Get("hello")

	if !found {
		t.Fatal("expected fresh entry to remain in cache")
	}

	if !bytes.Equal(actual, []byte("world")) {
		t.Errorf("expected %q, got %q", []byte("world"), actual)
	}
}

func TestOldEntryIsRemoved(t *testing.T) {
	interval := 50 * time.Millisecond
	cache := NewCache(interval)

	cache.Add("hello", []byte("world"))

	// Make the entry appear older than the cache interval
	cache.mu.Lock()
	entry := cache.cache["hello"]
	entry.createdAt = time.Now().Add(-2 * interval)
	cache.cache["hello"] = entry
	cache.mu.Unlock()

	// Allow time for the next ticker event
	deadline := time.Now().Add(time.Second)

	for time.Now().Before(deadline) {
		_, found := cache.Get("hello")

		if !found {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}

	t.Fatal("expected old entry to be removed from cache")
}
