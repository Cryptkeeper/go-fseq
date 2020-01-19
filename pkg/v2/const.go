package v2

var Identifier = []uint8("PESQ")

var ValidIdentifiers = [][]uint8{
	[]uint8("PESQ"),
	[]uint8("FSEQ"),
}

const (
	HeaderLength = 32
	MinorVersion = 0
	MajorVersion = 2
)
