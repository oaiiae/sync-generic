# Sync-Generic

A Go package providing generic wrappers around standard library synchronization primitives.

- Type safety at compile time
- No need for type assertions in your code
- Full API compatibility with standard library

## Features

- `Pool[T]` - Type-safe generic wrapper for `sync.Pool`
- `Map[K, V]` - Type-safe generic wrapper for `sync.Map`

## Usage

```go
// Pool example
func _() {
	pool := new(sync.Pool[[]byte]).New(func() []byte { return make([]byte, 1024) })
	buffer := pool.Get()
	defer pool.Put(buffer)
	// use buffer
}
```

```go
// Map example
func _() {
	cache := new(sync.Map[string, int])
	value, ok := cache.LoadOrStore("count", 42)
}
```
