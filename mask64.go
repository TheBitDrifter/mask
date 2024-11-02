//go:build !(m256 || m512 || m1024)

package mask

const (
	MaxBits  = 64
	maskSize = MaxBits / bitSize
)

type (
	Mask [maskSize]uint64
)
