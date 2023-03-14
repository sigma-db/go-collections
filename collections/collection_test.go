package collections

import "testing"

func TestChaining(t *testing.T) {
	c := Collection[int]{1, 2, 3, 4, 5, 6, 7, 8, 9}

	even := func(i int) bool { return i%2 == 0 }
	square := func(i int) int { return i * i }
	sum := func(a, b int) int { return a + b }

	got := c.Filter(even).Map(square).Reduce(sum)
	want := 120

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
