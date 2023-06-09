package collections

// A Predicate is a function that checks whether a value satisfies some condition.
type Predicate[T any] func(T) bool

// A Function is a function that transforms a value.
type Function[T, U any] func(T) U

// A Reducer is a function that combines two values into one.
type Reducer[T, U any] func(U, T) U

// A Sampler is a function that checks whether a value's index satisfies some condition.
type Sampler func(int) bool

// Find returns a poiner to the first element from `s` for which `f` evaluates to `true` or `nil` if no such element exists.
func Find[E any](s []E, f Predicate[E]) *E {
	for i := range s {
		if f(s[i]) {
			return &s[i]
		}
	}
	return nil
}

// Filter returns a new slice containing all elements from `s` for which `f` evaluates to `true`.
func Filter[E any](s []E, f Predicate[E]) []E {
	t := make([]E, 0, len(s))
	for i := range s {
		if f(s[i]) {
			t = append(t, s[i])
		}
	}
	return t
}

// Map applies `f` to each element of slice `s` and returns a new slice containing the results.
func Map[A, B any](s []A, f Function[A, B]) []B {
	t := make([]B, len(s))
	for i := range s {
		t[i] = f(s[i])
	}
	return t
}

// Fold applies `f` to each element of slice `s` and the current value of `a` and returns the final value of `a`.
func Fold[A, B any](s []A, a B, f Reducer[A, B]) B {
	for i := range s {
		a = f(a, s[i])
	}
	return a
}

// Reduce operates like Fold, but uses `s[0]` as the initial accumulator value and then folds `s[1:]`.
// If `s` is empty, Reduce panics.
func Reduce[E any](s []E, f Reducer[E, E]) E {
	if len(s) == 0 {
		panic(ErrEmptySliceParameter)
	}
	return Fold(s[1:], s[0], f)
}

// Sample returns a new slice containing all elements from `s` whose index satisfies `f`.
//
// Note that while Sample resembles Filter in that it returns a new slice containing a subset of the elements from `s`,
// Sample only considers the index of each element, whereas Filter considers the element's value itself.
func Sample[E any](s []E, f Sampler) []E {
	// collect indices of elements that satisfy f
	j := make([]int, 0, len(s))
	for i := range s {
		if f(i) {
			j = append(j, i)
		}
	}
	// copy elements from s that satisfy f into a new slice t
	t := make([]E, len(j))
	for i, j := range j {
		t[i] = s[j]
	}
	return t
}
