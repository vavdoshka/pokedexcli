package cache

import (
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {

	type testInput struct {
		key   string
		value []byte
	}

	type testCase struct {
		input    testInput
		expected []byte
	}

	cases := []testCase{
		{
			input: testInput{
				key:   "https://example.com/url",
				value: []byte("some text"),
			},
			expected: []byte("some text"),
		},
		{
			input: testInput{
				key:   "https://example.com/url",
				value: []byte("some other text"),
			},
			expected: []byte("some other text"),
		},
		{
			input: testInput{
				key:   "https://example.com/url_new",
				value: []byte(""),
			},
			expected: []byte(""),
		},
	}

	cache := NewCache(5 * time.Second)

	for _, v := range cases {

		cache.Add(v.input.key, v.input.value)

		actual, ok := cache.Get(v.input.key)
		if !ok {
			t.Errorf("%s was not found in cache %v", v.input.key, cache.cache)
			continue
		}

		if string(actual) != string(v.expected) {
			t.Errorf("actual '%v' does not equal to expected: '%v'", actual, v.expected)
		}
	}

}

func TestCacheAddGetMissing(t *testing.T) {

	cache := NewCache(5 * time.Second)
	cache.Add("some", []byte("some"))

	actual, ok := cache.Get("other")

	if ok {
		t.Errorf("found a key 'other' which was not set in map '%v'", cache.cache)
	}

	if actual != nil {
		t.Errorf("returned non default value '%v' for missing entry from map '%v'", actual, cache.cache)
	}
}

func TestCacheReap(t *testing.T) {

	cache := NewCache(5 * time.Second)
	cache.Add("some", []byte("some"))

	cache.cache["other"] = CacheEntry{
		time.Now().Add(-6 * time.Second),
		[]byte("other"),
	}

	now := time.Now()
	cache.reapNow(now)

	actual, ok := cache.Get("other")

	if ok {
		t.Errorf("found a key 'other' which had to be removed from map '%v'", cache.cache)
	}

	if actual != nil {
		t.Errorf("returned non default value '%v' for missing entry from map '%v'", actual, cache.cache)
	}

	actual, ok = cache.Get("some")

	if !ok {
		t.Errorf("did not found a key 'some' which had to stay in map '%v'", cache.cache)
	}

	if string(actual) != "some" {
		t.Errorf("wrong value for key 'some' returned '%v'", string(actual))
	}
}
