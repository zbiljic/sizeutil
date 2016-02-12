// Copyright 2016 Nemanja Zbiljic
//

package sizeutil

// A unit of size.
type SizeUnit int64

const (
	BYTES     SizeUnit = 8.0
	KILOBYTES SizeUnit = 1024 * BYTES
	MEGABYTES SizeUnit = 1024 * KILOBYTES
	GIGABYTES SizeUnit = 1024 * MEGABYTES
	TERABYTES SizeUnit = 1024 * GIGABYTES
	PETABYTES SizeUnit = 1024 * TERABYTES
)

// Converts a size of the given unit into the current unit.
func (su SizeUnit) Convert(size int64, unit SizeUnit) int64 {
	return (size * int64(unit)) / int64(su)
}

// Converts the given number of the current units into bytes.
func (su SizeUnit) ToBytes(size int64) int64 {
	return BYTES.Convert(size, su)
}

// Converts the given number of the current units into kilobytes.
func (su SizeUnit) ToKilobytes(size int64) int64 {
	return KILOBYTES.Convert(size, su)
}

// Converts the given number of the current units into megabytes.
func (su SizeUnit) ToMegabytes(size int64) int64 {
	return MEGABYTES.Convert(size, su)
}

// Converts the given number of the current units into gigabytes.
func (su SizeUnit) ToGigabytes(size int64) int64 {
	return GIGABYTES.Convert(size, su)
}

// Converts the given number of the current units into terabytes.
func (su SizeUnit) ToTerabytes(size int64) int64 {
	return TERABYTES.Convert(size, su)
}

// Converts the given number of the current units into petabytes.
func (su SizeUnit) ToPetabytes(size int64) int64 {
	return PETABYTES.Convert(size, su)
}

func (su SizeUnit) String() string {
	switch su {
	case BYTES:
		return "BYTES"
	case KILOBYTES:
		return "KILOBYTES"
	case MEGABYTES:
		return "MEGABYTES"
	case GIGABYTES:
		return "GIGABYTES"
	case TERABYTES:
		return "TERABYTES"
	case PETABYTES:
		return "PETABYTES"
	default:
		return "BYTES"
	}
}
