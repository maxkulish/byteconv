package byteconv

import (
	"fmt"
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

	if bytes <= 0 {
		return "0"
	}

	var	unit string
	var res float64

	switch {
	case bytes >= EiB:
		unit = "EiB"
		res = bytes / EiB
	case bytes >= PiB:
		unit = "PiB"
		res = bytes / PiB
	case bytes >= TiB:
		unit = "TiB"
		res = bytes / TiB
	case bytes >= GiB:
		unit = "GiB"
		res = bytes / GiB
	case bytes >= MiB:
		unit = "MiB"
		res = bytes / MiB
	case bytes >= KiB:
		unit = "KiB"
		res = bytes / KiB
	case bytes >= BYTE:
		unit = "B"
		res = bytes
	}

	strRes := strconv.FormatFloat(res, 'f', 1, 64)
	strRes = strings.TrimSuffix(strRes, ".0")

	return strRes + unit
}

// ToBytes parses a string formatted by ByteSize as bytes. Note binary-prefixed and SI prefixed units both mean a base-2 units
// KiB	= 1024
// MiB = 1024 * K
// GiB = 1024 * M
// TiB = 1024 * G
// PiB = 1024 * T
// EiB = 1024 * P
func StringBinaryToBytes(s string) float64 {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	i := strings.IndexFunc(s, unicode.IsLetter)

	if i == -1 {
		return 0
	}

	bytesString, multiple := s[:i], s[i:]
	bytes, err := strconv.ParseFloat(bytesString, 64)
	if err != nil || bytes <= 0 {
		return 0
	}

	switch multiple {
	case "EIB":
		return bytes * EiB
	case "PIB":
		return bytes * PiB
	case "TIB":
		return bytes * TiB
	case "GIB":
		return bytes * GiB
	case "MIB":
		return bytes * MiB
	case "KIB":
		return bytes * KiB
	case "B":
		return bytes
	default:
		return 0
	}
}

// BytesSize returns a human-readable in 2 formats: IEC (binary) or SI (decimal)
// Binary string of the form 10MiB, 12.5KiB
// Decimal string of the form 10MB, 12.5KB
// The precision prec controls the number of digits
//The special precision -1 uses the smallest number of digits
func BytesSize(bytes float64, format string, prec int) string {

	if bytes <= 0 {
		return "0"
	}

	// Default format is decimal: MB, GB
	value := 1000.0
	resFormat := ""

	// Binary format: MiB, GiB
	if format == "binary" {
		value = 1024.0
		resFormat = "i"
	}

	if bytes < value {
		strRes := strconv.FormatFloat(bytes, 'f', prec, 64)
		return strings.TrimSuffix(strRes, ".0") + "B"
	}

	divider, exp := value, 0
	for n := bytes / value; n >= value; n /= value {
		divider *= value
		exp++
	}

	strRes := strconv.FormatFloat(bytes/divider, 'f', prec, 64)
	if prec == 0 {
			strRes = strings.TrimSuffix(strRes, ".0")
	}

	return strRes + fmt.Sprintf("%c%sB", "KMGTPE"[exp], resFormat)
}