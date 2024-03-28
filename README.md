# gitctx

`gitctx` is a simple command line tool written in Golang, designed to facilitate the management and switching of multiple git configurations, which we refer to as "contexts." Inspired by `kubectx`, this tool works by storing mappings of git config files and ssh config files.

## Features

- **Context Switching**: Easily switch between different git configurations and ssh configurations.
- **Mapping Support**: Store and manage multiple configurations by creating mappings to git config files and ssh config files.
- **Inspired by kubectx**: If you're familiar with `kubectx`, `gitctx` operates in a similar fashion for git configurations.

## Release Notes

### Upgrade to v1.1.0
In this release the core functionality of the tool has changed, allowing you to keep track of both local repo contexts and global context. The format of the `.gitctx.current` has changed, therefore in order to be able to run `gitctx` without issues you need to run as a first step after upgrading your binary package the following command:
```
gitctx migrate
```


## Installation

To install `gitctx`, you can find the binary package for your OS distribution in the latest releases.

## Usage
```bash
USAGE:
  gitctx list      [-v]      : list the contexts
  gitctx <NAME>    [-v]      : switch to context <NAME>
  gitctx show                : show current context
  gitctx add       [-v]      : add a new context
  gitctx migrate             : migrate the current-context file to v1.1.0 format


FLAGS:
  -v                      : show verbose output


  gitctx -h               : show this message
```

### Switching Contexts
To switch to a different git configuration context, use the following command:

```bash
gitctx <context-name>
```
```bash
~ gitctx work

Updated context to work
```

### Show Current Context
To show the current git context:
```bash
~ gitctx show

work
```

### List Contexts
To list the configured contexts:
```bash
~ gitctx list

work
personal
```

### Adding a Context
To add a new git configuration context, use:

```bash
gitctx add
```
You will be prompted to provide the paths of the configs for a new mapping.

## Prerequisites
The program assumes you have different ssh keys for your Github accounts. Following that, you need to save separate ssh config files for each gitconfig file.

#### Example:
##### Github account #1:
A file containing git configuration
`~/.gitconfig-personal`


A file with the ssh config for this `~/.ssh/config-personal-git`
```bash
Host github.com
	HostName github.com
	User git
	IdentityFile <PATH TO SSH KEY USED WITH THIS ACCOUNT>
```
##### Github account #2:
A file containing git configuration `~/.gitconfig-work`

A file with the ssh config for this `~/.ssh/config-work-git`
```bash
Host github.com
	HostName github.com
	User git
	IdentityFile <PATH TO SSH KEY USED WITH THIS ACCOUNT>
```

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
Inspired by the excellent kubectx tool.
Feel free to contribute, report issues, or provide feedback!
