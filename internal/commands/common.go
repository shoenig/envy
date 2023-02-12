package commands

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set"
	"github.com/pkg/errors"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/go-conceal"
	"github.com/shoenig/regexplus"
)

var (
	argRe       = regexp.MustCompile(`^(?P<key>\w+)=(?P<secret>.+)$`)
	namespaceRe = regexp.MustCompile(`^[-:/\w]+$`)
)

func checkName(namespace string) error {
	if !namespaceRe.MatchString(namespace) {
		return errors.New("namespace uses non-word characters")
	}
	return nil
}

type Extractor interface {
	PreProcess(args []string) (string, *set.Set[string], *set.HashSet[*conceal.Text, int], error)
	Namespace(vars *set.HashSet[*conceal.Text, int]) (*safe.Namespace, error)
}

type extractor struct {
	ring keyring.Ring
}

func newExtractor(ring keyring.Ring) Extractor {
	return &extractor{
		ring: ring,
	}
}

// PreProcess returns
// - the namespace
// - the set of keys to be removed
// - the set of key/values to be added
// - any error
func (e *extractor) PreProcess(args []string) (string, *set.Set[string], *set.HashSet[*conceal.Text, int], error) {
	if len(args) < 2 {
		return "", nil, nil, fmt.Errorf("requires at least 2 arguments (namespace, <key,...>)")
	}
	ns := args[0]
	rm := set.New[string](4)
	add := set.NewHashSet[*conceal.Text, int](8)
	for i := 1; i < len(args); i++ {
		s := args[i]
		switch {
		case strings.HasPrefix(s, "-"):
			rm.Insert(strings.TrimPrefix(s, "-"))
		case strings.Contains(s, "="):
			add.Insert(conceal.New(s))
		default:
			return "", nil, nil, fmt.Errorf("argument must start with '-' or contain '='")
		}
	}
	return ns, rm, add, nil
}

func (e *extractor) Namespace(vars *set.HashSet[*conceal.Text, int]) (*safe.Namespace, error) {
	content, err := e.process(vars.Slice())
	if err != nil {
		return nil, err
	}
	return &safe.Namespace{
		Name:    "",
		Content: content,
	}, nil
}

func (e *extractor) process(args []*conceal.Text) (map[string]safe.Encrypted, error) {
	content := make(map[string]safe.Encrypted, len(args))
	for _, kv := range args {
		if key, secret, err := e.encryptEnvVar(kv); err != nil {
			return nil, err
		} else {
			content[key] = secret
		}
	}
	return content, nil
}

func (e *extractor) encryptEnvVar(kv *conceal.Text) (string, safe.Encrypted, error) {
	m := regexplus.FindNamedSubmatches(argRe, kv.Unveil())
	if len(m) == 2 {
		s := e.encrypt(conceal.New(m["secret"]))
		return m["key"], s, nil
	}
	return "", nil, errors.New("malformed environment variable pair")
}

func (e *extractor) encrypt(s *conceal.Text) safe.Encrypted {
	return e.ring.Encrypt(s)
}

func fsBool(fs *flag.FlagSet, name string) bool {
	b, err := strconv.ParseBool(fs.Lookup(name).Value.String())
	if err != nil {
		return false
	}
	return b
}
