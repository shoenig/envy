// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/go-conceal"
	"github.com/shoenig/test/must"
)

func TestSetCmd_ok(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	ring.EncryptMock.When(conceal.New("abc123")).Then(safe.Encrypted{8, 8, 8, 8, 8, 8})
	ring.EncryptMock.When(conceal.New("1234")).Then(safe.Encrypted{9, 9, 9, 9})

	box.SetMock.Expect(&safe.Profile{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": {8, 8, 8, 8, 8, 8},
			"bar": {9, 9, 9, 9},
		},
	}).Return(nil)

	tool := &setup.Tool{
		Writer: w,
		Ring:   ring,
		Box:    box,
	}

	rc := invoke([]string{"set", "myNS", "foo=abc123", "bar=1234"}, tool)

	must.Zero(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "", b.String())
}

func TestSetCmd_io_error(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	box.SetMock.Expect(&safe.Profile{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": {8, 8, 8, 8, 8, 8},
			"bar": {9, 9, 9, 9},
		},
	}).Return(errors.New("io error"))

	ring.EncryptMock.When(conceal.New("abc123")).Then(safe.Encrypted{8, 8, 8, 8, 8, 8})
	ring.EncryptMock.When(conceal.New("1234")).Then(safe.Encrypted{9, 9, 9, 9})

	tool := &setup.Tool{
		Writer: w,
		Ring:   ring,
		Box:    box,
	}

	rc := invoke([]string{"set", "myNS", "foo=abc123", "bar=1234"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to set variables: io error\n", b.String())
}

func TestSetCmd_bad_ns(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Ring:   ring,
		Box:    box,
	}

	// e.g. forgot to specify profile
	rc := invoke([]string{"set", "foo=abc123", "bar=1234"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: could not set profile: name uses non-word characters\n", b.String())
}

func TestSetCmd_no_vars(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	tool := &setup.Tool{
		Writer: w,
		Ring:   ring,
		Box:    box,
	}

	// e.g. reminder to use purge to remove profile
	rc := invoke([]string{"set", "ns1"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to parse args: requires at least 2 arguments (profile, <key,...>)\n", b.String())
}
