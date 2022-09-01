package cmd

import (
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}

func pack(cmd *cobra.Command, args []string) {
	err := manageFile(cmd, args, true)
	if err != nil {
		if err == ErrUnknownMethod {
			cmd.PrintErr(err)
		}
		handleErr(err)
	}
}
