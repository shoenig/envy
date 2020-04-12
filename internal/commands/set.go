package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/keyring"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/setup"
)

const (
	setCmdName     = "set"
	setCmdSynopsis = "Set environment variable(s) for namespace."
	setCmdUsage    = "set [namespace] [env=value, ...]"
)

func NewSetCmd(t *setup.Tool) subcommands.Command {
	return &setCmd{
		writer: t.Writer,
		ex:     newExtractor(t.Ring),
	}
}

type setCmd struct {
	writer output.Writer
	ring   keyring.Ring
	ex     Extractor
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

	ns, err := sc.ex.Namespace(args)
	if err != nil {
		sc.writer.Errorf("unable to parse args: %v", err)
		return subcommands.ExitUsageError
	}

	sc.writer.Directf("extracted namespace: %v", ns)

	return subcommands.ExitSuccess
}
