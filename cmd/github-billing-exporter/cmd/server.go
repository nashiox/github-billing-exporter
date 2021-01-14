package cmd

import (
	"log"

	"github.com/nashiox/github-billing-exporter/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/xerrors"
)

func serverCmd() *cobra.Command {
	var (
		serverArgs = &server.Args{}
	)

	serverCmd := &cobra.Command{
		Use:          "server",
		Short:        "Starts GitHubBillingExporter as a server",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return xerrors.Errorf("%q is an invalid argument", args[0])
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Run(serverArgs)
		},
	}

	serverCmd.PersistentFlags().IntVarP(
		&serverArgs.Port,
		"port",
		"p",
		9999,
		"Exporter Listen Port",
	)
	serverCmd.PersistentFlags().IntVarP(
		&serverArgs.Refresh,
		"refresh",
		"r",
		300,
		"Refresh Interval Secounds",
	)
	serverCmd.PersistentFlags().StringVarP(
		&serverArgs.Organization,
		"organization",
		"o",
		"",
		"GitHub Organization Name",
	)
	serverCmd.PersistentFlags().StringVarP(
		&serverArgs.User,
		"user",
		"u",
		"",
		"GitHub User Name",
	)
	serverCmd.PersistentFlags().StringVarP(
		&serverArgs.Token,
		"token",
		"t",
		"",
		"GitHub Token",
	)

	if err := viper.BindPFlags(serverCmd.PersistentFlags()); err != nil {
		log.Fatalf("Failed to bind flags: %v\n", err)
	}

	cobra.OnInitialize(func() {
		viper.AutomaticEnv()

		if err := viper.Unmarshal(&serverArgs); err != nil {
			log.Fatalf("Failed to unmarshal arguments: %v\n", err)
		}
	})

	return serverCmd
}
