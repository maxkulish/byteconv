# byteconv
Converts a size in bytes to a human-readable string in SI (decimal) or IEC (binary) format

**Note**: Inspired by https://code.cloudfoundry.org/bytefmt

###Examples
####Convert bytes to binary format
```go
byteconv.BytesSize(-456, "binary", 0)              // returns 0
byteconv.BytesSize(1099511627776, "binary", 0)     // returns 1TiB
byteconv.BytesSize(5048576, "binary", 0)           // returns 5MiB
byteconv.BytesSize(1024, "binary", 0)              // returns 1KiB
byteconv.BytesSize(1048576, "binary", 2)           // returns 1.00MiB
byteconv.BytesSize(1058576, "binary", -1)          // returns 1.0095367431640625MiB
```

####Convert bytes to decimal format
```go
byteconv.BytesSize(1058576, "decimal", 2)           // returns 1.06MB
byteconv.BytesSize(1125899906842624, "decimal", 2)  // returns 1PB
byteconv.BytesSize(1000, "decimal", 0)              // returns 1KB
byteconv.BytesSize(1058576, "decimal", -1)          // returns 1.058576MB
```