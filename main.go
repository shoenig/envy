package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"gophers.dev/cmds/envy/internal/commands"
	"gophers.dev/cmds/envy/internal/output"
)

const (
	usageGroup = "usage"
	envyGroup  = "envy"
)

func main() {
	writer := output.NewWriter(os.Stdout, os.Stderr)

	listCmd := commands.NewListCmd(writer)
	setCmd := commands.NewSetCmd(writer)
	addCmd := commands.NewAddCmd(writer)
	purgeCmd := commands.NewPurgeCmd(writer)
	showCmd := commands.NewShowCmd(writer)

	fs := flag.NewFlagSet(envyGroup, flag.ContinueOnError)
	subs := subcommands.NewCommander(fs, "envy commands")
	subs.Register(subs.HelpCommand(), usageGroup)
	subs.Register(subs.FlagsCommand(), usageGroup)
	subs.Register(listCmd, envyGroup)
	subs.Register(setCmd, envyGroup)
	subs.Register(addCmd, envyGroup)
	subs.Register(purgeCmd, envyGroup)
	subs.Register(showCmd, envyGroup)

	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	ctx := context.Background()
	rc := subs.Execute(ctx, fs.Args())
	os.Exit(int(rc))
}
