package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunk string
type HexChunk string
type HexChunks []HexChunk
type BinaryChunks []BinaryChunk

const chunksSize = 8
const hexChunksSeparator = " "

func NewHexChunks(str string) HexChunks {
	parts := strings.Split(str, hexChunksSeparator)
	res := make(HexChunks, 0, len(parts))
	for _, part := range parts {
		res = append(res, HexChunk(part))
	}
	return res
}

func (hcs HexChunks) ToString() string {
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}
	var buf strings.Builder
	for i, hc := range hcs {
		buf.WriteString(string(hc))
		if i+1 != len(hcs) {
			buf.WriteString(hexChunksSeparator)
		}
	}
	return buf.String()
}

func (hcs HexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(hcs))
	for _, chunk := range hcs {
		bChunk := chunk.ToBinary()
		res = append(res, bChunk)
	}
	return res
}

func (hc HexChunk) ToBinary() BinaryChunk {
	num, err := strconv.ParseUint(string(hc), 16, chunksSize)
	if err != nil {
		panic("can not parse hex chunk: " + err.Error())
	}
	return BinaryChunk(fmt.Sprintf("%08b", num))
}

// Joins chunk into one line
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, bc := range bcs {
		buf.WriteString(string(bc))
	}
	return buf.String()
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))
	for _, chunk := range bcs {
		HexChunk := chunk.ToHex()
		res = append(res, HexChunk)
	}
	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("cant parse binary chunk: " + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num)) // 2h or 3f or 1
	if len(res) == 1 {
		res = "0" + res
	}
	return HexChunk(res) // 2H or 3F or 01
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
