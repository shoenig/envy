// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package main

import (
	"os"

	"cattlecloud.net/go/babycli"
	"github.com/shoenig/envy/internal/commands"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/setup"
)

func main() {
	tool := setup.New(
		os.Getenv("ENVY_DB_FILE"),
		output.New(os.Stdout, os.Stderr),
	)

	args := babycli.Arguments()
	rc := commands.Invoke(args, tool)
	os.Exit(rc)
}
