// Package prefix adds prefixes to I/O primitives.
package prefix

import "io"

// Writer is an [io.Writer] that wraps another [io.Writer] and injects a prefix
// on the first write and after each subsequent newline.
type Writer struct {
	p string
	w io.Writer
	n bool
}

// NewWriter creates a new [Writer].
func NewWriter(prefix string, w io.Writer) *Writer {
	return &Writer{prefix, w, true}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	var a []byte
	for _, b := range p {
		if w.n {
			a = append(a, []byte(w.p)...)
			w.n = false
		}
		a = append(a, b)
		if b == '\n' {
			w.n = true
		}
	}
	return w.w.Write(a)
}
