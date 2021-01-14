package main

import (
	"os"

	"github.com/nashiox/github-billing-exporter/cmd/github-billing-exporter/cmd"
)

func main() {
	rootCmd := cmd.GetRootCmd(os.Args[1:])

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
