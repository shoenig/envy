package commands

import (
	"context"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/pkg/errors"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/test/must"
)

func TestListCmd_Ops(t *testing.T) {
	db := newDBFile(t)
	defer cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewListCmd(setup.New(db, w))

	must.Eq(t, listCmdName, cmd.Name())
	must.Eq(t, listCmdSynopsis, cmd.Synopsis())
	must.Eq(t, listCmdUsage, cmd.Usage())
}

func TestListCmd_Execute(t *testing.T) {
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

	// no arguments for list
	fs, args := setupFlagSet(t, []string{})
	lc.SetFlags(fs)
	ctx := context.Background()
	rc := lc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitSuccess, rc)
	must.Eq(t, "namespace1\nns2\nmy-ns\n", a.String())
	must.Eq(t, "", b.String())
}

func TestListCmd_Execute_listFails(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	lc := &listCmd{
		writer: w,
		box:    box,
	}

	box.ListMock.Expect().Return(nil, errors.New("io error"))

	// no arguments for list
	fs, args := setupFlagSet(t, []string{})
	lc.SetFlags(fs)
	ctx := context.Background()
	rc := lc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitFailure, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to list namespaces: io error\n", b.String())
}

func TestListCmd_Execute_extraArgs(t *testing.T) {
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

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: list command expects no args\n", b.String())
}
