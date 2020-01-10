package fseq

const (
	majorVersionIndex = 7
	minorVersionIndex = 6
)

func ReadVersion(b []byte) (minorVersion, majorVersion uint8) {
	_ = b[majorVersionIndex] // early bounds check to guarantee safety of reads below
	return b[minorVersionIndex], b[majorVersionIndex]
}
