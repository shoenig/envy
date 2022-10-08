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
	updateCmdName     = "update"
	updateCmdSynopsis = "Add or Update environment variable(s) in namespace."
	updateCmdUsage    = "update [namespace] [env=value,...]"
)

func NewUpdateCmd(t *setup.Tool) subcommands.Command {
	return &updateCmd{
		writer: t.Writer,
		ex:     newExtractor(t.Ring),
		box:    t.Box,
	}
}

type updateCmd struct {
	writer output.Writer
	ex     Extractor
	box    safe.Box
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
	ns, err := uc.ex.Namespace(args)
	if err != nil {
		uc.writer.Errorf("unable to parse args: %v", err)
		return subcommands.ExitUsageError
	}

	if len(ns.Content) == 0 {
		uc.writer.Errorf("use 'purge' to remove namespace")
		return subcommands.ExitUsageError
	}

	if err := uc.box.Update(ns); err != nil {
		uc.writer.Errorf("unable to update namespace: %v", err)
		return subcommands.ExitFailure
	}

	uc.writer.Directf("updated %d items in %q", len(ns.Content), ns.Name)
	return subcommands.ExitSuccess
}
