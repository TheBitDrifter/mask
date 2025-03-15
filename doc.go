/*
Package mask provides utilities for managing bitmasks in Go, supporting various mask sizes
through build configurations.

The mask package implements functionality to manipulate and compare bitmasks,
with support for configurable sizes (64, 256, 512, or 1024 bits) based on build tags.
A fixed-size Mask256 is also available regardless of build configuration.

# Mask Types

The package defines two primary types:
  - Mask: A variable-sized bitmask whose size depends on build tags
  - Mask256: A fixed 256-bit bitmask available in all configurations

# Basic Usage

Creating and manipulating masks:

	// Create a new mask
	var m mask.Mask

	// Set bits in the mask
	m.Mark(3)    // Set bit 3
	m.Mark(42)   // Set bit 42

	// Check if a bit is set
	if m.Contains(3) {
		// Bit 3 is set
	}

	// Clear a bit
	m.Unmark(3)

# Comparing Masks

Masks can be compared for various conditions:

	var m1, m2 mask.Mask
	m1.Mark(1)
	m1.Mark(2)

	m2.Mark(2)
	m2.Mark(3)

	// Check if m1 contains all bits in m2
	if m1.ContainsAll(m2) {
		// m1 has all bits set that are set in m2
	}

	// Check if m1 contains any bits in m2
	if m1.ContainsAny(m2) {
		// m1 and m2 share at least one bit
	}

	// Check if m1 contains none of the bits in m2
	if m1.ContainsNone(m2) {
		// m1 and m2 have no bits in common
	}

	// Check if a mask is empty (no bits set)
	if m1.IsEmpty() {
		// No bits are set in m1
	}

# Build Configurations

The package supports different bit sizes through build tags:

  - Default: 64-bit mask (no build tag required)
  - m256: 256-bit mask
  - m512: 512-bit mask
  - m1024: 1024-bit mask

Example of building with a specific configuration:

	go build -tags m256

The Mask256 type is always available with a fixed 256-bit size, regardless of
which build configuration is used.
*/
package mask
