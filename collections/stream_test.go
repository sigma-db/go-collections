package collections

import (
	"math"
	"testing"
)

func TestStreamFromIterable(t *testing.T) {
	s := FromValueIterable[int](Range(1, 10))
	s = Transform(s, FilterStream(func(x *int) bool { return *x > 1 }))
	result := Collect(s)

	if len(result) != 8 {
		t.Log(result)
		t.Errorf("Expected 8, got %d", len(result))
	}
	for i, v := range result {
		if v != i+2 {
			t.Errorf("Expected %d, got %d", i+2, v)
		}
	}
}

func TestStreamFromStreamer(t *testing.T) {
	s := asc(10).Stream()
	s = Transform(s, FilterStream(func(x *int) bool { return *x > 1 }))
	result := Collect(s)

	if len(result) != 8 {
		t.Log(result)
		t.Errorf("Expected 8, got %d", len(result))
	}
	for i, v := range result {
		if v != i+2 {
			t.Errorf("Expected %d, got %d", i+2, v)
		}
	}
}

func TestStreamComposition(t *testing.T) {
	s := FromValueIterable[float64](Range(0.5, 5.0, 0.25))
	s = Transform(s, FilterStream(func(x *float64) bool { _, f := math.Modf(*x); return f == 0.5 }))
	s = Transform(s, MapStream(func(x *float64) float64 { return *x * 2 }))
	s = Transform(s, ReduceStream(func(a float64, b *float64) float64 { return a + *b }))
	q := Collect(s)

	if l := len(q); l != 1 {
		t.Fatalf("Expected %d, got %d", 1, l)
	}
	if y := q[0]; y != 25 {
		t.Errorf("Expected %d, got %f", 25, y)
	}
}
