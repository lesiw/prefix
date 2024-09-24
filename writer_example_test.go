package prefix_test

import (
	"os"

	"lesiw.io/prefix"
)

// nolint:errcheck
func ExampleNewWriter() {
	w := prefix.NewWriter("> ", os.Stdout)
	w.Write([]byte("hello\n"))
	w.Write([]byte("world\n"))
	// Output:
	// > hello
	// > world
}
