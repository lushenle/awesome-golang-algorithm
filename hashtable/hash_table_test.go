package hashtable

import (
	"fmt"
	"testing"
)

func Test_hash(t *testing.T) {
	type args struct {
		key Key
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testHash",
			args: args{key: 1},
			want: 1009613986114487,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.key); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func populateHashtable(count, start int) *HashTable {
	dict := HashTable{}
	for i := start; i < start+count; i++ {
		dict.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	return &dict
}

func TestValueHashTable_Put(t *testing.T) {
	dict := populateHashtable(3, 0)
	if size := dict.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	dict.Put("key1", "value1") //should not add a new one, just change the existing one
	if size := dict.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	dict.Put("key4", "value4") //should add it
	if size := dict.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}

	if value := dict.Get("key1"); value != "value1" {
		t.Errorf("wrong Get(), expected value1 and got %d", value)
	}
}

func TestValueHashTable_Remove(t *testing.T) {
	dict := populateHashtable(3, 0)
	dict.Remove("key2")
	if size := dict.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}
