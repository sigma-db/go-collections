package collections

import "testing"

func TestBitSet(t *testing.T) {
	const capacity = 132
	const max = (capacity-2)/4*4 + 2

	bits := NewBitSet(capacity)
	if bits.Capacity() < capacity {
		t.Errorf("Expected capacity to be at least %d, got %d", capacity, bits.Capacity())
	}

	// set all bits i with i==2 (mod 4)
	for i := 0; i < capacity; i += 2 {
		bits.Set(i)
	}
	for i := 0; i < capacity; i += 4 {
		bits.Unset(i)
	}

	for i := 0; i < capacity; i++ {
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
