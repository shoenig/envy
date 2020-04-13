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

func (lc listCmd) Name() string {
	return listCmdName
}

func (lc listCmd) Synopsis() string {
	return listCmdSynopsis
}

func (lc listCmd) Usage() string {
	return listCmdUsage
}

func (lc listCmd) SetFlags(set *flag.FlagSet) {
	// no flags when listing namespaces
}

func (lc listCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	namespaces, err := lc.box.List()
	if err != nil {
		lc.writer.Errorf("unable to list namespaces: %v", err)
		return subcommands.ExitFailure
	}

	for _, ns := range namespaces {
		lc.writer.Directf("%s", ns)
	}

	return subcommands.ExitSuccess
}
