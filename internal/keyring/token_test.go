// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package keyring

import (
	"os"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/shoenig/test/must"
	"github.com/zalando/go-keyring"
)

func init() {
	// For tests only, use the mock implementation of the keyring provider.
	keyring.MockInit()
}

func setEnv(t *testing.T, key, value string) string {
	previous := os.Getenv(key)
	t.Setenv(key, value)
	return previous
}

func TestInit_user(t *testing.T) {
	// no parallel, sets environment

	t.Run("default", func(t *testing.T) {
		prevUser := setEnv(t, "USER", "")
		defer setEnv(t, "USER", prevUser)

		prevEnvyUser := setEnv(t, "ENVY_USER", "")
		defer setEnv(t, "ENVY_USER", prevEnvyUser)

		u := user()
		must.Eq(t, "default", u)
	})

	t.Run("user", func(t *testing.T) {
		prevUser := setEnv(t, "USER", "alice")
		defer setEnv(t, "USER", prevUser)

		prevEnvyUser := setEnv(t, "ENVY_USER", "")
		defer setEnv(t, "ENVY_USER", prevEnvyUser)

		u := user()
		must.Eq(t, "alice", u)
	})

	t.Run("envy_user", func(t *testing.T) {
		prevUser := setEnv(t, "USER", "alice")
		defer setEnv(t, "USER", prevUser)

		prevEnvyUser := setEnv(t, "ENVY_USER", "bob")
		defer setEnv(t, "ENVY_USER", prevEnvyUser)

		u := user()
		must.Eq(t, "bob", u)
	})
}

func isUUID(t *testing.T, id string) {
	_, err := uuid.ParseUUID(id)
	must.NoError(t, err)
}

func TestInit_bootstrap(t *testing.T) {
	// no parallel, sets environment

	id := bootstrap("envy.name", "alice")
	isUUID(t, id.Unveil())
}

func TestInit_init(t *testing.T) {
	// no parallel, sets environment

	prevEnvyUser := setEnv(t, "ENVY_USER", "alice")
	defer setEnv(t, "ENVY_USER", prevEnvyUser)

	// first time goes through bootstrap
	id := Init("envy.name")
	isUUID(t, id.Unveil())

	// subsequent time should already exist
	id2 := Init("envy.name")
	isUUID(t, id2.Unveil())

	// and the result should be the same
	must.Eq(t, id.Unveil(), id2.Unveil())
}
