package collections

type CollectionError uint

const (
	_ CollectionError = iota

	// ErrEmptySliceParameter occurs when a function is called with an empty slice parameter where a non-empty slice is required.
	ErrEmptySliceParameter

	// ErrNotEnoughRangeParameters occurs when Range is called with no parameters.
	ErrNotEnoughRangeParameters

	// ErrTooManyParameters occurs when Range is called with more than three parameters.
	ErrTooManyRangeParameters

	// ErrRangeStepZero occurs when Range is called with a step size of zero.
	ErrRangeStepIsZero
)

func (e CollectionError) Error() string {
	switch e {
	case ErrEmptySliceParameter:
		return "empty slice parameter where a non-empty slice is required"
	case ErrNotEnoughRangeParameters:
		return "not enough parameters for Range"
	case ErrTooManyRangeParameters:
		return "too many parameters for Range"
	case ErrRangeStepIsZero:
		return "step size of Range is zero"
	default:
		return "unknown error"
	}
}
