package main

import "fmt"

func showHelp() {
	helpMessage := `
USAGE:
  gitctx list   [-v]      : list the contexts
  gitctx <NAME> [-v]      : switch to context <NAME>
  gitctx show             : show current context
  gitctx add    [-v]      : add a new context


FLAGS:
  -v                      : show verbose output


  gitctx -h               : show this message
`
	fmt.Println(helpMessage)
}
