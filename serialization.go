package akka

import (
	"reflect"
)

type Serializer interface {
	Identifier() int
	ToBinary(v interface{}) []byte

	// Returns whether this serializer needs a manifest in the fromBinary method
	IncludeManifest() bool
	FromBinary(data []byte) (v interface{}, err error)
	FromBinaryWithType(data []byte, typ reflect.Type) (v interface{}, err error)
}
