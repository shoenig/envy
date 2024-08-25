// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/google/subcommands"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
)

const (
	execCmdName     = "exec"
	execCmdSynopsis = "Run command with environment variables from namespace."
	execCmdUsage    = "exec [namespace] [command] <args, ...>"

	flagInsulate = "insulate"
)

func NewExecCmd(t *setup.Tool) subcommands.Command {
	return &execCmd{
		writer:        t.Writer,
		ring:          t.Ring,
		box:           t.Box,
		execInputStd:  os.Stdin,
		execOutputStd: os.Stdout,
		execOutputErr: os.Stderr,
	}
}

type execCmd struct {
	writer        output.Writer
	ring          keyring.Ring
	box           safe.Box
	execInputStd  io.Reader
	execOutputStd io.Writer
	execOutputErr io.Writer
}

func (wc execCmd) Name() string {
	return execCmdName
}

func (wc execCmd) Synopsis() string {
	return execCmdSynopsis
}

func (wc execCmd) Usage() string {
	return execCmdUsage
}

func (wc execCmd) SetFlags(fs *flag.FlagSet) {
	// no flags when running command through exec
	_ = fs.Bool(flagInsulate, false, "insulate will run command without passing through environment")
}

func (wc execCmd) Execute(_ context.Context, fs *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	insulate := fsBool(fs, flagInsulate)

	if len(fs.Args()) < 2 {
		wc.writer.Errorf("expected namespace and command argument(s)")
		return subcommands.ExitUsageError
	}

	ns, err := wc.box.Get(fs.Arg(0))
	if err != nil {
		wc.writer.Errorf("could not retrieve namespace: %v", err)
		return subcommands.ExitUsageError
	}

	argVars, command, args := splitArgs(fs.Args()[1:])
	cmd := wc.newCmd(ns, insulate, argVars, command, args)
	if err := cmd.Run(); err != nil {
		wc.writer.Errorf("failed to exec: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

// splitArgs will split the list of flag.Args() into:
// in-lined env vars, command, command args
func splitArgs(flagArgs []string) (argVars []string, command string, commandArgs []string) {
	var (
		startIdx   = 0
		commandIdx = 0
	)

	// Special case for when variables are injected in the form: "env FOO=BAR command"
	if flagArgs[0] == "env" {
		startIdx++
		commandIdx++
	}

	for commandIdx < len(flagArgs) {
		_, err := exec.LookPath(flagArgs[commandIdx])
		if err != nil && strings.Contains(flagArgs[commandIdx], "=") {
			// Assume that arguments not in the path and with an equal sign are env vars.
			commandIdx++
		} else {
			break
		}
	}

	// Only detected environment-setting, fall back to assuming the first is the command.
	if commandIdx >= len(flagArgs) {
		return nil, flagArgs[0], flagArgs[1:]
	}

	command = flagArgs[commandIdx]
	argVars = flagArgs[startIdx:commandIdx]
	commandArgs = flagArgs[commandIdx+1:]

	return argVars, command, commandArgs
}

func (wc execCmd) newCmd(ns *safe.Namespace, insulate bool, argVars []string, command string, args []string) *exec.Cmd {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, command, args...)

	// Environment variables are injected in the following order:
	// 1. OS variables if insulate is false
	// 2. envy namespace vars
	// 3. Variables in input args
	cmd.Env = append(wc.env(ns, envContext(insulate)), argVars...)
	cmd.Stdout = wc.execOutputStd
	cmd.Stderr = wc.execOutputErr
	cmd.Stdin = wc.execInputStd
	return cmd
}

func envContext(insulate bool) []string {
	if insulate {
		return nil
	}
	return os.Environ()
}

func (wc execCmd) env(ns *safe.Namespace, environment []string) []string {
	for key, value := range ns.Content {
		secret := wc.ring.Decrypt(value).Unveil()
		environment = append(environment, fmt.Sprintf(
			"%s=%s", key, secret,
		))
	}
	return environment
}
