package collections

type rangeIterator[T Number] struct {
	i, count uint64
	v, step  T
}

func (it *rangeIterator[T]) Next() bool {
	if it.i > 0 {
		it.v += it.step
	}
	it.i++
	return it.i <= it.count
}

func (it *rangeIterator[T]) Iterator() (Iterator[*T], *T) {
	return it, &it.v
}

func Range[T Number](options ...T) IterableIterator[*T] {
	switch len(options) {
	case 0:
		panic(ErrNotEnoughParameters)
	case 1:
		if to := options[0]; 0 < to {
			return &rangeIterator[T]{count: uint64(to), step: 1}
		}
	case 2:
		if from, to := options[0], options[1]; from < to {
			return &rangeIterator[T]{v: from, count: uint64(to - from), step: 1}
		}
	case 3:
		switch from, to, step := options[0], options[1], options[2]; {
		case step == 0:
			panic(ErrRangeStepIsZero)
		case
			step > 0 && from >= to,
			step < 0 && from <= to:
			break
		default:
			return &rangeIterator[T]{v: from, count: uint64((to - from) / step), step: step}
		}
	default:
		panic(ErrTooManyParameters)
	}
	return &rangeIterator[T]{}
}
