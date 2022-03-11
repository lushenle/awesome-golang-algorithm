package hashtable

import (
	"fmt"
	"sync"

	"github.com/lushenle/awesome-golang-algorithm/generic"
)

// Key the key of the dictionary
type Key generic.Type

// Value the content of the dictionary
type Value generic.Type

// HashTable the set of Items
type HashTable struct {
	items map[int]Value
	lock  sync.RWMutex
}

// hash internal function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
func hash(key Key) int {
	k := fmt.Sprintf("%s", key)
	h := 0
	for i := 0; i < len(k); i++ {
		h = 31*h + int(k[i])
	}

	return h
}

// Put item with value `value` and key `key` into the hashtable
func (ht *HashTable) Put(key Key, value Value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(key)
	if ht.items == nil {
		ht.items = make(map[int]Value)
	}
	ht.items[i] = value
}

// Get item with key from the hashtable
func (ht *HashTable) Get(key Key) Value {
	ht.lock.RLock()
	defer ht.lock.RUnlock()

	i := hash(key)
	return ht.items[i]
}

// Remove item with key from hashtable
func (ht *HashTable) Remove(key Key) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	i := hash(key)
	delete(ht.items, i)
}

// Size return the number of the hashtable elements
func (ht *HashTable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()

	return len(ht.items)
}
