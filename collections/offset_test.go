package collections

import "testing"

func TestOffset(t *testing.T) {
	type schema struct {
		id    uint32
		title string
		ip    [4]byte
	}

	title := NewOffset(func(s *schema) *string { return &s.title })

	relation := []schema{
		{1, "abc", [4]byte{1, 2, 3, 4}},
		{2, "d", [4]byte{192, 168, 172, 1}},
		{3, "", [4]byte{9, 10, 11, 12}},
	}

	want := []string{"abc", "d", ""}
	for i, tuple := range relation {
		if got := *title.Get(&tuple); got != want[i] {
			t.Errorf("got %v, want %v", got, want[i])
		}
	}
}
