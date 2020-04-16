package keyring

import (
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/stretchr/testify/require"
	"gophers.dev/pkgs/secrets"
)

func TestRing_EncryptDecrypt(t *testing.T) {
	t.Parallel()

	id, err := uuid.GenerateUUID()
	require.NoError(t, err)

	password := "passw0rd"
	r := New(secrets.New(id))

	enc := r.Encrypt(secrets.New(password))
	require.NotNil(t, enc)
	require.NotEqual(t, []byte(password), enc)

	plain := r.Decrypt(enc)
	require.Equal(t, password, plain.Secret())
}
