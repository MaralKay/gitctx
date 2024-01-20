package main

import (
	"fmt"
	"os"
)

func listContexts() {
	config, err := readConfig()
	if err != nil {
		fmt.Printf("Error: Unable to list contexts.")
		os.Exit(1)
	}

	for k := range config {
		fmt.Println(k)
	}
}
