package safe

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gophers.dev/cmds/envy/internal/output"
	"gophers.dev/pkgs/ignore"
	"gophers.dev/pkgs/secrets"
)

func TestSafe_Path(t *testing.T) {
	t.Parallel()

	w := output.New(os.Stdout, os.Stdout)
	p, err := Path(w)
	require.NoError(t, err)
	require.True(t, strings.HasSuffix(p, "/envy/envy.safe"))
}

func newFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "-envoy.safe")
	require.NoError(t, err)
	defer ignore.Close(f)
	return f.Name()
}

func TestSafe_operations(t *testing.T) {
	t.Parallel()

	box, err := New(newFile(t))
	require.NoError(t, err)

	_, err = box.Get("does-not-exist")
	require.EqualError(t, err, "namespace \"does-not-exist\" does not exist")

	err = box.Set(&Namespace{
		Name: "ns1",
		Content: map[string]secrets.Text{
			"key1": secrets.New("value1"),
			"key2": secrets.New("value2"),
		},
	})
	require.NoError(t, err)

	err = box.Set(&Namespace{
		Name: "ns2",
		Content: map[string]secrets.Text{
			"keyA": secrets.New("foo"),
			"keyB": secrets.New("bar"),
		},
	})
	require.NoError(t, err)

	ns1, err := box.Get("ns1")
	require.NoError(t, err)
	require.Equal(t, &Namespace{
		Name: "ns1",
		Content: map[string]secrets.Text{
			"key1": secrets.New("value1"),
			"key2": secrets.New("value2"),
		},
	}, ns1)

	ns2, err := box.Get("ns2")
	require.NoError(t, err)
	require.Equal(t, &Namespace{
		Name: "ns2",
		Content: map[string]secrets.Text{
			"keyA": secrets.New("foo"),
			"keyB": secrets.New("bar"),
		},
	}, ns2)

	err = box.Close()
	require.NoError(t, err)
}
