package vlc

import (
	"archiver/lib/compression/vlc/table"
	"strings"
	"unicode"
)

type EncoderDecoder struct{}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (_ EncoderDecoder) Encode(str string) []byte {
	// prepare text A -> !a
	str = prepareText(str)
	// encode to binary: some text -> 10101010
	bStr := encodeBin(str)
	// split binary by chunks (8)  bits to bytes -> '101001 1010...
	chunks := splitByChunks(bStr, chunksSize)
	return chunks.Bytes()
}

func (_ EncoderDecoder) Decode(data []byte) string {
	// hex chunks -> binary chunks
	// binary chunks -> binary string
	bString := NewBinChunks(data).Join()
	// build decoding tree
	dTree := getEncodingTable().DecodingTree()
	// dTree -> text
	text := dTree.Decode(bString)
	return exportText(text)
}

// My lovely Text -> !my lovely !text
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

// !my lovely !text -> My lovely Text
func exportText(str string) string {
	var buf strings.Builder
	var isCapital bool
	for _, val := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(val))
			isCapital = false
			continue
		}
		if val == '!' {
			isCapital = true
			continue
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

func getEncodingTable() table.EncodingTable {
	return table.EncodingTable{
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
