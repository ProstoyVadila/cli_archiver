package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunk string
type HexChunk string
type HexChunks []HexChunk
type BinaryChunks []BinaryChunk

const chunksSize = 8

func Encode(str string) string {
	// prepare text A -> !a
	str = prepareText(str)
	// encode to binary: some text -> 10101010
	bStr := encodeBin(str)
	// split binary by chunks (8)  bits to bytes -> '101001 1010...
	chunks := splitByChunks(bStr, chunksSize)
	// bytes to hex
	// chunks.ToHex()
	return chunks.ToHex().ToString() // "2F 4A 3F ..."
}

func (hcs HexChunks) ToString() string {
	const sep = " "
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
			buf.WriteString(sep)
		}
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

// My Text -> !my !text
func prepareText(str string) string {
	var buf strings.Builder

	for _, val := range str {
		if unicode.IsUpper(val) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(val))
		} else {
			buf.WriteRune(val)
		}
	}
	return buf.String()
}

// str into binary str
func encodeBin(str string) string {
	var buf strings.Builder

	for _, r := range str {
		buf.WriteString(bin(r))
	}
	return buf.String()
}

func bin(r rune) string {
	table := getEncodingTable()
	res, ok := table[r]
	if !ok {
		panic("unknown character: " + string(r))
	}
	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
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
