# Golang Output Formatting
____

    Package fmt implements formatted I/O with functions analogous
    to C's printf and scanf. The format 'verbs' are derived from
    C's but are simpler.

### General

- `%v` the value in a default format
- `%T` a Go-syntax representation of the type of the value
- `%%` a literal percent sign; consumes no value

### Boolean

- `%t` the word true or false

### Integer

- `%b` base 2
- `%c` the character represented by the corresponding Unicode code point
- `%d` base 10
- `%o` base 8
- `%O` base 8 with 0o prefix
- `%q` a single-quoted character literal safely escaped with Go syntax.
- `%x` base 16, with lower-case letters for a-f
- `%X` base 16, with upper-case letters for A-F
- `%U` Unicode format: U+1234; same as "U+%04X"

### Floating Points

- `%e` scientific notation
- `%f` decimal point but no exponent
- `%g` for large exponents

### String

- `%s` the uninterpreted bytes of the string or slice
- `%q` a double-quoted string safely escaped with Go syntax
- `%x` base 16, lower-case, two characters per byte
- `%X` base 16, upper-case, two characters per byte

### Width And Precision

- `%f` default width, default precision
- `%9f` width 9, default precision
- `%.2f` default width, precision 2
- `%9.2f` width 9, precision 2
- `%9.f` width 9, precision 0

### Paddnig

- `%08d` pads digit to length 8 with preceeding 0's
- `%-5d` pads with spaces | width 5, left justified
___

*bibliography* 
- https://pkg.go.dev/fmt