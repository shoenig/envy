package keyring

import (
	"os"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/zalando/go-keyring"
	"github.com/shoenig/secrets"
)

const (
	userEnvOverride = "ENVY_USER"
	userEnvDefault  = "USER"
	userDefault     = "default"
)

func user() string {
	user := userDefault
	if u := os.Getenv(userEnvDefault); u != "" {
		user = u
	}
	if u := os.Getenv(userEnvOverride); u != "" {
		user = u
	}
	return user
}

func trim(id string) string {
	return strings.TrimSpace(strings.ReplaceAll(id, "-", ""))
}

func bootstrap(name, user string) secrets.Text {
	token, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}

	if err := keyring.Set(name, user, token); err != nil {
		panic(err)
	}
	return secrets.New(token)
}

// Init will acquire the secret that envy uses to encrypt values from the OS
// keyring provider.
func Init(name string) secrets.Text {
	user := user()
	token, err := keyring.Get(name, user)

	switch {
	case err == keyring.ErrNotFound:
		return bootstrap(name, user)
	case err != nil:
		panic(err)
	default:
		return secrets.New(token)
	}
}
