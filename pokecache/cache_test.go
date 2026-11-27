package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("first", []byte("original"))
	cache.Add("second", []byte("other"))
	cache.Add("first", []byte("updated"))

	for _, tc := range []struct {
		key   string
		want  string
		found bool
	}{
		{key: "first", want: "updated", found: true},
		{key: "second", want: "other", found: true},
		{key: "missing", found: false},
	} {
		t.Run(tc.key, func(t *testing.T) {
			got, ok := cache.Get(tc.key)
			if ok != tc.found || string(got) != tc.want {
				t.Fatalf("Get(%q) = %q, %v; want %q, %v", tc.key, got, ok, tc.want, tc.found)
			}
			if !ok && got != nil {
				t.Fatal("missing key should return nil")
			}
		})
	}
}

func TestReaping(t *testing.T) {
	cache := NewCache(5 * time.Millisecond)
	cache.Add("1", []byte("test"))
	cache.Add("2", []byte("test2"))
	time.Sleep(20 * time.Millisecond)
	cache.Add("3", []byte("test3"))

	for _, tc := range []struct {
		key    string
		exists bool
	}{
		{key: "1", exists: false},
		{key: "2", exists: false},
		{key: "3", exists: true},
	} {
		_, ok := cache.Get(tc.key)
		if ok && !tc.exists {
			t.Fatalf("%s is not cleared", tc.key)
		}
		if !ok && tc.exists {
			t.Fatalf("%s is should still exist", tc.key)
		}
	}
}
