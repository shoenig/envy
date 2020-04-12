package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/setup"
)

const (
	execCmdName     = "exec"
	execCmdSynopsis = "Run command with environment variables from namespace."
	execCmdUsage    = "exec [namespace] [command] <args, ...>"
)

func NewExecCmd(t *setup.Tool) subcommands.Command {
	return &execCmd{
		writer: t.Writer,
	}
}

type execCmd struct {
	writer output.Writer
}

func (wc execCmd) Name() string {
	return execCmdName
}

func (wc execCmd) Synopsis() string {
	return execCmdSynopsis
}

func (wc execCmd) Usage() string {
	return execCmdUsage
}

func (wc execCmd) SetFlags(set *flag.FlagSet) {
	// no flags when running command through exec
}

func (wc execCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	wc.writer.Directf("the exec command!")
	return subcommands.ExitSuccess
}
