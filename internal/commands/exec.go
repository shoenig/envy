// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"noxide.lol/go/babycli"
)

func newExecCmd(tool *setup.Tool) *babycli.Component {
	return &babycli.Component{
		Name: "exec",
		Help: "run a command using environment variables from profile",
		Flags: babycli.Flags{
			{
				Type:  babycli.BooleanFlag,
				Long:  "insulate",
				Short: "i",
				Help:  "disable child process from inheriting parent environment variables",
				Default: &babycli.Default{
					Value: false,
					Show:  false,
				},
			},
		},
		Function: func(c *babycli.Component) babycli.Code {
			if c.Nargs() < 2 {
				tool.Writer.Errorf("must specify profile and command argument(s)")
				return babycli.Failure
			}

			args := c.Arguments()
			p, err := tool.Box.Get(args[0])
			if err != nil {
				tool.Writer.Errorf("unable to read profile: %v", err)
				return babycli.Failure
			}

			insulate := c.GetBool("insulate")
			argVars, command, args := splitArgs(args[1:])
			cmd := newCmd(tool, p, insulate, argVars, command, args)

			if err := cmd.Run(); err != nil {
				tool.Writer.Errorf("failed to exec: %v", err)
				return babycli.Failure
			}

			return babycli.Success
		},
	}
}
func env(tool *setup.Tool, pr *safe.Profile, environment []string) []string {
	for key, value := range pr.Content {
		secret := tool.Ring.Decrypt(value).Unveil()
		environment = append(environment, fmt.Sprintf(
			"%s=%s", key, secret,
		))
	}
	return environment
}

func envContext(insulate bool) []string {
	if insulate {
		return nil
	}
	return os.Environ()
}

func newCmd(tool *setup.Tool, ns *safe.Profile, insulate bool, argVars []string, command string, args []string) *exec.Cmd {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, command, args...)

	// Environment variables are injected in the following order:
	// 1. OS variables if insulate is false
	// 2. envy profile vars
	// 3. Variables in input args
	cmd.Env = append(env(tool, ns, envContext(insulate)), argVars...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd
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
