# byteconv
Converts a size in bytes to a human-readable string in SI (decimal) or IEC (binary, base-2) format

**Note**: Inspired by https://code.cloudfoundry.org/bytefmt

#### Examples
Convert bytes to binary (base-2) format
```go
byteconv.BytesSize(-456, "binary", 0)              // returns 0
byteconv.BytesSize(1099511627776, "binary", 0)     // returns 1TiB
byteconv.BytesSize(5048576, "binary", 0)           // returns 5MiB
byteconv.BytesSize(1024, "binary", 0)              // returns 1KiB
byteconv.BytesSize(1048576, "binary", 2)           // returns 1.00MiB
byteconv.BytesSize(1058576, "binary", -1)          // returns 1.0095367431640625MiB
```

Convert bytes to decimal format
```go
byteconv.BytesSize(1058576, "decimal", 2)           // returns 1.06MB
byteconv.BytesSize(1125899906842624, "decimal", 2)  // returns 1PB
byteconv.BytesSize(1000, "decimal", 0)              // returns 1KB
byteconv.BytesSize(1058576, "decimal", -1)          // returns 1.058576MB
```

Binary string to bytes
```go
byteconv.StringToBytes("1KiB")                      // returns 1024
byteconv.StringToBytes("-3MiB")                     // returns 0
byteconv.StringToBytes("1GiB")                      // returns 1073741824
byteconv.StringToBytes("\t10.18TiB\n")              // returns 11193028370759.679688
```


Benchmark results
```shell script
BenchmarkBytesToBinarySize-12           10000000               200 ns/op              40 B/op          3 allocs/op
BenchmarkStringToBytesTiB-12            10000000               132 ns/op               8 B/op          1 allocs/op
BenchmarkStringToBytesTB-12             20000000                83.3 ns/op             0 B/op          0 allocs/op
BenchmarkBytesToDecimalSize-12           5000000               384 ns/op              42 B/op          4 allocs/op
BenchmarkBytesToBinarySize2-12           5000000               339 ns/op              64 B/op          5 allocs/op
```

coverage: 98.3% of statements