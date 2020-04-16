package commands

import (
	"context"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gophers.dev/cmds/envy/internal/keyring"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/safe"
	"gophers.dev/cmds/envy/internal/setup"
	"gophers.dev/pkgs/secrets"
)

func TestUpdateCmd_Ops(t *testing.T) {
	t.Parallel()

	db := newDBFile(t)
	cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewUpdateCmd(setup.New(db, w))

	t.Run("name", func(t *testing.T) {
		require.Equal(t, updateCmdName, cmd.Name())
	})
	t.Run("synopsis", func(t *testing.T) {
		require.Equal(t, updateCmdSynopsis, cmd.Synopsis())
	})
	t.Run("usage", func(t *testing.T) {
		require.Equal(t, updateCmdUsage, cmd.Usage())
	})
}

func TestUpdateCmd_Execute(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	uc := &updateCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	ring.EncryptMock.When(secrets.New("1")).Then(safe.Encrypted{0xA})
	ring.EncryptMock.When(secrets.New("2")).Then(safe.Encrypted{0xB})

	box.UpdateMock.Expect(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"a": safe.Encrypted{0xA},
			"b": safe.Encrypted{0xB},
		},
	}).Return(nil)

	fs, args := setupFlagSet(t, []string{"update", "myNS", "a=1", "b=2"})
	uc.SetFlags(fs)
	ctx := context.Background()
	rc := uc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitSuccess, rc)
	require.Equal(t, "updated 2 items in \"myNS\"\n", a.String())
	require.Empty(t, b.String())
}

func TestUpdateCmd_Execute_noArgs(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	uc := &updateCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{"update", "myNS"})
	uc.SetFlags(fs)
	ctx := context.Background()
	rc := uc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: use 'purge' to remove namespace\n", b.String())
}

func TestUpdateCmd_Execute_noNS(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	uc := &updateCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{"update"})
	uc.SetFlags(fs)
	ctx := context.Background()
	rc := uc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: unable to parse args: not enough arguments\n", b.String())
}

func TestUpdateCmd_Execute_ioError(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	uc := &updateCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	ring.EncryptMock.When(secrets.New("1")).Then(safe.Encrypted{0xA})
	ring.EncryptMock.When(secrets.New("2")).Then(safe.Encrypted{0xB})

	box.UpdateMock.Expect(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"a": safe.Encrypted{0xA},
			"b": safe.Encrypted{0xB},
		},
	}).Return(errors.New("io error"))

	fs, args := setupFlagSet(t, []string{"update", "myNS", "a=1", "b=2"})
	uc.SetFlags(fs)
	ctx := context.Background()
	rc := uc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitFailure, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: unable to update namespace: io error\n", b.String())
}
