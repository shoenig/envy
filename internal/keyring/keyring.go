package keyring

import (
	"fmt"

	"gophers.dev/cmds/envy/internal/safe"
	"gophers.dev/pkgs/secrets"
)

func Init(name string) secrets.Text {
	return secrets.New("abc123")
}

type Ring interface {
	Encrypt(secrets.Text) safe.Encrypted
	Decrypt(safe.Encrypted) secrets.Text
}

type ring struct {
	key secrets.Text
}

func New(key secrets.Text) Ring {
	return &ring{
		key: key,
	}
}

func (r *ring) Encrypt(s secrets.Text) safe.Encrypted {
	fmt.Printf("encrypt %s\n", s.Secret())
	// todo
	return safe.Encrypted(s.Secret())
}

func (r *ring) Decrypt(s safe.Encrypted) secrets.Text {
	fmt.Printf("decrypt %s\n", s)
	// todo
	return secrets.New(string(s))
}
