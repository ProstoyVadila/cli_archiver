package vlc

// import (
// 	"archiver/lib/compression/vlc/table"
// 	shannonfano "archiver/lib/compression/vlc/table/shannon_fano"
// 	"reflect"
// 	"testing"
// )

// func Test_encodeBin(t *testing.T) {
// 	ed := New(shannonfano.NewGenerator())
// 	tests := []struct {
// 		name string
// 		str  string
// 		tbl  table.EncodingTable
// 		want string
// 	}{
// 		{
// 			name: "base test",
// 			str:  "!ted",
// 			tbl:  ed.tableGenerator.NewTable("!ted"),
// 			want: "001000100110100101",
// 		},
// 		{
// 			name: "second test",
// 			str:  "ass",
// 			tbl:  ed.tableGenerator.NewTable("ass"),
// 			want: "01101010101",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := encodeBin(tt.str, tt.tbl); got != tt.want {
// 				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestEncode(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		str  string
// 		want []byte
// 	}{
// 		{
// 			name: "base test",
// 			str:  "My name is Ted",
// 			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			encoder := New(shannonfano.NewGenerator())
// 			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ToHex() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestDecode(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		et   []byte
// 		want string
// 	}{
// 		{
// 			name: "base test",
// 			et:   []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
// 			want: "My name is Ted",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			decoder := New(shannonfano.NewGenerator())
// 			if got := decoder.Decode(tt.et); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ToHex() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
