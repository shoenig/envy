package keyring

import (
	"errors"
	"os"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/shoenig/go-conceal"
	"github.com/zalando/go-keyring"
)

const (
	userEnvOverride = "ENVY_USER"
	userEnvDefault  = "USER"
	userDefault     = "default"
)

func user() string {
	usr := userDefault
	if envUsr := os.Getenv(userEnvDefault); envUsr != "" {
		usr = envUsr
	}
	if envUsr := os.Getenv(userEnvOverride); envUsr != "" {
		usr = envUsr
	}
	return usr
}

func trim(id string) string {
	return strings.TrimSpace(strings.ReplaceAll(id, "-", ""))
}

func bootstrap(name, user string) *conceal.Text {
	token, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}

	if err = keyring.Set(name, user, token); err != nil {
		panic(err)
	}
	return conceal.New(token)
}

// Init will acquire the secret that envy uses to encrypt values from the OS
// keyring provider.
func Init(name string) *conceal.Text {
	usr := user()
	token, err := keyring.Get(name, usr)

	switch {
	case errors.Is(err, keyring.ErrNotFound):
		return bootstrap(name, usr)
	case err != nil:
		panic(err)
	default:
		return conceal.New(token)
	}
}
