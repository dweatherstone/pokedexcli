package pokecache

import (
	"bytes"
	"testing"
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
			cache := NewCache()
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
			cache := NewCache()
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
	cache := NewCache()

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
