package commands

import (
	"bytes"
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

func TestExecCmd_Ops(t *testing.T) {
	db := newDBFile(t)
	cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewExecCmd(setup.New(db, w))

	t.Run("name", func(t *testing.T) {
		require.Equal(t, execCmdName, cmd.Name())
	})
	t.Run("synopsis", func(t *testing.T) {
		require.Equal(t, execCmdSynopsis, cmd.Synopsis())
	})
	t.Run("usage", func(t *testing.T) {
		require.Equal(t, execCmdUsage, cmd.Usage())
	})
}

func TestExecCmd_Execute(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()
	var c, d bytes.Buffer

	ec := &execCmd{
		writer:        w,
		ring:          ring,
		box:           box,
		execOutputStd: &c,
		execOutputErr: &d,
	}

	box.GetMock.Expect("myNS").Return(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"a": safe.Encrypted{0x1},
			"b": safe.Encrypted{0x2},
		},
	}, nil)

	ring.DecryptMock.When(safe.Encrypted{0x1}).Then(secrets.New("passw0rd"))
	ring.DecryptMock.When(safe.Encrypted{0x2}).Then(secrets.New("hunter2"))

	fs, args := setupFlagSet(t, []string{"myNS", "./testing/a.sh"})
	ec.SetFlags(fs)
	ctx := context.Background()
	rc := ec.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitSuccess, rc)
	require.Empty(t, a.String())
	require.Empty(t, b.String())
	require.Equal(t, "a is passw0rd\nb is hunter2\n", c.String())
	require.Empty(t, d.String())
}

func TestExecCmd_Execute_noCommand(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()
	var c, d bytes.Buffer

	ec := &execCmd{
		writer:        w,
		ring:          ring,
		box:           box,
		execOutputStd: &c,
		execOutputErr: &d,
	}

	fs, args := setupFlagSet(t, []string{"myNS"})
	ec.SetFlags(fs)
	ctx := context.Background()
	rc := ec.Execute(ctx, fs, args)

	require.Equal(t, subcommands.ExitUsageError, rc)
	require.Empty(t, a.String())
	require.Equal(t, "envy: expected namespace and command argument(s)\n", b.String())
	require.Empty(t, c.String())
	require.Empty(t, d.String())
}
