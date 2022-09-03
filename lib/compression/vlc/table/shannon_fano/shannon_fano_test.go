package shannonfano

import (
	"reflect"
	"testing"
)

func Test_build(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "case 1: empty",
			text: "",
			want: encodingTable{},
		},
		{
			name: "case 2: abbbcc",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Size:     2,
					Bits:     3, // 11b
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Size:     1,
					Bits:     0, // 0b
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Size:     2,
					Bits:     2, // 10b
				},
			},
		},
		{
			name: "case 3: aabbcc",
			text: "aabbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 2,
					Size:     1,
					Bits:     0, // 0b
				},
				'b': code{
					Char:     'b',
					Quantity: 2,
					Size:     2,
					Bits:     2, // 0b
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Size:     2,
					Bits:     3, // 10b
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(newCharStat(tt.text)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "case 1: two elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "case 2: three elements, certain position",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1}, // Bits 0b
				{Quantity: 1, Bits: 2, Size: 2}, // Bits 10b
				{Quantity: 1, Bits: 3, Size: 2}, // Bits 11b
			},
		},
		{
			name: "case 3: three elements, uncertain position",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1}, // Bits 0b
				{Quantity: 1, Bits: 2, Size: 2}, // Bits 10b
				{Quantity: 1, Bits: 3, Size: 2}, // Bits 11b
			},
		},
		{
			name: "case 4: one element",
			codes: []code{
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1},
			},
		},
		{
			name:  "case 5: empty",
			codes: []code{},
			want:  []code{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)
			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("assignCodes() = %v, want %v", tt.codes, tt.want)
			}
		})
	}

}

func Test_bestDividerPosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "case 1",
			codes: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "case 2",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "case 3",
			codes: []code{
				{Quantity: 3},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "case 4",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
		{
			name: "case 5",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "case 6",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDividerPosition(tt.codes); got != tt.want {
				t.Errorf("bestDividerPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_code_addBit(t *testing.T) {
	tests := []struct {
		name string
		code code
		want code
	}{
		{
			name: "case 1",
			code: code{},
			want: code{
				Size: 1,
				Bits: 0,
			},
		},
		{
			name: "case 2",
			code: code{
				Size: 1,
				Bits: 1, // 1b
			},
			want: code{
				Size: 2,
				Bits: 2, // 10b
			},
		},
		{
			name: "case 3",
			code: code{
				Size: 2,
				Bits: 3, // 11b
			},
			want: code{
				Size: 3,
				Bits: 6, // 100b
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.code.addBit()
			if !reflect.DeepEqual(tt.code, tt.want) {
				t.Errorf(
					"code.addBit() changes to Bits:%v Size:%v, want Bits:%v Size:%v",
					tt.code.Bits,
					tt.code.Size,
					tt.want.Bits,
					tt.want.Size,
				)
			}
		})
	}
}
