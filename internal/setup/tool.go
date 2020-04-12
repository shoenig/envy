package setup

import (
	"os"

	"gophers.dev/cmds/envy/internal/output"
)

type Tool struct {
	Writer output.Writer
}

func New() *Tool {
	return &Tool{
		Writer: output.New(os.Stdout, os.Stderr),
	}
}
