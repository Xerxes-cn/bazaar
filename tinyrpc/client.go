package tinyrpc

import (
	"github.com/Xerxes-cn/bazaar/tinyrpc/compressor"
	"github.com/Xerxes-cn/bazaar/tinyrpc/serializer"
)

type Option func(o *options)

type options struct {
	compressType compressor.CompressType
	serializer   serializer.Serializer
}
