package fvar

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Var struct {
	Code [2]uint8
	Data []uint8
}

func (v Var) Len() uint16 {
	// +2 bytes since len includes the uint16 length
	return uint16(2 + len(v.Code) + len(v.Data))
}

func (v Var) String() string {
	return fmt.Sprintf("(code: %s, data: %s)", v.Code, v.Data)
}

func (v Var) CodeString() string {
	return string(v.Code[:])
}

func (v Var) DataString() string {
	// Omit the last byte if the string is null terminated
	if v.Data[len(v.Data)-1] == 0x00 {
		return string(v.Data[:len(v.Data)-1])
	}
	return string(v.Data)
}

func Write(w io.Writer, v Var) (int, error) {
	var b = make([]byte, v.Len())

	binary.LittleEndian.PutUint16(b[:2], v.Len())
	copy(b[2:4], v.Code[:])
	copy(b[4:v.Len()], v.Data)

	return w.Write(b)
}

func Read(r io.Reader) (Var, error) {
	var v Var

	// Read the first 2 bytes to determine length
	var head = make([]byte, 2)
	if n, err := r.Read(head); err != nil {
		return v, err
	} else if n != len(head) {
		return v, fmt.Errorf("expected %d, read %d", len(head), n)
	}

	// Read the length and then re-read the full body
	var length = binary.LittleEndian.Uint16(head) - 2
	var body = make([]byte, length)
	if n, err := r.Read(body); err != nil {
		return v, err
	} else if n != len(body) {
		return v, fmt.Errorf("expected %d, read %d", len(body), n)
	}

	// Copy body data into Var instance
	copy(v.Code[:], body[0:2])
	v.Data = body[2:length]

	return v, nil
}
