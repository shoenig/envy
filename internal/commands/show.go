// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"context"
	"flag"
	"sort"

	"github.com/google/subcommands"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
)

const (
	showCmdName     = "show"
	showCmdSynopsis = "Show environment variable(s) in namespace."
	showCmdUsage    = "show [--decrypt] [--rm] [namespace]"

	flagDecrypt = "decrypt"
)

func NewShowCmd(t *setup.Tool) subcommands.Command {
	return &showCmd{
		writer: t.Writer,
		ring:   t.Ring,
		box:    t.Box,
	}
}

type showCmd struct {
	writer output.Writer
	ring   keyring.Ring
	box    safe.Box
}

func (sc showCmd) Name() string {
	return showCmdName
}

func (sc showCmd) Synopsis() string {
	return showCmdSynopsis
}

func (sc showCmd) Usage() string {
	return showCmdUsage
}

func (sc showCmd) SetFlags(fs *flag.FlagSet) {
	_ = fs.Bool(flagDecrypt, false, "decrypt will print secrets")
}

func (sc showCmd) Execute(_ context.Context, fs *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	decrypt := fsBool(fs, flagDecrypt)

	if len(fs.Args()) != 1 {
		sc.writer.Errorf("expected only namespace argument")
		return subcommands.ExitUsageError
	}

	ns, err := sc.box.Get(fs.Arg(0))
	if err != nil {
		sc.writer.Errorf("could not retrieve namespace: %v", err)
		return subcommands.ExitFailure
	}

	keys := make([]string, 0, len(ns.Content))
	for k := range ns.Content {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if decrypt {
			value := ns.Content[key]
			secret := sc.ring.Decrypt(value)
			sc.writer.Printf("%s=%s", key, secret.Unveil())
		} else {
			sc.writer.Printf("%s", key)
		}
	}

	return subcommands.ExitSuccess
}
