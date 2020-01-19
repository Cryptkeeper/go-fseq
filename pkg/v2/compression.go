package v2

type Compression uint8

func (c Compression) String() string {
	switch c {
	case None:
		return "none"
	case Zstd:
		return "zstd"
	case Zlib:
		return "zlib"
	default:
		return "unknown"
	}
}

const (
	None Compression = 0
	Zstd Compression = 1
	Zlib Compression = 2
)

type CompressionBlock struct {
	FrameNumber uint32
	Length      uint32
}
