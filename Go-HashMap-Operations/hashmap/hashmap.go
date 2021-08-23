package hashmap

import (
	"fmt"
	"reflect"
)

// HashMap is the implementation of a hash map in Go
type HashMap map[interface{}]interface{}

// Item holds a single key-value pair
type Item struct {
	Key interface{}
	Val interface{}
}

var err error

// NewHashMap constructs new HashMap and returns reference
func NewHashMap() *HashMap {
	h := make(HashMap)
	return &h
}

// Print prints the key-value mappings alone with the data types of a HashMap
func (h *HashMap) Print() {
	if len(*h) == 0 {
		fmt.Println("Empty HashMap")
	}

	for k, v := range *h {
		fmt.Printf("\t%v (%T) -> %v (%T)\n", k, k, v, v)
	}
}

// NewItem constructs new Item and returns reference
func NewItem(k, v interface{}) *Item {
	return &Item{
		Key: k,
		Val: v,
	}
}

// Print prints the key-value mappings alone with the data types of a HashMap
func (i *Item) Print() {
	fmt.Printf(">>> %v (%T) -> %v (%T)\n\n", i.Key, i.Key, i.Val, i.Val)
}

// Put inserts key into HashMap and assigns the value. If key already exists in HashMap, it just replaces the value.
func (h *HashMap) Put(key, val interface{}) {
	(*h)[key] = val
}

// Put inserts key into HashMap and assigns the value only if  key doesn't exist in HashMap.
func (h *HashMap) PutIfAbsent(key, val interface{}) bool {
	if _, ok := (*h)[key]; !ok {

		// When key is not found
		(*h)[key] = val
		return true
	}

	// When key is found
	return false
}

// Merge merges given HashMap with the calling HashMap
// If key already exists in origin (calling) HashMap, values are replaced with newer ones.
// This is similar to MergeMissing, but replaces repeated keys.
func (h *HashMap) Merge(i *HashMap) {
	for k, v := range *i {
		(*h)[k] = v
	}
}

// MergeMissing merges given HashMap with the original (calling) HashMap
// If key already exists in original (calling) HashMap, values of given HashMap ate ignored.
// This is similar to Merge, but ignores repeated keys.
func (h *HashMap) MergeMissing(i *HashMap) {
	for k, v := range *i {
		if _, ok := (*h)[k]; !ok {
			(*h)[k] = v
		}
	}
}

// Clear clears all the items in HashMap
func (h *HashMap) Clear() {
	newHashMap := make(HashMap)
	*h = newHashMap
}

// IsEmpty returns true if HashMap is empty. Else, false.
func (h *HashMap) IsEmpty() bool {
	return len(*h) == 0
}

// Items iterates over items in the HashMap in a different goroutine. Items can be read from returned channel.
// In other words, this creates an iterator to read items one-by-one.
func (h *HashMap) Items() chan *Item {
	ch := make(chan *Item)
	go func() {
		for k, v := range *h {
			ch <- &Item{k, v}
		}
		close(ch)
	}()
	return ch
}

// Keys returns slice of keys in hashMap.
func (h *HashMap) Keys() []interface{} {
	keys := make([]interface{}, 0)
	for k := range *h {
		keys = append(keys, k)
	}
	return keys
}

// Values returns slice of values in hashMap.
func (h *HashMap) Values() []interface{} {
	vals := make([]interface{}, 0)
	for _, v := range *h {
		vals = append(vals, v)
	}
	return vals
}

// Clone creates a duplicate of the current HashMap with state.
func (h *HashMap) Clone() *HashMap {
	newHashMap := make(HashMap)
	newHashMap = *h
	return &newHashMap
}

// Size gives the number of items in HashMap.
func (h *HashMap) Size() int {
	return len(*h)
}

// Get returns the value of given key if found. Else, nil value.
func (h *HashMap) Get(k interface{}) interface{} {
	v, ok := (*h)[k]
	if !ok {

		// When key exists
		return nil
	}

	// When key doesn't exist
	return v
}

// Get returns the value of given key if found. Else, given default value.
func (h *HashMap) GetOrDefault(k, def interface{}) interface{} {
	v, ok := (*h)[k]
	if !ok {

		// When key exists
		return def
	}

	// When key doesn't exist
	return v
}

// Replace replaces the value to give key, with given new value.
// Replace() is similar to Put() and contrasting PutIfAbsent(). Replace() is like PutIfPresent.
func (h *HashMap) Replace(k, newVal interface{}) bool {
	if _, ok := (*h)[k]; !ok {

		// When key doesn't exist
		return false
	}

	// When key exists
	(*h)[k] = newVal
	return true
}

// ReplaceAll replaces all the values of keys from original (calling) HashMap with that in given HashMap.
// If key from given HashMap misses in original (calling) HashMap, ReplaceAll() ignores to write.
// If replacing + creating missing keys is required, use Merge() and MergeMissing()
func (h *HashMap) ReplaceAll(i *HashMap) {
	for k, v := range *i {
		if _, ok := (*h)[k]; ok {

			// When key exists
			(*h)[k] = v
		}
	}
}

// Equals compares original (calling) HashMap and given HashMap and returns true if equal.
func (h *HashMap) Equals(i *HashMap) bool {
	if len(*h) != len(*i) {
		return false
	}

	for k, v := range *i {
		if w, ok := (*h)[k]; ok {
			if !reflect.DeepEqual(v, w) {

				// When key exists in both but values doesn't match
				return false
			}
		} else {
			// When key from given HashMap is missing in original (calling) HashMap
			return false
		}
	}

	// When All keys and values of given HashMap match that or original (calling) HashMap
	return true
}

// ContainsKey checks if HashMap has the given key. Returns true if found. Else, false.
func (h *HashMap) ContainsKey(key interface{}) bool {
	for k := range *h {
		if k == key {

			// When given key is found in HashMap
			return true
		}
	}

	return false
}

// ContainsValue checks if HashMap has the given value. Returns true if found. Else, false.
func (h *HashMap) ContainsValue(val interface{}) bool {
	for _, v := range *h {
		if v == val {

			// When given value is found in HashMap
			return true
		}
	}

	return false
}

// Remove removes an item (key-value pair) from HashMap. If found and removed, it returns true. Else, false.
// This is to remove single item. To remove all items, use Clear().
func (h *HashMap) Remove(i *Item) bool {
	if val, ok := (*h)[i.Key]; !ok {

		// When key from given item is missing in HashMap
		return false
	} else {
		if val != i.Val {

			// When key from given item is found in HashMap, but it's corresponding value doesn't match that of item.
			return false
		} else {

			// When key from given item is found in HashMap, but it's corresponding value matches that of item.
			delete(*h, i.Key)
			return true
		}
	}
}

// ForEach runs given function with all the values in the HashMap.
// This is similar to Compute(), but this doesn't update values back to HashMap.
func (h *HashMap) ForEach(fun func(interface{})) {
	for _, v := range *h {
		fun(v)
	}
}

// Compute runs given function with all the values in the HashMap.
// This is similar to ForEach(), but this updates values back that given function returns, to HashMap.
func (h *HashMap) Compute(fun func(interface{}) (interface{}, error)) error {
	for k, v := range *h {
		(*h)[k], err = fun(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// ComputeIfAbsent runs given function if key is missing in the HashMap, \
// and adds key+computed value in HashMap.
// This is similar to ForEach(), but this updates values back that given function returns, to HashMap.
func (h *HashMap) ComputeIfAbsent(key interface{}, fun func() interface{}) (bool, error) {
	if _, ok := (*h)[key]; !ok {

		// When given key is missing in HashMap
		(*h)[key] = fun()
		return true, nil
	} else {

		// When given key is present in HashMap
		return false, nil
	}
}

// ComputeIfPresent runs given function if key is present in the HashMap, \
// and replaces actual value of key in HashMap with computed value.
// This is similar to ForEach(), but this updates values back that given function returns, to HashMap.
func (h *HashMap) ComputeIfPresent(key interface{}, fun func() interface{}) (bool, error) {
	if _, ok := (*h)[key]; ok {

		// When given key is present in HashMap
		(*h)[key] = fun()
		return true, nil
	} else {

		// When given key is missing in HashMap
		return false, nil
	}
}
