envy
====

Use `envy` to manage sensitive environment variables when running commands.

![GitHub](https://img.shields.io/github/license/shoenig/envy.svg)

# Project Overview

`github.com/shoenig/envy` provides a command-line utility for managing
secretive environment variables when running commands.

`envy` builds on ideas from [envchain](https://github.com/sorah/envchain) and [schain](https://github.com/evanphx/schain). It makes use of the [go-keyring](https://github.com/zalando/go-keyring) library for multi-platform keyring management. Encryption is based on Go's built-in `crypto/aes` library. Persistent storage is managed through [boltdb](https://github.com/etcd-io/bbolt).

Supports **Linux**, **macOS**, and **Windows**

# Getting Started

#### Build from source

The `envy` command can be compiled by running
```bash
$ go install github.com/shoenig/envy@latest
```

# Example Usages

#### usage overview
```bash
Subcommands for envy:
	exec             Run command with environment variables from namespace.
	list             List all namespaces.
	purge            Purge a namespace.
	set              Set environment variable(s) for namespace.
	show             Show environment variable(s) in namespace.
	update           Add or Update environment variable(s) in namespace.
```

#### set a namespace
```bash
$ envy set example a=foo b=bar c=baz
stored 3 items in "example"
```

#### execute command
```bash
$ envy exec example hack/test.sh
a: is foo, b is: bar
```

#### list namespaces
```bash
$ envy list
consul/connect-acls:no_tls
example
nomad/e2e
test
```

#### show namespace
```bash
$ envy show test
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
```

#### show namespace w/ values
```bash
$ envy show --decrypt test
AWS_ACCESS_KEY_ID=aaabbbccc
AWS_SECRET_ACCESS_KEY=233kjsdf309jfsd
```

#### update variable in namespace
```bash
$ envy update test AWS_ACCESS_KEY_ID=xxxxyyyyzzz
updated 1 items in "test"
```

#### remove namespace
```bash
$ envy purge test
purged namespace "test"
```

# Contributing

The `github.com/shoenig/envy` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file
an issue.

# LICENSE

The `github.com/shoenig/envy` module is open source under the [MIT](LICENSE) license.
