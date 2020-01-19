package v2

import "bytes"

var Identifier = []uint8("PESQ")

var ValidIdentifiers = [][]uint8{
	[]uint8("PESQ"),
	[]uint8("FSEQ"),
}

func ValidIdentifier(id [4]uint8) bool {
	for _, v := range ValidIdentifiers {
		if bytes.Equal(v, id[:]) {
			return true
		}
	}
	return false
}

const (
	HeaderLength = 32
	MinorVersion = 0
	MajorVersion = 2
)
