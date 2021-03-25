package main

import (
	"log"
	"os"

	"github.com/nashiox/github-billing-exporter/cmd/github-billing-exporter/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rootCmd := cmd.GetRootCmd(os.Args[1:])

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
