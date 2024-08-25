// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package safe

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestNamespace_String(t *testing.T) {
	ns := &Namespace{
		Name: "ns1",
		Content: map[string]Encrypted{
			"foo": []byte{1, 1, 1, 1, 1},
			"bar": []byte{2, 2, 2, 2, 2},
		},
	}
	s := ns.String()
	exp := "(ns1 [bar foo])"
	must.Eq(t, exp, s)
}
