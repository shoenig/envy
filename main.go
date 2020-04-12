package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/commands"
	"gophers.dev/cmds/envy/internal/setup"
)

const (
	usageGroup = "usage"
	envyGroup  = "envy"
)

func main() {
	tool := setup.New()

	fs := flag.NewFlagSet(envyGroup, flag.ContinueOnError)
	subs := subcommands.NewCommander(fs, "envy commands")
	subs.Register(subs.HelpCommand(), usageGroup)
	subs.Register(subs.FlagsCommand(), usageGroup)
	subs.Register(commands.NewListCmd(tool), envyGroup)
	subs.Register(commands.NewSetCmd(tool), envyGroup)
	subs.Register(commands.NewUpdateCmd(tool), envyGroup)
	subs.Register(commands.NewPurgeCmd(tool), envyGroup)
	subs.Register(commands.NewShowCmd(tool), envyGroup)
	subs.Register(commands.NewExecCmd(tool), envyGroup)

	if err := fs.Parse(os.Args[1:]); err != nil {
		tool.Writer.Errorf("unable to parse args: %v", err)
		os.Exit(1)
	}

	ctx := context.Background()
	rc := subs.Execute(ctx, fs.Args())
	os.Exit(int(rc))
}
