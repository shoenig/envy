package commands

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
		verb, ns, argv, err := extract(wrap([]string{"a", "b", "c", "d"}))
		require.Equal(t, "a", verb)
		require.Equal(t, "b", ns)
		require.Equal(t, []string{"c", "d"}, argv)
		require.NoError(t, err)
	})
}
