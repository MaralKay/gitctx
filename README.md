# gitctx

`gitctx` is a simple command line tool written in Golang, designed to facilitate the management and switching of multiple git configurations, which we refer to as "contexts." Inspired by `kubectx`, this tool works by storing mappings of git config files and ssh config files.

## Features

- **Context Switching**: Easily switch between different git configurations and ssh configurations.
- **Mapping Support**: Store and manage multiple configurations by creating mappings to git config files and ssh config files.
- **Inspired by kubectx**: If you're familiar with `kubectx`, `gitctx` operates in a similar fashion for git configurations.

## Installation

To install `gitctx`, you can use the following command:

```bash
# Assuming you have Go installed
go get -u github.com/MaralKay/gitctx
```

Make sure that your Go binary directory is in your system's PATH.

## Usage
### Switching Contexts
To switch to a different git configuration context, use the following command:

```bash
gitctx <context-name>
```

### Show Current Context
To show the current git context:
```bash
gitctx show
```

### Adding a Context
To add a new git configuration context, use:

```bash
gitctx add
```
You will be prompted to provide the paths of the configs for a new mapping.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
Inspired by the excellent kubectx tool.
Feel free to contribute, report issues, or provide feedback!
