package collections

import "testing"

func asc(limit int) Generator[int] {
	return func(yield Yield[int]) {
		for i := 0; i < limit; i++ {
			yield(&i)
		}
	}
}

func TestGenerator(t *testing.T) {
	var result []int
	for it, v := Generator[int](asc(10)).Iterator(); it.Next(); {
		result = append(result, *v)
	}
	if len(result) != 10 {
		t.Errorf("Expected 10, got %d", len(result))
	}
	for i, v := range result {
		if v != i {
			t.Errorf("Expected %d, got %d", i, v)
		}
	}
}
