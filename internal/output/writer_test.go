package output

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriter_Directf(t *testing.T) {
	t.Parallel()

	var a, b bytes.Buffer

	w := New(&a, &b)
	w.Directf("foo: %d", 42)
	require.Equal(t, "foo: 42\n", a.String())
	require.Equal(t, "", b.String())
}

func TestWriter_Errorf(t *testing.T) {
	t.Parallel()

	var a, b bytes.Buffer

	w := New(&a, &b)
	w.Errorf("foo: %d", 99)
	require.Equal(t, "", a.String())
	require.Equal(t, "envy: foo: 99\n", b.String())
}
