package collections

import "testing"

func TestBitSet(t *testing.T) {
	const (
		cap = 132
		max = (cap-2)/4*4 + 2
	)

	bits := NewBitSet(cap)
	if bits.Capacity() < cap {
		t.Errorf("Expected capacity to be at least %d, got %d", cap, bits.Capacity())
	}

	// set all bits i with i==2 (mod 4)
	for i := 0; i < cap; i += 2 {
		bits.Set(i)
	}
	for i := 0; i < cap; i += 4 {
		bits.Unset(i)
	}

	for i := 0; i < cap; i++ {
		if bits.IsSet(i) != (i%4 == 2) {
			t.Errorf("Expected bit %d to be %v", i, i%4 == 2)
		}
	}

	i := 2
	for it, v := bits.Iterator(); it.Next(); {
		if *v != i {
			t.Errorf("Expected %d, got %d", i, *v)
		}
		i += 4
	}
	if i != max+4 {
		t.Errorf("Expected %d, got %d", max+4, i)
	}
}
