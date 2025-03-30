// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package commands

import (
	"sort"

	"cattlecloud.net/go/babycli"
	"github.com/shoenig/envy/internal/setup"
)

func newShowCmd(tool *setup.Tool) *babycli.Component {
	return &babycli.Component{
		Name: "show",
		Help: "show values in an environment variable profile",
		Flags: babycli.Flags{
			{
				Type:  babycli.BooleanFlag,
				Long:  "unveil",
				Short: "u",
				Help:  "show decrypted values",
				Default: &babycli.Default{
					Value: false,
					Show:  false,
				},
			},
		},
		Function: func(c *babycli.Component) babycli.Code {
			args := c.Arguments()

			if len(args) != 1 {
				tool.Writer.Errorf("must specify profile and command argument(s)")
				return babycli.Failure
			}

			name := args[0]
			p, err := tool.Box.Get(name)
			if err != nil {
				tool.Writer.Errorf("could not read profile: %v", err)
				return babycli.Failure
			}

			keys := make([]string, 0, len(p.Content))
			for k := range p.Content {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for _, key := range keys {
				if c.GetBool("unveil") {
					value := p.Content[key]
					secret := tool.Ring.Decrypt(value)
					tool.Writer.Printf("%s=%s", key, secret.Unveil())
				} else {
					tool.Writer.Printf("%s", key)
				}
			}

			return babycli.Success
		},
	}
}
