// Copyright 2016 Nemanja Zbiljic
//

package sizeutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertsToPetabytes(t *testing.T) {
	assert.Equal(t, int64(2), Petabytes(2).ToPetabytes())
}

func TestConvertsToTerabytes(t *testing.T) {
	assert.Equal(t, int64(2048), Petabytes(2).ToTerabytes())
}

func TestConvertsToGigabytes(t *testing.T) {
	assert.Equal(t, int64(2048), Terabytes(2).ToGigabytes())
}

func TestConvertsToMegabytes(t *testing.T) {
	assert.Equal(t, int64(2048), Gigabytes(2).ToMegabytes())
}

func TestConvertsToKilobytes(t *testing.T) {
	assert.Equal(t, int64(2048), Megabytes(2).ToKilobytes())
}

func TestConvertsToBytes(t *testing.T) {
	assert.Equal(t, int64(2048), Kilobytes(2).ToBytes())
}

func TestParsesPetabytes(t *testing.T) {
	var s *Size
	var err error

	s, err = ParseSize("2PB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Petabytes(2), s)

	s, err = ParseSize("2PiB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Petabytes(2), s)

	s, err = ParseSize("1 petabyte")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Petabytes(1), s)

	s, err = ParseSize("2 petabytes")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Petabytes(2), s)

}

func TestParsesTerabytes(t *testing.T) {
	var s *Size
	var err error

	s, err = ParseSize("2TB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Terabytes(2), s)

	s, err = ParseSize("2TiB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Terabytes(2), s)

	s, err = ParseSize("1 terabyte")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Terabytes(1), s)

	s, err = ParseSize("2 terabytes")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Terabytes(2), s)

}

func TestParsesGigabytes(t *testing.T) {
	var s *Size
	var err error

	s, err = ParseSize("2GB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Gigabytes(2), s)

	s, err = ParseSize("2GiB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Gigabytes(2), s)

	s, err = ParseSize("1 gigabyte")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Gigabytes(1), s)

	s, err = ParseSize("2 gigabytes")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Gigabytes(2), s)

}

func TestParsesMegabytes(t *testing.T) {
	var s *Size
	var err error

	s, err = ParseSize("2MB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Megabytes(2), s)

	s, err = ParseSize("2MiB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Megabytes(2), s)

	s, err = ParseSize("1 megabyte")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Megabytes(1), s)

	s, err = ParseSize("2 megabytes")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Megabytes(2), s)

}

func TestParsesKilobytes(t *testing.T) {
	var s *Size
	var err error

	s, err = ParseSize("2KB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Kilobytes(2), s)

	s, err = ParseSize("2KiB")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Kilobytes(2), s)

	s, err = ParseSize("1 kilobyte")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Kilobytes(1), s)

	s, err = ParseSize("2 kilobytes")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Kilobytes(2), s)

}

func TestParsesBytes(t *testing.T) {
	var s *Size
	var err error

	s, err = ParseSize("2B")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Bytes(2), s)

	s, err = ParseSize("1 byte")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Bytes(1), s)

	s, err = ParseSize("2 bytes")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Bytes(2), s)

}

func TestParseSizeWithWhiteSpaces(t *testing.T) {
	s, err := ParseSize("64   kilobytes")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Kilobytes(64), s)
}

func TestUnableParseWrongSizeCount(t *testing.T) {
	_, err := ParseSize("three bytes")
	assert.Error(t, err)
}

func TestUnableParseWrongSizeUnit(t *testing.T) {
	_, err := ParseSize("1EB")
	assert.Error(t, err)
}

func TestUnableParseWrongSizeFormat(t *testing.T) {
	_, err := ParseSize("1 mega byte")
	assert.Error(t, err)
}

func TestIsHumanReadable(t *testing.T) {
	assert.Equal(t, Gigabytes(3).String(), "3 gigabytes")
	assert.Equal(t, Kilobytes(1).String(), "1 kilobyte")
}

func TestHasAQuantity(t *testing.T) {
	assert.Equal(t, Gigabytes(3).Quantity(), int64(3))
}

func TestHasAUnit(t *testing.T) {
	assert.Equal(t, Gigabytes(3).Unit(), GIGABYTES)
}

func TestCompare(t *testing.T) {
	// both zero
	assert.Equal(t, 0, Bytes(0).Compare(Bytes(0)))
	assert.Equal(t, 0, Bytes(0).Compare(Kilobytes(0)))
	assert.Equal(t, 0, Bytes(0).Compare(Megabytes(0)))
	assert.Equal(t, 0, Bytes(0).Compare(Gigabytes(0)))
	assert.Equal(t, 0, Bytes(0).Compare(Terabytes(0)))
	assert.Equal(t, 0, Bytes(0).Compare(Petabytes(0)))

	assert.Equal(t, 0, Kilobytes(0).Compare(Bytes(0)))
	assert.Equal(t, 0, Kilobytes(0).Compare(Kilobytes(0)))
	assert.Equal(t, 0, Kilobytes(0).Compare(Megabytes(0)))
	assert.Equal(t, 0, Kilobytes(0).Compare(Gigabytes(0)))
	assert.Equal(t, 0, Kilobytes(0).Compare(Terabytes(0)))
	assert.Equal(t, 0, Kilobytes(0).Compare(Petabytes(0)))

	assert.Equal(t, 0, Megabytes(0).Compare(Bytes(0)))
	assert.Equal(t, 0, Megabytes(0).Compare(Kilobytes(0)))
	assert.Equal(t, 0, Megabytes(0).Compare(Megabytes(0)))
	assert.Equal(t, 0, Megabytes(0).Compare(Gigabytes(0)))
	assert.Equal(t, 0, Megabytes(0).Compare(Terabytes(0)))
	assert.Equal(t, 0, Megabytes(0).Compare(Petabytes(0)))

	assert.Equal(t, 0, Gigabytes(0).Compare(Bytes(0)))
	assert.Equal(t, 0, Gigabytes(0).Compare(Kilobytes(0)))
	assert.Equal(t, 0, Gigabytes(0).Compare(Megabytes(0)))
	assert.Equal(t, 0, Gigabytes(0).Compare(Gigabytes(0)))
	assert.Equal(t, 0, Gigabytes(0).Compare(Terabytes(0)))
	assert.Equal(t, 0, Gigabytes(0).Compare(Petabytes(0)))

	assert.Equal(t, 0, Terabytes(0).Compare(Bytes(0)))
	assert.Equal(t, 0, Terabytes(0).Compare(Kilobytes(0)))
	assert.Equal(t, 0, Terabytes(0).Compare(Megabytes(0)))
	assert.Equal(t, 0, Terabytes(0).Compare(Gigabytes(0)))
	assert.Equal(t, 0, Terabytes(0).Compare(Terabytes(0)))
	assert.Equal(t, 0, Terabytes(0).Compare(Petabytes(0)))

	assert.Equal(t, 0, Petabytes(0).Compare(Bytes(0)))
	assert.Equal(t, 0, Petabytes(0).Compare(Kilobytes(0)))
	assert.Equal(t, 0, Petabytes(0).Compare(Megabytes(0)))
	assert.Equal(t, 0, Petabytes(0).Compare(Gigabytes(0)))
	assert.Equal(t, 0, Petabytes(0).Compare(Terabytes(0)))
	assert.Equal(t, 0, Petabytes(0).Compare(Petabytes(0)))

	// one zero, one negative
	assert.Equal(t, 1, Bytes(0).Compare(Bytes(-1)))
	assert.Equal(t, 1, Bytes(0).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Bytes(0).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Bytes(0).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Bytes(0).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Bytes(0).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Kilobytes(0).Compare(Bytes(-1)))
	assert.Equal(t, 1, Kilobytes(0).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Kilobytes(0).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Kilobytes(0).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Kilobytes(0).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Kilobytes(0).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Megabytes(0).Compare(Bytes(-1)))
	assert.Equal(t, 1, Megabytes(0).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Megabytes(0).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Megabytes(0).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Megabytes(0).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Megabytes(0).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Gigabytes(0).Compare(Bytes(-1)))
	assert.Equal(t, 1, Gigabytes(0).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Gigabytes(0).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Gigabytes(0).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Gigabytes(0).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Gigabytes(0).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Terabytes(0).Compare(Bytes(-1)))
	assert.Equal(t, 1, Terabytes(0).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Terabytes(0).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Terabytes(0).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Terabytes(0).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Terabytes(0).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Petabytes(0).Compare(Bytes(-1)))
	assert.Equal(t, 1, Petabytes(0).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Petabytes(0).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Petabytes(0).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Petabytes(0).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Petabytes(0).Compare(Petabytes(-1)))

	assert.Equal(t, -1, Bytes(-1).Compare(Bytes(0)))
	assert.Equal(t, -1, Bytes(-1).Compare(Kilobytes(0)))
	assert.Equal(t, -1, Bytes(-1).Compare(Megabytes(0)))
	assert.Equal(t, -1, Bytes(-1).Compare(Gigabytes(0)))
	assert.Equal(t, -1, Bytes(-1).Compare(Terabytes(0)))
	assert.Equal(t, -1, Bytes(-1).Compare(Petabytes(0)))

	assert.Equal(t, -1, Kilobytes(-1).Compare(Bytes(0)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Kilobytes(0)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Megabytes(0)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Gigabytes(0)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Terabytes(0)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Petabytes(0)))

	assert.Equal(t, -1, Megabytes(-1).Compare(Bytes(0)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Kilobytes(0)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Megabytes(0)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Gigabytes(0)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Terabytes(0)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Petabytes(0)))

	assert.Equal(t, -1, Gigabytes(-1).Compare(Bytes(0)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Kilobytes(0)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Megabytes(0)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Gigabytes(0)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Terabytes(0)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Petabytes(0)))

	assert.Equal(t, -1, Terabytes(-1).Compare(Bytes(0)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Kilobytes(0)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Megabytes(0)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Gigabytes(0)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Terabytes(0)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Petabytes(0)))

	assert.Equal(t, -1, Petabytes(-1).Compare(Bytes(0)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Kilobytes(0)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Megabytes(0)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Gigabytes(0)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Terabytes(0)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Petabytes(0)))

	// one zero, one positive
	assert.Equal(t, -1, Bytes(0).Compare(Bytes(1)))
	assert.Equal(t, -1, Bytes(0).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Bytes(0).Compare(Megabytes(1)))
	assert.Equal(t, -1, Bytes(0).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Bytes(0).Compare(Terabytes(1)))
	assert.Equal(t, -1, Bytes(0).Compare(Petabytes(1)))

	assert.Equal(t, -1, Kilobytes(0).Compare(Bytes(1)))
	assert.Equal(t, -1, Kilobytes(0).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Kilobytes(0).Compare(Megabytes(1)))
	assert.Equal(t, -1, Kilobytes(0).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Kilobytes(0).Compare(Terabytes(1)))
	assert.Equal(t, -1, Kilobytes(0).Compare(Petabytes(1)))

	assert.Equal(t, -1, Megabytes(0).Compare(Bytes(1)))
	assert.Equal(t, -1, Megabytes(0).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Megabytes(0).Compare(Megabytes(1)))
	assert.Equal(t, -1, Megabytes(0).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Megabytes(0).Compare(Terabytes(1)))
	assert.Equal(t, -1, Megabytes(0).Compare(Petabytes(1)))

	assert.Equal(t, -1, Gigabytes(0).Compare(Bytes(1)))
	assert.Equal(t, -1, Gigabytes(0).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Gigabytes(0).Compare(Megabytes(1)))
	assert.Equal(t, -1, Gigabytes(0).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Gigabytes(0).Compare(Terabytes(1)))
	assert.Equal(t, -1, Gigabytes(0).Compare(Petabytes(1)))

	assert.Equal(t, -1, Terabytes(0).Compare(Bytes(1)))
	assert.Equal(t, -1, Terabytes(0).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Terabytes(0).Compare(Megabytes(1)))
	assert.Equal(t, -1, Terabytes(0).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Terabytes(0).Compare(Terabytes(1)))
	assert.Equal(t, -1, Terabytes(0).Compare(Petabytes(1)))

	assert.Equal(t, -1, Petabytes(0).Compare(Bytes(1)))
	assert.Equal(t, -1, Petabytes(0).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Petabytes(0).Compare(Megabytes(1)))
	assert.Equal(t, -1, Petabytes(0).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Petabytes(0).Compare(Terabytes(1)))
	assert.Equal(t, -1, Petabytes(0).Compare(Petabytes(1)))

	assert.Equal(t, 1, Bytes(1).Compare(Bytes(0)))
	assert.Equal(t, 1, Bytes(1).Compare(Kilobytes(0)))
	assert.Equal(t, 1, Bytes(1).Compare(Megabytes(0)))
	assert.Equal(t, 1, Bytes(1).Compare(Gigabytes(0)))
	assert.Equal(t, 1, Bytes(1).Compare(Terabytes(0)))
	assert.Equal(t, 1, Bytes(1).Compare(Petabytes(0)))

	assert.Equal(t, 1, Kilobytes(1).Compare(Bytes(0)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Kilobytes(0)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Megabytes(0)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Gigabytes(0)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Terabytes(0)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Petabytes(0)))

	assert.Equal(t, 1, Megabytes(1).Compare(Bytes(0)))
	assert.Equal(t, 1, Megabytes(1).Compare(Kilobytes(0)))
	assert.Equal(t, 1, Megabytes(1).Compare(Megabytes(0)))
	assert.Equal(t, 1, Megabytes(1).Compare(Gigabytes(0)))
	assert.Equal(t, 1, Megabytes(1).Compare(Terabytes(0)))
	assert.Equal(t, 1, Megabytes(1).Compare(Petabytes(0)))

	assert.Equal(t, 1, Gigabytes(1).Compare(Bytes(0)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Kilobytes(0)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Megabytes(0)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Gigabytes(0)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Terabytes(0)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Petabytes(0)))

	assert.Equal(t, 1, Terabytes(1).Compare(Bytes(0)))
	assert.Equal(t, 1, Terabytes(1).Compare(Kilobytes(0)))
	assert.Equal(t, 1, Terabytes(1).Compare(Megabytes(0)))
	assert.Equal(t, 1, Terabytes(1).Compare(Gigabytes(0)))
	assert.Equal(t, 1, Terabytes(1).Compare(Terabytes(0)))
	assert.Equal(t, 1, Terabytes(1).Compare(Petabytes(0)))

	assert.Equal(t, 1, Petabytes(1).Compare(Bytes(0)))
	assert.Equal(t, 1, Petabytes(1).Compare(Kilobytes(0)))
	assert.Equal(t, 1, Petabytes(1).Compare(Megabytes(0)))
	assert.Equal(t, 1, Petabytes(1).Compare(Gigabytes(0)))
	assert.Equal(t, 1, Petabytes(1).Compare(Terabytes(0)))
	assert.Equal(t, 1, Petabytes(1).Compare(Petabytes(0)))

	// both negative
	assert.Equal(t, -1, Bytes(-2).Compare(Bytes(-1)))
	assert.Equal(t, 1, Bytes(-2).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Bytes(-2).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Bytes(-2).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Bytes(-2).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Bytes(-2).Compare(Petabytes(-1)))

	assert.Equal(t, -1, Kilobytes(-2).Compare(Bytes(-1)))
	assert.Equal(t, -1, Kilobytes(-2).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Kilobytes(-2).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Kilobytes(-2).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Kilobytes(-2).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Kilobytes(-2).Compare(Petabytes(-1)))

	assert.Equal(t, -1, Megabytes(-2).Compare(Bytes(-1)))
	assert.Equal(t, -1, Megabytes(-2).Compare(Kilobytes(-1)))
	assert.Equal(t, -1, Megabytes(-2).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Megabytes(-2).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Megabytes(-2).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Megabytes(-2).Compare(Petabytes(-1)))

	assert.Equal(t, -1, Gigabytes(-2).Compare(Bytes(-1)))
	assert.Equal(t, -1, Gigabytes(-2).Compare(Kilobytes(-1)))
	assert.Equal(t, -1, Gigabytes(-2).Compare(Megabytes(-1)))
	assert.Equal(t, -1, Gigabytes(-2).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Gigabytes(-2).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Gigabytes(-2).Compare(Petabytes(-1)))

	assert.Equal(t, -1, Terabytes(-2).Compare(Bytes(-1)))
	assert.Equal(t, -1, Terabytes(-2).Compare(Kilobytes(-1)))
	assert.Equal(t, -1, Terabytes(-2).Compare(Megabytes(-1)))
	assert.Equal(t, -1, Terabytes(-2).Compare(Gigabytes(-1)))
	assert.Equal(t, -1, Terabytes(-2).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Terabytes(-2).Compare(Petabytes(-1)))

	assert.Equal(t, -1, Petabytes(-2).Compare(Bytes(-1)))
	assert.Equal(t, -1, Petabytes(-2).Compare(Kilobytes(-1)))
	assert.Equal(t, -1, Petabytes(-2).Compare(Megabytes(-1)))
	assert.Equal(t, -1, Petabytes(-2).Compare(Gigabytes(-1)))
	assert.Equal(t, -1, Petabytes(-2).Compare(Terabytes(-1)))
	assert.Equal(t, -1, Petabytes(-2).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Bytes(-1).Compare(Bytes(-2)))
	assert.Equal(t, 1, Bytes(-1).Compare(Kilobytes(-2)))
	assert.Equal(t, 1, Bytes(-1).Compare(Megabytes(-2)))
	assert.Equal(t, 1, Bytes(-1).Compare(Gigabytes(-2)))
	assert.Equal(t, 1, Bytes(-1).Compare(Terabytes(-2)))
	assert.Equal(t, 1, Bytes(-1).Compare(Petabytes(-2)))

	assert.Equal(t, -1, Kilobytes(-1).Compare(Bytes(-2)))
	assert.Equal(t, 1, Kilobytes(-1).Compare(Kilobytes(-2)))
	assert.Equal(t, 1, Kilobytes(-1).Compare(Megabytes(-2)))
	assert.Equal(t, 1, Kilobytes(-1).Compare(Gigabytes(-2)))
	assert.Equal(t, 1, Kilobytes(-1).Compare(Terabytes(-2)))
	assert.Equal(t, 1, Kilobytes(-1).Compare(Petabytes(-2)))

	assert.Equal(t, -1, Megabytes(-1).Compare(Bytes(-2)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Kilobytes(-2)))
	assert.Equal(t, 1, Megabytes(-1).Compare(Megabytes(-2)))
	assert.Equal(t, 1, Megabytes(-1).Compare(Gigabytes(-2)))
	assert.Equal(t, 1, Megabytes(-1).Compare(Terabytes(-2)))
	assert.Equal(t, 1, Megabytes(-1).Compare(Petabytes(-2)))

	assert.Equal(t, -1, Gigabytes(-1).Compare(Bytes(-2)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Kilobytes(-2)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Megabytes(-2)))
	assert.Equal(t, 1, Gigabytes(-1).Compare(Gigabytes(-2)))
	assert.Equal(t, 1, Gigabytes(-1).Compare(Terabytes(-2)))
	assert.Equal(t, 1, Gigabytes(-1).Compare(Petabytes(-2)))

	assert.Equal(t, -1, Terabytes(-1).Compare(Bytes(-2)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Kilobytes(-2)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Megabytes(-2)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Gigabytes(-2)))
	assert.Equal(t, 1, Terabytes(-1).Compare(Terabytes(-2)))
	assert.Equal(t, 1, Terabytes(-1).Compare(Petabytes(-2)))

	assert.Equal(t, -1, Petabytes(-1).Compare(Bytes(-2)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Kilobytes(-2)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Megabytes(-2)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Gigabytes(-2)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Terabytes(-2)))
	assert.Equal(t, 1, Petabytes(-1).Compare(Petabytes(-2)))

	// both positive
	assert.Equal(t, -1, Bytes(1).Compare(Bytes(2)))
	assert.Equal(t, -1, Bytes(1).Compare(Kilobytes(2)))
	assert.Equal(t, -1, Bytes(1).Compare(Megabytes(2)))
	assert.Equal(t, -1, Bytes(1).Compare(Gigabytes(2)))
	assert.Equal(t, -1, Bytes(1).Compare(Terabytes(2)))
	assert.Equal(t, -1, Bytes(1).Compare(Petabytes(2)))

	assert.Equal(t, 1, Kilobytes(1).Compare(Bytes(2)))
	assert.Equal(t, -1, Kilobytes(1).Compare(Kilobytes(2)))
	assert.Equal(t, -1, Kilobytes(1).Compare(Megabytes(2)))
	assert.Equal(t, -1, Kilobytes(1).Compare(Gigabytes(2)))
	assert.Equal(t, -1, Kilobytes(1).Compare(Terabytes(2)))
	assert.Equal(t, -1, Kilobytes(1).Compare(Petabytes(2)))

	assert.Equal(t, 1, Megabytes(1).Compare(Bytes(2)))
	assert.Equal(t, 1, Megabytes(1).Compare(Kilobytes(2)))
	assert.Equal(t, -1, Megabytes(1).Compare(Megabytes(2)))
	assert.Equal(t, -1, Megabytes(1).Compare(Gigabytes(2)))
	assert.Equal(t, -1, Megabytes(1).Compare(Terabytes(2)))
	assert.Equal(t, -1, Megabytes(1).Compare(Petabytes(2)))

	assert.Equal(t, 1, Gigabytes(1).Compare(Bytes(2)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Kilobytes(2)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Megabytes(2)))
	assert.Equal(t, -1, Gigabytes(1).Compare(Gigabytes(2)))
	assert.Equal(t, -1, Gigabytes(1).Compare(Terabytes(2)))
	assert.Equal(t, -1, Gigabytes(1).Compare(Petabytes(2)))

	assert.Equal(t, 1, Terabytes(1).Compare(Bytes(2)))
	assert.Equal(t, 1, Terabytes(1).Compare(Kilobytes(2)))
	assert.Equal(t, 1, Terabytes(1).Compare(Megabytes(2)))
	assert.Equal(t, 1, Terabytes(1).Compare(Gigabytes(2)))
	assert.Equal(t, -1, Terabytes(1).Compare(Terabytes(2)))
	assert.Equal(t, -1, Terabytes(1).Compare(Petabytes(2)))

	assert.Equal(t, 1, Petabytes(1).Compare(Bytes(2)))
	assert.Equal(t, 1, Petabytes(1).Compare(Kilobytes(2)))
	assert.Equal(t, 1, Petabytes(1).Compare(Megabytes(2)))
	assert.Equal(t, 1, Petabytes(1).Compare(Gigabytes(2)))
	assert.Equal(t, 1, Petabytes(1).Compare(Terabytes(2)))
	assert.Equal(t, -1, Petabytes(1).Compare(Petabytes(2)))

	assert.Equal(t, 1, Bytes(2).Compare(Bytes(1)))
	assert.Equal(t, -1, Bytes(2).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Bytes(2).Compare(Megabytes(1)))
	assert.Equal(t, -1, Bytes(2).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Bytes(2).Compare(Terabytes(1)))
	assert.Equal(t, -1, Bytes(2).Compare(Petabytes(1)))

	assert.Equal(t, 1, Kilobytes(2).Compare(Bytes(1)))
	assert.Equal(t, 1, Kilobytes(2).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Kilobytes(2).Compare(Megabytes(1)))
	assert.Equal(t, -1, Kilobytes(2).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Kilobytes(2).Compare(Terabytes(1)))
	assert.Equal(t, -1, Kilobytes(2).Compare(Petabytes(1)))

	assert.Equal(t, 1, Megabytes(2).Compare(Bytes(1)))
	assert.Equal(t, 1, Megabytes(2).Compare(Kilobytes(1)))
	assert.Equal(t, 1, Megabytes(2).Compare(Megabytes(1)))
	assert.Equal(t, -1, Megabytes(2).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Megabytes(2).Compare(Terabytes(1)))
	assert.Equal(t, -1, Megabytes(2).Compare(Petabytes(1)))

	assert.Equal(t, 1, Gigabytes(2).Compare(Bytes(1)))
	assert.Equal(t, 1, Gigabytes(2).Compare(Kilobytes(1)))
	assert.Equal(t, 1, Gigabytes(2).Compare(Megabytes(1)))
	assert.Equal(t, 1, Gigabytes(2).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Gigabytes(2).Compare(Terabytes(1)))
	assert.Equal(t, -1, Gigabytes(2).Compare(Petabytes(1)))

	assert.Equal(t, 1, Terabytes(2).Compare(Bytes(1)))
	assert.Equal(t, 1, Terabytes(2).Compare(Kilobytes(1)))
	assert.Equal(t, 1, Terabytes(2).Compare(Megabytes(1)))
	assert.Equal(t, 1, Terabytes(2).Compare(Gigabytes(1)))
	assert.Equal(t, 1, Terabytes(2).Compare(Terabytes(1)))
	assert.Equal(t, -1, Terabytes(2).Compare(Petabytes(1)))

	assert.Equal(t, 1, Petabytes(2).Compare(Bytes(1)))
	assert.Equal(t, 1, Petabytes(2).Compare(Kilobytes(1)))
	assert.Equal(t, 1, Petabytes(2).Compare(Megabytes(1)))
	assert.Equal(t, 1, Petabytes(2).Compare(Gigabytes(1)))
	assert.Equal(t, 1, Petabytes(2).Compare(Terabytes(1)))
	assert.Equal(t, 1, Petabytes(2).Compare(Petabytes(1)))

	// one negative, one positive
	assert.Equal(t, -1, Bytes(-1).Compare(Bytes(1)))
	assert.Equal(t, -1, Bytes(-1).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Bytes(-1).Compare(Megabytes(1)))
	assert.Equal(t, -1, Bytes(-1).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Bytes(-1).Compare(Terabytes(1)))
	assert.Equal(t, -1, Bytes(-1).Compare(Petabytes(1)))

	assert.Equal(t, -1, Kilobytes(-1).Compare(Bytes(1)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Megabytes(1)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Terabytes(1)))
	assert.Equal(t, -1, Kilobytes(-1).Compare(Petabytes(1)))

	assert.Equal(t, -1, Megabytes(-1).Compare(Bytes(1)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Megabytes(1)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Terabytes(1)))
	assert.Equal(t, -1, Megabytes(-1).Compare(Petabytes(1)))

	assert.Equal(t, -1, Gigabytes(-1).Compare(Bytes(1)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Megabytes(1)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Terabytes(1)))
	assert.Equal(t, -1, Gigabytes(-1).Compare(Petabytes(1)))

	assert.Equal(t, -1, Terabytes(-1).Compare(Bytes(1)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Megabytes(1)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Terabytes(1)))
	assert.Equal(t, -1, Terabytes(-1).Compare(Petabytes(1)))

	assert.Equal(t, -1, Petabytes(-1).Compare(Bytes(1)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Kilobytes(1)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Megabytes(1)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Gigabytes(1)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Terabytes(1)))
	assert.Equal(t, -1, Petabytes(-1).Compare(Petabytes(1)))

	assert.Equal(t, 1, Bytes(1).Compare(Bytes(-1)))
	assert.Equal(t, 1, Bytes(1).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Bytes(1).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Bytes(1).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Bytes(1).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Bytes(1).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Kilobytes(1).Compare(Bytes(-1)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Kilobytes(1).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Megabytes(1).Compare(Bytes(-1)))
	assert.Equal(t, 1, Megabytes(1).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Megabytes(1).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Megabytes(1).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Megabytes(1).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Megabytes(1).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Gigabytes(1).Compare(Bytes(-1)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Gigabytes(1).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Terabytes(1).Compare(Bytes(-1)))
	assert.Equal(t, 1, Terabytes(1).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Terabytes(1).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Terabytes(1).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Terabytes(1).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Terabytes(1).Compare(Petabytes(-1)))

	assert.Equal(t, 1, Petabytes(1).Compare(Bytes(-1)))
	assert.Equal(t, 1, Petabytes(1).Compare(Kilobytes(-1)))
	assert.Equal(t, 1, Petabytes(1).Compare(Megabytes(-1)))
	assert.Equal(t, 1, Petabytes(1).Compare(Gigabytes(-1)))
	assert.Equal(t, 1, Petabytes(1).Compare(Terabytes(-1)))
	assert.Equal(t, 1, Petabytes(1).Compare(Petabytes(-1)))

}
