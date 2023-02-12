package commands

import (
	"bytes"
	"flag"
	"os"
	"testing"

	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/test/must"
	"github.com/zalando/go-keyring"
)

func init() {
	// For tests only, use the mock implementation of the keyring provider.
	keyring.MockInit()
}

func newDBFile(t *testing.T) string {
	f, err := os.CreateTemp("", "tool-")
	must.NoError(t, err)
	err = f.Close()
	must.NoError(t, err)
	return f.Name()
}

func cleanupFile(t *testing.T, name string) {
	err := os.Remove(name)
	must.NoError(t, err)
}

func newWriter() (*bytes.Buffer, *bytes.Buffer, output.Writer) {
	var a, b bytes.Buffer
	return &a, &b, output.New(&a, &b)
}

func setupFlagSet(t *testing.T, arguments []string) (*flag.FlagSet, interface{}) {
	fs := flag.NewFlagSet("test", flag.PanicOnError)
	err := fs.Parse(arguments)
	must.NoError(t, err)
	return fs, arguments
}
