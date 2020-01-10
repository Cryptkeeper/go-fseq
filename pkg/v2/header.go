package v2

type Header struct {
	Identifier             [4]uint8
	ChannelDataStartOffset uint16
	MinorVersion           uint8
	MajorVersion           uint8
	HeaderLength           uint16
	ChannelCount           uint32
	FrameCount             uint32
	StepTimeMilliseconds   uint8
	Flags1                 uint8 `fppignored:"true" fppdefault:"0"`
	Compression            Compression
	CompressionBlockCount  uint8
	SparseRangeCount       uint8
	Flags2                 uint8 `fppignored:"true" fppdefault:"0"`
	UniqueID               uint64
}
