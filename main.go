package main

import (
	"fmt"
	"github.com/mdkn/calc/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s; %v", os.Args[0], err)
		os.Exit(-1)
	}
}
