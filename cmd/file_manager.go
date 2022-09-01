package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const permissionLevel fs.FileMode = 0644
const packedExtension = "vlc"
const unpackedExtension = "txt"

var ErrEmptyPath = errors.New("path to file is not specified")
var ErrUnknownMethod = errors.New("unknown method")

func manageFile(cmd *cobra.Command, args []string, isPacking bool) error {
	if len(args) == 0 || args[0] == "" {
		return ErrEmptyPath
	}

	filePath := args[0]
	data, err := readFile(filePath)
	if err != nil {
		return err
	}

	method := cmd.Flag("method").Value.String()
	encoderDecoder, err := getCompressionMethod(method)
	if err != nil {
		return err
	}

	var performedData []byte
	if isPacking {
		performedData = encoderDecoder.Encode(string(data))
	} else {
		performedData = []byte(encoderDecoder.Decode(data))
	}
	err = writeFile(packedFileName(filePath, isPacking), performedData, permissionLevel)
	if err != nil {
		return err
	}
	return nil
}

func readFile(filePath string) ([]byte, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return []byte{}, err
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func writeFile(filepath string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(filepath, data, perm)
}

// convert /path/to/file.txt -> file.vlc
func packedFileName(path string, isPacking bool) string {
	var extension string
	if isPacking {
		extension = packedExtension
	} else {
		extension = unpackedExtension
	}
	fileName := filepath.Base(path)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return baseName + "." + extension
}

func getCompressionMethod(method string) (compression.EncoderDecoder, error) {
	switch method {
	case "vlc":
		return vlc.New(), nil
	default:
		return nil, ErrUnknownMethod
	}
}
