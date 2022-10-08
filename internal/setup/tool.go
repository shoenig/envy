package setup

import (
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
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

	return &Tool{
		Writer: w,
		Ring:   keyring.New(keyring.Init(envyKeyringName)),
		Box:    safe.New(dbFile),
	}
}
