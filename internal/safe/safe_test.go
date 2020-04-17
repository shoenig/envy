package safe

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gophers.dev/pkgs/ignore"
)

var _ Box = (*box)(nil)

func TestSafe_Path(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		p, err := Path("")
		require.NoError(t, err)
		require.True(t, strings.HasSuffix(p, "/envy/envy.safe"))
	})

	t.Run("non-default", func(t *testing.T) {
		p, err := Path("/my/custom/path")
		require.NoError(t, err)
		require.Equal(t, "/my/custom/path", p)
	})
}

func newFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "-envoy.safe")
	require.NoError(t, err)
	defer ignore.Close(f)
	return f.Name()
}

func TestSafe_Set(t *testing.T) {
	b, err := New(newFile(t))
	require.NoError(t, err)

	_, err = b.Get("does-not-exist")
	require.EqualError(t, err, "namespace \"does-not-exist\" does not exist")

	// set ns1 first time
	err = b.Set(&Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	})
	require.NoError(t, err)

	// set ns2 first time
	err = b.Set(&Namespace{
		Name: "ns2",
		Content: map[string]Encrypted{
			"keyA": []byte("foo"),
			"keyB": []byte("bar"),
		},
	})
	require.NoError(t, err)

	ns1, err := b.Get("ns1")
	require.NoError(t, err)
	require.Equal(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	}, ns1)

	ns2, err := b.Get("ns2")
	require.NoError(t, err)
	require.Equal(t, &Namespace{
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
	require.NoError(t, err)
	require.Equal(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key2": []byte("value3"),
			"key3": []byte("value4"),
		},
	}, ns1)

	err = b.(*box).Close()
	require.NoError(t, err)
}

func TestSafe_Purge(t *testing.T) {
	b, err := New(newFile(t))
	require.NoError(t, err)

	// set ns1
	err = b.Set(&Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	})
	require.NoError(t, err)

	// ensure ns1 is set
	ns1, err := b.Get("ns1")
	require.NoError(t, err)
	require.Equal(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	}, ns1)

	// purge ns1
	err = b.Purge("ns1")
	require.NoError(t, err)

	// ensure ns1 is not set anymore
	_, err = b.Get("ns1")
	require.EqualError(t, err, `namespace "ns1" does not exist`)

	err = b.(*box).Close()
	require.NoError(t, err)
}

func TestSafe_Update(t *testing.T) {
	b, err := New(newFile(t))
	require.NoError(t, err)

	// set ns1
	err = b.Set(&Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	})
	require.NoError(t, err)

	// ensure ns1 is set
	ns1, err := b.Get("ns1")
	require.NoError(t, err)
	require.Equal(t, &Namespace{
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
	require.NoError(t, err)

	// ensure ns1 is joined union
	ns1, err = b.Get("ns1")
	require.NoError(t, err)
	require.Equal(t, &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
			"key3": []byte("value3"),
		},
	}, ns1)

	err = b.(*box).Close()
	require.NoError(t, err)
}
