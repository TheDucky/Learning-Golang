# Golang Types
<hr>

## Numbers

### Integres

#### Unsigned Integres (no negatives)

- `uint8` (0 to 255)
- `uint16` (0 to 65535)
- `uint32` (0 to 4294967295)
- `uint64` (0 to 18446744073709551615)

#### Signed Integres

- `int8` (-128 to 127) 
- `int16` (-32768 to 32767)
- `int32` (-2147483648 to 2147483647)
- `int64` (-9223372036854775808 to 9223372036854775807)

#### Other Numeric Types

- `byte` (same as uint8)
- `rune` (same as int32)
- `uint` (32 or 64 bits. Depends on the system)
- `int` (same size as uint)
- `uintptr` (an unsigned integer to store the uninterpreted bits of a pointer value)

### Floating Points

#### Floats
- `float32` (IEEE-754 32-bit floating-point numbers)
- `float64` (IEEE-754 64-bit floating-point numbers)

#### complex (nerdy math and physics stuff)

- `	complex64` (Complex numbers with float32 real and imaginary parts)
- `complex128` (Complex numbers with float64 real and imaginary parts)

## Strings

- `string = "Store String Here"`

## Boolean

- `bool = true/false`