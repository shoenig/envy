// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MIT

package safe

import (
	"fmt"
	"sort"
	"strings"
)

type Encrypted []byte

type Namespace struct {
	Name    string
	Content map[string]Encrypted
}

func (ns *Namespace) String() string {
	keys := ns.Keys()
	return fmt.Sprintf("(%s [%s])", ns.Name, strings.Join(keys, " "))
}

func (ns *Namespace) Keys() []string {
	keys := make([]string, 0, len(ns.Content))
	for key := range ns.Content {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
