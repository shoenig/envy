// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"github.com/shoenig/envy/internal/setup"
	"noxide.lol/go/babycli"
)

func newPurgeCmd(tool *setup.Tool) *babycli.Component {
	return &babycli.Component{
		Name: "purge",
		Help: "purge an environment profile",
		Function: func(c *babycli.Component) babycli.Code {
			if c.Nargs() != 1 {
				tool.Writer.Errorf("must specify one profile")
				return babycli.Failure
			}

			args := c.Arguments()
			profile := args[0]

			if err := checkName(profile); err != nil {
				tool.Writer.Errorf("unable to purge profile: %v", err)
				return babycli.Failure
			}

			if err := tool.Box.Purge(profile); err != nil {
				tool.Writer.Errorf("unable to purge profile: %v", err)
				return babycli.Failure
			}

			tool.Writer.Printf("purged profile %q", profile)
			return babycli.Success
		},
	}
}
