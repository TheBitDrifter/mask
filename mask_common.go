package mask

// bitSize represents the number of bits in a uint64
const bitSize = 64

// Mark sets the specified bit in the mask
func (m *Mask) Mark(bit uint32) {
	idx := bit / bitSize
	bitPosition := bit % bitSize
	m[idx] |= (1 << bitPosition)
}

// Contains checks if the specified bit is set in the mask
func (m Mask) Contains(bit uint32) bool {
	idx := bit / bitSize
	offset := bit - (bitSize * idx)
	mask := uint64(1 << offset)
	return m[idx]&mask == mask
}

// ContainsAll checks if all bits set in the other mask are also set in this mask
func (m Mask) ContainsAll(other Mask) bool {
	for i, v := range other {
		if m[i]&v != v {
			return false
		}
	}
	return true
}

// ContainsAny checks if any bit set in the other mask is also set in this mask
func (m Mask) ContainsAny(other Mask) bool {
	for i, v := range other {
		if m[i]&v != 0 {
			return true
		}
	}
	return false
}

// ContainsNone checks if none of the bits set in the other mask are set in this mask
// Returns false if the other mask is empty
func (m Mask) ContainsNone(other Mask) bool {
	if other.IsEmpty() {
		return false
	}
	for i, v := range other {
		if m[i]&v != 0 {
			return false
		}
	}
	return true
}

// IsEmpty checks if the mask has no bits set
func (m Mask) IsEmpty() bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// Unmark clears the specified bit in the mask
func (m *Mask) Unmark(bit uint32) {
	idx := bit / bitSize
	bitPosition := bit % bitSize
	m[idx] &^= (1 << bitPosition) // &^ is AND NOT
}
