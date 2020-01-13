package main

import (
   "bytes"
   "encoding/binary"
   "flag"
   "fmt"
   "github.com/Cryptkeeper/go-fseq"
   "github.com/Cryptkeeper/go-fseq/pkg/v1"
   "github.com/Cryptkeeper/go-fseq/pkg/v2"
   "io"
   "io/ioutil"
   "log"
   "os"
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

      log.Printf("%+v\n", h)

      // If the file is compressed, dump the Compression Blocks into files
      // These will not be decompressed, simply read and separated
      if h.Compression != v2.None {
         var blocks = make([]v2.CompressionBlock, h.CompressionBlockCount)

         for i := 0; i < int(h.CompressionBlockCount); i++ {
            if err := binary.Read(r, binary.LittleEndian, &blocks[i]); err != nil {
               log.Fatal(err)
            }

            // fpp source code ignores 0 length Compression Blocks (no idea why they're encoded, alignment?)
            if blocks[i].Length == 0 {
               // Trim array since it isn't full
               blocks = blocks[:i]

               break
            }
         }

         log.Printf("%+v\n", blocks)

         // Compression Block offsets are relative to Channel Data Start Offset
         if _, err := r.Seek(int64(h.ChannelDataStartOffset), io.SeekStart); err != nil {
            log.Fatal(err)
         }

         for _, block := range blocks {
            var b = make([]uint8, block.Length)

            if n, err := r.Read(b); err != nil {
               log.Fatal(err)
            } else if n != int(block.Length) {
               log.Fatalf("n = %d, expected %d", n, block.Length)
            }

            var out = fmt.Sprintf("frame_%d.bin", block.FrameNumber)

            log.Println("writing:", out)

            // Write the Compression Block out to a file
            if err := ioutil.WriteFile(out, b, os.ModePerm); err != nil {
               log.Fatal(err)
            }
         }

         // Ensure the full file has been read
         if r.Len() != 0 {
            log.Fatalf("%d bytes remaining, should be empty!", r.Len())
         }
      }

   default:
      log.Fatal("unknown format version!")
   }
}
