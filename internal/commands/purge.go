package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/output"
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
	}
}

type purgeCmd struct {
	writer output.Writer
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

func (pc purgeCmd) SetFlags(set *flag.FlagSet) {
	// no flags when purging namespace
}

func (pc purgeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	pc.writer.Directf("the purge command!")
	return subcommands.ExitSuccess
}
