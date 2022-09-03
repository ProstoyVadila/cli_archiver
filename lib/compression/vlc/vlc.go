package vlc

import (
	"archiver/lib/compression/vlc/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"
)

const (
	encodedTableSizeBytesCount = 4
	encodedDataSizeBytesCount  = 4
)

type EncoderDecoder struct {
	tableGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tableGenerator: tblGenerator}
}

func (ed EncoderDecoder) Encode(str string) []byte {
	tbl := ed.tableGenerator.NewTable(str)
	bStr := encodeBin(str, tbl)
	return buildEncodedFile(tbl, bStr)
}

func (ed EncoderDecoder) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)
	return tbl.Decode(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	tableSizeBinary, data := data[:encodedTableSizeBytesCount], data[encodedTableSizeBytesCount:]
	dataSizeBinary, data := data[:encodedDataSizeBytesCount], data[encodedDataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]
	tbl := decodeTable(tblBinary)
	body := NewBinChunks(data).Join()
	return tbl, body[:dataSize]
}

func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	var buff bytes.Buffer
	encodedTable := encodeTable(tbl)

	buff.Write(encodeInt(len(encodedTable)))
	buff.Write(encodeInt(len(data)))
	buff.Write(encodedTable)
	buff.Write([]byte(splitByChunks(data, chunksSize).Bytes()))
	return buff.Bytes()
}

func encodeInt(num int) []byte {
	res := make([]byte, encodedTableSizeBytesCount)
	binary.BigEndian.PutUint32(res, uint32(num))
	return res
}

func decodeTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable
	r := bytes.NewReader((tblBinary))
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal("can not serialize encode table from encdoded file", err)
	}
	return tbl
}

func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer
	if err := gob.NewEncoder(&tableBuf).Encode(tbl); err != nil {
		log.Fatal("can not serialize encode table to encoded file", err)
	}
	return tableBuf.Bytes()
}

// str into binary str
func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, r := range str {
		buf.WriteString(bin(r, table))
	}
	return buf.String()
}

func bin(r rune, table table.EncodingTable) string {
	res, ok := table[r]
	if !ok {
		panic("unknown character: " + string(r))
	}
	return res
}

// func getEncodingTable() table.EncodingTable {
// 	return table.EncodingTable{
// 		' ': "11",
// 		't': "1001",
// 		'n': "10000",
// 		's': "0101",
// 		'r': "01000",
// 		'd': "00101",
// 		'!': "001000",
// 		'c': "000101",
// 		'm': "000011",
// 		'g': "0000100",
// 		'b': "0000010",
// 		'v': "00000001",
// 		'k': "0000000001",
// 		'q': "000000000001",
// 		'e': "101",
// 		'o': "10001",
// 		'a': "011",
// 		'i': "01001",
// 		'h': "0011",
// 		'l': "001001",
// 		'u': "00011",
// 		'f': "000100",
// 		'p': "0000101",
// 		'w': "0000011",
// 		'y': "0000001",
// 		'j': "000000001",
// 		'x': "00000000001",
// 		'z': "000000000000",
// 	}
// }
