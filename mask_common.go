package mask

const bitSize = 64

func (m *Mask) Mark(bit uint32) {
	idx := bit / bitSize
	bitPosition := bit % bitSize

	m[idx] |= (1 << bitPosition)
}

func (m Mask) Contains(bit uint32) bool {
	idx := bit / bitSize
	offset := bit - (bitSize * idx)
	mask := uint64(1 << offset)
	return m[idx]&mask == mask
}

func (m Mask) ContainsAll(other Mask) bool {
	for i, v := range other {
		if m[i]&v != v {
			return false
		}
	}
	return true
}

func (m Mask) ContainsAny(other Mask) bool {
	for i, v := range other {
		if m[i]&v != 0 {
			return true
		}
	}
	return false
}

func (m Mask) ContainsNone(other Mask) bool {
	if isEmpty(other) {
		return false
	}
	for i, v := range other {
		if m[i]&v != 0 {
			return false
		}
	}
	return true
}

func isEmpty(m Mask) bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
