package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunk string
type BinaryChunks []BinaryChunk

const chunksSize = 8

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))
	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))
	for _, chunk := range bcs {
		res = append(res, chunk.Byte())
	}
	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("can not parse binary chunk " + err.Error())
	}
	return byte(num)
}

// Joins chunk into one line
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}
	return buf.String()
}

// split binary by chunks (8)  bits to bytes -> '1010010 10101111'
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	var buf strings.Builder
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount++
	}
	res := make(BinaryChunks, 0, chunksCount)
	for i, chunk := range bStr {
		buf.WriteString(string(chunk))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}
