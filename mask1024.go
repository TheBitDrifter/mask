//go:build m1024

package mask

const (
	MaxBits  = 1024
	maskSize = MaxBits / bitSize
)

type Mask [maskSize]uint64
