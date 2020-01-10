package v2

type Compression uint8

const (
	None Compression = 0
	Zstd Compression = 1
	Zlib Compression = 2
)

type CompressionBlock struct {
	FrameNumber uint32
	Length      uint32
}
