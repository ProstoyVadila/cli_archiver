package shannonfano

import (
	"archiver/lib/compression/vlc/table"
	"fmt"
	"math"
	"sort"
	"strings"
)

type CharStat map[rune]int
type encodingTable map[rune]code

type Generator struct{}
type code struct {
	Char     rune
	Quantity int
	Size     int
	Bits     uint32
}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) table.EncodingTable {
	//char quantity stats
	stats := newCharStat(text)
	//encoding table and change from map[rune]code to map[rune]string
	return build(stats).Export()
}

func (et encodingTable) Export() map[rune]string {
	res := make(map[rune]string)
	for k, v := range et {
		byteStr := fmt.Sprintf("%b", v.Bits)

		if lenDiff := v.Size - len(byteStr); lenDiff > 0 {
			byteStr = strings.Repeat("0", lenDiff) + byteStr
		}

		res[k] = byteStr
	}
	return res
}

func (c *code) addBit() {
	c.Bits <<= 1 // 00000111 -> 000011111
	c.Size++     // size 3 -> 4
}

func newCharStat(text string) CharStat {
	res := make(CharStat)
	for _, char := range text {
		res[char]++
	}
	return res
}

func build(stats CharStat) encodingTable {
	codes := make([]code, 0, len(stats))
	for char, qty := range stats {
		codes = append(codes, code{
			Char:     char,
			Quantity: qty,
		})
	}
	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity
		}
		return codes[i].Char < codes[j].Char
	})

	assignCodes(codes)
	res := make(encodingTable)
	for _, code := range codes {
		res[code.Char] = code
	}
	return res
}

func assignCodes(codes []code) {
	// TODO: fix case with len(codes) == 1
	if len(codes) == 0 {
		return
	}
	if len(codes) == 1 {
		if codes[0].Size == 0 {
			codes[0].addBit()
		}
		return
	}
	divider := bestDividerPosition(codes)
	for i := 0; i < len(codes); i++ {
		codes[i].addBit()
		if i >= divider {
			codes[i].Bits |= 1 // from left to right if 1 then 1 and if first is 0 then 1 too
		}
	}
	assignCodes(codes[:divider])
	assignCodes(codes[divider:])

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func bestDividerPosition(codes []code) int {
	total := 0
	for _, code := range codes {
		total += code.Quantity
	}

	left := 0
	prevDiff := math.MaxInt
	bestPosition := 0

	for i := 0; i < len(codes)-1; i++ {
		left += codes[0].Quantity

		right := total - left

		diff := abs(right - left)
		if diff >= prevDiff {
			break
		}
		prevDiff = diff
		bestPosition = i + 1
	}
	return bestPosition
}
