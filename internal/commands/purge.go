package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/safe"
	"gophers.dev/cmds/envy/internal/setup"
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

func (pc purgeCmd) Name() string {
	return purgeCmdName
}

func (pc purgeCmd) Synopsis() string {
	return purgeCmdSynopsis
}

func (pc purgeCmd) Usage() string {
	return purgeCmdUsage
}

func (pc purgeCmd) SetFlags(_ *flag.FlagSet) {
	// no flags when purging namespace
}

func (pc purgeCmd) Execute(ctx context.Context, fs *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if len(fs.Args()) != 1 {
		pc.writer.Errorf("expected only namespace argument")
		return subcommands.ExitUsageError
	}

	ns, err := pc.box.Get(fs.Arg(0))
	if err != nil {
		pc.writer.Errorf("could not retrieve namespace: %v", err)
		return subcommands.ExitFailure
	}

	if err := pc.box.Purge(ns.Name); err != nil {
		pc.writer.Errorf("could not purge namespace: %v", err)
		return subcommands.ExitFailure
	}

	pc.writer.Directf("purged %s", ns.Name)
	return subcommands.ExitSuccess
}
