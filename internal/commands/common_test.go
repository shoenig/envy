package commands

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gophers.dev/cmds/envy/internal/output"
)

func newDBFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "tool-")
	require.NoError(t, err)
	err = f.Close()
	require.NoError(t, err)
	return f.Name()
}

func cleanupFile(t *testing.T, name string) {
	err := os.Remove(name)
	require.NoError(t, err)
}

func newWriter() (*bytes.Buffer, *bytes.Buffer, output.Writer) {
	var a, b bytes.Buffer
	return &a, &b, output.New(&a, &b)
}

func setupFlagSet(t *testing.T, arguments []string) (*flag.FlagSet, interface{}) {
	fs := flag.NewFlagSet("test", flag.PanicOnError)
	err := fs.Parse(arguments)
	require.NoError(t, err)
	return fs, arguments
}

func TestCommon_args(t *testing.T) {
	t.Parallel()

	// google/subcommands passes args wrapped like this
	wrap := func(a []string) []interface{} {
		return []interface{}{a}
	}

	t.Run("no arguments", func(t *testing.T) {
		_, _, _, err := extract(wrap([]string{}))
		require.EqualError(t, err, "not enough arguments")
	})

	t.Run("one argument", func(t *testing.T) {
		_, _, _, err := extract(wrap([]string{"foo"}))
		require.EqualError(t, err, "not enough arguments")
	})

	t.Run("two arguments", func(t *testing.T) {
		verb, ns, argv, err := extract(wrap([]string{"foo", "bar"}))
		require.Equal(t, "foo", verb)
		require.Equal(t, "bar", ns)
		require.Empty(t, argv)
		require.NoError(t, err)
	})

	t.Run("four arguments", func(t *testing.T) {
		verb, ns, secrets, err := extract(wrap([]string{"a", "b", "c", "d"}))
		require.Equal(t, "a", verb)
		require.Equal(t, "b", ns)
		require.Equal(t, 2, len(secrets))
		require.Equal(t, "c", secrets[0].Secret())
		require.Equal(t, "d", secrets[1].Secret())
		require.NoError(t, err)
	})
}
