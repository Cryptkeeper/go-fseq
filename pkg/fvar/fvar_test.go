package fvar

import (
	"bytes"
	"testing"
)

var encoded = []byte{
	0x08, // low 8 bits of length
	0x00, // high 8 bits of length
	't',  // code[0]
	't',  // code[1]
	'd',  // data...
	'a',
	't',
	'a',
}

func TestRead(t *testing.T) {
	var r = bytes.NewReader(encoded)
	v, err := Read(r)
	if err != nil {
		t.Fatal(err)
	}

	if len(encoded) != int(v.Len()) {
		t.Fatalf("expected len %d, read %d", len(encoded), v.Len())
	}
	if !bytes.Equal(encoded[2:4], v.Code[:]) {
		t.Fatalf("expected code %s, read %s", encoded[2:4], v.Code)
	}
	if !bytes.Equal(encoded[4:], v.Data) {
		t.Fatalf("expected data %s, read %s", encoded[4:], v.Data)
	}
}

var decoded = Var{
	Code: [2]uint8{'t', 't'},
	Data: []uint8("data"),
}

func TestWrite(t *testing.T) {
	var b = &bytes.Buffer{}
	n, err := Write(b, decoded)
	if err != nil {
		t.Fatal(err)
	}

	if n != int(decoded.Len()) {
		t.Fatalf("expected len %d, wrote %d", decoded.Len(), n)
	}
	if !bytes.Equal(b.Bytes(), encoded) {
		t.Fatalf("expected %s, wrote %s", encoded, b.Bytes())
	}
}
