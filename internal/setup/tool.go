package setup

import (
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

func New(file string, w output.Writer) *Tool {
	dbFile, err := safe.Path(file)
	if err != nil {
		panic(err)
	}

	box, err := safe.New(dbFile)
	if err != nil {
		panic(err)
	}

	return &Tool{
		Writer: w,
		Ring:   keyring.New(keyring.Init(envyKeyringName)),
		Box:    box,
	}
}
