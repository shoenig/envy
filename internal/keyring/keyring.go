// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package keyring

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/go-conceal"
)

// A Ring is used to encrypt and decrypt secrets.
//
//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock@v3.0.10 -g -i Ring -s _mock.go
type Ring interface {
	Encrypt(*conceal.Text) safe.Encrypted
	Decrypt(safe.Encrypted) *conceal.Text
}

type ring struct {
	key *conceal.Bytes
}

func New(key *conceal.Text) Ring {
	return &ring{
		key: uuidToLen32(key),
	}
}

func uuidToLen32(id *conceal.Text) *conceal.Bytes {
	return conceal.NewBytes([]byte(trim(id.Unveil())))
}

func (r *ring) Encrypt(s *conceal.Text) safe.Encrypted {
	bCipher, err := aes.NewCipher(r.key.Unveil())
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(bCipher)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		panic(err)
	}

	return safe.Encrypted(gcm.Seal(nonce, nonce, []byte(s.Unveil()), nil))
}

func (r *ring) Decrypt(s safe.Encrypted) *conceal.Text {
	bCipher, err := aes.NewCipher(r.key.Unveil())
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(bCipher)
	if err != nil {
		panic(err)
	}

	nonce, cText := s[:gcm.NonceSize()], s[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cText, nil)
	if err != nil {
		panic(err)
	}

	return conceal.New(string(plainText))
}
