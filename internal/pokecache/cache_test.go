package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
		{
			key: "",
			val: []byte{},
		},
	}

	cache := NewCache(time.Minute)

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			cache.Add(c.key, c.val)

			actual, ok := cache.Get(c.key)
			if !ok {
				t.Fatalf("expected to find key %q", c.key)
			}
			if string(actual) != string(c.val) {
				t.Errorf("expected value %q, got %q", c.val, actual)
			}
		})
	}
}

func TestGetMissing(t *testing.T) {
	cache := NewCache(time.Minute)
	cache.Add("present", []byte("value"))

	if _, ok := cache.Get("absent"); ok {
		t.Error("expected absent key to report ok == false")
	}
}

func TestAddOverwrites(t *testing.T) {
	cache := NewCache(time.Minute)
	cache.Add("key", []byte("first"))
	cache.Add("key", []byte("second"))

	actual, ok := cache.Get("key")
	if !ok {
		t.Fatal("expected to find key")
	}
	if string(actual) != "second" {
		t.Errorf("expected value %q, got %q", "second", actual)
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 10 * time.Millisecond
	const waitTime = baseTime * 5

	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	if _, ok := cache.Get("https://example.com"); !ok {
		t.Fatal("expected to find key immediately after adding it")
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("https://example.com"); ok {
		t.Error("expected key to be reaped after the interval elapsed")
	}
}

func TestReapLeavesFreshEntries(t *testing.T) {
	cache := NewCache(time.Hour)
	cache.Add("fresh", []byte("value"))

	// Reap with an interval far longer than the entry's age.
	cache.reap(time.Hour)

	if _, ok := cache.Get("fresh"); !ok {
		t.Error("expected a fresh entry to survive reaping")
	}
}

// TestConcurrent is only meaningful under `go test -race`.
func TestConcurrent(t *testing.T) {
	const goroutines = 50
	const opsPerGoroutine = 100

	cache := NewCache(time.Minute)

	var wg sync.WaitGroup
	for i := range goroutines {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range opsPerGoroutine {
				key := fmt.Sprintf("key-%d", j%10)
				cache.Add(key, []byte(fmt.Sprintf("val-%d-%d", id, j)))
				cache.Get(key)
			}
		}(i)
	}
	wg.Wait()

	if _, ok := cache.Get("key-0"); !ok {
		t.Error("expected key-0 to be present after concurrent writes")
	}
}
