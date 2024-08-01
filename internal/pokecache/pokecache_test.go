package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)

	tests := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("this is a test"),
		},
		{
			inputKey: "",
			inputVal: []byte(""),
		},
	}

	for _, test := range tests {
		cache.Add(test.inputKey, test.inputVal)

		actual, ok := cache.Get(test.inputKey)

		if !ok {
			t.Errorf("the key %s was not found", test.inputKey)
			continue
		}

		if string(actual) != string(test.inputVal) {
			t.Errorf("%q doesn't match %q",
				string(test.inputVal),
				string(actual),
			)
			continue
		}
	}

}

func TestCacheReap(t *testing.T) {
	const baseTime = time.Millisecond * 10
	const waitTime = baseTime * 2
	cache := NewCache(baseTime)
	cache.Add("https://boot.dev", []byte("testdata"))

	_, ok := cache.Get("https://boot.dev")
	if !ok {
		t.Error("the key was not found")
		return
	}

	time.Sleep(5 * time.Millisecond)

	_, ok = cache.Get("https://boot.dev")
	if !ok {
		t.Error("the key was not found")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://boot.dev")
	if ok {
		t.Error("the key was found")
		return
	}

}
