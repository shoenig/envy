package keyring

import (
	"os"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

func init() {
	// For tests only, use the mock implementation of the keyring provider.
	keyring.MockInit()
}

func setEnv(t *testing.T, key, value string) string {
	previous := os.Getenv(key)
	err := os.Setenv(key, value)
	require.NoError(t, err)
	return previous
}

func TestInit_user(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		prevUser := setEnv(t, "USER", "")
		defer setEnv(t, "USER", prevUser)

		prevEnvyUser := setEnv(t, "ENVY_USER", "")
		defer setEnv(t, "ENVY_USER", prevEnvyUser)

		u := user()
		require.Equal(t, "default", u)
	})

	t.Run("user", func(t *testing.T) {
		prevUser := setEnv(t, "USER", "alice")
		defer setEnv(t, "USER", prevUser)

		prevEnvyUser := setEnv(t, "ENVY_USER", "")
		defer setEnv(t, "ENVY_USER", prevEnvyUser)

		u := user()
		require.Equal(t, "alice", u)
	})

	t.Run("envy_user", func(t *testing.T) {
		prevUser := setEnv(t, "USER", "alice")
		defer setEnv(t, "USER", prevUser)

		prevEnvyUser := setEnv(t, "ENVY_USER", "bob")
		defer setEnv(t, "ENVY_USER", prevEnvyUser)

		u := user()
		require.Equal(t, "bob", u)
	})
}

func isUUID(t *testing.T, id string) {
	_, err := uuid.ParseUUID(id)
	require.NoError(t, err)
}

func TestInit_bootstrap(t *testing.T) {
	id := bootstrap("envy.name", "alice")
	isUUID(t, id.Secret())
}

func TestInit_init(t *testing.T) {
	prevEnvyUser := setEnv(t, "ENVY_USER", "alice")
	defer setEnv(t, "ENVY_USER", prevEnvyUser)

	// first time goes through bootstrap
	id := Init("envy.name")
	isUUID(t, id.Secret())

	// subsequent time should already exist
	id2 := Init("envy.name")
	isUUID(t, id2.Secret())

	// and the result should be the same
	require.Equal(t, id.Secret(), id2.Secret())
}
