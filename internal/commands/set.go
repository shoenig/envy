// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/hashicorp/go-set/v2"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
)

const (
	setCmdName     = "set"
	setCmdSynopsis = "Set environment variable(s) for namespace."
	setCmdUsage    = "set [namespace] [env=value, ...]"
)

func NewSetCmd(t *setup.Tool) subcommands.Command {
	return &setCmd{
		writer: t.Writer,
		ex:     newExtractor(t.Ring),
		box:    t.Box,
	}
}

type setCmd struct {
	writer output.Writer
	ex     Extractor
	box    safe.Box
}

func (sc setCmd) Name() string {
	return setCmdName
}

func (sc setCmd) Synopsis() string {
	return setCmdSynopsis
}

func (sc setCmd) Usage() string {
	return setCmdUsage
}

func (sc setCmd) SetFlags(fs *flag.FlagSet) {
}

func (sc setCmd) Execute(ctx context.Context, fs *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	namespace, remove, add, err := sc.ex.PreProcess(fs.Args())
	if err != nil {
		sc.writer.Errorf("unable to parse args: %v", err)
		return subcommands.ExitUsageError
	}

	if err = checkName(namespace); err != nil {
		sc.writer.Errorf("could not set namespace: %v", err)
		return subcommands.ExitUsageError
	}

	if err = sc.rm(namespace, remove); err != nil {
		sc.writer.Errorf("unable to remove variables: %v", err)
		return subcommands.ExitFailure
	}

	n, err := sc.ex.Namespace(add)
	if err != nil {
		sc.writer.Errorf("unable to parse args: %v", err)
		return subcommands.ExitUsageError
	}
	n.Name = namespace

	if err = sc.set(n); err != nil {
		sc.writer.Errorf("unable to set variables: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (sc setCmd) rm(namespace string, keys *set.Set[string]) error {
	if keys.Empty() {
		return nil
	}
	return sc.box.Delete(namespace, keys)
}

func (sc setCmd) set(ns *safe.Namespace) error {
	return sc.box.Set(ns)
}
