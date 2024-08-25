// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/google/subcommands"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/go-conceal"
	"github.com/shoenig/test/must"
)

func skipOS(t *testing.T) {
	switch runtime.GOOS {
	case "windows":
		t.Skip("skipping on windows")
	}
}

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
	skipOS(t)

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
		execInputStd:  strings.NewReader("ok\n"),
	}

	box.GetMock.Expect("myNS").Return(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"a": {0x1},
			"b": {0x2},
		},
	}, nil)

	ring.DecryptMock.When(safe.Encrypted{0x1}).Then(conceal.New("passw0rd"))
	ring.DecryptMock.When(safe.Encrypted{0x2}).Then(conceal.New("hunter2"))

	fs, args := setupFlagSet(t, []string{"myNS", "./testing/a.sh"})
	ec.SetFlags(fs)
	ctx := context.Background()
	rc := ec.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitSuccess, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "", b.String())
	must.Eq(t, "a is passw0rd\nb is hunter2\nok\n", c.String())
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

func TestExecCmd_Execute_badCommand(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()
	var c, d bytes.Buffer

	box.GetMock.Expect("myNS").Return(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"a": {0x1},
			"b": {0x2},
		},
	}, nil)

	ring.DecryptMock.When(safe.Encrypted{0x1}).Then(conceal.New("passw0rd"))
	ring.DecryptMock.When(safe.Encrypted{0x2}).Then(conceal.New("hunter2"))

	ec := &execCmd{
		writer:        w,
		ring:          ring,
		box:           box,
		execOutputStd: &c,
		execOutputErr: &d,
	}

	fs, args := setupFlagSet(t, []string{"myNS", "/does/not/exist"})
	ec.SetFlags(fs)
	ctx := context.Background()
	rc := ec.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitFailure, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "", c.String())
	must.Eq(t, "", d.String())

	switch runtime.GOOS {
	case "windows":
		must.Eq(t, "envy: failed to exec: exec: \"/does/not/exist\": file does not exist\n", b.String())
	default:
		must.Eq(t, "envy: failed to exec: fork/exec /does/not/exist: no such file or directory\n", b.String())
	}
}

func Test_splitArgs(t *testing.T) {
	skipOS(t)

	type args struct {
		flagArgs []string
	}
	tests := map[string]struct {
		setup           func()
		args            args
		wantArgVars     []string
		wantCommand     string
		wantCommandArgs []string
	}{
		"no env vars": {
			args: args{
				flagArgs: []string{"cat", "log.out"},
			},
			wantArgVars:     []string{},
			wantCommand:     "cat",
			wantCommandArgs: []string{"log.out"},
		},
		"env vars with command but no args": {
			args: args{
				flagArgs: []string{"FOO=BAR", "ZIP=ZAP", "ls"},
			},
			wantArgVars:     []string{"FOO=BAR", "ZIP=ZAP"},
			wantCommand:     "ls",
			wantCommandArgs: []string{},
		},
		"env vars with command and args": {
			args: args{
				flagArgs: []string{"FOO=BAR", "ZIP=ZAP", "curl", "-k", "localhost:8501"},
			},
			wantArgVars:     []string{"FOO=BAR", "ZIP=ZAP"},
			wantCommand:     "curl",
			wantCommandArgs: []string{"-k", "localhost:8501"},
		},
		"explicit env": {
			args: args{
				flagArgs: []string{"env", "FOO=BAR", "ZIP=ZAP", "curl", "-k", "localhost:8501"},
			},
			wantArgVars:     []string{"FOO=BAR", "ZIP=ZAP"},
			wantCommand:     "curl",
			wantCommandArgs: []string{"-k", "localhost:8501"},
		},
		"only env": {
			args: args{
				flagArgs: []string{"env"},
			},
			wantArgVars:     nil,
			wantCommand:     "env",
			wantCommandArgs: []string{},
		},
		"all vars and none in path": {
			args: args{
				flagArgs: []string{"FOO=BAR", "ZIP=ZAP", "BIP=BOP"},
			},
			wantArgVars:     nil,
			wantCommand:     "FOO=BAR",
			wantCommandArgs: []string{"ZIP=ZAP", "BIP=BOP"},
		},
		"cmd with equal in path is allowed": {
			setup: func() {
				tempDir, err := os.MkdirTemp("", "envy")
				must.NoError(t, err)
				t.Cleanup(func() { os.RemoveAll(tempDir) })

				f, err := os.OpenFile(filepath.Join(tempDir, "my=cmd"), os.O_CREATE|os.O_EXCL, 0700)
				must.NoError(t, err)
				must.NoError(t, f.Close())

				err = os.Setenv("PATH", tempDir)
				must.NoError(t, err)
			},
			args: args{
				flagArgs: []string{"FOO=BAR", "my=cmd", "ZIP=ZAP"},
			},
			wantArgVars:     []string{"FOO=BAR"},
			wantCommand:     "my=cmd",
			wantCommandArgs: []string{"ZIP=ZAP"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			gotArgVars, gotCommand, gotCommandArgs := splitArgs(tt.args.flagArgs)

			must.Eq(t, tt.wantArgVars, gotArgVars)
			must.Eq(t, tt.wantCommand, gotCommand)
			must.Eq(t, tt.wantCommandArgs, gotCommandArgs)
		})
	}
}
