package commands

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/stretchr/testify/require"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/safe"
	"gophers.dev/cmds/envy/internal/setup"
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

func TestListCmd_Ops(t *testing.T) {
	db := newDBFile(t)
	cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewListCmd(setup.New(db, w))

	t.Run("name", func(t *testing.T) {
		require.Equal(t, listCmdName, cmd.Name())
	})
	t.Run("synopsis", func(t *testing.T) {
		require.Equal(t, listCmdSynopsis, cmd.Synopsis())
	})
	t.Run("usage", func(t *testing.T) {
		require.Equal(t, listCmdUsage, cmd.Usage())
	})
}

func TestListCmd_Execute(t *testing.T) {
	box := safe.NewBoxMock(t)
	a, b, w := newWriter()

	lc := &listCmd{
		writer: w,
		box:    box,
	}

	box.ListMock.Expect().Return([]string{
		"namespace1", "ns2", "my-ns",
	}, nil)

	ctx := context.Background()
	rc := lc.Execute(ctx, nil, nil)

	require.Equal(t, subcommands.ExitSuccess, rc)
	require.Equal(t, "namespace1\nns2\nmy-ns\n", a.String())
	require.Empty(t, b.String())

}
