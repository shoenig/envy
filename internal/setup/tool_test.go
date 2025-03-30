// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package setup

import (
	"os"
	"testing"

	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/test/must"
	"github.com/shoenig/test/util"
	"github.com/zalando/go-keyring"
)

func init() {
	// For tests only, use the mock implementation of the keyring provider.
	keyring.MockInit()
}

func TestTool_New(t *testing.T) {
	db := util.TempFile(t)

	tool := New(db, output.New(os.Stdout, os.Stdout))
	must.NotNil(t, tool.Box)
	must.NotNil(t, tool.Ring)
	must.NotNil(t, tool.Writer)
}
