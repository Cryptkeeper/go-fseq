package fseq

import "bytes"

const (
	majorVersionIndex = 7
	minorVersionIndex = 6
)

func ReadVersion(b []byte) (minorVersion, majorVersion uint8) {
	_ = b[majorVersionIndex] // early bounds check to guarantee safety of reads below
	return b[minorVersionIndex], b[majorVersionIndex]
}

var ValidIdentifiers = [][]uint8{
	[]uint8("PSEQ"),
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