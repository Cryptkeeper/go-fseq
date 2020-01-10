package uint24

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestUint24(t *testing.T) {
	var buf = &bytes.Buffer{}

	// Test that encoding and decoding match for valid range
	for i := uint32(0); i < MaxUint24; i++ {
		var u = NewUint24(i)
		if err := binary.Write(buf, binary.LittleEndian, u); err != nil {
			t.Fatal(err)
		}

		if buf.Len() != 3 {
			t.Fatalf("expected buf len 3, read %d", buf.Len())
		}

		var out Uint24
		if err := binary.Read(buf, binary.LittleEndian, &out); err != nil {
			t.Fatal(err)
		}

		if buf.Len() != 0 {
			t.Fatalf("expected buf len 0, read %d", buf.Len())
		}

		if out.Uint32() != i {
			t.Fatalf("expected %d, unmarshaled %d", i, out)
		}

		// Re-use the buffer
		buf.Reset()
	}

	// Ensure out of bounds values panic
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("out of bounds value did not panic")
		}
	}()

	NewUint24(MaxUint24 + 1)
}
