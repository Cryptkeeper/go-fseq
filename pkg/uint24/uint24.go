package uint24

import (
	"encoding/binary"
	"errors"
	"strconv"
)

const MaxUint24 = 1<<24 - 1

var errTooLarge = errors.New("cannot marshal/unmarshal Uint24 larger than MaxUint24")

type Uint24 uint32

func (u Uint24) MarshalBinary() (data []byte, err error) {
	if u > MaxUint24 {
		return nil, errTooLarge
	}
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], uint32(u))
	return b[:3], nil
}

func (u *Uint24) UnmarshalBinary(data []byte) error {
	if *u > MaxUint24 {
		return errTooLarge
	}
	val := binary.LittleEndian.Uint32(data)
	*u = Uint24(val)
	return nil
}

func (u Uint24) String() string {
	return strconv.Itoa(int(u))
}
