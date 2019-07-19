package byteconv

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const (
	BYTE = 1 << (10 * iota)
	KiB  // 1024
	MiB  // 1048576
	GiB  // 1073741824
	TiB  // 1099511627776 (exceeds 1 << 32)
	PiB  // 1125899906842624
	EiB  // 1152921504606846976
)

var invalidByteNumber = errors.New(
	"byte quantity must be a positive integer with a unit of measurement like M, MB, MiB, G, GiB, or GB")

// BytesToBinarySize returns a human-readable IEC (binary) format string of the form 10MiB, 12.5KiB.
// History https://en.wikipedia.org/wiki/Mebibyte
// The following units are available:
//	EiB: exbibyte
//	PiB: pebibyte
//	TiB: tebibyte
//	GiB: gibibyte
//	MiB: mebibyte
//	KiB: kibibyte
//	B: Byte
// The unit that results in the smallest number greater than or equal to 1 is always chosen.
func BytesToBinarySize(bytes float64) string {
	unit := ""
	var result float64

	switch {
	case bytes >= EiB:
		unit = "EiB"
		result = bytes / EiB
	case bytes >= PiB:
		unit = "PiB"
		result = bytes / PiB
	case bytes >= TiB:
		unit = "TiB"
		result = bytes / TiB
	case bytes >= GiB:
		unit = "GiB"
		result = bytes / GiB
	case bytes >= MiB:
		unit = "MiB"
		result = bytes / MiB
	case bytes >= KiB:
		unit = "KiB"
		result = bytes / KiB
	case bytes >= BYTE:
		unit = "B"
		result = bytes
	case bytes == 0:
		return "0"
	}

	strRes := strconv.FormatFloat(result, 'f', 1, 64)
	strRes = strings.TrimSuffix(strRes, ".0")

	return strRes + unit
}

// ToMebibyte parses a string formatted by StringBinaryToBytes as mebibytes.
func ToMebibyte(s string) (float64, error) {
	bytes, err := StringBinaryToBytes(s)
	if err != nil {
		return 0, err
	}

	return bytes / MiB, nil
}

// ToBytes parses a string formatted by ByteSize as bytes. Note binary-prefixed and SI prefixed units both mean a base-2 units
// KiB	= 1024
// MiB = 1024 * K
// GiB = 1024 * M
// TiB = 1024 * G
// PiB = 1024 * T
// EiB = 1024 * P
func StringBinaryToBytes(s string) (float64, error) {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	i := strings.IndexFunc(s, unicode.IsLetter)

	if i == -1 {
		return 0, invalidByteNumber
	}

	bytesString, multiple := s[:i], s[i:]
	bytes, err := strconv.ParseFloat(bytesString, 64)
	if err != nil || bytes <= 0 {
		return 0, invalidByteNumber
	}

	switch multiple {
	case "EIB":
		return bytes * EiB, nil
	case "PIB":
		return bytes * PiB, nil
	case "TIB":
		return bytes * TiB, nil
	case "GIB":
		return bytes * GiB, nil
	case "MIB":
		return bytes * MiB, nil
	case "KIB":
		return bytes * KiB, nil
	case "B":
		return bytes, nil
	default:
		return 0, invalidByteNumber
	}
}
