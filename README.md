# lesiw.io/prefix

[![Go Reference](https://pkg.go.dev/badge/lesiw.io/prefix.svg)](https://pkg.go.dev/lesiw.io/prefix)

Package prefix adds prefixes to I/O primitives.

``` go
package main

import (
    "os"

    "lesiw.io/prefix"
)

func ExampleNewWriter() {
    w := prefix.NewWriter("> ", os.Stdout)
    w.Write([]byte("hello\n"))
    w.Write([]byte("world\n"))
    // Output:
    // > hello
    // > world
}
```

[▶️ Run this example on the Go playground](https://go.dev/play/p/ywU4qElXeSe)
