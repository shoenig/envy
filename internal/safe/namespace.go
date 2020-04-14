package safe

import (
	"fmt"
	"strings"
)

type Encrypted []byte

type Namespace struct {
	Name    string
	Content map[string]Encrypted
}

func (ns *Namespace) String() string {
	keys := make([]string, 0, len(ns.Content))
	for key := range ns.Content {
		keys = append(keys, key)
	}
	return fmt.Sprintf("(%s [%s])", ns.Name, strings.Join(keys, " "))
}
