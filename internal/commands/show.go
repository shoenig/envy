package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/output"
)

const (
	showCmdName     = "show"
	showCmdSynopsis = "Show environment variable(s) in namespace."
	showCmdUsage    = "show [namespace]"
)

func NewShowCmd(w output.Writer) subcommands.Command {
	return &showCmd{
		writer: w,
	}
}

type showCmd struct {
	writer output.Writer
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

func (sc showCmd) SetFlags(set *flag.FlagSet) {
	// no flags when showing environment variables of namespace
}

func (sc showCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	sc.writer.Directf("the show command!")
	return subcommands.ExitSuccess
}
