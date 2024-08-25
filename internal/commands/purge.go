// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
)

const (
	purgeCmdName     = "purge"
	purgeCmdSynopsis = "Purge a namespace."
	purgeCmdUsage    = "purge [namespace]"
)

func NewPurgeCmd(t *setup.Tool) subcommands.Command {
	return &purgeCmd{
		writer: t.Writer,
		box:    t.Box,
	}
}

type purgeCmd struct {
	writer output.Writer
	box    safe.Box
}

func (purgeCmd) Name() string {
	return purgeCmdName
}

func (purgeCmd) Synopsis() string {
	return purgeCmdSynopsis
}

func (purgeCmd) Usage() string {
	return purgeCmdUsage
}

func (purgeCmd) SetFlags(_ *flag.FlagSet) {
	// no flags when purging namespace
}

func (pc purgeCmd) Execute(_ context.Context, fs *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if fs.NArg() != 1 {
		pc.writer.Errorf("expected one namespace argument")
		return subcommands.ExitUsageError
	}

	namespace := fs.Arg(0)
	if err := checkName(namespace); err != nil {
		pc.writer.Errorf("could not purge namespace: %v", err)
		return subcommands.ExitUsageError
	}

	if err := pc.box.Purge(namespace); err != nil {
		pc.writer.Errorf("could not purge namespace: %v", err)
		return subcommands.ExitFailure
	}

	pc.writer.Printf("purged namespace %q", namespace)
	return subcommands.ExitSuccess
}
