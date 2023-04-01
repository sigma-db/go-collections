package collections

type cartesianProduct[T, U any] struct {
	s []T
	t []U
}

type cartesianProductStream[T, U any] struct {
	*cartesianProduct[T, U]
	i, j int
	r    ReferencePair[T, U] // points to s[i] and t[j]
	v    Pair[T, U]
}

func (s *cartesianProductStream[T, U]) read() *Pair[T, U] {
	if s.i >= len(s.s) {
		return nil
	}
	s.r.First = &s.s[s.i]
	s.r.Second = &s.t[s.j]
	s.j++
	if s.j == len(s.t) {
		s.i++
		s.j = 0
	}
	s.v = s.r.Value()
	return &s.v
}

func (cp cartesianProduct[T, U]) Stream() Stream[Pair[T, U]] {
	return &cartesianProductStream[T, U]{cartesianProduct: &cp}
}
