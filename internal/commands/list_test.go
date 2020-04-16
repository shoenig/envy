package commands

import (
	"context"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/safe"
	"gophers.dev/cmds/envy/internal/setup"
)

func TestListCmd_Ops(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

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

func TestListCmd_Execute_listFails(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	lc := &listCmd{
		writer: w,
		box:    box,
	}

	box.ListMock.Expect().Return(nil, errors.New("io error"))

	ctx := context.Background()
	rc := lc.Execute(ctx, nil, nil)

	require.Equal(t, subcommands.ExitFailure, rc)
	require.Empty(t, "", a.String())
	require.Equal(t, "envy: unable to list namespaces: io error\n", b.String())
}

func TestListCmd_Execute_extraArgs(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	lc := &listCmd{
		writer: w,
		box:    box,
	}

	// nonsense args for list
	fs, args := setupFlagSet(t, []string{"a=b", "c=d"})
	lc.SetFlags(fs)
	ctx := context.Background()
	rc := lc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, "", a.String())
	require.Equal(t, "envy: list command expects no args\n", b.String())
}
