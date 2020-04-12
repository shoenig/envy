package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/setup"
)

const (
	setCmdName     = "set"
	setCmdSynopsis = "Set environment variable(s) for namespace."
	setCmdUsage    = "set [namespace] [env=value,...]"
)

func NewSetCmd(t *setup.Tool) subcommands.Command {
	return &setCmd{
		writer: t.Writer,
	}
}

type setCmd struct {
	writer output.Writer
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

func (sc setCmd) SetFlags(set *flag.FlagSet) {
	// no flags when setting environment
}

func (sc setCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	sc.writer.Directf("the set command!")
	return subcommands.ExitSuccess
}
