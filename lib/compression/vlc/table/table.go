package table

import (
	"strings"
)

type Generator interface {
	NewTable(text string) EncodingTable
}

type DecodingTree struct {
	Value string
	Zero  *DecodingTree
	One   *DecodingTree
}

type EncodingTable map[rune]string

func (et EncodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}

	for runeKey, code := range et {
		res.add(code, runeKey)
	}
	return res
}

func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder
	currentNode := dt

	writeBufIfValue := func(currNode *DecodingTree, buf *strings.Builder) *DecodingTree {
		if currNode.Value != "" {
			buf.WriteString(currNode.Value)
			currNode = dt
		}
		return currNode
	}

	for _, char := range str {
		currentNode = writeBufIfValue(currentNode, &buf)
		switch char {
		case '0':
			currentNode = currentNode.Zero
		case '1':
			currentNode = currentNode.One
		}
	}
	_ = writeBufIfValue(currentNode, &buf)
	return buf.String()
}

func (dt *DecodingTree) add(code string, value rune) {
	currentNode := dt
	for _, char := range code {
		switch char {
		case '0':
			if currentNode.Zero == nil {
				currentNode.Zero = &DecodingTree{}
			}
			currentNode = currentNode.Zero
		case '1':
			if currentNode.One == nil {
				currentNode.One = &DecodingTree{}
			}
			currentNode = currentNode.One
		}
	}
	currentNode.Value = string(value)
}
