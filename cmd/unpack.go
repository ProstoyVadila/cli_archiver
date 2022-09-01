package cmd

import (
	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "compression")
	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}

func unpack(cmd *cobra.Command, args []string) {
	err := manageFile(cmd, args, false)
	if err != nil {
		if err == ErrUnknownMethod {
			cmd.PrintErr(err)
		}
		handleErr(err)
	}
}
