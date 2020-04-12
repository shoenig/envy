package safe

import (
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/pkgs/secrets"
)

const (
	filename = "envy.safe"
)

// Path returns the filepath to the state store (boltdb) used by envoy for
// persisting encrypted secrets.
//
// For now this defers to 'os.UserConfigDir/envy/envoy.safe', but should be
// made configurable later on.
func Path(w output.Writer) (string, error) {
	configs, err := os.UserConfigDir()
	if err != nil {
		return "", errors.Wrap(err, "no user config directory")
	}

	dir := filepath.Join(configs, "envy")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", errors.Wrap(err, "unable to not create config directory")
	}

	return filepath.Join(dir, filename), nil
}

// Box represents the persistent storage of encrypted secrets.
//
// todo: keep track of envy storage schema version, in case a newer version of
//  envy needs to modify the way secrets are stored
type Box struct {
	database *bbolt.DB
}

type Namespace struct {
	Name    string
	Content map[string]secrets.Text
}

func New(file string) (*Box, error) {
	options := &bbolt.Options{
		Timeout: 1 * time.Second,
	}

	db, err := bbolt.Open(file, 0600, options)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open persistent storage")
	}

	return &Box{
		database: db,
	}, nil
}

func (b *Box) Close() error {
	return b.database.Close()
}

func bucket(create bool, tx *bbolt.Tx, namespace string) (*bbolt.Bucket, error) {
	if create {
		return tx.CreateBucketIfNotExists([]byte(namespace))
	}
	if bkt := tx.Bucket([]byte(namespace)); bkt != nil {
		return bkt, nil
	}
	return nil, errors.Errorf("namespace %q does not exist", namespace)
}

func put(bkt *bbolt.Bucket, ns *Namespace) error {
	for k, v := range ns.Content {
		if err := bkt.Put([]byte(k), []byte(v.Secret())); err != nil {
			return err
		}
	}
	return nil
}

func wipe(bkt *bbolt.Bucket, namespace string) error {
	return bkt.ForEach(func(k, _ []byte) error {
		return bkt.Delete(k)
	})
}

// Purge will delete the namespace, including any existing content.
func (b *Box) Purge(namespace string) error {
	return b.database.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(namespace))
	})
}

// Set will set the contents of ns, eliminating or overwriting anything that
// existed in that namespace before.
func (b *Box) Set(ns *Namespace) error {
	return b.database.Update(func(tx *bbolt.Tx) error {
		bkt, err := bucket(true, tx, ns.Name)
		if err != nil {
			return err
		}
		if err := wipe(bkt, ns.Name); err != nil {
			return err
		}
		return put(bkt, ns)
	})
}

// Update will set the contents of ns, overwriting anything that existed in that
// namespace before. Pre-existing non-overlapping values will remain.
func (b *Box) Update(ns *Namespace) error {
	return b.database.Update(func(tx *bbolt.Tx) error {
		bkt, err := bucket(true, tx, ns.Name)
		if err != nil {
			return err
		}
		return put(bkt, ns)
	})
}

// Get will return the contents of namespace.
func (b *Box) Get(namespace string) (*Namespace, error) {
	content := make(map[string]secrets.Text)

	if err := b.database.View(func(tx *bbolt.Tx) error {
		bkt, err := bucket(false, tx, namespace)
		if err != nil {
			return err
		}

		if err := bkt.ForEach(func(k []byte, v []byte) error {
			content[string(k)] = secrets.New(string(v))
			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &Namespace{
		Name:    namespace,
		Content: content,
	}, nil
}
