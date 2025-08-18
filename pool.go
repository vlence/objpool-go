package objpool

// A PoolObj is a reference to an object in the pool. It keeps
// track of the object and whether it is in use or not.
type PoolObj[T any] struct {
	obj   T
	inUse bool
}

// Obj returns the object referenced by this PoolObj.
func (poolObj *PoolObj[T]) Obj() *T {
	return &poolObj.obj
}

// Free makes the object referenced by this PoolObj for use by
// other clients.
func (poolObj *PoolObj[T]) Free() error {
	// TODO: probably not safe for concurrent use
	poolObj.inUse = false
	return nil
}

// A Pool is a collection of objects that can be reused.
type Pool[T any] struct {
	requestPoolObj chan bool
	poolObj        chan *PoolObj[T]
}

// NewPool creates a new pool of size T objects.
func NewPool[T any](size int) *Pool[T] {
	pool := &Pool[T]{}
	pool.poolObj = make(chan *PoolObj[T])
	pool.requestPoolObj = make(chan bool)

	go pool.server(size)

	return pool
}

// server accepts requests for pool objects. This method should be
// run as a separate goroutine.
func (pool *Pool[T]) server(size int) {
	var p *PoolObj[T]

	// TODO: is everything allocated in the stack?
	poolObjs := make([]PoolObj[T], size)

	for {
		select {
		case <-pool.requestPoolObj:
			p = nil

			for i := range size {
				poolObj := &poolObjs[i]

				if !poolObj.inUse {
					fmt.Printf("pool: free object at index %d\n", i)
					poolObj.inUse = true
					p = poolObj
					break
				}
			}

			if p == nil {
				panic(fmt.Errorf("pool: no free objects"))
			}

			pool.poolObj <- p
		}
	}
}

// Get returns a pool object that's free to use.
func (pool *Pool[T]) Get() *PoolObj[T] {
	pool.requestPoolObj <- true
	return <-pool.poolObj
}
