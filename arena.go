package objpool

// A simple arena allocator.
type Arena []byte

// NewArena creates a new arena allocator of size cap.
func NewArena(cap int) Arena {
	if cap < 0 {
		return make(Arena, 0)
	}

	return make(Arena, 0, cap)
}

// Alloc allocates size bytes in arena and returns them.
// If size bytes isn't available then an empty slice is
// returned. The bytes returned are zeroed.
func (arena *Arena) Alloc(size int) []byte {
	i := len(*arena)
	c := cap(*arena)
	free := c-i

	if size > free {
		return (*arena)[0:0]
	}

	for range size {
		*arena = append(*arena, 0)
	}

	return (*arena)[i:i+size]
}

// Free frees up the arena's memory for future allocations.
func (arena *Arena) Free() int {
	used := len(*arena)

	*arena = (*arena)[0:0]

	return used
}