package collections

type CollectionError uint

const (
	_ CollectionError = iota

	// ErrEmptySliceParameter occurs when a function is called with an empty slice parameter where a non-empty slice is required.
	ErrEmptySliceParameter

	// ErrNotEnoughParameters occurs when Range is called with no parameters.
	ErrNotEnoughParameters

	// ErrTooManyParameters occurs when Range is called with more than three parameters.
	ErrTooManyParameters

	// ErrRangeStepZero occurs when Range is called with a step size of zero.
	ErrRangeStepIsZero

	// ErrEmptyStream occurs when an empty stream is passed where a non-empty stream is required.
	ErrEmptyStream

	// ErrSliceRangeOutOfBounds occurs when either end of a sub-slice is outside the source slide's range.
	ErrSliceRangeOutOfBounds

	// ErrIndexOutOfBounds occurs when an accessed index refers to memory outside the structure's allocated memory.
	ErrIndexOutOfBounds

	// ErrNegativeCapacity occurs when a negative capacity is passed to a function that requires a non-negative capacity.
	ErrNegativeCapacity
)

func (e CollectionError) Error() string {
	switch e {
	case ErrEmptySliceParameter:
		return "empty slice parameter where a non-empty slice is required"
	case ErrNotEnoughParameters:
		return "not enough parameters"
	case ErrTooManyParameters:
		return "too many parameters"
	case ErrRangeStepIsZero:
		return "step size of Range is zero"
	case ErrEmptyStream:
		return "empty stream where a non-empty stream is required"
	case ErrSliceRangeOutOfBounds:
		return "either end of a sub-slice is outside the source slide's range"
	case ErrIndexOutOfBounds:
		return "accessed index refers to memory outside the structure's allocated memory"
	case ErrNegativeCapacity:
		return "negative capacity where a non-negative capacity is required"
	default:
		return "unknown error"
	}
}
