package mask

// Mask256 represents a 256-element bit mask using an array of uint64
type (
	Mask256 [256]uint64
)

// Mark sets the specified bit in the mask
func (m *Mask256) Mark(bit uint32) {
	idx := bit / bitSize
	bitPosition := bit % bitSize
	m[idx] |= (1 << bitPosition)
}

// Contains checks if the specified bit is set in the mask
func (m Mask256) Contains(bit uint32) bool {
	idx := bit / bitSize
	offset := bit - (bitSize * idx)
	mask := uint64(1 << offset)
	return m[idx]&mask == mask
}

// ContainsAll returns true if this mask contains all bits set in the other mask
func (m Mask256) ContainsAll(other Mask256) bool {
	for i, v := range other {
		if m[i]&v != v {
			return false
		}
	}
	return true
}

// ContainsAny returns true if this mask contains any bits set in the other mask
func (m Mask256) ContainsAny(other Mask256) bool {
	for i, v := range other {
		if m[i]&v != 0 {
			return true
		}
	}
	return false
}

// ContainsNone returns true if this mask contains none of the bits set in the other mask
func (m Mask256) ContainsNone(other Mask256) bool {
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

// IsEmpty returns true if no bits are set in this mask
func (m Mask256) IsEmpty() bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// Unmark clears the specified bit in the mask
func (m *Mask256) Unmark(bit uint32) {
	idx := bit / bitSize
	bitPosition := bit % bitSize
	m[idx] &^= (1 << bitPosition) // &^ is AND NOT
}
