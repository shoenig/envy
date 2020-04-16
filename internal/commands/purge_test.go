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

func TestPurgeCmd_Ops(t *testing.T) {
	t.Parallel()

	db := newDBFile(t)
	cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewPurgeCmd(setup.New(db, w))

	t.Run("name", func(t *testing.T) {
		require.Equal(t, purgeCmdName, cmd.Name())
	})
	t.Run("synopsis", func(t *testing.T) {
		require.Equal(t, purgeCmdSynopsis, cmd.Synopsis())
	})
	t.Run("usage", func(t *testing.T) {
		require.Equal(t, purgeCmdUsage, cmd.Usage())
	})
}

func TestPurgeCmdExecute(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	pc := &purgeCmd{
		writer: w,
		box:    box,
	}

	box.PurgeMock.Expect("myNS").Return(nil)

	fs, args := setupFlagSet(t, []string{"myNS"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitSuccess, rc)
	require.Equal(t, "purged namespace \"myNS\"\n", a.String())
	require.Empty(t, b.String())
}

func TestPurgeCmd_Execute_purgeFails(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	a, b, w := newWriter()

	pc := &purgeCmd{
		writer: w,
		box:    box,
	}

	box.PurgeMock.Expect("myNS").Return(errors.New("does not exist"))

	fs, args := setupFlagSet(t, []string{"myNS"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitFailure, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: could not purge namespace: does not exist\n", b.String())
}

func TestPurgeCmd_Execute_noArg(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	pc := &purgeCmd{
		writer: w,
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: expected one namespace argument\n", b.String())
}

func TestPurgeCmd_Execute_twoArg(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	pc := &purgeCmd{
		writer: w,
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{"ns1", "ns2"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: expected one namespace argument\n", b.String())
}
