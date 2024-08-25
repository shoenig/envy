// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"github.com/shoenig/envy/internal/setup"
	"noxide.lol/go/babycli"
)

func newListCmd(tool *setup.Tool) *babycli.Component {
	return &babycli.Component{
		Name: "list",
		Help: "list environment profiles",
		Function: func(c *babycli.Component) babycli.Code {
			if c.Nargs() > 0 {
				tool.Writer.Errorf("list command expects no args")
				return babycli.Failure
			}
			profiles, err := tool.Box.List()
			if err != nil {
				tool.Writer.Errorf("unable to list profiles: %v", err)
				return babycli.Failure
			}

			for _, profile := range profiles {
				tool.Writer.Printf("%s", profile)
			}
			return babycli.Success
		},
	}
}
