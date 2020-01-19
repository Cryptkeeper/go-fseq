package v2

import (
	"math"
	"time"
)

const (
	FppMaxCompressionBlockLength = 64 * 2014 // 2014 is not a typo, copied from https://github.com/FalconChristmas/fpp/blob/master/src/fseq/FSEQFile.cpp#L774
	FppMaxCompressionBlockCount  = math.MaxUint8
	FppWithheldFrameCount        = 10
)

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

func (h Header) Duration() time.Duration {
	var framesPerSecond = time.Duration(h.StepTimeMilliseconds) * time.Millisecond
	return framesPerSecond * time.Duration(h.FrameCount)
}

func (h Header) FramesPerSecond() int {
	return 1000 / int(h.StepTimeMilliseconds)
}

func (h Header) ComputeMaxBlocks(maxBlockLength, maxBlockCount, withheldFrameCount int) (blockCount, withheldFrameLength, framesPerBlock int) {
	withheldFrameLength = int(h.ChannelCount) * withheldFrameCount

	if h.FrameCount == 0 {
		blockCount = 0
		framesPerBlock = 0
		return
	}

	// Instantly return if frameCount fits within the withheldFrameCount
	var availableFrameCount = int(h.FrameCount) - withheldFrameCount
	if availableFrameCount <= 0 {
		blockCount = 1
		framesPerBlock = int(h.FrameCount)
		return
	}

	// blockCount does not include withheldFrameCount
	blockCount = int(math.Ceil(float64(int(h.ChannelCount)*availableFrameCount) / float64(maxBlockLength)))

	// If withholding frames, lower the maxBlockCount to ensure available space
	var withheldBlockCount = 0
	if withheldFrameCount > 0 {
		withheldBlockCount = 1
	}

	var availableBlockCount = maxBlockCount - withheldBlockCount
	if blockCount > availableBlockCount {
		blockCount = availableBlockCount
	}

	// Determine framesPerBlock, ignoring withheldFrameCount
	framesPerBlock = int(math.Ceil(float64(availableFrameCount) / float64(blockCount)))

	return
}
