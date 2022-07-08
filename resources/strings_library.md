# strings library in Golang..  
### `import "strings"`
---

![go-string-index](/images/go-string-index.png)

<u>code</u> ↴

```go
package main

import "fmt"

func main() {

    // Refer to image above to understant index and output values

    var example := "golang"

    fmt.Printf("%c\n", example[0])
    fmt.Printf("%c\n", example[3])
    fmt.Printf("%c\n", example[len(example) -2])
    fmt.Printf("%c\n", example[len(example) -5])
}
```

<u>OUTPUT:</u> ↴

```
~>> go run stringIndexExample.go
g
a
n
o
~>>
```
---

To find the length of a string In Go, we use the `len()` function. (used in the above example) (this function is not part of the strings library)

| Functions | Descriptions |
| :---: | :--- |
| `.Compare()` | compares two strings
| `.Contains()` | checks if a substring is present inside a string
| `.Replaces()` | replaces a substring with another substring
| `.ToLower()`	| converts a string to lowercase
| `.ToUpper()`	| converts a string to uppercase
| `.Split()` | splits a string into multiple substrings
| `.Compare()` | compare 2 strings to eachother

```go
import (
  "fmt"
  "strings" //This is the strings library :)
)
```

---

*bibliography*
- https://golangr.com/strings/
- https://www.programiz.com/golang/string