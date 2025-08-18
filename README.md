# objpool-go
Simple object pool in Go.

Go provides an excellent garbage collector and makes memory management easy, but if you know
how much memory or how many objects you will need at runtime it makes sense to allocate them
at bulk ahead of time and reuse them whenever possible.

`Pool` guarantees that the `PoolObj` it gives you is not in use at the time you receive it. It
does not guarantee, however, that the object referenced by the `PoolObj` is not in use. Since
you define the objects that are stored in the pool it is therefore your responsibility to
guarantee that your objects do not cause race conditions.