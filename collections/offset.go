package collections

import "unsafe"

// A Selector[T, U any] is a function that returns a pointer to a field of type U in a struct of type T.
type Selector[T, U any] func(*T) *U

// A Offset[T, U any] is the offset of a field of type U in a struct of type T.
type Offset[T, U any] uint

// Get returns a pointer to the field of type U in struct t of type T that is referred to by this Offset.
//
// That is, a call to o.Get(&t) on an Offset[T, U] o created by a call to newOffset[T, U](func(*T) *U { return &t.u }) returns &t.u.
//
//	type Type struct {
//	    // ...
//		id uint
//	    // ...
//	}
//	id := newOffset(func(*Type) *uint { return &t.id })
//
//	t := Type{id: 42}
//	_ = *id.Get(&t) == t.id // true
//
// Note that the Offset is not validated prior to accessing the referenced field and thus may cause a panic if it does not point to a field of type U inside the given struct t of type T.
func (o Offset[T, U]) Get(t *T) *U {
	return (*U)(unsafe.Add(unsafe.Pointer(t), o))
}

// NewOffset returns the offset of the field of type U in a struct of type T that is referred to by the given Selector.
//
// If the address returned by the Selector is not within the bounds of the struct, NewOffset panics with ErrSelectorOffsetOutOfBounds.
func NewOffset[T, U any](s Selector[T, U]) Offset[T, U] {
	var x T
	y := s(&x)
	xp := uintptr(unsafe.Pointer(&x))
	yp := uintptr(unsafe.Pointer(y))
	if xp <= yp && yp < xp+unsafe.Sizeof(x) {
		return Offset[T, U](yp - xp)
	}
	panic(ErrSelectorOffsetOutOfBounds)
}
