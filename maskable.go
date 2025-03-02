package mask

// Maskable represents any type that can be converted to a Mask
type Maskable interface {
	Mask() Mask
}
