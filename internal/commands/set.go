// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"cattlecloud.net/go/babycli"
	"github.com/shoenig/envy/internal/setup"
)

func newSetCmd(tool *setup.Tool) *babycli.Component {
	return &babycli.Component{
		Name: "set",
		Help: "set environment variable(s) in a profile",
		Function: func(c *babycli.Component) babycli.Code {
			args := c.Arguments()
			extractor := newExtractor(tool.Ring)
			profile, remove, add, err := extractor.Process(args)
			if err != nil {
				tool.Writer.Errorf("unable to parse args: %v", err)
				return babycli.Failure
			}

			if err = checkName(profile); err != nil {
				tool.Writer.Errorf("could not set profile: %v", err)
				return babycli.Failure
			}

			if !remove.Empty() {
				if err := tool.Box.Delete(profile, remove); err != nil {
					tool.Writer.Errorf("coult not remove variables: %v", err)
					return babycli.Failure
				}
			}

			n, err := extractor.Profile(add)
			if err != nil {
				tool.Writer.Errorf("unable to parse args: %v", err)
				return babycli.Failure
			}
			n.Name = profile

			if err = tool.Box.Set(n); err != nil {
				tool.Writer.Errorf("unable to set variables: %v", err)
				return babycli.Failure
			}

			return babycli.Success
		},
	}
}
