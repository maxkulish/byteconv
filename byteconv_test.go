package byteconv

import (
	"testing"
)

func TestBytesToBinarySize(t *testing.T) {
	f := func(b float64, resExp string) {
		t.Helper()
		res := BytesToBinarySize(b)

		if res != resExp {
			t.Fatalf("unexpected result for BytesToBinarySize(%v); got %s; want %s", b, res, resExp)
		}
	}

	f(999, "999B")
	f(1024, "1KiB")
	f(1048576, "1MiB")
	f(1073741824, "1GiB")
	f(1099511627776, "1TiB")
	f(1125899906842624, "1PiB")
	f(1152921504606846976, "1EiB")
}

func TestStringBinaryToBytes(t *testing.T) {
	f := func(s string, resExp float64) {
		t.Helper()
		res, err := StringBinaryToBytes(s)
		if err != nil {
			t.Fatalf("unexpected error in StringBinaryToBytes(%s)", s)
		}

		if res != resExp {
			t.Fatalf("unexpected result for StringBinaryToBytes(%s); got %f; want %f", s, res, resExp)
		}
	}

	f("999B", 999)
	f("1KiB", 1024)
	f("1MiB", 1048576)
	f("1GiB", 1073741824)
	f("1TiB", 1099511627776)
	f("1PiB", 1125899906842624)
	f("1EiB", 1152921504606846976)
}

