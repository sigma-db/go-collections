package util

import (
	"testing"
)

func TestMemoize(t *testing.T) {
	f0 := Memoize0(func() int { return 42 })
	f1 := Memoize1(func(a int) int { return a })
	f2 := Memoize2(func(a int, b int) int { return a + b })
	f3 := Memoize3(func(a int, b int, c int) int { return a + b + c })
	f4 := Memoize4(func(a int, b int, c int, d int) int { return a + b + c + d })

	// before memoization
	if f0() != 42 {
		t.Error("f0() != 42")
	}
	if f1(1) != 1 {
		t.Error("f1(1) != 1")
	}
	if f2(1, 2) != 3 {
		t.Error("f2(1, 2) != 3")
	}
	if f3(1, 2, 3) != 6 {
		t.Error("f3(1, 2, 3) != 6")
	}
	if f4(1, 2, 3, 4) != 10 {
		t.Error("f4(1, 2, 3, 4) != 10")
	}

	// after memoization
	if f0() != 42 {
		t.Error("f0() != 42")
	}
	if f1(1) != 1 {
		t.Error("f1(1) != 1")
	}
	if f2(1, 2) != 3 {
		t.Error("f2(1, 2) != 3")
	}
	if f3(1, 2, 3) != 6 {
		t.Error("f3(1, 2, 3) != 6")
	}
	if f4(1, 2, 3, 4) != 10 {
		t.Error("f4(1, 2, 3, 4) != 10")
	}
}
