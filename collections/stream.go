package collections

type Stream[T any] interface {
	read() *T
}

type Streamer[T any] interface {
	Stream() Stream[T]
}

type TransformerConstructor[T, U any] Function[Stream[T], Stream[U]]

type transformStream[T, U any] struct {
	r Stream[T]
	f Function[*T, *U]
}

func Transform[T, U any](r Stream[T], ctor TransformerConstructor[T, U]) Stream[U] {
	return ctor(r)
}

func Collect[T any](r Stream[T]) []T {
	var result []T
	for v := r.read(); v != nil; v = r.read() {
		result = append(result, *v)
	}
	return result
}

func (s *transformStream[T, U]) read() *U {
	if v := s.r.read(); v != nil {
		return s.f(v)
	}
	return nil
}

type filterStream[T any] struct {
	r Stream[T]
	f Predicate[*T]
}

func FilterStream[T any](f Predicate[*T]) TransformerConstructor[T, T] {
	return func(r Stream[T]) Stream[T] {
		return &filterStream[T]{r, f}
	}
}

func (s *filterStream[T]) read() *T {
	for v := s.r.read(); v != nil; {
		if s.f(v) {
			return v
		}
		v = s.r.read()
	}
	return nil
}

type mapStream[T, U any] struct {
	s Stream[T]
	f Function[*T, U]
	v U
}

func MapStream[T, U any](f Function[*T, U]) TransformerConstructor[T, U] {
	return func(s Stream[T]) Stream[U] {
		return &mapStream[T, U]{s: s, f: f}
	}
}

func (s *mapStream[T, U]) read() *U {
	if v := s.s.read(); v != nil {
		s.v = s.f(v)
		return &s.v
	}
	return nil
}

type reduceStream[T any] struct {
	r Stream[T]
	f Reducer[*T, T]
}

func ReduceStream[T any](f Reducer[*T, T]) TransformerConstructor[T, T] {
	return func(s Stream[T]) Stream[T] {
		return &reduceStream[T]{s, f}
	}
}

func (s *reduceStream[T]) read() *T {
	if r := s.r.read(); r != nil {
		v := *r
		for r = s.r.read(); r != nil; r = s.r.read() {
			v = s.f(v, r)
		}
		return &v
	}
	return nil
}
