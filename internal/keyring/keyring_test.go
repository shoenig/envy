// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package keyring

import (
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/shoenig/go-conceal"
	"github.com/shoenig/test/must"
)

func TestRing_EncryptDecrypt(t *testing.T) {
	id, err := uuid.GenerateUUID()
	must.NoError(t, err)

	password := "passw0rd"
	r := New(conceal.New(id))

	enc := r.Encrypt(conceal.New(password))
	must.NotNil(t, enc)
	must.NotEq(t, []byte(password), enc)

	plain := r.Decrypt(enc)
	must.Eq(t, password, plain.Unveil())
}
