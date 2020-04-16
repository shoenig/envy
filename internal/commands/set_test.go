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

func TestSetCmd_Ops(t *testing.T) {
	t.Parallel()

	db := newDBFile(t)
	cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewSetCmd(setup.New(db, w))

	t.Run("name", func(t *testing.T) {
		require.Equal(t, setCmdName, cmd.Name())
	})
	t.Run("synopsis", func(t *testing.T) {
		require.Equal(t, setCmdSynopsis, cmd.Synopsis())
	})
	t.Run("usage", func(t *testing.T) {
		require.Equal(t, setCmdUsage, cmd.Usage())
	})
}

func TestSetCmd_Execute(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	ring.EncryptMock.When(secrets.New("abc123")).Then(safe.Encrypted{8, 8, 8, 8, 8, 8})
	ring.EncryptMock.When(secrets.New("1234")).Then(safe.Encrypted{9, 9, 9, 9})

	box.SetMock.Expect(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": safe.Encrypted{8, 8, 8, 8, 8, 8},
			"bar": safe.Encrypted{9, 9, 9, 9},
		},
	}).Return(nil)

	pc := &setCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{"set", "myNS", "foo=abc123", "bar=1234"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitSuccess, rc)
	require.Equal(t, "stored 2 items in \"myNS\"\n", a.String())
	require.Empty(t, b.String())
}

func TestSetCmd_Execute_ioError(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	box.SetMock.Expect(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": safe.Encrypted{8, 8, 8, 8, 8, 8},
			"bar": safe.Encrypted{9, 9, 9, 9},
		},
	}).Return(errors.New("io error"))

	ring.EncryptMock.When(secrets.New("abc123")).Then(safe.Encrypted{8, 8, 8, 8, 8, 8})
	ring.EncryptMock.When(secrets.New("1234")).Then(safe.Encrypted{9, 9, 9, 9})

	pc := &setCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{"set", "myNS", "foo=abc123", "bar=1234"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitFailure, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: unable to update namespace: io error\n", b.String())
}

func TestSetCmd_Execute_badNS(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	pc := &setCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	// e.g. forgot to specify namespace
	fs, args := setupFlagSet(t, []string{"set", "foo=abc123", "bar=1234"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: unable to parse args: namespace uses non-word characters\n", b.String())
}

func TestSetCmd_Execute_noVars(t *testing.T) {
	t.Parallel()

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	pc := &setCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	// e.g. reminder to use purge to remove namespace
	fs, args := setupFlagSet(t, []string{"set", "ns1"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: use 'purge' to remove namespace\n", b.String())
}
