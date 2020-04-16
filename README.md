envy
====

Use `envy` to manage sensitive environment variables when running commands.

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/cmds/envy)](https://goreportcard.com/report/gophers.dev/cmds/envy)
[![Build Status](https://travis-ci.org/shoenig/envy.svg?branch=master)](https://travis-ci.org/shoenig/envy)
[![GoDoc](https://godoc.org/gophers.dev/cmds/envy?status.svg)](https://godoc.org/gophers.dev/cmds/envy)
![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/envy.svg)
![GitHub](https://img.shields.io/github/license/shoenig/envy.svg)

# Project Overview

Module `gophers.dev/cmds/envy` provides a command-line utility for managing
secretive environment variables when running commands.

# Getting Started

#### Build from source

The `envy` command can be compiled by running
```bash
$ go get gophers.dev/cmds/envy
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

The `gophers.dev/cmds/envy` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file
an issue.

# LICENSE

The `gophers.dev/cmds/envy` module is open source under the [MIT](LICENSE) license.
