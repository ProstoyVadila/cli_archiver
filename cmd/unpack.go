package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"

	"io"
	"os"
	"path/filepath"
	"strings"

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

const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]
	method := cmd.Flag("method").Value.String()
	var decoder compression.Decoder
	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("unknown unpack method")
	}

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := decoder.Decode(data) // TODO: change to decode
	packed += "\n"
	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

// convert /path/to/file.txt -> file.vlc
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return baseName + "." + unpackedExtension
}
