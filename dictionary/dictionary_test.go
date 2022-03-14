package dictionary

import (
	"fmt"
	"testing"
)

func populateDictionary(count, start int) *ValueDictionary {
	dict := ValueDictionary{}
	for i := start; i < (start + count); i++ {
		dict.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	return &dict
}

func TestValueDictionary_Dictionary(t *testing.T) {
	dict := populateDictionary(3, 0)
	if size := dict.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	if value := dict.Get("key1"); value != "value1" {
		t.Errorf("wrong Get(), expected value1 and got %d", value)
	}

	if dict.Delete("key1") {
		if size := dict.Size(); size != 2 {
			t.Errorf("wrong count, expected 2 and got %d", size)
		}
	}

	if !dict.Has("key2") {
		t.Error("get key error, key2 not exist")
	}

	dict.Clear()
	if dict.Size() != 0 {
		t.Error("count size error")
	}
}

func TestValueDictionary_Keys(t *testing.T) {
	dict := populateDictionary(3, 0)
	items := dict.Keys()
	if len(items) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(items))
	}
	dict = populateDictionary(520, 0)
	items = dict.Keys()
	if len(items) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(items))
	}
}

func TestValueDictionary_Values(t *testing.T) {
	dict := populateDictionary(3, 0)
	items := dict.Values()
	if len(items) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(items))
	}
	dict = populateDictionary(520, 0)
	items = dict.Values()
	if len(items) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(items))
	}
}
