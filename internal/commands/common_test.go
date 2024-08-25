// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"bytes"

	"github.com/shoenig/envy/internal/output"
	"github.com/zalando/go-keyring"
)

func init() {
	// For tests only, use the mock implementation of the keyring provider.
	keyring.MockInit()
}

func newWriter() (*bytes.Buffer, *bytes.Buffer, output.Writer) {
	var a, b bytes.Buffer
	return &a, &b, output.New(&a, &b)
}
