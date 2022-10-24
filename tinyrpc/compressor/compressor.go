package compressor

type CompressType uint16

const (
	Raw CompressType = iota
	Gzip
	Snappy
	Zlib
)
