package collections

type join[T, U any] struct {
	s []T
	t []U
	f JoinPredicate[T, U]
}

type joinIterator[T, U any] struct {
	*join[T, U]
	i, j int
	v    ReferencePair[T, U] // refers to join.s[i] and join.t[j]
	// vp   *Reference[Pair[T, U]] // refers to join.s[i] and join.t[j]
}

func (j join[T, U]) Iterator() (Iterator[Reference[Pair[T, U]]], Reference[Pair[T, U]]) {
	it := &joinIterator[T, U]{join: &j}
	return it, &it.v
}

func (it *joinIterator[T, U]) Next() bool {
	for it.i < len(it.s) {
		s, t := &it.s[it.i], &it.t[it.j]
		it.j++
		if it.j == len(it.t) {
			it.i++
			it.j = 0
		}
		if it.f(s, t) {
			it.v.First = s
			it.v.Second = t
			return true
		}
	}
	return false
}
