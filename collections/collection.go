package collections

type Collection[T any] []T

type collectionIterator[T any] struct {
	*Collection[T]
	i int
	v T // == (*Collection)[i]
}

func (it *collectionIterator[T]) Next() bool {
	if it.i >= len(*it.Collection) {
		return false
	}
	it.v = (*it.Collection)[it.i]
	it.i++
	return true
}

func (it *collectionIterator[T]) Iterator() (Iterator[T], *T) {
	return it, &it.v
}

func (c Collection[T]) Iterator() (Iterator[T], *T) {
	it := &collectionIterator[T]{Collection: &c}
	return it, &it.v
}

func (c Collection[T]) Find(f Predicate[T]) *T {
	return Find(c, f)
}

func (c Collection[T]) Filter(f Predicate[T]) Collection[T] {
	return Filter(c, f)
}

func (c Collection[T]) Map(f Function[T, T]) Collection[T] {
	return Map(c, f)
}

func (c Collection[T]) Fold(a T, f Reducer[T, T]) T {
	return Fold(c, a, f)
}

func (c Collection[T]) Reduce(f Reducer[T, T]) T {
	return Reduce(c, f)
}
