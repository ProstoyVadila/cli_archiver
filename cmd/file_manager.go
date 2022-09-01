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

var ErrEmptyPath = errors.New("path to file is not specified")
var ErrUnknownMethod = errors.New("unknown method")

type CompressionMethod struct {
	name           string
	encoderDecoder compression.EncoderDecoder
	extension      FileExtension
}
type FileExtension struct {
	packed   string
	unpacked string
}

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
	compressionMethod, err := getCompressionMethod(method)
	if err != nil {
		return err
	}

	var performedData []byte
	if isPacking {
		performedData = compressionMethod.encoderDecoder.Encode(string(data))
	} else {
		performedData = []byte(compressionMethod.encoderDecoder.Decode(data))
	}
	newFileName := archivedFileName(filePath, isPacking, compressionMethod.extension)
	err = writeFile(newFileName, performedData, permissionLevel)
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
func archivedFileName(path string, isPacking bool, fileExtension FileExtension) string {
	var extension string
	if isPacking {
		extension = fileExtension.packed
	} else {
		extension = fileExtension.unpacked
	}
	fileName := filepath.Base(path)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return baseName + "." + extension
}

func getCompressionMethod(method string) (CompressionMethod, error) {
	switch method {
	case "vlc":
		fileExtension := FileExtension{packed: method, unpacked: "txt"}
		archiveMethod := CompressionMethod{name: method, encoderDecoder: vlc.New(), extension: fileExtension}
		return archiveMethod, nil
	default:
		return CompressionMethod{}, ErrUnknownMethod
	}
}
