// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/test/must"
)

func TestListCmd_ok(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Box:    box,
	}

	box.ListMock.Expect().Return([]string{
		"namespace1", "ns2", "my-ns",
	}, nil)

	// no arguments for list
	rc := invoke([]string{"list"}, tool)

	must.Zero(t, rc)
	must.Eq(t, "namespace1\nns2\nmy-ns\n", a.String())
	must.Eq(t, "", b.String())
}

func TestListCmd_list_fails(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Box:    box,
	}

	box.ListMock.Expect().Return(nil, errors.New("io error"))

	// no arguments for list
	rc := invoke([]string{"list"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to list profiles: io error\n", b.String())
}

func TestListCmd_extra_args(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Box:    box,
	}

	// nonsense args for list
	rc := invoke([]string{"list", "a=b", "c=d"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: list command expects no args\n", b.String())
}
