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

func TestPurgeCmd_ok(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Box:    box,
	}

	box.PurgeMock.Expect("myNS").Return(nil)

	rc := invoke([]string{"purge", "myNS"}, tool)

	must.Zero(t, rc)
	must.Eq(t, "purged profile \"myNS\"\n", a.String())
	must.Eq(t, "", b.String())
}

func TestPurgeCmd_fails(t *testing.T) {
	box := safe.NewBoxMock(t)
	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Box:    box,
	}

	box.PurgeMock.Expect("myNS").Return(errors.New("does not exist"))

	rc := invoke([]string{"purge", "myNS"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to purge profile: does not exist\n", b.String())
}

func TestPurgeCmd_no_arg(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Box:    box,
	}

	rc := invoke([]string{"purge"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: must specify one profile\n", b.String())
}

func TestPurgeCmd_bad_profile(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Box:    box,
	}

	// namespace must be valid
	rc := invoke([]string{"purge", "foo=bar"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to purge profile: namespace uses non-word characters\n", b.String())
}

func TestPurgeCmd_two_arg(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Box:    box,
	}

	rc := invoke([]string{"purge", "ns1", "ns2"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: must specify one profile\n", b.String())
}
