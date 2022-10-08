package safe

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/shoenig/test/must"
	"github.com/shoenig/ignore"
)

var _ Box = (*box)(nil)

func TestSafe_Path(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		p, err := Path("")
		must.NoError(t, err)
		must.True(t, strings.HasSuffix(p, "/envy/envy.safe"))
	})

	t.Run("non-default", func(t *testing.T) {
		p, err := Path("/my/custom/path")
		must.NoError(t, err)
		must.Eq(t, "/my/custom/path", p)
	})
}

func newFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "-envoy.safe")
	must.NoError(t, err)
	defer ignore.Close(f)
	return f.Name()
}

func TestSafe_Set(t *testing.T) {
	b := New(newFile(t))

	_, err := b.Get("does-not-exist")
	must.EqError(t, err, "namespace \"does-not-exist\" does not exist")

	// set ns1 first time
	err = b.Set(&Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	})
	must.NoError(t, err)

	// set ns2 first time
	err = b.Set(&Namespace{
		Name: "ns2",
		Content: map[string]Encrypted{
			"keyA": []byte("foo"),
			"keyB": []byte("bar"),
		},
	})
	must.NoError(t, err)

	ns1, err := b.Get("ns1")
	must.NoError(t, err)
	must.Eq(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	}, ns1)

	ns2, err := b.Get("ns2")
	must.NoError(t, err)
	must.Eq(t, &Namespace{
		Name: "ns2",
		Content: map[string]Encrypted{
			"keyA": []byte("foo"),
			"keyB": []byte("bar"),
		},
	}, ns2)

	// set ns2 second time, ensure total replacement
	err = b.Set(&Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key2": []byte("value3"),
			"key3": []byte("value4"),
		},
	})

	ns1, err = b.Get("ns1")
	must.NoError(t, err)
	must.Eq(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key2": []byte("value3"),
			"key3": []byte("value4"),
		},
	}, ns1)
}

func TestSafe_Purge(t *testing.T) {
	b := New(newFile(t))

	// set ns1
	err := b.Set(&Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	})
	must.NoError(t, err)

	// ensure ns1 is set
	ns1, err := b.Get("ns1")
	must.NoError(t, err)
	must.Eq(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	}, ns1)

	// purge ns1
	err = b.Purge("ns1")
	must.NoError(t, err)

	// ensure ns1 is not set anymore
	_, err = b.Get("ns1")
	must.EqError(t, err, `namespace "ns1" does not exist`)
}

func TestSafe_Update(t *testing.T) {
	b := New(newFile(t))

	// set ns1
	err := b.Set(&Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	})
	must.NoError(t, err)

	// ensure ns1 is set
	ns1, err := b.Get("ns1")
	must.NoError(t, err)
	must.Eq(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	}, ns1)

	// update ns1
	err = b.Update(&Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key2": []byte("value2"),
			"key3": []byte("value3"),
		},
	})
	must.NoError(t, err)

	// ensure ns1 is joined union
	ns1, err = b.Get("ns1")
	must.NoError(t, err)
	must.Eq(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
			"key3": []byte("value3"),
		},
	}, ns1)
}
