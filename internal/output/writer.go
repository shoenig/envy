package output

import (
	"fmt"
	"io"
)

type Writer interface {
	Directf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// NewWriter creates a new Writer with the given output sinks. Typically one
// would plug normal into os.Stdout and failure into os.Stderr, but other
// outputs may be provided, for example in use of test cases.
//
// todo: make tracing configurable
func NewWriter(normal, failure io.Writer) Writer {
	return &writer{
		normal:  normal,
		failure: failure,
		traces:  false,
	}
}

type writer struct {
	normal  io.Writer
	failure io.Writer
	traces  bool
}

func (w *writer) Directf(format string, args ...interface{}) {
	tweaked := format + "\n"
	s := fmt.Sprintf(tweaked, args...)
	w.write(s)
}

func (w *writer) Errorf(format string, args ...interface{}) {
	tweaked := "envy: " + format + "\n"
	s := fmt.Sprintf(tweaked, args...)
	w.error(s)
}

func (w *writer) write(s string) {
	_, _ = io.WriteString(w.normal, s)
}

func (w *writer) error(s string) {
	_, _ = io.WriteString(w.failure, s)
}
