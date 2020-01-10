# go-fseq [![Go Report Card](https://goreportcard.com/badge/github.com/Cryptkeeper/go-fseq)](https://goreportcard.com/report/github.com/Cryptkeeper/go-fseq) [![GoDoc](https://godoc.org/github.com/Cryptkeeper/go-fseq?status.svg)](https://godoc.org/github.com/Cryptkeeper/go-lightorama/)
A bare-bones Go library for encoding and decoding the v1 and v2 versions of the [FSEQ file format](https://github.com/FalconChristmas/fpp/blob/master/docs/FSEQ_Sequence_File_Format.txt) used by the [Falcon Player](https://github.com/FalconChristmas/fpp) ("fpp") software. Given the file format's outdated and rather lacking documentation, much of this library has been reserve engineered from the [FSEQ C++ implementation](https://github.com/FalconChristmas/fpp/blob/master/src/fseq/FSEQFile.cpp) and may contain errors or not work with all files. Use at your own risk.

This is merely a third-party implementation of the FSEQ file format and I am not responsible for managing it as a standard.

## Usage
### Installation
Install using `go get github.com/Cryptkeeper/go-fseq`

### Example Usage
See [example/example.go](example/example.go)

### Tests
Tests have been provided for the `uint24` and `fvar` types in their corresponding packages. Use `go test ./...` in the root directory to run the tests.

## Notes
- The `v1` and `v2` packages correspond to the FSEQ file format versions, not `go-fseq` versions.
- Additional data structures used by the file format have been provided as structs. Their usage is left as an exercise to the reader.
- Although the FSEQ file format is effectively maintained by the [fpp](https://github.com/FalconChristmas/fpp) project which originated it, the fpp does not use all features of the file format. As such, some fields include tags (`fppignored` or `fppdefault`) which specify that field's behavior within fpp.
- ESEQ is a specialized subset of the FSEQ file format and is not supported by the library. Some documentation for it has been provided for it by the [fpp](https://github.com/FalconChristmas/fpp/blob/master/docs/ESEQ_Effect_Sequence_file_format.txt) project.
- While FSEQ files have a `.fseq` file extension, they may have an identifier value of "PSEQ" ([documentation reference](https://github.com/FalconChristmas/fpp/blob/master/docs/FSEQ_Sequence_File_Format.txt)). Older implementations may instead expect or use a "FSEQ" identifier value.