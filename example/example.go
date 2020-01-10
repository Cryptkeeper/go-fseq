package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"github.com/Cryptkeeper/go-fseq"
	"github.com/Cryptkeeper/go-fseq/pkg/v1"
	"github.com/Cryptkeeper/go-fseq/pkg/v2"
	"io/ioutil"
	"log"
)

func main() {
	var path = flag.String("path", "in.fseq", "input file to read")
	flag.Parse()

	b, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatal(err)
	}

	_, majorVersion := fseq.ReadVersion(b) // discard minorVersion return
	switch majorVersion {
	case v1.MajorVersion:
		// go-fseq supports the v1 format, this is simply an example usage of version matching
		log.Fatal("outdated format version, v1!")

	case v2.MajorVersion:
		var r = bytes.NewReader(b)

		// fseq files are encoded as little-endian
		var h v2.Header
		if err := binary.Read(r, binary.LittleEndian, &h); err != nil {
			log.Fatal(err)
		}

		log.Println("fseq v2 unique id:", h.UniqueID)

	default:
		log.Fatal("unknown format version!")
	}
}
