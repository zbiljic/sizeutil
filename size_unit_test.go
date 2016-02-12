// Copyright 2016 Nemanja Zbiljic
//

package sizeutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BYTES

func TestOneByteInBytes(t *testing.T) {
	assert.Equal(t, int64(1), BYTES.Convert(1, BYTES))
	assert.Equal(t, int64(1), BYTES.ToBytes(1))
}

func TestOneByteInKilobytes(t *testing.T) {
	assert.Equal(t, int64(0), KILOBYTES.Convert(1, BYTES))
	assert.Equal(t, int64(0), BYTES.ToKilobytes(1))
}

func TestOneByteInMegabytes(t *testing.T) {
	assert.Equal(t, int64(0), MEGABYTES.Convert(1, BYTES))
	assert.Equal(t, int64(0), BYTES.ToMegabytes(1))
}

func TestOneByteInGigabytes(t *testing.T) {
	assert.Equal(t, int64(0), GIGABYTES.Convert(1, BYTES))
	assert.Equal(t, int64(0), BYTES.ToGigabytes(1))
}

func TestOneByteInTerabytes(t *testing.T) {
	assert.Equal(t, int64(0), TERABYTES.Convert(1, BYTES))
	assert.Equal(t, int64(0), BYTES.ToTerabytes(1))
}

func TestOneByteInPetabytes(t *testing.T) {
	assert.Equal(t, int64(0), PETABYTES.Convert(1, BYTES))
	assert.Equal(t, int64(0), BYTES.ToPetabytes(1))
}

// KILOBYTES

func TestOneKilobyteInBytes(t *testing.T) {
	assert.Equal(t, int64(1024), BYTES.Convert(1, KILOBYTES))
	assert.Equal(t, int64(1024), KILOBYTES.ToBytes(1))
}

func TestOneKilobyteInKilobytes(t *testing.T) {
	assert.Equal(t, int64(1), KILOBYTES.Convert(1, KILOBYTES))
	assert.Equal(t, int64(1), KILOBYTES.ToKilobytes(1))
}

func TestOneKilobyteInMegabytes(t *testing.T) {
	assert.Equal(t, int64(0), MEGABYTES.Convert(1, KILOBYTES))
	assert.Equal(t, int64(0), KILOBYTES.ToMegabytes(1))
}

func TestOneKilobyteInGigabytes(t *testing.T) {
	assert.Equal(t, int64(0), GIGABYTES.Convert(1, KILOBYTES))
	assert.Equal(t, int64(0), KILOBYTES.ToGigabytes(1))
}

func TestOneKilobyteInTerabytes(t *testing.T) {
	assert.Equal(t, int64(0), TERABYTES.Convert(1, KILOBYTES))
	assert.Equal(t, int64(0), KILOBYTES.ToTerabytes(1))
}

func TestOneKilobyteInPetabytes(t *testing.T) {
	assert.Equal(t, int64(0), PETABYTES.Convert(1, KILOBYTES))
	assert.Equal(t, int64(0), KILOBYTES.ToPetabytes(1))
}

// MEGABYTES

func TestOneMegabyteInBytes(t *testing.T) {
	assert.Equal(t, int64(1048576), BYTES.Convert(1, MEGABYTES))
	assert.Equal(t, int64(1048576), MEGABYTES.ToBytes(1))
}

func TestOneMegabyteInKilobytes(t *testing.T) {
	assert.Equal(t, int64(1024), KILOBYTES.Convert(1, MEGABYTES))
	assert.Equal(t, int64(1024), MEGABYTES.ToKilobytes(1))
}

func TestOneMegabyteInMegabytes(t *testing.T) {
	assert.Equal(t, int64(1), MEGABYTES.Convert(1, MEGABYTES))
	assert.Equal(t, int64(1), MEGABYTES.ToMegabytes(1))
}

func TestOneMegabyteInGigabytes(t *testing.T) {
	assert.Equal(t, int64(0), GIGABYTES.Convert(1, MEGABYTES))
	assert.Equal(t, int64(0), MEGABYTES.ToGigabytes(1))
}

func TestOneMegabyteInTerabytes(t *testing.T) {
	assert.Equal(t, int64(0), TERABYTES.Convert(1, MEGABYTES))
	assert.Equal(t, int64(0), MEGABYTES.ToTerabytes(1))
}

func TestOneMegabyteInPetabytes(t *testing.T) {
	assert.Equal(t, int64(0), PETABYTES.Convert(1, MEGABYTES))
	assert.Equal(t, int64(0), MEGABYTES.ToPetabytes(1))
}

// GIGABYTES

func TestOneGigabyteInBytes(t *testing.T) {
	assert.Equal(t, int64(1073741824), BYTES.Convert(1, GIGABYTES))
	assert.Equal(t, int64(1073741824), GIGABYTES.ToBytes(1))
}

func TestOneGigabyteInKilobytes(t *testing.T) {
	assert.Equal(t, int64(1048576), KILOBYTES.Convert(1, GIGABYTES))
	assert.Equal(t, int64(1048576), GIGABYTES.ToKilobytes(1))
}

func TestOneGigabyteInMegabytes(t *testing.T) {
	assert.Equal(t, int64(1024), MEGABYTES.Convert(1, GIGABYTES))
	assert.Equal(t, int64(1024), GIGABYTES.ToMegabytes(1))
}

func TestOneGigabyteInGigabytes(t *testing.T) {
	assert.Equal(t, int64(1), GIGABYTES.Convert(1, GIGABYTES))
	assert.Equal(t, int64(1), GIGABYTES.ToGigabytes(1))
}

func TestOneGigabyteInTerabytes(t *testing.T) {
	assert.Equal(t, int64(0), TERABYTES.Convert(1, GIGABYTES))
	assert.Equal(t, int64(0), GIGABYTES.ToTerabytes(1))
}

func TestOneGigabyteInPetabytes(t *testing.T) {
	assert.Equal(t, int64(0), PETABYTES.Convert(1, GIGABYTES))
	assert.Equal(t, int64(0), GIGABYTES.ToPetabytes(1))
}

// TERABYTES

func TestOneTerabyteInBytes(t *testing.T) {
	assert.Equal(t, int64(1099511627776), BYTES.Convert(1, TERABYTES))
	assert.Equal(t, int64(1099511627776), TERABYTES.ToBytes(1))
}

func TestOneTerabyteInKilobytes(t *testing.T) {
	assert.Equal(t, int64(1073741824), KILOBYTES.Convert(1, TERABYTES))
	assert.Equal(t, int64(1073741824), TERABYTES.ToKilobytes(1))
}

func TestOneTerabyteInMegabytes(t *testing.T) {
	assert.Equal(t, int64(1048576), MEGABYTES.Convert(1, TERABYTES))
	assert.Equal(t, int64(1048576), TERABYTES.ToMegabytes(1))
}

func TestOneTerabyteInGigabytes(t *testing.T) {
	assert.Equal(t, int64(1024), GIGABYTES.Convert(1, TERABYTES))
	assert.Equal(t, int64(1024), TERABYTES.ToGigabytes(1))
}

func TestOneTerabyteInTerabytes(t *testing.T) {
	assert.Equal(t, int64(1), TERABYTES.Convert(1, TERABYTES))
	assert.Equal(t, int64(1), TERABYTES.ToTerabytes(1))
}

func TestOneTerabyteInPetabytes(t *testing.T) {
	assert.Equal(t, int64(0), PETABYTES.Convert(1, TERABYTES))
	assert.Equal(t, int64(0), TERABYTES.ToPetabytes(1))
}

// PETABYTES

func TestOnePetabyteInBytes(t *testing.T) {
	assert.Equal(t, int64(1125899906842624), BYTES.Convert(1, PETABYTES))
	assert.Equal(t, int64(1125899906842624), PETABYTES.ToBytes(1))
}

func TestOnePetabyteInKilobytes(t *testing.T) {
	assert.Equal(t, int64(1099511627776), KILOBYTES.Convert(1, PETABYTES))
	assert.Equal(t, int64(1099511627776), PETABYTES.ToKilobytes(1))
}

func TestOnePetabyteInMegabytes(t *testing.T) {
	assert.Equal(t, int64(1073741824), MEGABYTES.Convert(1, PETABYTES))
	assert.Equal(t, int64(1073741824), PETABYTES.ToMegabytes(1))
}

func TestOnePetabyteInGigabytes(t *testing.T) {
	assert.Equal(t, int64(1048576), GIGABYTES.Convert(1, PETABYTES))
	assert.Equal(t, int64(1048576), PETABYTES.ToGigabytes(1))
}

func TestOnePetabyteInTerabytes(t *testing.T) {
	assert.Equal(t, int64(1024), TERABYTES.Convert(1, PETABYTES))
	assert.Equal(t, int64(1024), PETABYTES.ToTerabytes(1))
}

func TestOnePetabyteInPetabytes(t *testing.T) {
	assert.Equal(t, int64(1), PETABYTES.Convert(1, PETABYTES))
	assert.Equal(t, int64(1), PETABYTES.ToPetabytes(1))
}
