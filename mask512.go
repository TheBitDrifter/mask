//go:build m512

package mask

// MaxBits defines the maximum number of bits supported by this mask implementation
const (
	MaxBits  = 512
	maskSize = MaxBits / bitSize
)

type Mask [maskSize]uint64
