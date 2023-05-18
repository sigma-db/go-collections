package collections

type Stream[T any] interface {
	Read() *T
}

type Streamer[T any] interface {
	Stream() Stream[T]
}

type Transformer[T, U any] Function[Stream[T], Stream[U]]

type Collector[T, U any] Function[Stream[T], U]

func Transform[T, U any](s Stream[T], t Transformer[T, U]) Stream[U] {
	return t(s)
}

func Collect[T, U any](s Stream[T], c Collector[T, U]) U {
	return c(s)
}

func SliceCollector[T any](s Stream[T]) []T {
	var result []T
	for v := s.Read(); v != nil; v = s.Read() {
		result = append(result, *v)
	}
	return result
}

func FoldCollector[T, U any](a U, f Reducer[*T, U]) Collector[T, U] {
	return func(s Stream[T]) U {
		for {
			if r := s.Read(); r != nil {
				a = f(a, r)
			} else {
				break
			}
		}
		return a
	}
}

func ReduceCollector[T any](f Reducer[*T, T]) Collector[T, T] {
	return func(s Stream[T]) T {
		if r := s.Read(); r != nil {
			v := *r
			for {
				if r := s.Read(); r != nil {
					v = f(v, r)
				} else {
					break
				}
			}
			return v
		}
		panic(ErrEmptyStream)
	}
}

type filterStream[T any] struct {
	in Stream[T]
	f  Predicate[*T]
}

func FilterStream[T any](f Predicate[*T]) Transformer[T, T] {
	return func(r Stream[T]) Stream[T] {
		return &filterStream[T]{r, f}
	}
}

func (s *filterStream[T]) Read() *T {
	for v := s.in.Read(); v != nil; {
		if s.f(v) {
			return v
		}
		v = s.in.Read()
	}
	return nil
}

type mapStream[T, U any] struct {
	in Stream[T]
	f  Function[*T, U]
	v  U
}

func MapStream[T, U any](f Function[*T, U]) Transformer[T, U] {
	return func(s Stream[T]) Stream[U] {
		return &mapStream[T, U]{in: s, f: f}
	}
}

func (s *mapStream[T, U]) Read() *U {
	if v := s.in.Read(); v != nil {
		s.v = s.f(v)
		return &s.v
	}
	return nil
}

type foldStream[T, U any] struct {
	in Stream[T]
	f  Reducer[*T, U]
	v  U
}

func FoldStream[T, U any](a U, f Reducer[*T, U]) Transformer[T, U] {
	return func(s Stream[T]) Stream[U] {
		return &foldStream[T, U]{s, f, a}
	}
}

func (s *foldStream[T, U]) Read() *U {
	for {
		if r := s.in.Read(); r != nil {
			s.v = s.f(s.v, r)
		} else {
			break
		}
	}
	return &s.v
}

type reduceStream[T any] struct {
	in Stream[T]
	f  Reducer[*T, T]
}

func ReduceStream[T any](f Reducer[*T, T]) Transformer[T, T] {
	return func(s Stream[T]) Stream[T] {
		return &reduceStream[T]{s, f}
	}
}

func (s *reduceStream[T]) Read() *T {
	if r := s.in.Read(); r != nil {
		v := *r
		for {
			if r := s.in.Read(); r != nil {
				v = s.f(v, r)
			} else {
				break
			}
		}
		return &v
	}
	return nil
}

type sampleStream[T any] struct {
	in Stream[T]
	f  Predicate[int]
	i  int
}

func SampleStream[T any](f Predicate[int]) Transformer[T, T] {
	return func(s Stream[T]) Stream[T] {
		return &sampleStream[T]{in: s, f: f}
	}
}

func (s *sampleStream[T]) Read() *T {
	r := s.in.Read()
	for ; r != nil; r = s.in.Read() {
		if s.f(s.i) {
			break
		}
		s.i++
	}
	s.i++
	return r
}
