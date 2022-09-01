package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"archiver/lib/vlc"
)

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

var vlcPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := vlc.Encode(string(data))
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

func init() {
	packCmd.AddCommand(vlcPackCmd)
}
