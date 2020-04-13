package setup

import (
	"os"

	"gophers.dev/cmds/envy/internal/keyring"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/cmds/envy/internal/safe"
)

const (
	envyKeyringName = "envy.secure.env.vars"
)

type Tool struct {
	Writer output.Writer
	Ring   keyring.Ring
	Box    safe.Box
}

func New() *Tool {
	dbFile, err := safe.Path()
	if err != nil {
		panic(err)
	}

	box, err := safe.New(dbFile)
	if err != nil {
		panic(err)
	}

	return &Tool{
		Writer: output.New(os.Stdout, os.Stderr),
		Ring:   keyring.New(keyring.Init(envyKeyringName)),
		Box:    box,
	}
}
