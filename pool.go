package sync

import "sync"

// Pool is a generic wrapper of [sync.Pool].
type Pool[T any] struct {
	pool sync.Pool
}

// New optionally specifies a function to generate a value when [Pool.Get]
// would otherwise have no item to return. It may not be changed concurrently
// with calls to [Pool.Get].
func (p *Pool[T]) New(f func() T) *Pool[T] {
	if f != nil {
		p.pool.New = func() any { return f() }
	} else {
		p.pool.New = nil
	}
	return p
}

// Get removes an arbitrary item from the [Pool] and returns it to the caller.
// Get may choose to ignore the pool and treat it as empty. Callers should not
// assume any relation between values passed to [Pool.Put] and the values
// returned by Get. If Get would be unable to return an item and [Pool.New] was
// called with a non-nil function, Get returns the result of calling this
// function. If no item can be returned, Get returns the zero value.
func (p *Pool[T]) Get() (x T) {
	x, _ = p.pool.Get().(T)
	return
}

// Put adds x to the Pool.
func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}
