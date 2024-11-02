//go:build m512

package mask

const (
	MaxBits  = 512
	maskSize = MaxBits / bitSize
)

type Mask [maskSize]uint64
