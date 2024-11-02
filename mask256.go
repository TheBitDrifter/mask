//go:build m256

package mask

const (
	MaxBits  = 256
	maskSize = MaxBits / bitSize
)

type Mask [maskSize]uint64
