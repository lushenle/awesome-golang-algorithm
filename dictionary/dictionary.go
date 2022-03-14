package dictionary

import (
	"sync"

	"github.com/lushenle/awesome-golang-algorithm/generic"
)

// Key the key of the dictionary
type Key generic.Type

// Value the content of the dictionary
type Value generic.Type

// ValueDictionary the set of Items
type ValueDictionary struct {
	items map[Key]Value
	lock  sync.RWMutex
}

// Set adds a new item to the dictionary
func (d *ValueDictionary) Set(key Key, value Value) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.items == nil {
		d.items = make(map[Key]Value)
	}
	d.items[key] = value
}

// Delete remove a value from the dictionary, given it's key
func (d *ValueDictionary) Delete(key Key) bool {
	d.lock.Lock()
	defer d.lock.Unlock()

	_, ok := d.items[key]
	if ok {
		delete(d.items, key)
	}

	return ok
}

// Has return true if the key exists in the dictionary
func (d *ValueDictionary) Has(key Key) bool {
	d.lock.Lock()
	defer d.lock.Unlock()

	_, ok := d.items[key]
	return ok
}

// Get return the value associated with the key
func (d *ValueDictionary) Get(key Key) Value {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.items[key]
}

// Clear remove all the items from the dictionary
func (d *ValueDictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.items = make(map[Key]Value)
}

// Size return the amount of elements in the dictionary
func (d *ValueDictionary) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return len(d.items)
}

// Keys return a slice of all the keys present
func (d *ValueDictionary) Keys() []Key {
	d.lock.RLock()
	defer d.lock.RUnlock()

	var keys []Key
	for i := range d.items {
		keys = append(keys, i)
	}

	return keys
}

// Values return a slice of all the values present
func (d *ValueDictionary) Values() []Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	var values []Value
	for i := range d.items {
		values = append(values, d.items[i])
	}

	return values
}
