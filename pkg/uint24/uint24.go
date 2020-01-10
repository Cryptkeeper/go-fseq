package uint24

import (
	"strconv"
)

const MaxUint24 = 1<<24 - 1

func NewUint24(val uint32) *Uint24 {
	var u = new(Uint24)
	u.Set(val)
	return u
}

type Uint24 [3]uint8

func (u *Uint24) Set(val uint32) {
	// panic since this is closer to how a compiler/runtime would treat an overflow compared to err returns
	if val > MaxUint24 {
		panic("cannot set Uint24 larger than uint24.MaxUint24")
	}
	(*u)[0] = uint8(val & 0xFF)
	(*u)[1] = uint8((val >> 8) & 0xFF)
	(*u)[2] = uint8((val >> 16) & 0xFF)
}

func (u Uint24) Uint32() uint32 {
	return uint32(u[0]) | uint32(u[1])<<8 | uint32(u[2])<<16
}

func (u Uint24) String() string {
	return strconv.Itoa(int(u.Uint32()))
}
