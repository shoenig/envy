package setup

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
	"gophers.dev/cmds/envy/internal/output"
)

func init() {
	// For tests only, use the mock implementation of the keyring provider.
	keyring.MockInit()
}

func newDBFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "tool-")
	require.NoError(t, err)
	err = f.Close()
	require.NoError(t, err)
	return f.Name()
}

func cleanupFile(t *testing.T, name string) {
	err := os.Remove(name)
	require.NoError(t, err)
}

func TestTool_New(t *testing.T) {
	

	db := newDBFile(t)
	defer cleanupFile(t, db)

	tool := New(db, output.New(os.Stdout, os.Stdout))
	require.NotNil(t, tool.Box)
	require.NotNil(t, tool.Ring)
	require.NotNil(t, tool.Writer)
}
