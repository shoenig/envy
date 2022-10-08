package commands

import (
	"context"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/pkg/errors"
	"github.com/shoenig/test/must"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/secrets"
)

func TestUpdateCmd_Ops(t *testing.T) {
	db := newDBFile(t)
	defer cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewUpdateCmd(setup.New(db, w))

	must.Eq(t, updateCmdName, cmd.Name())
	must.Eq(t, updateCmdSynopsis, cmd.Synopsis())
	must.Eq(t, updateCmdUsage, cmd.Usage())
}

func TestUpdateCmd_Execute(t *testing.T) {
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

	must.Eq(t, subcommands.ExitSuccess, rc)
	must.Eq(t, "updated 2 items in \"myNS\"\n", a.String())
	must.Eq(t, "", b.String())
}

func TestUpdateCmd_Execute_noArgs(t *testing.T) {
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

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: use 'purge' to remove namespace\n", b.String())
}

func TestUpdateCmd_Execute_noNS(t *testing.T) {
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

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to parse args: not enough arguments\n", b.String())
}

func TestUpdateCmd_Execute_ioError(t *testing.T) {
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

	must.Eq(t, subcommands.ExitFailure, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to update namespace: io error\n", b.String())
}
