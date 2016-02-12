// Copyright 2016 Nemanja Zbiljic
//

package sizeutil

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var sizePattern *regexp.Regexp = regexp.MustCompile(`(\d+)\s*(\S+)`)

var suffixes = map[string]SizeUnit{
	"B":         BYTES,
	"byte":      BYTES,
	"bytes":     BYTES,
	"KB":        KILOBYTES,
	"KiB":       KILOBYTES,
	"kilobyte":  KILOBYTES,
	"kilobytes": KILOBYTES,
	"MB":        MEGABYTES,
	"MiB":       MEGABYTES,
	"megabyte":  MEGABYTES,
	"megabytes": MEGABYTES,
	"GB":        GIGABYTES,
	"GiB":       GIGABYTES,
	"gigabyte":  GIGABYTES,
	"gigabytes": GIGABYTES,
	"TB":        TERABYTES,
	"TiB":       TERABYTES,
	"terabyte":  TERABYTES,
	"terabytes": TERABYTES,
	"PB":        PETABYTES,
	"PiB":       PETABYTES,
	"petabyte":  PETABYTES,
	"petabytes": PETABYTES,
}

type Size struct {
	count int64
	unit  SizeUnit
}

func Bytes(count int64) *Size {
	return newSize(count, BYTES)
}

func Kilobytes(count int64) *Size {
	return newSize(count, KILOBYTES)
}

func Megabytes(count int64) *Size {
	return newSize(count, MEGABYTES)
}

func Gigabytes(count int64) *Size {
	return newSize(count, GIGABYTES)
}

func Terabytes(count int64) *Size {
	return newSize(count, TERABYTES)
}

func Petabytes(count int64) *Size {
	return newSize(count, PETABYTES)
}

func ParseSize(size string) (*Size, error) {
	parts := sizePattern.FindStringSubmatch(strings.TrimSpace(size))

	if len(parts) < 3 {
		return nil, errors.New("Byte quantity must be a positive integer with a unit of measurement.")
	}

	count, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil || count < 1 {
		return nil, errors.New("Byte quantity must be a positive integer")
	}

	unit, ok := suffixes[parts[2]]
	if !ok {
		return nil, errors.New("Invalid size: " + size + ". Wrong size unit")
	}

	return newSize(count, unit), nil
}

func newSize(count int64, unit SizeUnit) *Size {
	return &Size{count: count, unit: unit}
}

func (s *Size) Quantity() int64 {
	return s.count
}

func (s *Size) Unit() SizeUnit {
	return s.unit
}

func (s *Size) ToBytes() int64 {
	return BYTES.Convert(s.count, s.unit)
}

func (s *Size) ToKilobytes() int64 {
	return KILOBYTES.Convert(s.count, s.unit)
}

func (s *Size) ToMegabytes() int64 {
	return MEGABYTES.Convert(s.count, s.unit)
}

func (s *Size) ToGigabytes() int64 {
	return GIGABYTES.Convert(s.count, s.unit)
}

func (s *Size) ToTerabytes() int64 {
	return TERABYTES.Convert(s.count, s.unit)
}

func (s *Size) ToPetabytes() int64 {
	return PETABYTES.Convert(s.count, s.unit)
}

func (s *Size) String() string {
	units := strings.ToLower(s.unit.String())
	if s.count == 1 {
		units = units[:len(units)-1]
	}
	return fmt.Sprintf("%d %s", s.count, units)
}

// Compare compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func (x *Size) Compare(y *Size) int {
	if x.unit == y.unit {
		return compare(x.count, y.count)
	}
	return compare(x.ToBytes(), y.ToBytes())
}

func compare(x, y int64) int {
	switch {
	case x > y:
		return 1
	case x < y:
		return -1
	default:
		return 0
	}
}
