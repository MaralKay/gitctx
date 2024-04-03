package main

import "fmt"

func showHelp() {
	helpMessage := `
USAGE:
  gitctx list           [-v]  : list the contexts
  gitctx <NAME>         [-v]  : switch to context <NAME>
  gitctx show                 : show current context
  gitctx add            [-v]  : add a new context
  gitctx delete <NAME>  [-v]  : delete context called <NAME>
  gitctx migrate              : migrate the current-context file from v1.0.x to v1.1.0 format
  gitctx version              : show cli version


FLAGS:
  -v                      : show verbose output


  gitctx -h               : show this message
`
	fmt.Println(helpMessage)
}
