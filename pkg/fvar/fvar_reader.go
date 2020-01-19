package fvar

import "bytes"

func ReadAll(b []byte) ([]Var, error) {
	var r = bytes.NewReader(b)
	var out = make([]Var, 0)

	// 4 is derived from the roundTo4 behavior that occurs when encoding FSEQ files
	for r.Len() > 4 {
		v, err := Read(r)
		if err != nil {
			return nil, err
		}

		out = append(out, v)
	}

	return out, nil
}
