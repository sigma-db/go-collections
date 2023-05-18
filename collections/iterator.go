package collections

type Iterator[T any] interface {
	Next() bool
}

type Iterable[T any] interface {
	Iterator() (Iterator[T], T)
}

type IterableIterator[T any] interface {
	Iterator[T]
	Iterable[T]
}

type Reference[T any] interface {
	Value() T
}

type valueIteratorStream[T any] struct {
	it Iterator[T]
	v  *T
}

type referenceIteratorStream[T any] struct {
	it Iterator[Reference[T]]
	r  Reference[T]
	v  T
}

func (s *referenceIteratorStream[T]) Read() *T {
	if s.it.Next() {
		s.v = s.r.Value()
		return &s.v
	}
	return nil
}

func (s *valueIteratorStream[T]) Read() *T {
	if s.it.Next() {
		return s.v
	}
	return nil
}

func FromValueIterable[T any](it Iterable[*T]) Stream[T] {
	it2, v := it.Iterator()
	return &valueIteratorStream[T]{it2, v}
}

func FromInterfaceIterable[T any](it Iterable[*T]) Stream[T] {
	it2, v := it.Iterator()
	return &valueIteratorStream[T]{it2, v}
}

func FromReferenceIterable[T any](it Iterable[Reference[T]]) Stream[T] {
	it2, v := it.Iterator()
	return &referenceIteratorStream[T]{it: it2, r: v}
}

func Channel[T any](i Iterable[T]) <-chan T {
	ch := make(chan T)
	go func() {
		for it, v := i.Iterator(); it.Next(); {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
