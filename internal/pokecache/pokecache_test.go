package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()
	cache.Add("key1", []byte("this is a test"))

	tests := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("this is a test"),
		},
	}

	for _, test := range tests {
		cache.Add(test.inputKey, test.inputVal)

		actual, ok := cache.Get(test.inputKey)

		if !ok {
			t.Errorf("the key %s was not found", test.inputKey)
			continue
		}

		if string(actual) != string(test.inputKey) {
			t.Errorf("%s doesn't match %s", string(test.inputVal), string(actual))
		}
	}

}
