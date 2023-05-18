package collections

import "testing"

func TestRangeInt(t *testing.T) {
	const from, to, step = 2, 12, 3
	t.Run("Range[int]/1", func(t *testing.T) {
		i := 0
		for it, v := Range(to).Iterator(); it.Next(); {
			if *v != i {
				t.Errorf("Expected %d, got %d", i, *v)
			}
			i += 1
		}
		if i != to {
			t.Errorf("Expected %d, got %d", to, i)
		}
	})

	t.Run("Range[int]/2", func(t *testing.T) {
		i := from
		for it, v := Range(from, to).Iterator(); it.Next(); {
			if *v != i {
				t.Errorf("Expected %d, got %d", i, *v)
			}
			i += 1
		}
		if i != to {
			t.Errorf("Expected %d, got %d", to, i)
		}
	})

	t.Run("Range[int]/3", func(t *testing.T) {
		i := from
		for it, v := Range(from, to, step).Iterator(); it.Next(); {
			if *v >= to {
				t.Errorf("Value %d exceeds limit %d", *v, to)
			}
			if *v != i {
				t.Errorf("Expected %d, got %d", i, *v)
			}
			i += step
		}
		last := from + (to-from)/step*step
		if i < last {
			t.Errorf("Expected %d, got %d", last, i)
		}
	})
}
