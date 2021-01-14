package cmd

import "github.com/spf13/cobra"

func GetRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "github-billing-exporter",
		Short:        "GitHubBillingExporter is a PrometheusExporter that collects billing reports on GitHub.",
		SilenceUsage: true,
	}

	rootCmd.SetArgs(args)
	rootCmd.AddCommand(serverCmd())

	return rootCmd
}
