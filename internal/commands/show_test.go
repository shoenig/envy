// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"testing"

	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/go-conceal"
	"github.com/shoenig/test/must"
)

func TestShowCmd_Execute(t *testing.T) {
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

	box.GetMock.Expect("myNS").Return(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": {1, 1, 1},
			"bar": {2, 2, 2},
		},
	}, nil)

	rc := invoke([]string{"show", "myNS"}, tool)

	must.Zero(t, rc)
	must.Eq(t, "bar\nfoo\n", a.String())
	must.Eq(t, "", b.String())
}

func TestShowCmd_Execute_unveil(t *testing.T) {
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

	box.GetMock.Expect("myNS").Return(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": {1, 1, 1},
			"bar": {2, 2, 2},
		},
	}, nil)

	ring.DecryptMock.When([]byte{2, 2, 2}).Then(conceal.New("passw0rd"))
	ring.DecryptMock.When([]byte{1, 1, 1}).Then(conceal.New("hunter2"))

	rc := invoke([]string{"show", "--unveil", "myNS"}, tool)

	must.Zero(t, rc)
	must.Eq(t, "bar=passw0rd\nfoo=hunter2\n", a.String())
	must.Eq(t, "", b.String())
}

func TestShowCmd_Execute_noNS(t *testing.T) {
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

	rc := invoke([]string{"show"}, tool)

	must.One(t, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: must specify profile and command argument(s)\n", b.String())
}
