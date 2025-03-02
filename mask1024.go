//go:build m1024

package mask

// MaxBits defines the maximum number of bits supported by this mask implementation
const (
	MaxBits  = 1024
	maskSize = MaxBits / bitSize
)

type Mask [maskSize]uint64
