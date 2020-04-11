package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/output"
)

const (
	addCmdName     = "add"
	addCmdSynopsis = "Add environment variable(s) to namespace."
	addCmdUsage    = "add [namespace] [env=value,...]"
)

func NewAddCmd(w output.Writer) subcommands.Command {
	return &addCmd{
		writer: w,
	}
}

type addCmd struct {
	writer output.Writer
}

func (ac addCmd) Name() string {
	return addCmdName
}

func (ac addCmd) Synopsis() string {
	return addCmdSynopsis
}

func (ac addCmd) Usage() string {
	return addCmdUsage
}

func (ac addCmd) SetFlags(set *flag.FlagSet) {
	// no flags when adding environment variable to namespace
}

func (ac addCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	ac.writer.Directf("the add command!")
	return subcommands.ExitSuccess
}
