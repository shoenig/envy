package commands

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/shoenig/test/must"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/secrets"
)

func TestExecCmd_Ops(t *testing.T) {
	db := newDBFile(t)
	defer cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewExecCmd(setup.New(db, w))

	must.Eq(t, execCmdName, cmd.Name())
	must.Eq(t, execCmdSynopsis, cmd.Synopsis())
	must.Eq(t, execCmdUsage, cmd.Usage())
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

	must.Eq(t, subcommands.ExitSuccess, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "", b.String())
	must.Eq(t, "a is passw0rd\nb is hunter2\n", c.String())
	must.Eq(t, "", d.String())
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

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: expected namespace and command argument(s)\n", b.String())
	must.Eq(t, "", c.String())
	must.Eq(t, "", d.String())
}
