package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/output"
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
	}
}

type listCmd struct {
	writer output.Writer
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
	lc.writer.Directf("the list command!")
	return subcommands.ExitSuccess
}
