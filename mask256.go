//go:build m256

package mask

// MaxBits defines the maximum number of bits supported by this mask implementation
const (
	MaxBits  = 256
	maskSize = MaxBits / bitSize
)

type Mask [maskSize]uint64
