package objpool

type Arena interface {
	// Alloc allocates size bytes in the arena and returns a pointer to it.
	// If size bytes aren't available then an empty slice is returned.
	Alloc(size int) []byte

	// Free frees this arena and returns the number of bytes freed.
	Free() int
}

// An arena allocator that zeroes the memory before returning it.
type ZeroedArena []byte

// NewZeroedArena creates a new arena allocator of size cap.
func NewZeroedArena(cap int) ZeroedArena {
	if cap < 0 {
		return make(ZeroedArena, 0)
	}

	return make(ZeroedArena, 0, cap)
}

// Alloc allocates size bytes in arena and returns them.
// If size bytes isn't available then an empty slice is
// returned. The bytes returned are zeroed.
func (arena *ZeroedArena) Alloc(size int) []byte {
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

func (arena *ZeroedArena) Free() int {
	used := len(*arena)

	*arena = (*arena)[0:0]

	return used
}