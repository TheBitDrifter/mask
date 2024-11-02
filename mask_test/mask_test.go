package mask_test

import (
	"testing"

	"github.com/TheBitDrifter/mask"
)

func TestMask(t *testing.T) {
	var maxBit uint32 = mask.MaxBits - 1

	tests := []struct {
		name      string
		initial   mask.Mask
		toMark    []uint32
		checkAll  mask.Mask
		checkAny  mask.Mask
		checkNone mask.Mask
		wantAll   bool
		wantAny   bool
		wantNone  bool
		getBits   map[uint32]bool
	}{
		{
			name:    "Mark lowest bit (0) and check",
			initial: mask.Mask{},
			toMark:  []uint32{0},
			checkAll: func() mask.Mask {
				var m mask.Mask
				m.Mark(0)
				return m
			}(),
			checkAny: func() mask.Mask {
				var m mask.Mask
				m.Mark(0)
				return m
			}(),
			checkNone: mask.Mask{},
			wantAll:   true,
			wantAny:   true,
			wantNone:  false,
			getBits: map[uint32]bool{
				0: true,
				1: false,
			},
		},
		{
			name:    "Mark middle and last bit",
			initial: mask.Mask{},
			toMark:  []uint32{1, maxBit},
			checkAll: func() mask.Mask {
				var m mask.Mask
				m.Mark(1)
				m.Mark(maxBit)
				return m
			}(),
			checkAny: func() mask.Mask {
				var m mask.Mask
				m.Mark(maxBit)
				return m
			}(),
			checkNone: mask.Mask{},
			wantAll:   true,
			wantAny:   true,
			wantNone:  false,
			getBits: map[uint32]bool{
				1:      true,
				2:      false,
				maxBit: true,
				32:     false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.initial
			for _, bit := range tt.toMark {
				m.Mark(bit)
			}
			if got := m.ContainsAll(tt.checkAll); got != tt.wantAll {
				t.Errorf("Mask.IncludesAll() = %v, expected %v", got, tt.wantAll)
			}
			if got := m.ContainsAny(tt.checkAny); got != tt.wantAny {
				t.Errorf("Mask.IncludesAny() = %v, expected %v", got, tt.wantAny)
			}
			if got := m.ContainsNone(tt.checkNone); got != tt.wantNone {
				t.Errorf("Mask.IncludesNone() = %v, expected %v", got, tt.wantNone)
			}
			for bit, want := range tt.getBits {
				if got := m.Contains(bit); got != want {
					t.Errorf("Mask.Marked(%d) = %v, expected %v", bit, got, want)
				}
			}
		})
	}
}
