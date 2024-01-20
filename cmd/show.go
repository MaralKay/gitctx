package main

import (
	"fmt"
	"os"
)

func showContext() {

	context, err := os.ReadFile(currentContextPath)
	if err != nil {
		fmt.Printf("Couldn't read current context: %s", err)
		os.Exit(1)
	}

	fmt.Printf("\n%s\n", context)
}
