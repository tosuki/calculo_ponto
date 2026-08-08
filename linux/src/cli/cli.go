package cli

import (
	"github.com/4mti/ponto/src/cli/cmd/overlay"
	"github.com/spf13/cobra"
)

func RunCli() error {
	rootCmd := &cobra.Command{
		Use:   "ponto",
		Short: "Um aplicativo pra facilitar a vida de trabalhadores com horas pra pagar.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(overlay.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
