package commands

import (
	"context"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/stretchr/testify/require"
	"gophers.dev/cmds/envy/internal/keyring"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/safe"
	"gophers.dev/cmds/envy/internal/setup"
	"gophers.dev/pkgs/secrets"
)

func TestShowCmd_Ops(t *testing.T) {
	db := newDBFile(t)
	cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewShowCmd(setup.New(db, w))

	t.Run("name", func(t *testing.T) {
		require.Equal(t, showCmdName, cmd.Name())
	})
	t.Run("synopsis", func(t *testing.T) {
		require.Equal(t, showCmdSynopsis, cmd.Synopsis())
	})
	t.Run("usage", func(t *testing.T) {
		require.Equal(t, showCmdUsage, cmd.Usage())
	})
}

func TestShowCmd_Execute(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	sc := &showCmd{
		writer: w,
		ring:   ring,
		box:    box,
	}

	box.GetMock.Expect("myNS").Return(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": safe.Encrypted{1, 1, 1},
			"bar": safe.Encrypted{2, 2, 2},
		},
	}, nil)

	fs, args := setupFlagSet(t, []string{"myNS"})
	sc.SetFlags(fs)
	ctx := context.Background()
	rc := sc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitSuccess, rc)
	require.Equal(t, "bar\nfoo\n", a.String())
	require.Empty(t, b.String())
}

func TestShowCmd_Execute_decrypt(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	sc := &showCmd{
		writer: w,
		ring:   ring,
		box:    box,
	}

	box.GetMock.Expect("myNS").Return(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": safe.Encrypted{1, 1, 1},
			"bar": safe.Encrypted{2, 2, 2},
		},
	}, nil)

	ring.DecryptMock.When([]byte{2, 2, 2}).Then(secrets.New("passw0rd"))
	ring.DecryptMock.When([]byte{1, 1, 1}).Then(secrets.New("hunter2"))

	fs, args := setupFlagSet(t, []string{"myNS"})
	sc.SetFlags(fs)
	require.NoError(t, fs.Set("decrypt", "true"))
	ctx := context.Background()
	rc := sc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitSuccess, rc)
	require.Equal(t, "bar=passw0rd\nfoo=hunter2\n", a.String())
	require.Empty(t, b.String())
}

func TestShowCmd_Execute_noNS(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	sc := &showCmd{
		writer: w,
		ring:   ring,
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{})
	sc.SetFlags(fs)
	ctx := context.Background()
	rc := sc.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: expected only namespace argument\n", b.String())
}
