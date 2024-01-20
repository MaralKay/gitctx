package main

import "fmt"

func showHelp() {
	helpMessage := `
USAGE:
  gitctx list             : list the contexts
  gitctx <NAME>           : switch to context <NAME>
  gitctx show             : show current context
  gitctx add              : add a new context


  gitctx -h               : show this message
`
	fmt.Println(helpMessage)
}
