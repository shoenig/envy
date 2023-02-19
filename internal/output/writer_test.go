package output

import (
	"bytes"
	"testing"

	"github.com/shoenig/test/must"
)

func TestWriter_Directf(t *testing.T) {
	var a, b bytes.Buffer

	w := New(&a, &b)
	w.Printf("foo: %d", 42)
	must.Eq(t, "foo: 42\n", a.String())
	must.Eq(t, "", b.String())
}

func TestWriter_Errorf(t *testing.T) {
	var a, b bytes.Buffer

	w := New(&a, &b)
	w.Errorf("foo: %d", 99)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: foo: 99\n", b.String())
}
