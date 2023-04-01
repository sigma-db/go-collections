package collections

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

type rangeIterator[T Number] struct {
	i, limit    uint64
	value, step T
}

func (it *rangeIterator[T]) Next() bool {
	if it.i > 0 {
		it.value += it.step
	}
	it.i++
	return it.i <= it.limit
}

func (it *rangeIterator[T]) Iterator() (Iterator[*T], *T) {
	return it, &it.value
}

func Range[T Number](options ...T) IterableIterator[*T] {
	switch len(options) {
	case 0:
		panic(ErrNotEnoughRangeParameters)
	case 1:
		if to := options[0]; to > 0 {
			return &rangeIterator[T]{limit: uint64(options[0]), step: 1}
		}
	case 2:
		if from, to := options[0], options[1]; from < to {
			return &rangeIterator[T]{value: from, limit: uint64(to - from), step: 1}
		}
	case 3:
		switch from, to, step := options[0], options[1], options[2]; {
		case step == 0:
			panic(ErrRangeStepIsZero)
		case step > 0 && from >= to:
			break
		case step < 0 && from <= to:
			break
		default:
			return &rangeIterator[T]{value: from, limit: uint64((to - from) / step), step: step}
		}
	default:
		panic(ErrTooManyRangeParameters)
	}
	return &rangeIterator[T]{}
}
