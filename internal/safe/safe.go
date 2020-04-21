package safe

import (
	"os"
	"path/filepath"
	"sync"
	"time"

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
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", errors.Wrap(err, "unable to not create config directory")
	}

	return filepath.Join(dir, filename), nil
}

// A Box represents the persistent storage of encrypted secrets.
//
//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -g -i Box -s _mock.go
type Box interface {
	Set(*Namespace) error
	Purge(string) error
	Update(*Namespace) error
	Get(string) (*Namespace, error)
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
		if err := bkt.Put([]byte(k), []byte(v)); err != nil {
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
func (b *box) Purge(namespace string) error {
	defer b.close(b.open())

	return b.database.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(namespace))
	})
}

// Set will set the contents of ns, eliminating or overwriting anything that
// existed in that namespace before.
func (b *box) Set(ns *Namespace) error {
	defer b.close(b.open())

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
func (b *box) Update(ns *Namespace) error {
	defer b.close(b.open())

	return b.database.Update(func(tx *bbolt.Tx) error {
		bkt, err := bucket(true, tx, ns.Name)
		if err != nil {
			return err
		}
		return put(bkt, ns)
	})
}

func duplicate(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// Get will return the contents of namespace.
func (b *box) Get(namespace string) (*Namespace, error) {
	defer b.close(b.open())

	content := make(map[string]Encrypted)
	if err := b.database.View(func(tx *bbolt.Tx) error {
		bkt, err := bucket(false, tx, namespace)
		if err != nil {
			return err
		}

		if err := bkt.ForEach(func(k []byte, v []byte) error {
			content[string(k)] = Encrypted(duplicate(v))
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

// List will return a list of namespaces that have been created.
func (b *box) List() ([]string, error) {
	defer b.close(b.open())

	var namespaces []string
	if err := b.database.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(ns []byte, _ *bbolt.Bucket) error {
			namespaces = append(namespaces, string(ns))
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return namespaces, nil
}
