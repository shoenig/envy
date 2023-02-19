envy
====

Use `envy` to manage sensitive environment variables when running commands.

![GitHub](https://img.shields.io/github/license/shoenig/envy.svg)

# Project Overview

`github.com/shoenig/envy` provides a CLI utility for running commands with secret
environment variables like `GITHUB_TOKEN`, etc.

`envy` builds on ideas from [envchain](https://github.com/sorah/envchain) and [schain](https://github.com/evanphx/schain).
It makes use of the [go-keyring](https://github.com/zalando/go-keyring) library for multi-platform keyring management.
Encryption is based on Go's built-in `crypto/aes` library.
Persistent storage is managed through [boltdb](https://github.com/etcd-io/bbolt).

Supports **Linux**, **macOS**, and **Windows**

# Getting Started

#### Install

The `envy` command is available to download from the [Releases](https://github.com/shoenig/envy/releases) page.

Multiple operating systems and architectures are supported, including

- Linux
- macOS
- Windows

#### Install from Go module

The `envy` command can be installed from source by running

```bash
$ go install github.com/shoenig/envy@latest
```

# Example Usages

#### usage overview

```bash
Subcommands for envy:
	exec             Run a command with environment variables from namespace.
	  use `-insulate` to exclude external environment variables
	list             List available namespaces.
	purge            Remove a namespace and all of its environment variables.
	set              Set/Update/Remove environment variable(s) for namespace.
	show             Show environment variable(s) in namespace.
	  use `-decrypt` to include decrypted values
```

#### set a namespace

```bash
$ envy set example FOO=1 BAR=2 BAZ=3
```

#### update existing variable in a namespace

```bash
$ envy set example FOO=4
```

#### remove variable from namespace

```bash
$ envy set example -FOO
```

#### execute command

```bash
$ envy exec example env
BAR=2
BAZ=3
... <many more from user> ...
```

#### execute command excluding external environment

```bash
$ envy exec -insulate example env
BAR=2
BAZ=3
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
