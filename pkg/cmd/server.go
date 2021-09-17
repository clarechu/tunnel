package cmd

import "github.com/spf13/cobra"

// GetRootCmd returns the root of the cobra command-tree.
func GetRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "tunnel",
		Short:             "tunnel command  control interface.",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `tunnel configuration command line utility .`,
		PersistentPreRunE: persistentPreRunE,
	}
	return rootCmd
}

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	return nil
}
