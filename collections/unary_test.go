package collections

import (
	"fmt"
	"strconv"
	"testing"

	"golang.org/x/exp/slices"
)

func newEqualityPredicate[T comparable](c T) Predicate[T] {
	return func(v T) bool { return c == v }
}

func TestFind(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := []int{0, 1, 5, 9, 10}
	want := []*int{nil, &s[0], &s[4], &s[8], nil}
	for i, e := range v {
		t.Run(fmt.Sprintf("Find %d", e), func(t *testing.T) {
			if got := Find(s, newEqualityPredicate(e)); got != want[i] {
				t.Errorf("got %v, want %v", got, want[i])
			}
		})
	}
}

func TestMap(t *testing.T) {
	s := []Pair[int, int]{{1, 2}, {3, 4}, {5, 6}}
	f := func(p Pair[int, int]) string { return strconv.Itoa(p.First + p.Second) }
	want := []string{"3", "7", "11"}
	if got := Map(s, f); !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFold(t *testing.T) {
	asc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	empty := []int{}

	sum := func(a, b int) int { return a + b }
	cat := func(a string, b int) string { return a + strconv.Itoa(b) }

	type input[T, U any] struct {
		s []T
		a U
		f Reducer[T, U]
	}

	tcintint := []Pair[input[int, int], int]{
		{input[int, int]{asc, 0, sum}, 45},
		{input[int, int]{asc, 5, sum}, 50},
		{input[int, int]{empty, 58, sum}, 58},
	}

	tcintstr := []Pair[input[int, string], string]{
		{input[int, string]{asc, "0", cat}, "0123456789"},
	}

	for _, tc := range tcintint {
		input, want := tc.Unpack()
		t.Run(fmt.Sprintf("Fold(%v, %v, %v)", input.s, input.a, input.f), func(t *testing.T) {
			if got := Fold(input.s, input.a, input.f); got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}

	for _, tc := range tcintstr {
		input, want := tc.Unpack()
		t.Run(fmt.Sprintf("Fold(%v, %v, %v)", input.s, input.a, input.f), func(t *testing.T) {
			if got := Fold(input.s, input.a, input.f); got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	asc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	empty := []int{}
	single := []int{58}

	sum := func(a, b int) int { return a + b }
	mul := func(a, b int) int { return a * b }

	input := []Pair[[]int, Reducer[int, int]]{
		{asc, sum},
		{asc, mul},
		{single, sum},
		{single, mul},
	}
	want := []int{45, 362880, 58, 58}

	for i, input := range input {
		s, f := input.Unpack()
		t.Run(fmt.Sprintf("Reduce(%v, %v)", s, f), func(t *testing.T) {
			if got := Reduce(s, f); got != want[i] {
				t.Errorf("got %v, want %v", got, want[i])
			}
		})
	}

	t.Run("Reduce(s, f) panics for empty s", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Reduce(s, f) should panic for all f when len(s) == 0")
			}
		}()
		Reduce(empty, sum)
	})
}
