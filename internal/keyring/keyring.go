package keyring

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"gophers.dev/cmds/envy/internal/safe"
	"gophers.dev/pkgs/secrets"
)

type Ring interface {
	Encrypt(secrets.Text) safe.Encrypted
	Decrypt(safe.Encrypted) secrets.Text
}

type ring struct {
	key secrets.Bytes
}

func New(key secrets.Text) Ring {
	return &ring{
		key: uuidToLen32(key),
	}
}

func uuidToLen32(id secrets.Text) secrets.Bytes {
	return secrets.NewBytes([]byte(trim(id.Secret())))
}

func (r *ring) Encrypt(s secrets.Text) safe.Encrypted {
	bCipher, err := aes.NewCipher(r.key.Secret())
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(bCipher)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	return safe.Encrypted(gcm.Seal(nonce, nonce, []byte(s.Secret()), nil))
}

func (r *ring) Decrypt(s safe.Encrypted) secrets.Text {
	bCipher, err := aes.NewCipher(r.key.Secret())
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

	return secrets.New(string(plainText))
}
