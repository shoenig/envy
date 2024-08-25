// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package safe

import (
	"os"
	"path/filepath"
	"slices"
	"sync"
	"time"

	"github.com/hashicorp/go-set/v3"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

const (
	filename = "envy.safe"
)

// Path returns the filepath to the state store (boltdb) used by envoy for
// persisting encrypted secrets.
func Path(dbFile string) (string, error) {
	if dbFile != "" {
		return dbFile, nil
	}
	configs, err := os.UserConfigDir()
	if err != nil {
		return "", errors.Wrap(err, "no user config directory")
	}

	dir := filepath.Join(configs, "envy")
	if err = os.MkdirAll(dir, 0700); err != nil {
		return "", errors.Wrap(err, "unable to not create config directory")
	}

	return filepath.Join(dir, filename), nil
}

// A Box represents the persistent storage of encrypted secrets.
//
//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock@v3.0.10 -g -i Box -s _mock.go
type Box interface {
	Set(*Profile) error
	Delete(string, *set.Set[string]) error
	Purge(string) error
	Get(string) (*Profile, error)
	List() ([]string, error)
}

type box struct {
	file string

	lock     sync.Mutex
	database *bbolt.DB
}

func New(file string) Box {
	return &box{
		file: file,
	}
}

func (b *box) open() error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.database != nil {
		return errors.New("database already open")
	}

	options := &bbolt.Options{
		Timeout: 3 * time.Second,
	}
	db, err := bbolt.Open(b.file, 0600, options)
	if err != nil {
		return errors.Wrap(err, "unable to open persistent storage")
	}
	b.database = db
	return nil
}

func (b *box) close(openErr error) {
	if openErr != nil {
		panic(openErr)
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	if b.database == nil {
		panic("database already closed")
	}

	if err := b.database.Close(); err != nil {
		panic(err)
	}

	b.database = nil
}

func bucket(create bool, tx *bbolt.Tx, profile string) (*bbolt.Bucket, error) {
	if create {
		return tx.CreateBucketIfNotExists([]byte(profile))
	}
	if bkt := tx.Bucket([]byte(profile)); bkt != nil {
		return bkt, nil
	}
	return nil, errors.Errorf("profile %q does not exist", profile)
}

func put(bkt *bbolt.Bucket, pr *Profile) error {
	for k, v := range pr.Content {
		if err := bkt.Put([]byte(k), []byte(v)); err != nil {
			return err
		}
	}
	return nil
}

// Purge will delete the profile, including any existing content.
func (b *box) Purge(profile string) error {
	defer b.close(b.open())

	return b.database.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(profile))
	})
}

// Set will amend the content of ns. Any overlapping pre-existing values will be
// overwritten.
func (b *box) Set(pr *Profile) error {
	defer b.close(b.open())

	return b.database.Update(func(tx *bbolt.Tx) error {
		bkt, err := bucket(true, tx, pr.Name)
		if err != nil {
			return err
		}
		return put(bkt, pr)
	})
}

// Delete will remove keys from profile.
func (b *box) Delete(profile string, keys *set.Set[string]) error {
	defer b.close(b.open())

	return b.database.Update(func(tx *bbolt.Tx) error {
		bkt, err := bucket(true, tx, profile)
		if err != nil {
			return err
		}

		for _, key := range keys.Slice() {
			if err = bkt.Delete([]byte(key)); err != nil {
				return err
			}
		}
		return nil
	})
}

// Get will return the contents of profile.
func (b *box) Get(profile string) (*Profile, error) {
	defer b.close(b.open())

	content := make(map[string]Encrypted)
	if err := b.database.View(func(tx *bbolt.Tx) error {
		bkt, err := bucket(false, tx, profile)
		if err != nil {
			return err
		}

		if err = bkt.ForEach(func(k []byte, v []byte) error {
			content[string(k)] = Encrypted(slices.Clone(v))
			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &Profile{
		Name:    profile,
		Content: content,
	}, nil
}

// List will return a list of profile that have been created.
func (b *box) List() ([]string, error) {
	defer b.close(b.open())

	var profiles []string
	if err := b.database.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(ns []byte, _ *bbolt.Bucket) error {
			profiles = append(profiles, string(ns))
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return profiles, nil
}
