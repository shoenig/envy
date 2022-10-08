package commands

import (
	"bytes"
	"flag"
	"io/ioutil"
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
	f, err := ioutil.TempFile("", "tool-")
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

func TestCommon_args(t *testing.T) {

	// google/subcommands passes args wrapped like this
	wrap := func(a []string) []interface{} {
		return []interface{}{a}
	}

	t.Run("no arguments", func(t *testing.T) {
		_, _, _, err := extract(wrap([]string{}))
		must.EqError(t, err, "not enough arguments")
	})

	t.Run("one argument", func(t *testing.T) {
		_, _, _, err := extract(wrap([]string{"foo"}))
		must.EqError(t, err, "not enough arguments")
	})

	t.Run("two arguments", func(t *testing.T) {
		verb, ns, argv, err := extract(wrap([]string{"foo", "bar"}))
		must.Eq(t, "foo", verb)
		must.Eq(t, "bar", ns)
		must.SliceEmpty(t, argv)
		must.NoError(t, err)
	})

	t.Run("four arguments", func(t *testing.T) {
		verb, ns, secrets, err := extract(wrap([]string{"a", "b", "c", "d"}))
		must.Eq(t, "a", verb)
		must.Eq(t, "b", ns)
		must.Eq(t, 2, len(secrets))
		must.Eq(t, "c", secrets[0].Secret())
		must.Eq(t, "d", secrets[1].Secret())
		must.NoError(t, err)
	})
}
