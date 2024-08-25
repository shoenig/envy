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
	listCmdName     = "list"
	listCmdSynopsis = "List all namespaces."
	listCmdUsage    = "list"
)

func NewListCmd(t *setup.Tool) subcommands.Command {
	return &listCmd{
		writer: t.Writer,
		box:    t.Box,
	}
}

type listCmd struct {
	writer output.Writer
	box    safe.Box
}

func (listCmd) Name() string {
	return listCmdName
}

func (listCmd) Synopsis() string {
	return listCmdSynopsis
}

func (listCmd) Usage() string {
	return listCmdUsage
}

func (listCmd) SetFlags(*flag.FlagSet) {
	// no flags when listing namespaces
}

func (lc listCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if f.NArg() != 0 {
		lc.writer.Errorf("list command expects no args")
		return subcommands.ExitUsageError
	}

	namespaces, err := lc.box.List()
	if err != nil {
		lc.writer.Errorf("unable to list namespaces: %v", err)
		return subcommands.ExitFailure
	}

	for _, ns := range namespaces {
		lc.writer.Printf("%s", ns)
	}

	return subcommands.ExitSuccess
}
