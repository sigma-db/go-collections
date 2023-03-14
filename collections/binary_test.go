package collections

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestCartesianProduct(t *testing.T) {
	a := []int{1, 2, 3}
	b := []string{"a", "b", "c"}

	want := []Pair[int, string]{
		{1, "a"}, {1, "b"}, {1, "c"},
		{2, "a"}, {2, "b"}, {2, "c"},
		{3, "a"}, {3, "b"}, {3, "c"},
	}
	got := CartesianProduct(a, b)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestJoin(t *testing.T) {
	r := []string{"a", "b", "c"}
	s := []int{1, 2, 3}
	condition := func(a *string, b *int) bool { return *a == string(rune(*b-1+'a')) }

	want := []Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}}
	got := Join(r, s, condition)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
