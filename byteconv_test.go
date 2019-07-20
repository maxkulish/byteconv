package byteconv

import (
	"fmt"
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

	f(0, "0")
	f(-456, "0")
	f(999, "999B")
	f(1024, "1KiB")
	f(1048576, "1MiB")
	f(5048576, "4.8MiB")
	f(1073741824, "1GiB")
	f(1099511627776, "1TiB")
	f(1125899906842624, "1PiB")
	f(1152921504606846976, "1EiB")
}

func BenchmarkBytesToBinarySize(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		BytesToBinarySize(1152921504606846976)
	}
}

func ExampleBytesToBinarySize() {
	fmt.Println(BytesToBinarySize(1073741824))
	fmt.Println(BytesToBinarySize(1152921504606846976))
	fmt.Println(BytesToBinarySize(5048576))
	// Output:
	// 1GiB
	// 1EiB
	// 4.8MiB
}

func TestStringBinaryToBytes(t *testing.T) {
	f := func(s string, resExp float64) {
		t.Helper()
		res:= StringBinaryToBytes(s)

		if res != resExp {
			t.Fatalf("unexpected result for StringBinaryToBytes(%v); got %f; want %f", s, res, resExp)
		}
	}

	f("0", 0)
	f("-456", 0)
	f("999B", 999)
	f("1KiB", 1024)
	f("1MiB", 1048576)
	f("003MiB", 3145728)
	f("-3MiB", 0)
	f("1GiB", 1073741824)
	f("1TiB", 1099511627776)
	f("1PiB", 1125899906842624)
	f("1EiB", 1152921504606846976)
	f("\n\n\r\t10.18TiB\n\n\r\t", 11193028370759.679688)
}

func BenchmarkStringBinaryToBytes(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++  {
		StringBinaryToBytes("\n\n\r\t10.18TiB\n\n\r\t")
	}
}

func ExampleStringBinaryToBytes() {
	fmt.Println(StringBinaryToBytes("55.6MiB"))
	fmt.Println(StringBinaryToBytes("12TiB"))
	// Output:
	// 5.83008256e+07
	// 1.3194139533312e+13
}

func TestBytesToDecimalSize(t *testing.T) {
	f := func(b float64, format string, prec int, resExp string) {
		t.Helper()
		res := BytesSize(b, format, prec)

		if res != resExp {
			t.Errorf("unexpected result for BytesToBinarySize(%v); got %s; want %s", b, res, resExp)
		}
	}

	f(0, "decimal", 1, "0")
	f(-456, "decimal", 1, "0")
	f(999,  "decimal", -1, "999B")
	f(1000, "decimal", 0, "1KB")
	f(1058576, "decimal", 2, "1.06MB")
	f(1058576, "decimal", -1, "1.058576MB")
	f(5000000, "decimal", 0, "5MB")
	f(1000000000, "decimal", 0, "1GB")
	f(1000000000000, "decimal", 1, "1.0TB")
	f(1125899906842624, "decimal", 0, "1PB")
	f(1152921504606846976, "decimal", 6,"1.152922EB")
}

func BenchmarkBytesToDecimalSize(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		BytesSize(1152921504606846976, "decimal", 1)
	}
}

func TestBytesToBinarySize2(t *testing.T) {
		f := func(b float64, format string, prec int, resExp string) {
		t.Helper()
		res := BytesSize(b, format, prec)

		if res != resExp {
			t.Errorf("unexpected result for BytesToBinarySize(%v); got %s; want %s", b, res, resExp)
		}
	}

	f(0, "binary", 1, "0")
	f(-456, "binary", 1, "0")
	f(999,  "binary", -1, "999B")
	f(1024, "binary", 0, "1KiB")
	f(1048576, "binary", 2, "1.00MiB")
	f(1058576, "binary", -1, "1.0095367431640625MiB")
	f(5048576, "binary", 0, "5MiB")
	f(1073741824, "binary", 0, "1GiB")
	f(1099511627776, "binary", 0, "1TiB")
	f(1125899906842624, "binary", 4, "1.0000PiB")
	f(1152921504606846976, "binary", 6,"1.000000EiB")
}

func BenchmarkBytesToBinarySize2(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		BytesSize(1152921504606846976, "binary", 1)
	}
}

func ExampleBytesSize() {
	fmt.Println(BytesSize(999,  "binary", -1))
	fmt.Println(BytesSize(1058576, "decimal", 2))
	fmt.Println(BytesSize(1099511627776, "binary", 0))
	fmt.Println(BytesSize(1152921504606846976, "decimal", 6,))
	// Output:
	// 999B
	// 1.06MB
	// 1TiB
	// 1.152922EB
}