package setup

import (
	"github.com/shoenig/test/must"
	"io/ioutil"
	"os"
	"testing"

	"github.com/zalando/go-keyring"
	"gophers.dev/cmds/envy/internal/output"
)

func init() {
	// For tests only, use the mock implementation of the keyring provider.
	keyring.MockInit()
}

func newDBFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "tool-")
	must.NoError(t, err)
	err = f.Close()
	must.NoError(t, err)
	return f.Name()
}

func cleanupFile(t *testing.T, name string) {
	err := os.Remove(name)
	must.NoError(t, err)
}

func TestTool_New(t *testing.T) {
	db := newDBFile(t)
	defer cleanupFile(t, db)

	tool := New(db, output.New(os.Stdout, os.Stdout))
	must.NotNil(t, tool.Box)
	must.NotNil(t, tool.Ring)
	must.NotNil(t, tool.Writer)
}
