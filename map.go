package sync

import "sync"

// Map is a generic wrapper of [sync.Map].
type Map[K, V any] struct {
	smap sync.Map
}

// Clear deletes all the entries, resulting in an empty Map.
func (m *Map[K, V]) Clear() {
	m.smap.Clear()
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete
// returns false (even if the old value is the nil interface value).
func (m *Map[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.smap.CompareAndDelete(key, old)
}

// CompareAndSwap swaps the old and new values for key if the value stored in
// the map is equal to old. The old value must be of a comparable type.
func (m *Map[K, V]) CompareAndSwap(key K, old V, new V) (swapped bool) { //nolint: revive,predeclared // keep name new to match original implementation
	return m.smap.CompareAndSwap(key, old, new)
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	m.smap.Delete(key)
}

// Load returns the value stored in the map for a key, or nil if no value is
// present. The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	var a any
	a, ok = m.smap.Load(key)
	value, _ = a.(V)
	return
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	var a any
	a, loaded = m.smap.LoadAndDelete(key)
	value, _ = a.(V)
	return
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	var a any
	a, loaded = m.smap.LoadOrStore(key, value)
	actual, _ = a.(V)
	return
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.smap.Range(func(k, v any) bool {
		key, _ := k.(K)
		value, _ := v.(V)
		return f(key, value)
	})
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) {
	m.smap.Store(key, value)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	var a any
	a, loaded = m.smap.Swap(key, value)
	previous, _ = a.(V)
	return
}
