// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"github.com/shoenig/envy/internal/setup"
	"noxide.lol/go/babycli"
)

func newSetCmd(tool *setup.Tool) *babycli.Component {
	return &babycli.Component{
		Name: "set",
		Help: "set environment variable(s) in a profile",
		Function: func(c *babycli.Component) babycli.Code {
			args := c.Arguments()
			extractor := newExtractor(tool.Ring)
			namespace, remove, add, err := extractor.PreProcess(args)
			if err != nil {
				tool.Writer.Errorf("unable to parse args: %v", err)
				return babycli.Failure
			}

			if err = checkName(namespace); err != nil {
				tool.Writer.Errorf("could not set namespace: %v", err)
				return babycli.Failure
			}

			if !remove.Empty() {
				if err := tool.Box.Delete(namespace, remove); err != nil {
					tool.Writer.Errorf("coult not remove variables: %v", err)
					return babycli.Failure
				}
			}

			n, err := extractor.Namespace(add)
			if err != nil {
				tool.Writer.Errorf("unable to parse args: %v", err)
				return babycli.Failure
			}
			n.Name = namespace

			if err = tool.Box.Set(n); err != nil {
				tool.Writer.Errorf("unable to set variables: %v", err)
				return babycli.Failure
			}

			return babycli.Success
		},
	}
}
