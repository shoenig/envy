package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/setup"
)

const (
	updateCmdName     = "update"
	updateCmdSynopsis = "Add or Update environment variable(s) in namespace."
	updateCmdUsage    = "update [namespace] [env=value,...]"
)

func NewUpdateCmd(t *setup.Tool) subcommands.Command {
	return &updateCmd{
		writer: t.Writer,
	}
}

type updateCmd struct {
	writer output.Writer
}

func (uc updateCmd) Name() string {
	return updateCmdName
}

func (uc updateCmd) Synopsis() string {
	return updateCmdSynopsis
}

func (uc updateCmd) Usage() string {
	return updateCmdUsage
}

func (uc updateCmd) SetFlags(set *flag.FlagSet) {
	// no flags when adding environment variable to namespace
}

func (uc updateCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	uc.writer.Directf("the add command!")
	return subcommands.ExitSuccess
}
