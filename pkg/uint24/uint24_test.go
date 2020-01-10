package uint24

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestUint24(t *testing.T) {
	var buf = &bytes.Buffer{}

	// Test that encoding and decoding match for valid range
	for i := 0; i < MaxUint24; i++ {
		if err := binary.Write(buf, binary.LittleEndian, Uint24(i)); err != nil {
			t.Fatal(err)
		}

		var out Uint24
		if err := binary.Read(buf, binary.LittleEndian, &out); err != nil {
			t.Fatal(err)
		}

		if int(out) != i {
			t.Fatalf("expected %d, unmarshaled %d", i, out)
		}

		// Re-use the buffer
		buf.Reset()
	}

	// Ensure out of bounds values return error
	if err := binary.Write(&bytes.Buffer{}, binary.LittleEndian, MaxUint24+1); err == nil {
		t.Fatalf("expected err on out of bounds marshal")
	}
}
