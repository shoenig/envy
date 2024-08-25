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
			namespaces, err := tool.Box.List()
			if err != nil {
				tool.Writer.Errorf("unable to list namespaces: %v", err)
				return babycli.Failure
			}

			for _, ns := range namespaces {
				tool.Writer.Printf("%s", ns)
			}
			return babycli.Success
		},
	}
}
