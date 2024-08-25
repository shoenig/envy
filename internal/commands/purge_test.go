// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

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

func TestPurgeCmd_Ops(t *testing.T) {
	db := newDBFile(t)
	defer cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewPurgeCmd(setup.New(db, w))

	must.Eq(t, purgeCmdName, cmd.Name())
	must.Eq(t, purgeCmdSynopsis, cmd.Synopsis())
	must.Eq(t, purgeCmdUsage, cmd.Usage())
}

func TestPurgeCmdExecute(t *testing.T) {
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

	must.Eq(t, subcommands.ExitSuccess, rc)
	must.Eq(t, "purged namespace \"myNS\"\n", a.String())
	must.Eq(t, "", b.String())
}

func TestPurgeCmd_Execute_purgeFails(t *testing.T) {
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

	must.Eq(t, subcommands.ExitFailure, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: could not purge namespace: does not exist\n", b.String())
}

func TestPurgeCmd_Execute_noArg(t *testing.T) {
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

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: expected one namespace argument\n", b.String())
}

func TestPurgeCmd_Execute_badNS(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	pc := &purgeCmd{
		writer: w,
		box:    box,
	}

	// namespace must be valid
	fs, args := setupFlagSet(t, []string{"foo=bar"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: could not purge namespace: namespace uses non-word characters\n", b.String())
}

func TestPurgeCmd_Execute_twoArg(t *testing.T) {
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

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: expected one namespace argument\n", b.String())
}
