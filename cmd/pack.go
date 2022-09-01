package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"

	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

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

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]
	method := cmd.Flag("method").Value.String()
	var encoder compression.Encoder
	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown pack method")
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
	packed := encoder.Encode(string(data))
	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}
}

// convert /path/to/file.txt -> file.vlc
func packedFileName(path string) string {
	fileName := filepath.Base(path)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return baseName + "." + packedExtension
}
