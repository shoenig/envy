package commands

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/google/subcommands"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
)

const (
	execCmdName     = "exec"
	execCmdSynopsis = "Run command with environment variables from namespace."
	execCmdUsage    = "exec [namespace] [command] <args, ...>"

	flagInsulate = "insulate"
)

func NewExecCmd(t *setup.Tool) subcommands.Command {
	return &execCmd{
		writer:        t.Writer,
		ring:          t.Ring,
		box:           t.Box,
		execOutputStd: os.Stdout,
		execOutputErr: os.Stderr,
	}
}

type execCmd struct {
	writer        output.Writer
	ring          keyring.Ring
	box           safe.Box
	execOutputStd io.Writer
	execOutputErr io.Writer
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

func (wc execCmd) SetFlags(fs *flag.FlagSet) {
	// no flags when running command through exec
	_ = fs.Bool(flagInsulate, false, "insulate will run command without passing through environment")
}

func (wc execCmd) Execute(ctx context.Context, fs *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	insulate := fsBool(fs, flagInsulate)

	if len(fs.Args()) < 2 {
		wc.writer.Errorf("expected namespace and command argument(s)")
		return subcommands.ExitUsageError
	}

	ns, err := wc.box.Get(fs.Arg(0))
	if err != nil {
		wc.writer.Errorf("could not retrieve namespace: %v", err)
		return subcommands.ExitUsageError
	}

	cmd := wc.newCmd(ns, insulate, fs.Arg(1), fs.Args()[2:])
	if err := cmd.Run(); err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (wc execCmd) newCmd(ns *safe.Namespace, insulate bool, command string, args []string) *exec.Cmd {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Env = wc.env(ns, envContext(insulate))
	cmd.Stdout = wc.execOutputStd
	cmd.Stderr = wc.execOutputErr
	return cmd
}

func envContext(insulate bool) []string {
	if insulate {
		return nil
	}
	return os.Environ()
}

func (wc execCmd) env(ns *safe.Namespace, environment []string) []string {
	for key, value := range ns.Content {
		secret := wc.ring.Decrypt(value).Secret()
		environment = append(environment, fmt.Sprintf(
			"%s=%s", key, secret,
		))
	}
	return environment
}
