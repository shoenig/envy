// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"context"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/go-conceal"
	"github.com/shoenig/test/must"
)

func TestShowCmd_Ops(t *testing.T) {
	db := newDBFile(t)
	defer cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewShowCmd(setup.New(db, w))

	must.Eq(t, showCmdName, cmd.Name())
	must.Eq(t, showCmdSynopsis, cmd.Synopsis())
	must.Eq(t, showCmdUsage, cmd.Usage())
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
			"foo": {1, 1, 1},
			"bar": {2, 2, 2},
		},
	}, nil)

	fs, args := setupFlagSet(t, []string{"myNS"})
	sc.SetFlags(fs)
	ctx := context.Background()
	rc := sc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitSuccess, rc)
	must.Eq(t, "bar\nfoo\n", a.String())
	must.Eq(t, "", b.String())
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
			"foo": {1, 1, 1},
			"bar": {2, 2, 2},
		},
	}, nil)

	ring.DecryptMock.When([]byte{2, 2, 2}).Then(conceal.New("passw0rd"))
	ring.DecryptMock.When([]byte{1, 1, 1}).Then(conceal.New("hunter2"))

	fs, args := setupFlagSet(t, []string{"myNS"})
	sc.SetFlags(fs)
	must.NoError(t, fs.Set("decrypt", "true"))

	ctx := context.Background()
	rc := sc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitSuccess, rc)
	must.Eq(t, "bar=passw0rd\nfoo=hunter2\n", a.String())
	must.Eq(t, "", b.String())
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

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: expected only namespace argument\n", b.String())
}
