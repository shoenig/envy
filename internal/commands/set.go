package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
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
		box:    t.Box,
	}
}

type setCmd struct {
	writer output.Writer
	ex     Extractor
	box    safe.Box
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
	ns, err := sc.ex.Namespace(args)
	if err != nil {
		sc.writer.Errorf("unable to parse args: %v", err)
		return subcommands.ExitUsageError
	}

	if len(ns.Content) == 0 {
		sc.writer.Errorf("use 'purge' to remove namespace")
		return subcommands.ExitUsageError
	}

	if err := sc.box.Set(ns); err != nil {
		sc.writer.Errorf("unable to update namespace: %v", err)
		return subcommands.ExitFailure
	}

	sc.writer.Directf("stored %d items in %q", len(ns.Content), ns.Name)
	return subcommands.ExitSuccess
}
