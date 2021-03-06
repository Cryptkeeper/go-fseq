package v1

import "time"

type Header struct {
	Identifier             [4]uint8
	ChannelDataStartOffset uint16
	MinorVersion           uint8
	MajorVersion           uint8
	HeaderLength           uint16
	ChannelCount           uint32
	FrameCount             uint32
	StepTimeMilliseconds   uint8
	Flags                  uint8         `fppignored:"true" fppdefault:"0"`
	UniverseCount          uint16        `fppignored:"true"`
	UniverseSize           uint16        `fppignored:"true"`
	Gamma                  uint8         `fppignored:"true" fppdefault:"1"`
	ColorEncoding          ColorEncoding `fppignored:"true" fppdefault:"2"`
	Reserved               [2]uint8      `fppignored:"true" fppdefault:"0"`
}

func (h Header) Duration() time.Duration {
	var framesPerSecond = time.Duration(h.StepTimeMilliseconds) * time.Millisecond
	return framesPerSecond * time.Duration(h.FrameCount)
}

func (h Header) FramesPerSecond() int {
	return 1000 / int(h.StepTimeMilliseconds)
}
