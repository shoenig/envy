package commands

import (
	"flag"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
	"gophers.dev/cmds/envy/internal/keyring"
	"gophers.dev/cmds/envy/internal/safe"
	"gophers.dev/pkgs/regexplus"
	"gophers.dev/pkgs/secrets"
)

var (
	argRe       = regexp.MustCompile(`^(?P<key>[\w]+)=(?P<secret>.+)$`)
	namespaceRe = regexp.MustCompile(`^[-\w]+$`)
)

func checkName(namespace string) error {
	if !namespaceRe.MatchString(namespace) {
		return errors.New("namespace uses non-word characters")
	}
	return nil
}

type Extractor interface {
	Namespace(args []interface{}) (*safe.Namespace, error)
}

type extractor struct {
	ring keyring.Ring
}

func newExtractor(ring keyring.Ring) Extractor {
	return &extractor{
		ring: ring,
	}
}

func (e *extractor) Namespace(args []interface{}) (*safe.Namespace, error) {
	_, namespace, argv, err := extract(args)
	if err != nil {
		return nil, err
	}

	content, err := e.process(argv)
	if err != nil {
		return nil, err
	}

	return &safe.Namespace{
		Name:    namespace,
		Content: content,
	}, nil
}

func extract(args []interface{}) (string, string, []secrets.Text, error) {
	arguments := make([]secrets.Text, 0)
	for _, arg := range args[0].([]string) {
		arguments = append(arguments, secrets.New(arg))
	}

	if len(arguments) < 2 {
		return "", "", nil, errors.New("not enough arguments")
	}

	command := arguments[0].Secret()
	namespace := arguments[1].Secret()

	if err := checkName(namespace); err != nil {
		return "", "", nil, err
	}

	return command, namespace, arguments[2:], nil
}

func (e *extractor) process(args []secrets.Text) (map[string]safe.Encrypted, error) {
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

func (e *extractor) encryptEnvVar(kv secrets.Text) (string, safe.Encrypted, error) {
	m := regexplus.FindNamedSubmatches(argRe, kv.Secret())
	if len(m) == 2 {
		s := e.encrypt(secrets.New(m["secret"]))
		return m["key"], s, nil
	}
	return "", nil, errors.New("malformed environment variable pair")
}

func (e *extractor) encrypt(s secrets.Text) safe.Encrypted {
	return e.ring.Encrypt(s)
}

func fsBool(fs *flag.FlagSet, name string) bool {
	b, err := strconv.ParseBool(fs.Lookup(name).Value.String())
	if err != nil {
		return false
	}
	return b
}
