package collections

const (
	// UIntSizeEncodingLength is the number of bits required to encode the size of a uint.
	UIntSizeEncodingLength = 5 + (^uint(0) >> 63)

	// UIntSize is the size of a uint in bits.
	UIntSize = 1 << UIntSizeEncodingLength
)

type BitSet []uint

func NewBitSet(cap int) BitSet {
	quo, rem := cap>>UIntSizeEncodingLength, cap&(UIntSize-1)
	if rem != 0 {
		quo += 1
	}
	return make([]uint, quo)
}

func (bits BitSet) Capacity() int {
	return cap(bits) << UIntSizeEncodingLength
}

func (bits BitSet) IsSet(i int) bool {
	return bits[i>>UIntSizeEncodingLength]&mask(i) != 0
}

func (bits BitSet) Set(i int) {
	bits[i>>UIntSizeEncodingLength] |= mask(i)
}

func (bits BitSet) Unset(i int) {
	bits[i>>UIntSizeEncodingLength] &^= mask(i)
}

func (bits BitSet) Iterator() (Iterator[*int], *int) {
	it := &bitSetIterator{bits: bits}
	return it, &it.v
}

type bitSetIterator struct {
	bits BitSet
	i, v int
}

func (it *bitSetIterator) Next() bool {
	for ; it.i < len(it.bits)*UIntSize; it.i++ {
		i, j := it.i/UIntSize, it.i&(UIntSize-1)
		if it.bits[i]&(1<<j) != 0 {
			it.v = it.i
			it.i++
			return true
		}
	}
	return false
}

func mask(i int) uint {
	return 1 << (i & (UIntSize - 1))
}
