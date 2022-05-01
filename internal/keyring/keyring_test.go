package keyring

import (
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/shoenig/test/must"
	"gophers.dev/pkgs/secrets"
)

func TestRing_EncryptDecrypt(t *testing.T) {
	id, err := uuid.GenerateUUID()
	must.NoError(t, err)

	password := "passw0rd"
	r := New(secrets.New(id))

	enc := r.Encrypt(secrets.New(password))
	must.NotNil(t, enc)
	must.NotEq(t, []byte(password), enc)

	plain := r.Decrypt(enc)
	must.Eq(t, password, plain.Secret())
}
